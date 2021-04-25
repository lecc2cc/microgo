package homework

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	xerrors "github.com/pkg/errors"
)

type UserDao struct {
	Table string
	db    sqlx.ExtContext
}

type User struct {
	UID  uint64 `db:"uid"`
	Name string `db:"username"`
}

func (c *UserDao) GetOne(uid uint64) (v *User, err error) {
	v = &User{}

	querySQL := fmt.Sprintf(`SELECT * FROM %s WHERE uid=? LIMIT 1;`,
		c.Table,
	)

	err = sqlx.GetContext(
		context.Background(),
		c.db,
		v,
		querySQL,
		appModule,
	)

	if err != nil {
		err = xerrors.Wrap(err, "UserDao GetOne err: "+querySQL)
	}
	return
}
