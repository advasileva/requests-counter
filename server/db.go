package main

import (
	"fmt"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type User struct {
	Id       int64
	Requests int64
}

func GetRequests(id int64) int64 {
	db := connect()
	defer db.Close()
	user := &User{Id: id}
	err := db.Model(user).WherePK().Select()
	if err != nil {
		panic(err)
	}
	return user.Requests
}

func IncRequests(id int64) {
	db := connect()
	defer db.Close()
	user := &User{Id: id}
	exists, err := db.Model(user).WherePK().Exists()
	if err != nil {
		panic(err)
	}
	if exists {
		err = db.Model(user).WherePK().Select()
		if err != nil {
			panic(err)
		}
		user.Requests++
		_, err = db.Model(user).WherePK().Update(user)
		if err != nil {
			panic(err)
		}
	} else {
		user = &User{
			Id:       id,
			Requests: 1,
		}
		_, err = db.Model(user).Insert()
		if err != nil {
			panic(err)
		}
	}

}

func connect() (db *pg.DB) {
	addr := fmt.Sprintf("%s:%s", os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"))
	db = pg.Connect(&pg.Options{
		Database: os.Getenv("POSTGRES_DB"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Addr:     addr,
	})

	err := createSchema(db)
	if err != nil {
		panic(err)
	}

	return
}

func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*User)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
