package product

type RequestProduct struct {
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Stok     uint    `json:"stok"`
}
