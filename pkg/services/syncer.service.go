package services

import (
	"log"

	"github.com/harleywinston/x-operators/pkg/consts"
	"github.com/harleywinston/x-operators/pkg/models"
	"github.com/harleywinston/x-operators/xui/pkg/driver"
)

type SyncerServices struct {
	xuiDriver driver.DriverServices
}

func (s *SyncerServices) getMasterState() (models.MasterState, error) {
	var res models.MasterState

	return res, nil
}

func (s *SyncerServices) Sync() error {
	log.Println("hello from syncer")
	inbounds, err := s.xuiDriver.ListInbounds()
	log.Println(inbounds)
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
				err := s.xuiDriver.AddClientService(user, i)
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
				err := s.xuiDriver.DeleteClientService(client, i)
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
