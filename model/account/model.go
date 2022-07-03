package account

import "gorm.io/gorm"

type Init struct {
	CustomerXID string `json:"customer_xid" validate:"required"`
	Token       string `json:"token"`
}

type Detail struct {
	CustomerXID string `json:"customer_xid"`
}

type User struct {
	gorm.Model        // adds ID, created_at etc.
	Name       string `json:"name"`
}
