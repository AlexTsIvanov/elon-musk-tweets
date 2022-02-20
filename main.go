package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/AlexTsIvanov/elon-musk-twitter/pkg/api/handlers"
	"github.com/AlexTsIvanov/elon-musk-twitter/pkg/api/repository"
	"github.com/AlexTsIvanov/elon-musk-twitter/pkg/api/service"
	"github.com/AlexTsIvanov/elon-musk-twitter/pkg/logic/dbpopulation"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	newEntries := flag.Bool("importdb", false, "Set true if you want to import entries in the DB")
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env present")
	}

	url := os.Getenv("URL")
	fileName := os.Getenv("FILENAME")
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("'MONGDB_URI' must set")
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("elon-musk")

	if *newEntries {
		err = dbpopulation.InsertEntriesInDB(db, url, fileName)
		if err != nil {
			log.Fatal(err)
		}
	}

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handlers := handlers.NewHandler(service)

	router := mux.NewRouter()
	router.HandleFunc("/api/tweetsperday", handlers.TweetsPerDay)
	router.HandleFunc("/api/retweetsperday", handlers.RetweetsPerDay)
	router.HandleFunc("/api/mostlikedtweet", handlers.MostLikedTweet)
	router.HandleFunc("/api/leastlikedtweet", handlers.LeastLikedTweet)
	router.HandleFunc("/api/tweetshourly", handlers.TweetsPerHour)
	router.Use(handlers.Cors)

	l := log.New(os.Stdout, "elon-musk-twitter", log.LstdFlags)
	s := http.Server{
		Addr:         ":9090",
		Handler:      router,            // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}
	// start the server
	go func() {
		l.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
