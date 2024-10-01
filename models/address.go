package models

type Address struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	UserID  string `json:"user_id"`
	User    *User  `json:"user"`
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode string `json:"zip_code"`
	Country string `json:"country"`
}

func (Address) TableName() string {
	return "addresses"
}
