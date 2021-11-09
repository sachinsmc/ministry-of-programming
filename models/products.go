package models

type Product struct {
	ID        string `json:"name" sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Title     string `json:"title"`
	Quantity  string `json:"quantity"`
	Price     string `json:"price"`
	CreatedAt int64  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64  `json:"updated_at" gorm:"autoUpdateTime"`
}
