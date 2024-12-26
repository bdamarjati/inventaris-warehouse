package db

import "time"

type User struct {
	Username  string
	Password  string
	Role      string
	CreatedAt time.Time
}

type RefCategories struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type RefStatus struct {
	ID          int64     `json:"id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}

type Inventories struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Quantity   int64     `json:"quantity"`
	CategoryId int64     `json:"category_id"`
	Condition  int       `json:"condition"`
	Status     int64     `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Consumables struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Quantity   int64     `json:"quantity"`
	CategoryId int64     `json:"category_id"`
	Status     int64     `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
