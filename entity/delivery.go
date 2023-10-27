package entity

type Delivery struct {
	OrderUID   string `json:"order_uid" db:"order_uid"`
	DeliveryID string `json:"delivery_id" db:"delivery_id"`
	Name       string `json:"name" db:"name"`
	Phone      string `json:"phone" db:"phone"`
	Zip        string `json:"zip" db:"zip"`
	City       string `json:"city" db:"city"`
	Address    string `json:"address" db:"address"`
	Region     string `json:"region" db:"region"`
	Email      string `json:"email" db:"email"`
}
