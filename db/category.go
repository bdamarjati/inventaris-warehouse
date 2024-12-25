package db

func (q *Queries) CreateCategory(name string) (RefCategories, error) {
	category := RefCategories{
		Name: name,
	}
	tx := q.db.Create(&category)
	return category, tx.Error
}

func (q *Queries) GetCategory(id int64) (RefCategories, error) {
	category := RefCategories{}
	tx := q.db.First(&category, "id = ?", id)
	return category, tx.Error
}

type ListCategoriesParams struct {
	Limit  int
	Offset int
}

func (q *Queries) ListCategories(arg ListCategoriesParams) ([]RefCategories, error) {
	categories := []RefCategories{}
	tx := q.db.Limit(arg.Limit).Offset(arg.Offset).Find(&categories)
	return categories, tx.Error
}

type UpdateCategoryParams struct {
	ID   int64
	Name string
}

func (q *Queries) UpdateCategory(arg UpdateCategoryParams) (int64, error) {
	category := RefCategories{
		ID:   arg.ID,
		Name: arg.Name,
	}
	tx := q.db.Model(category).Where("id = ?", category.ID).Updates(category)
	return tx.RowsAffected, tx.Error
}

func (q *Queries) DeleteCategory(id int64) (int64, error) {
	category := RefCategories{
		ID: id,
	}
	tx := q.db.Model(category).Where("id = ?", category.ID).Delete(category)
	return tx.RowsAffected, tx.Error
}
