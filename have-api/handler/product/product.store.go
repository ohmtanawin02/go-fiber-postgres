package product

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ohmtanawin02/go-postgres-basic/have-api/config"
	"github.com/ohmtanawin02/go-postgres-basic/have-api/models"
)

func GetProducts(ctx context.Context, name string, page, limit int) (models.ProductsResponse, error) {
	offset := (page - 1) * limit
	var rows *sql.Rows
	var err error

	if name != "" {
		query := "SELECT id, name, price FROM products WHERE name ILIKE $1 ORDER BY id LIMIT $2 OFFSET $3;"
		rows, err = config.DB.QueryContext(ctx, query, "%"+name+"%", limit, offset)
	} else {
		query := "SELECT id, name, price FROM products ORDER BY id LIMIT $1 OFFSET $2;"
		rows, err = config.DB.QueryContext(ctx, query, limit, offset)
	}

	if err != nil {
		return models.ProductsResponse{}, err
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return models.ProductsResponse{}, err
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return models.ProductsResponse{}, err
	}

	var totalDocs int
	if name != "" {
		err = config.DB.QueryRowContext(ctx, "SELECT COUNT(*) FROM products WHERE name ILIKE $1;", "%"+name+"%").Scan(&totalDocs)
	} else {
		err = config.DB.QueryRowContext(ctx, "SELECT COUNT(*) FROM products;").Scan(&totalDocs)
	}
	if err != nil {
		return models.ProductsResponse{}, err
	}

	return models.ProductsResponse{
		TotalDocs: totalDocs,
		Page:      page,
		Limit:     limit,
		Data:      products,
	}, nil
}

func GetProduct(ctx context.Context, id int) (models.Product, error) {
	var p models.Product
	row := config.DB.QueryRowContext(ctx, "SELECT id, name, price FROM products WHERE id=$1;", id)
	err := row.Scan(&p.ID, &p.Name, &p.Price)
	if err == sql.ErrNoRows {
		return models.Product{}, errors.New("product not found")
	} else if err != nil {
		return models.Product{}, err
	}
	return p, nil
}

func CreateProduct(ctx context.Context, product *models.Product) error {
	if productExistsByName(ctx, product.Name) {
		return errors.New("product name already exists")
	}
	_, err := config.DB.ExecContext(ctx,
		"INSERT INTO public.products(name, price) VALUES ($1, $2);",
		product.Name,
		product.Price,
	)
	return err
}

func UpdateProduct(ctx context.Context, id int, product *models.Product) error {
	if !productExistsByID(ctx, id) {
		return errors.New("product not found")
	}
	if productExistsByNameExceptID(ctx, product.Name, id) {
		return errors.New("product name already exists")
	}
	_, err := config.DB.ExecContext(ctx,
		"UPDATE public.products SET name=$2, price=$3 WHERE id=$1;", id,
		product.Name,
		product.Price,
	)
	return err
}

func DeleteProduct(ctx context.Context, id int) error {
	if !productExistsByID(ctx, id) {
		return errors.New("product not found")
	}
	_, err := config.DB.ExecContext(ctx, "DELETE FROM public.products WHERE id=$1;", id)
	return err
}
