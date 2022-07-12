package database

import (
	"clean_arch/domain"
	"fmt"
)

type UserRepository struct {
	SqlHandler
}

func (r *UserRepository) Store(u domain.User) (uint, error) {
	result, err := r.Execute(
		"INSERT INTO `users` (`first_name`, `last_name`, `avatar_url`) VALUES (?, ?)", u.FirstName, u.LastName, u.AvatarURL,
	)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint(id), nil
}

func (r *UserRepository) FindById(id uint) (*domain.User, error) {
	row, err := r.Query("SELECT `id`, `first_name`, `last_name`, `avatar_url` FROM `users WHERE `id` = ?", id)
	defer row.Close()
	if err != nil {
		return nil, nil
	}

	var firstName string
	var lastName string
	var avatarURL string
	if err := row.Scan(&id, &firstName, &lastName, &avatarURL); err != nil {
		return nil, err
	}

	return &domain.User{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		AvatarURL: avatarURL,
	}, nil
}

func (r *UserRepository) FindAll() (*domain.Users, error) {
	rows, err := r.Query("SELECT `id`, `first_name`, `last_name`, `avatar_url` FROM `users")
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var id uint
	var firstName string
	var lastName string
	var avatarURL string
	var users domain.Users
	for rows.Next() {
		if err := rows.Scan(&id, &firstName, &lastName, &avatarURL); err != nil {
			return nil, err
		}
		user := domain.User{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
			AvatarURL: avatarURL,
		}

		users = append(users, user)

	}

	return &users, nil
}
