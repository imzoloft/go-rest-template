package product

import (
	"net/http"

	"github.com/imzoloft/go-rest-api/httputil"
	"github.com/imzoloft/go-rest-api/response"
)

type service struct {
	store productStore
}

func NewService(store productStore) *service {
	return &service{
		store: store,
	}
}

func (s *service) validatePayload(r *http.Request) (*productPayload, response.Map) {
	payload := new(productPayload)
	if err := httputil.ParseJSON(r, &payload); err != nil {
		return nil, response.Map{
			"status": "fail",
			"data": response.Map{
				"product": err.Error(),
			},
		}
	}

	if payload.Name == "" {
		return nil, response.ErrNoProductNameProvided
	}

	return payload, nil
}

func (s *service) createProduct(name string) (int, response.Map) {
	_, err := s.store.getByName(name)
	if err == nil {
		return http.StatusBadRequest, response.Map{
			"status": "fail",
			"data": response.Map{
				"product": "product already exists",
			},
		}
	}

	err = s.store.create(name)
	if err != nil {
		return http.StatusInternalServerError, response.Map{
			"status":  "error",
			"message": err.Error(),
		}
	}
	return http.StatusCreated, nil
}
