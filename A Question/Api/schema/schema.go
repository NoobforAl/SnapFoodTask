package schema

type Food struct {
	Order_id uint   `json:"order_id"`
	Price    uint   `json:"price"`
	Title    string `json:"title"`
}
