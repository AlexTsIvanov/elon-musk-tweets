package dbpopulation

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/AlexTsIvanov/elon-musk-twitter/pkg/api/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DownloadFile(url, fileName string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(fmt.Sprintf("./%s", fileName))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	return nil
}

func InsertEntriesInDB(db *mongo.Database, url, fileName string) error {
	err := DownloadFile(url, fileName)
	if err != nil {
		log.Printf("Couldn't download %s", fileName)
		return err
	}

	jsonFile, err := os.Open(fileName)
	if err != nil {
		log.Printf("Couldn't open %s", fileName)
		return err
	}
	defer jsonFile.Close()

	bsonTweets := bson.A{}
	scanner := bufio.NewScanner(jsonFile)
	for scanner.Scan() {
		var line entity.Tweet
		err = json.Unmarshal([]byte(scanner.Text()), &line)
		if err != nil {
			log.Println("Cannot unmarshal into json")
			return err
		}
		bsonTweets = append(bsonTweets, line)
	}

	err = db.Collection("tweets").Drop(context.Background())
	if err != nil {
		log.Println("Cannot unmarshal into json")
		return err
	}
	_, err = db.Collection("tweets").InsertMany(context.Background(), bsonTweets)
	if err != nil {
		log.Println("Cannot insert into database")
		return err
	}
	return nil
}
