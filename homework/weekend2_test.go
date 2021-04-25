package homework

import (
	"database/sql"
	"errors"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	xerrors "github.com/pkg/errors"
)

func TestUserDao_GetOne(t *testing.T) {
	db, err := initDB()
	if err != nil {
		t.Fatalf("init: %s", err.Error())
	}
	dao := &UserDao{db: db, Table: "user"}

	user, err := dao.GetOne(1)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			t.Log("sql no rows")
		} else {
			t.Fatalf("get err: %+v", err)
		}
	}

	t.Log(user)
}

func initDB() (db *sqlx.DB, err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		err = xerrors.Wrap(err, "initDB error")
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}
