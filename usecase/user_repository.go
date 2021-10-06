package usecase

import "clean_arch/domain"

type UserRepository interface {
	Store(domain.User) (uint, error)
	FindById(uint) (*domain.User, error)
	FindAll() (*domain.Users, error)
}
