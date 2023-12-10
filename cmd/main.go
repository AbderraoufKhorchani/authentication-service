package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AbderraoufKhorchani/authentification-service/internal/helpers"
	"github.com/AbderraoufKhorchani/authentification-service/web"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	webPort = "8080"
	dsn     = "host=localhost port=5432 user=postgres password=password dbname=users sslmode=disable timezone=UTC connect_timeout=5"
)

var counts int64

func main() {
	log.Println("Starting authentication service")

	conn := connectToDB()
	if conn == nil {
		log.Panic("Can't connect to Postgres!")
	}

	helpers.New(conn)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: web.Routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}
func connectToDB() *gorm.DB {

	for {
		connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Println("Postgres not yet ready ...")
			log.Println(err)
			counts++
		} else {
			log.Println("Connected to Postgres!")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds....")
		time.Sleep(2 * time.Second)
		continue
	}
}
