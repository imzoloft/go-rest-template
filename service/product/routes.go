package product

import (
	"net/http"

	"github.com/imzoloft/go-rest-api/httputil"
	"github.com/imzoloft/go-rest-api/middleware"
)

type Handler struct {
	service productService
}

func NewHandler(service productService) Handler {
	return Handler{
		service: service,
	}
}

func (h *Handler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /products", h.getProducts())
	router.HandleFunc("GET /products/{id}", h.getProduct())
	router.HandleFunc("POST /products", middleware.Auth(http.HandlerFunc(h.addProduct())))
}

func (h *Handler) addProduct() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := h.service.validatePayload(r)
		if err != nil {
			httputil.WriteError(w, http.StatusBadRequest, err)
			return
		}

		status, err := h.service.createProduct(payload.Name)
		if err != nil {
			httputil.WriteError(w, status, err)
			return
		}
		httputil.WriteJSON(w, http.StatusCreated, nil)
	}
}

func (h *Handler) getProducts() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		//json.NewEncoder(w).Encode(p.productService.GetProducts())
	}
}

func (h *Handler) getProduct() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Content-Type", "application/json")

		// productId := r.PathValue("id")
		// product := p.productService.GetProduct(productId)

		// if product.ID == "" {
		// 	w.WriteHeader(http.StatusNotFound)
		// 	json.NewEncoder(w).Encode("Product not found")
		// 	return
		//}
		w.WriteHeader(http.StatusOK)
		//json.NewEncoder(w).Encode(p.productService.GetProduct(productId))
	}
}
