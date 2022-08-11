package database

import (
	"context"
	"errors"
	"log"
	"net/url"
	"sync"
	"time"

	"github.com/e-politica/api/config"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

var ErrDatabaseDown = errors.New("database down")

type Db struct {
	Once       *sync.Once
	Ctx        *context.Context
	ReopenConn chan bool
	IsWaiting  bool
	mu         sync.Mutex
	Conn       *pgx.Conn
}

func New() (db *Db) {
	err := db.Connect()
	if err != nil {
		log.Println("could not connect on the database", err)
		db.WaitForConnection()
	}

	return db
}

func (db *Db) Connect() (err error) {
	db.Once.Do(func() {
		url := "postgres://" +
			url.UserPassword(config.DbUser, config.DbPassword).String() +
			"@" + config.DbHost +
			":" + config.DbPort +
			"/" + config.DbName

		var conn *pgx.Conn
		conn, err = pgx.Connect(*db.Ctx, url)
		db.Conn = conn
	})
	return
}

func (db *Db) WaitForConnection() {
	if !config.DbWaitStart {
		return
	}

	db.mu.Lock()
	defer db.mu.Unlock()

	db.IsWaiting = true
	log.Println("=== WAITING FOR CONNECTION ===")
	log.Println("Start")
	for {
		log.Println("... trying to restore Db Db ...")
		if db.Conn == nil || db.Conn.Ping(*db.Ctx) != nil {
			time.Sleep(time.Second * time.Duration(config.DbReconnectSec))
			db.Once = &sync.Once{}
			db.Connect()
			continue
		}
		break
	}
	log.Println("End")
	log.Println("=== CONNECTION RESTORED ===")
	db.IsWaiting = false
}

func (db *Db) LoopCheckConnection() {
	for reopen := range db.ReopenConn {
		if reopen {
			go db.WaitForConnection()
		}
	}
}

func (db *Db) Reconnect() {
	if db.IsWaiting {
		return
	}
	db.ReopenConn <- true
}

func (db *Db) Exec(sql string, args ...interface{}) (pgconn.CommandTag, error) {
	if db.Conn == nil || db.Conn.Ping(*db.Ctx) != nil {
		go db.Reconnect()
		return nil, ErrDatabaseDown
	}
	return db.Conn.Exec(*db.Ctx, sql, args...)
}

func (db *Db) Query(sql string, args ...interface{}) (pgx.Rows, error) {
	if db.Conn == nil || db.Conn.Ping(*db.Ctx) != nil {
		go db.Reconnect()
		return nil, ErrDatabaseDown
	}
	return db.Conn.Query(*db.Ctx, sql, args...)
}

func (db *Db) QueryRow(sql string, args ...interface{}) (pgx.Row, error) {
	if db.Conn == nil || db.Conn.Ping(*db.Ctx) != nil {
		go db.Reconnect()
		return nil, ErrDatabaseDown
	}
	return db.Conn.QueryRow(*db.Ctx, sql, args...), nil
}

func (db *Db) QueryFunc(sql string, args []interface{}, scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error) {
	if db.Conn == nil || db.Conn.Ping(*db.Ctx) != nil {
		go db.Reconnect()
		return nil, ErrDatabaseDown
	}
	return db.Conn.QueryFunc(*db.Ctx, sql, args, scans, f)
}
