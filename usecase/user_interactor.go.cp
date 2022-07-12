package usecase

import "clean_arch/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (i *UserInteractor) Add(u domain.User) error {
	_, err := i.UserRepository.Store(u)
	if err != nil {
		return err
	}
	return nil
}

func (i *UserInteractor) Users() (*domain.Users, error) {
	users, err := i.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (i *UserInteractor) UserById(id uint) (*domain.User, error) {
	user, err := i.UserRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
