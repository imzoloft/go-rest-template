package product

import (
	"net/http"

	"github.com/imzoloft/go-rest-api/response"
)

type product struct {
	ID   string
	Name string
}

type productPayload struct {
	Name string `json:"name,omitempty"`
}

type productStore interface {
	getByName(name string) (*product, error)
	getAll() ([]product, error)
	getByID(id string) (*product, error)
	create(name string) error
}

type productService interface {
	createProduct(name string) (int, response.Map)
	validatePayload(r *http.Request) (*productPayload, response.Map)
}
