package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"lab4/internal/handler"
	"lab4/internal/repository"
	"lab4/internal/repository/postgres"
	"lab4/internal/service"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	if err := godotenv.Load("configs/db.env", "configs/start.env"); err != nil {
		log.Fatal("FAILED START", err)
	}

	config := postgres.Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
		SSLMode:  "disable",
	}
	log.Println(config)

	db, err := postgres.NewPostgresDB(config)

	if err != nil {
		log.Fatal("FAILED TO CREATE DB INSTANCE", err)
	}

	repo := repository.NewPostgresRepository(db)
	srvc := service.NewService(repo)
	mux := handler.NewHandler(srvc).InitRoutes()

	err = run(os.Getenv("PORT"), mux)
	if err != nil {
		log.Fatal("FAILED TO START SERVER", err)
	}
}

func run(port string, handler http.Handler) error {
	server := &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return server.ListenAndServe()
}
