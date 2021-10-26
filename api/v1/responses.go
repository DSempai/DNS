package api

import (
	"encoding/json"
	"net/http"
)

const (
	failed = "failed"
)

type Response struct {
	Status      string      `json:"status,omitempty"`
	Code        int         `json:"code,omitempty"`
	Description string      `json:"description,omitempty"`
	Result      interface{} `json:"result,omitempty"`
}

func OKResponse(w http.ResponseWriter, payload interface{}) {
	body, _ := json.Marshal(payload)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(body)
}

func BadRequestResponse(w http.ResponseWriter, err error) {
	payload := Response{
		Status:      failed,
		Code:        http.StatusBadRequest,
		Description: err.Error(),
	}
	body, _ := json.Marshal(payload)
	w.WriteHeader(http.StatusBadRequest)
	w.Write(body)
}

func NotFoundResponse(w http.ResponseWriter, err error) {
	payload := Response{
		Status:      failed,
		Code:        http.StatusNotFound,
		Description: err.Error(),
	}
	body, _ := json.Marshal(payload)
	w.WriteHeader(http.StatusNotFound)
	w.Write(body)
}
