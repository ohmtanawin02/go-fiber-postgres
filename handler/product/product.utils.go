package product

import (
	"context"
	"log"

	"github.com/ohmtanawin02/go-postgres-basic/config"
)

func productExistsByID(ctx context.Context, id int) bool {
	var count int
	err := config.DB.QueryRowContext(ctx, "SELECT COUNT(*) FROM products WHERE id=$1;", id).Scan(&count)
	if err != nil {
		log.Fatalf("Error checking if product exists by ID: %v\n", err)
	}
	return count > 0
}

func productExistsByName(ctx context.Context, name string) bool {
	var count int
	err := config.DB.QueryRowContext(ctx, "SELECT COUNT(*) FROM products WHERE name=$1;", name).Scan(&count)
	if err != nil {
		log.Fatalf("Error checking if product exists by name: %v\n", err)
	}
	return count > 0
}

func productExistsByNameExceptID(ctx context.Context, name string, id int) bool {
	var count int
	err := config.DB.QueryRowContext(ctx, "SELECT COUNT(*) FROM products WHERE name=$1 AND id!=$2;", name, id).Scan(&count)
	if err != nil {
		log.Fatalf("Error checking if product exists by name except ID: %v\n", err)
	}
	return count > 0
}
