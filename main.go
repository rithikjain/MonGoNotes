package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
	"os"
	"time"
)

func DbConnect() (*mongo.Database, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DATABASE_URL")))
	if err != nil {
		return nil, err
	}
	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	db := client.Database("notesapi")
	return db, nil
}

func GetPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		fmt.Println("INFO: No PORT environment variable detected, defaulting to 3000")
		return "localhost:3000"
	}
	return ":" + port
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := DbConnect()
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err.Error())
	}
	fmt.Println("Connected to DB...")

	notesCollection := db.Collection("notes")

	r := http.NewServeMux()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Yep, you hit the api endpoint!"))
		return
	})

	fmt.Println("Serving...")
	log.Fatal(http.ListenAndServe(GetPort(), r))
}
