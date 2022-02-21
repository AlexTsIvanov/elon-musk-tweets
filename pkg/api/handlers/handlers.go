package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AlexTsIvanov/elon-musk-twitter/pkg/api/service"
)

type Handler struct {
	s *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{s: s}
}

func (h *Handler) TweetsPerDay(rw http.ResponseWriter, r *http.Request) {
	data, err := h.s.TweetsPerDay()
	if err != nil {
		fmt.Println(err)
		// http.Error(rw, "Unable to fetch from database", http.StatusInternalServerError)
		return
	}
	e := json.NewEncoder(rw)
	err = e.Encode(data)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) RetweetsPerDay(rw http.ResponseWriter, r *http.Request) {
	data, err := h.s.RetweetsPerDay()
	if err != nil {
		http.Error(rw, "Unable to fetch from database", http.StatusInternalServerError)
		return
	}
	e := json.NewEncoder(rw)
	err = e.Encode(data)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) MostLikedTweet(rw http.ResponseWriter, r *http.Request) {
	data, err := h.s.MostLikedTweet()
	if err != nil {
		http.Error(rw, "Unable to fetch from database", http.StatusInternalServerError)
		return
	}
	e := json.NewEncoder(rw)
	err = e.Encode(data)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) LeastLikedTweet(rw http.ResponseWriter, r *http.Request) {
	data, err := h.s.LeastLikedTweet()
	if err != nil {
		http.Error(rw, "Unable to fetch from database", http.StatusInternalServerError)
		return
	}
	e := json.NewEncoder(rw)
	err = e.Encode(data)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) TweetsPerHour(rw http.ResponseWriter, r *http.Request) {
	data, err := h.s.TweetsPerHour()
	if err != nil {
		http.Error(rw, "Unable to fetch from database", http.StatusInternalServerError)
		return
	}
	e := json.NewEncoder(rw)
	err = e.Encode(data)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}
