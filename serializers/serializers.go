package serializers

import "time"

type UserSerializer struct {
	ID        uint      `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  []byte    `json:"password_hash"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProductSerializer struct {
	ID           uint      `json:"id"`
	Name         string    `json:"name"`
	SerialNumber string    `json:"serial_number"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
