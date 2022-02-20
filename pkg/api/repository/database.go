package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{db: db}
}

func (r *Repository) TweetsPerDay() (*mongo.Cursor, error) {
	groupStage := bson.D{{"$group", bson.D{{"_id", "$date"}, {"count", bson.D{{"$sum", 1}}}}}}
	sortStage := bson.D{{"$sort", bson.D{{"_id", -1}}}}
	data, err := r.db.Collection("tweets").Aggregate(context.Background(), mongo.Pipeline{groupStage, sortStage})
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) RetweetsPerDay() (*mongo.Cursor, error) {
	groupStage := bson.D{{"$group", bson.D{{"_id", "$date"}, {"count", bson.D{{"$sum", "$retweetscount"}}}}}}
	sortStage := bson.D{{"$sort", bson.D{{"_id", -1}}}}
	data, err := r.db.Collection("tweets").Aggregate(context.Background(), mongo.Pipeline{groupStage, sortStage})
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) MostLikedTweet() (*mongo.Cursor, error) {
	opts := options.Find()
	opts.SetSort(bson.D{{"likescount", -1}})
	opts.SetLimit(1)
	data, err := r.db.Collection("tweets").Find(context.Background(), bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) LeastLikedTweet() (*mongo.Cursor, error) {
	opts := options.Find()
	opts.SetSort(bson.D{{"likescount", 1}})
	opts.SetLimit(1)
	data, err := r.db.Collection("tweets").Find(context.Background(), bson.D{}, opts)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *Repository) TweetsPerHour(start, end, key string) (*mongo.Cursor, error) {
	matchStage := bson.D{{"$match", bson.D{{"time", bson.D{{"$gte", start}, {"$lt", end}}}}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", key}, {"count", bson.D{{"$sum", 1}}}}}}
	data, err := r.db.Collection("tweets").Aggregate(context.Background(), mongo.Pipeline{matchStage, groupStage})
	if err != nil {
		return nil, err
	}
	return data, nil
}
