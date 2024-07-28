package models

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type ProductsResponse struct {
	TotalDocs int       `json:"totalDocs"`
	Page      int       `json:"page"`
	Limit     int       `json:"limit"`
	Data      []Product `json:"data"`
}
