package models

type Admin struct {
	ID        string `json:"id" sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt int64  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64  `json:"updated_at" gorm:"autoUpdateTime"`
}
