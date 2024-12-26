package db

import "gorm.io/gorm"

type Queries struct {
	db gorm.DB
}

func New(db gorm.DB) *Queries {
	return &Queries{db: db}
}

type Store struct {
	*Queries
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(*db),
	}
}
