package routes

import (
	"github.com/e-politica/api/pkg/database"
	"github.com/e-politica/api/pkg/log"
)

type Tools struct {
	Db     *database.Db
	Logger *log.Logger
}
