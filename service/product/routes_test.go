package product

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddProductHandler(t *testing.T) {
	store := &mockStore{}
	service := NewService(store)
	handler := NewHandler(service)

	t.Run("failed to add the product on empty payload", func(t *testing.T) {
		payload := new(productPayload)
		marshalled, _ := json.Marshal(payload)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/api/v1/products", bytes.NewBuffer(marshalled))
		handler.addProduct()(w, r)

		if w.Code != http.StatusBadRequest {
			t.Errorf("expected %d, got %d", http.StatusBadRequest, w.Code)
		}
	})

	t.Run("add the product on valid payload", func(t *testing.T) {
		payload := &productPayload{Name: "breed"}
		marshalled, _ := json.Marshal(payload)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodPost, "/api/v1/products", bytes.NewBuffer(marshalled))
		handler.addProduct()(w, r)

		if w.Code != http.StatusCreated {
			t.Errorf("expected %d, got %d", http.StatusBadRequest, w.Code)
		}
	})
}

type mockStore struct{}

func (s *mockStore) getAll() ([]product, error) {
	return nil, nil
}

func (s *mockStore) getByID(id string) (*product, error) {
	return nil, nil
}

func (s *mockStore) getByName(name string) (*product, error) {

	return nil, errors.New("product not found")
}

func (s *mockStore) create(name string) error {
	return nil
}
