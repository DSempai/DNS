package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	failed = "failed"
)

type Response struct {
	Status string `json:"status,omitempty"`
	Code int `json:"code,omitempty"`
	Description string `json:"description,omitempty"`
	Result interface{} `json:"result,omitempty"`
}

func OKResponse(w http.ResponseWriter, payload interface{}) {
	body, _ := json.Marshal(payload)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(body)
}

func BadRequestResponse(w http.ResponseWriter, err error){
	payload := Response{
		Status:      failed,
		Code:        http.StatusBadRequest,
		Description: fmt.Sprintf("%s", err.Error()),
	}
	body, _ := json.Marshal(payload)
	w.WriteHeader(http.StatusBadRequest)
	_, _ = w.Write(body)
}

func NotFoundResponse(w http.ResponseWriter, err error) {
	payload := Response{
		Status:      failed,
		Code:        http.StatusNotFound,
		Description: fmt.Sprintf("%s", err.Error()),
	}
	body, _ := json.Marshal(payload)
	w.WriteHeader(http.StatusNotFound)
	_, _ = w.Write(body)
}