package database

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var pool *sql.DB
var inited bool = false

func Init() error {
	if inited {
		return nil
	} else {
		p, err := sql.Open("mysql", "root:cjq19990704@(localhost:3306)/test")
		if err != nil {
			pool = nil
			println(err.Error())
			return err
		}
		inited = true
		pool = p
		return nil
	}
}

func Ping(ctx context.Context) {
	if Init() != nil {
		println("init error")
		return
	}

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err := pool.PingContext(ctx); err != nil {
		println("ping error")
	}
}

func Query() {

}
