package services

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/harleywinston/x-operators/pkg/consts"
	"github.com/harleywinston/x-operators/pkg/models"
	"github.com/harleywinston/x-operators/xui/pkg/driver"
)

type SyncerServices struct {
	xuiDriver driver.DriverServices
}

func (s *SyncerServices) getMasterState() (models.MasterState, error) {
	var res models.MasterState

	HTTPClient := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, os.Getenv("MANAGER_BASE_URL")+"/state", nil)
	if err != nil {
		return models.MasterState{}, &consts.CustomError{
			Message: consts.CREATE_HTTP_REQUEST_ERROR.Message,
			Code:    consts.CREATE_HTTP_REQUEST_ERROR.Code,
			Detail:  err.Error(),
		}
	}

	resp, err := HTTPClient.Do(req)
	if err != nil {
		return models.MasterState{}, &consts.CustomError{
			Message: consts.CLIENT_DO_ERROR.Message,
			Code:    consts.CLIENT_DO_ERROR.Code,
			Detail:  err.Error(),
		}
	}

	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return models.MasterState{}, &consts.CustomError{
			Message: consts.JSON_UNMARSHAL_ERROR.Message,
			Code:    consts.JSON_UNMARSHAL_ERROR.Code,
			Detail:  err.Error(),
		}
	}

	return res, nil
}

func (s *SyncerServices) Sync() error {
	inbounds, err := s.xuiDriver.ListInbounds()
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

	if len(masterState.Users) < 1 {
		return consts.MANAGER_USER_COUNT_ERROR
	}
	if len(inbounds) < 1 {
		return consts.INBOUND_COUNT_ERROR
	}

	for _, user := range masterState.Users {
		for _, i := range inbounds {
			hasClient := false
			for _, client := range i.ClientStats {
				if strings.HasPrefix(client.Email, user.Email) {
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
				if strings.HasPrefix(client.Email, user.Email) {
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
