package db

import "time"

type User struct {
	Username  string
	Password  string
	Role      string
	CreatedAt time.Time
}

type RefCategories struct {
	ID        int64
	Name      string
	CreatedAt time.Time
}

type RefStatus struct {
	ID          int64
	Description string
	CreatedAt   time.Time
}

type Inventories struct {
	ID         int64
	Name       string
	Quantity   int64
	CategoryId int64
	Condition  int
	Status     int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Consumables struct {
	ID         int64
	Name       string
	Quantity   int64
	CategoryId int64
	Status     int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
