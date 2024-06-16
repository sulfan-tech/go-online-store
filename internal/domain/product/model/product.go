package model

type Product struct {
	ID       uint    `json:"id" gorm:"column:id;not null"`
	Name     string  `json:"name" gorm:"column:name;not null"`
	Category string  `json:"category" gorm:"column:category;not null"`
	Price    float64 `json:"price" gorm:"column:price;not null"`
	Stok     uint    `json:"stok" gorm:"column:stok;not null"`
}

func (Product) TableName() string {
	return "Product"
}
