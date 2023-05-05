package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/google/uuid"

	"github.com/harleywinston/x-operators/xui/configs"
	"github.com/harleywinston/x-operators/xui/consts"
	"github.com/harleywinston/x-operators/xui/internal/models"
)

type SetupServices struct{}

func (s *SetupServices) getFlow() string {
	return "xtls-rprx-vision-udp443"
}

func (s *SetupServices) getUUID() string {
	return uuid.New().String()
}

func (s *SetupServices) setClientVless(
	client *models.ClientModel,
	user models.UserModel,
	inboundID int,
) error {
	client.ID = inboundID
	client.Settings = &models.VlessClientSettings{
		Enable:     true,
		ID:         s.getUUID(),
		Flow:       s.getFlow(),
		Email:      user.Email,
		LimitIP:    2,
		TotalGB:    107 * 10e9,
		ExpiryTime: user.ExpiryTime,
		TgID:       "",
		SubID:      "",
	}
	return nil
}

func (s *SetupServices) setClientTrojan(
	client *models.ClientModel,
	user models.UserModel,
	inboundID int,
) error {
	client.ID = inboundID
	client.Settings = &models.TrojanClientSettings{
		Enable:     true,
		Password:   user.Passwd,
		Flow:       s.getFlow(),
		Email:      user.Email,
		LimitIP:    2,
		TotalGB:    107 * 10e9,
		ExpiryTime: user.ExpiryTime,
		TgID:       "",
		SubID:      "",
	}
	return nil
}

func (s *SetupServices) getReqClientJson(client models.ClientModel) ([]byte, error) {
	settingsString, err := client.Settings.GetSettingsString()
	if err != nil {
		return []byte{}, err
	}
	jsonData, err := json.Marshal(map[string]interface{}{
		"id":       client.ID,
		"settings": settingsString,
	})
	if err != nil {
		return []byte{}, &consts.CustomError{
			Message: consts.JSON_MARSHAL_ERROR.Message,
			Code:    consts.JSON_MARSHAL_ERROR.Code,
			Detail:  err.Error(),
		}
	}
	return jsonData, nil
}

func (s *SetupServices) ListInbounds() ([]models.InboundStatsModel, error) {
	var list []models.InboundStatsModel

	apiURL := configs.BaseURL.ResolveReference(&url.URL{Path: "/xui/API/inbounds/list"})
	req, err := http.NewRequest(http.MethodGet, apiURL.String(), nil)
	if err != nil {
		return []models.InboundStatsModel{}, &consts.CustomError{
			Message: consts.CREATE_HTTP_REQUEST_ERROR.Message,
			Code:    consts.CREATE_HTTP_REQUEST_ERROR.Code,
			Detail:  err.Error(),
		}
	}
	resp, err := configs.Clinet.Do(req)
	if resp.StatusCode != 200 {
		return []models.InboundStatsModel{}, &consts.CustomError{
			Message: consts.XUI_API_ERROR.Message,
			Code:    consts.XUI_API_ERROR.Code,
			Detail:  err.Error(),
		}
	}

	type ApiOutput struct {
		Obj []models.InboundStatsModel `json:"obj"`
	}

	var respJson ApiOutput
	err = json.NewDecoder(resp.Body).Decode(&respJson)
	if err != nil {
		return []models.InboundStatsModel{}, &consts.CustomError{
			Message: consts.JSON_UNMARSHAL_ERROR.Message,
			Code:    consts.JSON_UNMARSHAL_ERROR.Code,
			Detail:  err.Error(),
		}
	}
	list = respJson.Obj

	return list, nil
}

func (s *SetupServices) AddClientService(user models.UserModel) error {
	var client models.ClientModel
	if err := s.setClientVless(&client, user, 1); err != nil {
		return err
	}

	jsonReqData, err := s.getReqClientJson(client)
	if err != nil {
		return err
	}

	apiUrl := configs.BaseURL.ResolveReference(&url.URL{Path: "xui/API/inbounds/addClient"})
	req, err := http.NewRequest(http.MethodPost, apiUrl.String(), bytes.NewBuffer(jsonReqData))
	if err != nil {
		return &consts.CustomError{
			Message: consts.CREATE_HTTP_REQUEST_ERROR.Message,
			Code:    consts.CREATE_HTTP_REQUEST_ERROR.Code,
			Detail:  err.Error(),
		}
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := configs.Clinet.Do(req)
	if err != nil {
		return &consts.CustomError{
			Message: consts.CLIENT_DO_ERROR.Message,
			Code:    consts.CLIENT_DO_ERROR.Code,
			Detail:  err.Error(),
		}
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return &consts.CustomError{
			Message: consts.XUI_API_ERROR.Message,
			Code:    consts.XUI_API_ERROR.Code,
			Detail:  err.Error(),
		}
	}

	return nil
}

func (s *SetupServices) DeleteClientService(user models.UserModel) error {
	return nil
}
