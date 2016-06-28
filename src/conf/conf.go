package conf

import (
	"gopkg.in/pg.v4"
)

func GetDbConf() *pg.Options {
	return &pg.Options{
		Addr: "localhost:5432",
		User: "postgres",
		Password: "123456",
		Database: "webdb",
	}
}
