package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AbderraoufKhorchani/authentification-service/internal/helpers"
	"github.com/AbderraoufKhorchani/authentification-service/web"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const webPort = "8080"

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
	if err := godotenv.Load(); err != nil {
		log.Println("Can't load DSN")
	}

	dsn := os.Getenv("DSN")

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
