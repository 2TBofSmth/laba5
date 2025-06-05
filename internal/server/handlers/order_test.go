package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	db "github.com/stellafff25/Lab5/db/sqlc"
)

type DummyStore struct{}

func (ds *DummyStore) CreateOrder(ctx context.Context, arg db.CreateOrderParams) (db.Order, error) {
	return db.Order{
		ID:     1,
		Name:   arg.Name,
		Amount: arg.Amount,
	}, nil
}

func (ds *DummyStore) GetOrder(ctx context.Context, id int64) (db.Order, error) {
	return db.Order{
		ID:     id,
		Name:   "TestOrder",
		Amount: 100,
	}, nil
}

func (ds *DummyStore) UpdateOrder(ctx context.Context, arg db.UpdateOrderParams) (db.Order, error) {
	return db.Order{
		ID:     arg.ID,
		Name:   arg.Name,
		Amount: arg.Amount,
	}, nil
}

func (ds *DummyStore) DeleteOrder(ctx context.Context, id int64) error {
	return nil
}

func (ds *DummyStore) GetAllOrders(ctx context.Context) ([]db.Order, error) {
	return []db.Order{
		{ID: 1, Name: "Order1", Amount: 100},
		{ID: 2, Name: "Order2", Amount: 200},
	}, nil
}

func TestCreateOrder(t *testing.T) {
	dummyStore := &DummyStore{}
	handler := NewOrderHandler(dummyStore)

	body, _ := json.Marshal(map[string]interface{}{
		"name":   "TestOrder",
		"amount": 100,
	})

	req := httptest.NewRequest("POST", "/orders", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.CreateOrder(rec, req)

	if rec.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, rec.Code)
	}

	var result db.Order
	json.NewDecoder(rec.Body).Decode(&result)

	if result.Name != "TestOrder" {
		t.Errorf("Expected name 'TestOrder', got %v", result.Name)
	}
}

func TestGetOrder(t *testing.T) {
	dummyStore := &DummyStore{}
	handler := NewOrderHandler(dummyStore)

	req := httptest.NewRequest("GET", "/orders/1", nil)
	rec := httptest.NewRecorder()

	handler.GetOrder(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}
}

func TestUpdateOrder(t *testing.T) {
	dummyStore := &DummyStore{}
	handler := NewOrderHandler(dummyStore)

	body, _ := json.Marshal(map[string]interface{}{
		"name":   "UpdatedOrder",
		"amount": 150,
	})

	req := httptest.NewRequest("PUT", "/orders/1", bytes.NewReader(body))
	rec := httptest.NewRecorder()

	handler.UpdateOrder(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rec.Code)
	}
}

func TestDeleteOrder(t *testing.T) {
	dummyStore := &DummyStore{}
	handler := NewOrderHandler(dummyStore)

	req := httptest.NewRequest("DELETE", "/orders/1", nil)
	rec := httptest.NewRecorder()

	handler.DeleteOrder(rec, req)

	if rec.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, got %d", http.StatusNoContent, rec.Code)
	}
}
