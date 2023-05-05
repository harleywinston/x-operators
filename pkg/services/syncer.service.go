package services

import (
	"github.com/harleywinston/x-operators/pkg/consts"
	"github.com/harleywinston/x-operators/pkg/models"
	"github.com/harleywinston/x-operators/xui/pkg/service"
)

type SyncerServices struct {
	xuiService service.SetupServices
}

func (s *SyncerServices) getMasterState() (models.MasterState, error) {
	var res models.MasterState

	return res, nil
}

func (s *SyncerServices) Sync() error {
	inbounds, err := s.xuiService.ListInbounds()
	if err != nil {
		return &consts.CustomError{
			Message: consts.XUI_DRIVER_ERROR.Message,
			Code:    consts.XUI_DRIVER_ERROR.Code,
			Detail:  err.Error(),
		}
	}

	masterState, err := s.getMasterState()
	if err != nil {
		return err
	}

	for _, user := range masterState.Users {
		for _, i := range inbounds {
			hasClient := false
			for _, client := range i.ClientStats {
				if user.Email == client.Email {
					hasClient = true
				}
			}

			if !hasClient {
				err := s.xuiService.AddClientService(user, i)
				if err != nil {
					return &consts.CustomError{
						Message: consts.XUI_DRIVER_ERROR.Message,
						Code:    consts.XUI_DRIVER_ERROR.Code,
						Detail:  err.Error(),
					}
				}
			}
		}
	}

	for _, i := range inbounds {
		for _, client := range i.ClientStats {
			hasUser := false
			for _, user := range masterState.Users {
				if user.Email == client.Email {
					hasUser = true
				}
			}

			if !hasUser {
				err := s.xuiService.DeleteClientService(client, i)
				if err != nil {
					return &consts.CustomError{
						Message: consts.XUI_DRIVER_ERROR.Message,
						Code:    consts.XUI_DRIVER_ERROR.Code,
						Detail:  err.Error(),
					}
				}
			}
		}
	}
	return nil
}
