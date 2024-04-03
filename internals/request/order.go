package request

type CreateOrder struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
