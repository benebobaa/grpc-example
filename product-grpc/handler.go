package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	pb "simple-grpc-2/proto"
)

type Handler struct {
	userClient pb.AuthUserClient
}

func NewHandler(userClientGrpc pb.AuthUserClient) *Handler {
	return &Handler{
		userClient: userClientGrpc,
	}
}

type CheckoutRequest struct {
	ProductID string `json:"product_id"`
	Stock     int    `json:"stock"`
}

type Response struct {
	IsTokenValid bool            `json:"token_valid"`
	Message      string          `json:"message"`
	Data         CheckoutRequest `json:"data"`
}

func (h *Handler) CheckoutProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid HttpMethod", http.StatusMethodNotAllowed)
		return
	}

	token := r.Header.Get("token")

	var req CheckoutRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if req.Stock == 0 || req.ProductID == "" {
		http.Error(w, "Invalid Request, stock cannot 0 and product id cannot empty", http.StatusBadRequest)
		return
	}

	response, err := h.userClient.CheckToken(r.Context(), &pb.TokenRequest{
		Token: token,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}

	if !response.GetStatus() {
		http.Error(w, fmt.Sprintf("check token failed: %s", response.GetMessage()), http.StatusUnauthorized)
		return
	}

	sendJSONResponse(w, &Response{
		IsTokenValid: response.GetStatus(),
		Message:      "checkout product success",
		Data:         req,
	})
}

func sendJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Could not encode response to JSON", http.StatusInternalServerError)
	}
}
