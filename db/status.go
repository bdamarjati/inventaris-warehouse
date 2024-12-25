package db

type Tabler interface {
	TableName() string
}

func (RefStatus) TableName() string {
	return "ref_status"
}

func (q *Queries) CreateStatus(description string) (RefStatus, error) {
	status := RefStatus{
		Description: description,
	}
	tx := q.db.Create(&status)
	return status, tx.Error
}

func (q *Queries) GetStatus(id int64) (RefStatus, error) {
	status := RefStatus{}
	tx := q.db.First(&status, "id = ?", id)
	return status, tx.Error
}

type ListStatusParams struct {
	Limit  int
	Offset int
}

func (q *Queries) ListStatus(arg ListStatusParams) ([]RefStatus, error) {
	statuses := []RefStatus{}
	tx := q.db.Limit(arg.Limit).Offset(arg.Offset).Find(&statuses)
	return statuses, tx.Error
}

type UpdateStatusParams struct {
	ID          int64
	Description string
}

func (q *Queries) UpdateStatus(arg UpdateStatusParams) (int64, error) {
	status := RefStatus{
		ID:          arg.ID,
		Description: arg.Description,
	}
	tx := q.db.Model(status).Where("id = ?", status.ID).Updates(status)
	return tx.RowsAffected, tx.Error
}

func (q *Queries) DeleteStatus(id int64) (int64, error) {
	status := RefStatus{
		ID: id,
	}
	tx := q.db.Model(status).Where("id = ?", status.ID).Delete(status)
	return tx.RowsAffected, tx.Error
}
