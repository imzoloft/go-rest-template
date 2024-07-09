package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/imzoloft/go-rest-api/service/product"
)

type Server struct {
	Addr string
	DB   *sql.DB
}

func NewServer(addr string, db *sql.DB) *Server {
	return &Server{
		Addr: addr,
		DB:   db,
	}
}

func (s *Server) Run() error {
	router := http.NewServeMux()

	productStore := product.NewStore(s.DB)
	productSrv := product.NewService(productStore)
	productHandler := product.NewHandler(productSrv)

	productHandler.RegisterRoutes(router)

	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", router))
	v1.HandleFunc("GET /helloworld", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	// chain := middleware.Chain(
	// 	middleware.Logger,
	// )

	log.Printf("Server listening on %s", s.Addr)
	return http.ListenAndServe(s.Addr, v1)
}
