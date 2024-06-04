package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dysdimas/internal/usecase"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type FizzBuzzHandler struct {
	Usecase usecase.FizzBuzzUsecase
}

func NewFizzBuzzHandler(u usecase.FizzBuzzUsecase) *FizzBuzzHandler {
	return &FizzBuzzHandler{Usecase: u}
}

// Fizzbuzz handler for return the response.
func (h *FizzBuzzHandler) RangeFizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	fromStr := r.URL.Query().Get("from")
	toStr := r.URL.Query().Get("to")

	from, err1 := strconv.Atoi(fromStr)
	to, err2 := strconv.Atoi(toStr)

	if err1 != nil || err2 != nil {
		response := Response{
			Status:  http.StatusBadRequest,
			Message: "Invalid range parameters",
			Data:    nil,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	results, err := h.Usecase.RangeFizzBuzz(from, to)
	if err != nil {
		response := Response{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := Response{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    strings.Join(results, " "),
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)

	latency := time.Since(start)
	log.Printf("Request: from=%d to=%d, Response: %s, Latency: %s\n", from, to, response.Data, latency)
}
