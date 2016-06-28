package root

import (
"conf"
"fmt"
"net/http"
"gopkg.in/pg.v4"
)

func createSchema(db *pg.DB) error {
	queries := []string{
		`CREATE TYPE relation AS ENUM ('liked', 'disliked', 'matched')`,
		`CREATE TABLE users (
			id serial primary key, 
			name text, 
			type text
			)`,
		`CREATE TABLE relationships (
			id integer references users(id), 
			user_id integer references users(id),
			state relation, 
			type text,
			PRIMARY KEY (id, user_id)
			)`,
	}
	for _, q := range queries {
		_, err := db.Exec(q)
		if err != nil {
			return err
		}
	}
	return nil
}

func DefaultRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Println("welcome")
	w.Write([]byte("welcome!\n"))
}

func InitDatabase(w http.ResponseWriter, r *http.Request) {
	db := pg.Connect(conf.GetDbConf())
	err := createSchema(db)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Write([]byte("Datatable create ok!\n"))
	}
}