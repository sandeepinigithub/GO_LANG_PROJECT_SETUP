package repository

import (
	"GO_LANG_PROJECT_SETUP/config"
	model "GO_LANG_PROJECT_SETUP/models"
)

func GetAllUsers() ([]model.User, error) {
	rows, err := config.DB.Query("SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		err := rows.Scan(&u.ID, &u.Name, &u.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}

func GetUserByID(id int) (model.User, error) {
	var u model.User
	err := config.DB.QueryRow("SELECT id, name, email FROM users WHERE id = ?", id).
		Scan(&u.ID, &u.Name, &u.Email)
	return u, err
}

func CreateUser(u model.User) (int64, error) {
	res, err := config.DB.Exec("INSERT INTO users (name, email) VALUES (?, ?)", u.Name, u.Email)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func UpdateUser(id int, u model.User) error {
	_, err := config.DB.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", u.Name, u.Email, id)
	return err
}

func DeleteUser(id int) error {
	_, err := config.DB.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}
