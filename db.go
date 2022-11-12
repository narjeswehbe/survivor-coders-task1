package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345678"
	dbname   = "Students"
)

func connect() *gorm.DB {

	//connecting db
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("errorrrrr")
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)

	}
	//connecting gorm
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		fmt.Println("gorm error")
	} else {
		fmt.Println("gorm success")
	}
	err2 := gormDB.Migrator().AutoMigrate()
	if err2 != nil {

	}
	fmt.Println("Successfully connected!")
	return gormDB
}
