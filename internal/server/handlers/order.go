package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	db "github.com/stellafff25/Lab5/db/sqlc"
)

type OrderHandler struct {
	store db.Store
}

func NewOrderHandler(store db.Store) *OrderHandler {
	return &OrderHandler{store: store}
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name   string `json:"name"`
		Amount int32  `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if input.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	order, err := h.store.CreateOrder(r.Context(), db.CreateOrderParams{
		Name:   input.Name,
		Amount: input.Amount,
	})

	if err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	order, err := h.store.GetOrder(r.Context(), id)
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(order)
}

func (h *OrderHandler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
	orders, err := h.store.GetAllOrders(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch orders", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

func (h *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var input struct {
		Name   string `json:"name"`
		Amount int32  `json:"amount"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if input.Name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return
	}

	order, err := h.store.UpdateOrder(r.Context(), db.UpdateOrderParams{
		ID:     id,
		Name:   input.Name,
		Amount: input.Amount,
	})

	if err != nil {
		http.Error(w, "Failed to update order", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(order)
}

func (h *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err = h.store.DeleteOrder(r.Context(), id)
	if err != nil {
		http.Error(w, "Failed to delete order", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
