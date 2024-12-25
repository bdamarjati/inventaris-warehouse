package db

import "time"

type CreateConsumableParams struct {
	Name       string
	Quantity   int64
	CategoryId int64
	Status     int64
}

func (q *Queries) CreateConsumable(arg CreateConsumableParams) (Consumables, error) {
	consumable := Consumables{
		Name:       arg.Name,
		Quantity:   arg.Quantity,
		CategoryId: arg.CategoryId,
		Status:     arg.Status,
	}
	tx := q.db.Create(&consumable)
	return consumable, tx.Error
}

func (q *Queries) GetConsumable(id int64) (Consumables, error) {
	consumable := Consumables{}
	tx := q.db.First(&consumable, "id = ?", id)
	return consumable, tx.Error
}

type ListConsumableParams struct {
	Limit  int
	Offset int
}

func (q *Queries) ListConsumables(arg ListConsumableParams) ([]Consumables, error) {
	consumables := []Consumables{}
	tx := q.db.Limit(arg.Limit).Offset(arg.Offset).Find(&consumables)
	return consumables, tx.Error
}

type UpdateConsumableParams struct {
	ID         int64
	Name       string
	Quantity   int64
	CategoryId int64
	Status     int64
}

func (q *Queries) UpdateConsumable(arg UpdateConsumableParams) (int64, error) {
	consumable := Consumables{
		ID:         arg.ID,
		Name:       arg.Name,
		Quantity:   arg.Quantity,
		CategoryId: arg.CategoryId,
		Status:     arg.Status,
		UpdatedAt:  time.Now().Local(),
	}
	tx := q.db.Model(consumable).Where("id = ?", arg.ID).Updates(consumable)
	return tx.RowsAffected, tx.Error
}

func (q *Queries) DeleteConsumable(id int64) (int64, error) {
	consumanble := Consumables{
		ID: id,
	}
	tx := q.db.Model(consumanble).Where("id = ?", id).Delete(consumanble)
	return tx.RowsAffected, tx.Error
}
