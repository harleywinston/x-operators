package service

import "github.com/harleywinston/x-operators/xui/internal/models"

type SetupServices struct{}

func (s *SetupServices) AddClientService(user models.UserModel) error {
	return nil
}

func (s *SetupServices) DeleteClientService(user models.UserModel) error {
	return nil
}
