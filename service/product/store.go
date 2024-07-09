package product

import (
	"database/sql"
	"errors"
)

type store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *store {
	return &store{
		db: db,
	}
}

func (s *store) getAll() ([]product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]product, 0)
	for rows.Next() {
		p, err := scanRow(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *p)
	}
	return products, nil
}

func (s *store) getByID(id string) (*product, error) {
	rows, err := s.db.Query("SELECT * FROM products WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	p := new(product)
	for rows.Next() {
		p, err = scanRow(rows)
		if err != nil {
			return nil, err
		}
	}

	if p.ID == "" {
		return nil, errors.New("product not found")
	}
	return p, nil
}

func (s *store) getByName(name string) (*product, error) {
	rows, err := s.db.Query("SELECT * FROM products WHERE name = ?", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	p := new(product)
	for rows.Next() {
		p, err = scanRow(rows)
		if err != nil {
			return nil, err
		}
	}

	if p.ID == "" {
		return nil, errors.New("product not found")
	}
	return p, nil
}

func (s *store) create(name string) error {
	_, err := s.db.Exec("INSERT INTO products (name) VALUES (?)", name)
	return err
}

func scanRow(rows *sql.Rows) (*product, error) {
	p := new(product)
	err := rows.Scan(&p.ID, &p.Name)
	return p, err
}
