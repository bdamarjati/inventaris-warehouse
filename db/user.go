package db

type CreateUserParams struct {
	Username string
	Password string
	Role     string
}

func (q Queries) CreateUser(arg CreateUserParams) (int64, error) {
	tx := q.db.Create(&User{
		Username: arg.Username,
		Password: arg.Password,
		Role:     arg.Role,
	})
	err := tx.Error
	rowAffected := tx.RowsAffected
	return rowAffected, err
}

func (q *Queries) GetUser(username string) (User, error) {
	user := User{}
	tx := q.db.First(&user, "username = ?", username)
	return user, tx.Error
}

type ListUserParams struct {
	Limit  int
	Offset int
}

func (q *Queries) ListUser(arg ListUserParams) ([]User, error) {
	users := []User{}
	tx := q.db.Limit(arg.Limit).Offset(arg.Offset).Find(&users)
	return users, tx.Error
}

type UpdateUserParams struct {
	Username string
	Password string
	Role     string
}

func (q *Queries) UpdateUser(arg UpdateUserParams) (int64, error) {
	user := User{
		Username: arg.Username,
		Password: arg.Password,
		Role:     arg.Role,
	}
	tx := q.db.Model(user).Where("username = ?", user.Username).Updates(user)
	return tx.RowsAffected, tx.Error
}

func (q *Queries) DeleteUser(username string) (int64, error) {
	user := User{
		Username: username,
	}
	tx := q.db.Model(user).Where("username = ?", username).Delete(user)
	return tx.RowsAffected, tx.Error
}
