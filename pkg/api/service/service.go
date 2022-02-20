package service

import (
	"context"
	"fmt"

	"github.com/AlexTsIvanov/elon-musk-twitter/pkg/api/repository"
	"go.mongodb.org/mongo-driver/bson"
)

type Service struct {
	r *repository.Repository
}

func NewService(r *repository.Repository) *Service {
	return &Service{r: r}
}

func (s *Service) TweetsPerDay() ([]bson.M, error) {
	data, err := s.r.TweetsPerDay()
	if err != nil {
		return nil, err
	}
	var results []bson.M
	err = data.All(context.Background(), &results)
	return results, err
}

func (s *Service) RetweetsPerDay() ([]bson.M, error) {
	data, err := s.r.RetweetsPerDay()
	if err != nil {
		return nil, err
	}
	var results []bson.M
	err = data.All(context.Background(), &results)
	return results, err
}

func (s *Service) MostLikedTweet() ([]bson.M, error) {
	data, err := s.r.MostLikedTweet()
	if err != nil {
		return nil, err
	}
	var results []bson.M
	err = data.All(context.Background(), &results)
	return results, err
}

func (s *Service) LeastLikedTweet() ([]bson.M, error) {
	data, err := s.r.LeastLikedTweet()
	if err != nil {
		return nil, err
	}
	var results []bson.M
	err = data.All(context.Background(), &results)
	return results, err
}

func (s *Service) TweetsPerHour() ([]bson.M, error) {
	var finalResults []bson.M
	for i := 0; i <= 23; i++ {
		var results []bson.M
		var start, end string

		if i <= 9 {
			start = fmt.Sprintf("0%d:00:00", i)
			end = fmt.Sprintf("0%d:00:00", i+1)
		} else {
			start = fmt.Sprintf("%d:00:00", i)
			end = fmt.Sprintf("%d:00:00", i+1)
		}

		key := fmt.Sprintf("%d-%d", i, i+1)

		data, err := s.r.TweetsPerHour(start, end, key)
		if err != nil {
			return nil, err
		}
		err = data.All(context.Background(), &results)
		if err != nil {
			return nil, err
		}
		finalResults = append(finalResults, results...)
	}
	return finalResults, nil
}
