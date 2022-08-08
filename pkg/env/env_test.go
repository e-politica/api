package env

import (
	"os"
	"testing"
)

func TestGet(t *testing.T) {
	envName := "test_env"

	t.Run("test_get_string", func(t *testing.T) {
		exp := "abc"
		if err := os.Setenv(envName, exp); err != nil {
			t.Fatal(err)
		}
		if recv := Get(envName, ""); recv != exp {
			t.Errorf("recv: %v, exp: %v", recv, exp)
		}
	})

	t.Run("test_get_int", func(t *testing.T) {
		exp := 123
		if err := os.Setenv(envName, "123"); err != nil {
			t.Fatal(err)
		}
		if recv := Get(envName, 0); recv != exp {
			t.Errorf("recv: %v, exp: %v", recv, exp)
		}
	})

	t.Run("test_get_bool", func(t *testing.T) {
		exp := true
		if err := os.Setenv(envName, "true"); err != nil {
			t.Fatal(err)
		}
		if recv := Get(envName, false); recv != exp {
			t.Errorf("recv: %v, exp: %v", recv, exp)
		}
	})
}

func TestGetSpecific(t *testing.T) {
	envName := "test_env"

	t.Run("test_get_specific_string", func(t *testing.T) {
		exp := "abc"
		if err := os.Setenv(envName, exp); err != nil {
			t.Fatal(err)
		}
		if recv := GetString(envName, ""); recv != exp {
			t.Errorf("recv: %v, exp: %v", recv, exp)
		}
	})

	t.Run("test_get_specific_int", func(t *testing.T) {
		exp := 123
		if err := os.Setenv(envName, "123"); err != nil {
			t.Fatal(err)
		}
		if recv := GetInt(envName, 0); recv != exp {
			t.Errorf("recv: %v, exp: %v", recv, exp)
		}
	})

	t.Run("test_get_specific_bool", func(t *testing.T) {
		exp := true
		if err := os.Setenv(envName, "true"); err != nil {
			t.Fatal(err)
		}
		if recv := GetBool(envName, false); recv != exp {
			t.Errorf("recv: %v, exp: %v", recv, exp)
		}
	})
}

func BenchmarkGet(b *testing.B) {
	envName := "test_env"

	b.Run("test_get_string", func(b *testing.B) {
		exp := "abc"
		if err := os.Setenv(envName, exp); err != nil {
			b.Fatal(err)
		}
		if recv := Get(envName, ""); recv != exp {
			b.Errorf("recv: %v, exp: %v", recv, exp)
		}
	})

	b.Run("test_get_int", func(b *testing.B) {
		exp := 123
		if err := os.Setenv(envName, "123"); err != nil {
			b.Fatal(err)
		}
		if recv := Get(envName, 0); recv != exp {
			b.Errorf("recv: %v, exp: %v", recv, exp)
		}
	})

	b.Run("test_get_bool", func(b *testing.B) {
		exp := true
		if err := os.Setenv(envName, "true"); err != nil {
			b.Fatal(err)
		}
		if recv := Get(envName, false); recv != exp {
			b.Errorf("recv: %v, exp: %v", recv, exp)
		}
	})
}

func BenchmarkGetSpecific(b *testing.B) {
	envName := "test_env"

	b.Run("test_get_specific_string", func(b *testing.B) {
		exp := "abc"
		if err := os.Setenv(envName, exp); err != nil {
			b.Fatal(err)
		}
		if recv := GetString(envName, ""); recv != exp {
			b.Errorf("recv: %v, exp: %v", recv, exp)
		}
	})

	b.Run("test_get_specific_int", func(b *testing.B) {
		exp := 123
		if err := os.Setenv(envName, "123"); err != nil {
			b.Fatal(err)
		}
		if recv := GetInt(envName, 0); recv != exp {
			b.Errorf("recv: %v, exp: %v", recv, exp)
		}
	})

	b.Run("test_get_specific_bool", func(b *testing.B) {
		exp := true
		if err := os.Setenv(envName, "true"); err != nil {
			b.Fatal(err)
		}
		if recv := GetBool(envName, false); recv != exp {
			b.Errorf("recv: %v, exp: %v", recv, exp)
		}
	})
}
