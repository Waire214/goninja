package models

type ProductsDetails struct {
	ProductName        string `json:"product_name"`
	ProductDescription string `json:"product_description"`
	StockQuantity      int    `json:"stock_quantity"`
	DisplayImage       string `json:"photo"`
	Price              string `json:"price"`
	Industry           string `json:"industry"`
	ProductID          string `json:"product_id"`
	Tags               string `json:"tags"`
	BusinessID         string `json:"business_id"`
}
