package api

import (
	"github.com/crisguitar/paraules-noves/internal/infrastructure"
	"github.com/crisguitar/paraules-noves/internal/words"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"os"
	"strconv"
)

func NewRouter() *chi.Mux {
	dbConfig := getDbConfig()
	wordsRepository := words.NewRepository(dbConfig)

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Method("POST", "/words", words.NewCreateWordHandler(wordsRepository))
	r.Method("GET", "/words", words.NewGetAllWordsHandler(wordsRepository))

	return r
}

func getDbConfig() infrastructure.DbConfig {
	port := os.Getenv("DB_PORT")
	portNumber, _ := strconv.Atoi(port)
	return infrastructure.DbConfig{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		Host:     os.Getenv("DB_HOST"),
		Port:     portNumber,
	}
}
