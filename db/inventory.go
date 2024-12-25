package db

import "time"

type CreateInventoryParams struct {
	Name       string
	Quantity   int64
	CategoryId int64
	Condition  int
	Status     int64
}

func (q *Queries) CreateInventory(arg CreateInventoryParams) (Inventories, error) {
	inventory := Inventories{
		Name:       arg.Name,
		Quantity:   arg.Quantity,
		CategoryId: arg.CategoryId,
		Condition:  arg.Condition,
		Status:     arg.Status,
	}
	tx := q.db.Create(&inventory)
	return inventory, tx.Error
}

func (q *Queries) GetInventory(id int64) (Inventories, error) {
	inventory := Inventories{}
	tx := q.db.First(&inventory, "id = ?", id)
	return inventory, tx.Error
}

type ListInventoryParams struct {
	Limit  int
	Offset int
}

func (q *Queries) ListInventories(arg ListInventoryParams) ([]Inventories, error) {
	inventories := []Inventories{}
	tx := q.db.Limit(arg.Limit).Offset(arg.Offset).Find(&inventories)
	return inventories, tx.Error
}

type UpdateInventoryParams struct {
	ID         int64
	Name       string
	Quantity   int64
	CategoryId int64
	Condition  int
	Status     int64
}

func (q *Queries) UpdateInventory(arg UpdateInventoryParams) (int64, error) {
	inventory := Inventories{
		ID:         arg.ID,
		Name:       arg.Name,
		Quantity:   arg.Quantity,
		CategoryId: arg.CategoryId,
		Condition:  arg.Condition,
		Status:     arg.Status,
		UpdatedAt:  time.Now().Local(),
	}
	tx := q.db.Model(inventory).Where("id = ?", inventory.ID).Updates(inventory)
	return tx.RowsAffected, tx.Error
}

func (q *Queries) DeleteInventory(id int64) (int64, error) {
	Inventory := Inventories{
		ID: id,
	}
	tx := q.db.Model(Inventory).Where("id = ?", id).Delete(Inventory)
	return tx.RowsAffected, tx.Error
}
