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

func getFlow() string {
	return "xtls-rprx-vision-udp443"
}

func getUUID() string {
	return uuid.New().String()
}

// func getExpiryTime() int64 {
// 	now := time.Now()
// 	future := now.AddDate(0, 0, 30)
// 	return future.UnixNano() / int64(time.Millisecond)
// }

func setClientVless(client *models.ClientModel, user models.UserModel) error {
	var clientSetting models.ClientSettingsModel
	client.ID = 1

	clientSetting.ID = getUUID()
	clientSetting.Flow = getFlow()
	clientSetting.Email = user.Email
	clientSetting.LimitIP = 2
	clientSetting.TotalGB = 107 * 1000000000
	clientSetting.ExpiryTime = user.ExpiryTime
	clientSetting.TgID = ""
	clientSetting.SubID = ""
	clientSetting.Enable = true

	settingData, err := json.Marshal(map[string][]interface{}{"clients": {clientSetting}})
	if err != nil {
		return &consts.CustomError{
			Message: consts.JSON_MARSHAL_ERROR.Message,
			Code:    consts.JSON_MARSHAL_ERROR.Code,
			Detail:  err.Error(),
		}
	}
	client.Settings = string(settingData)
	return nil
}

func (s *SetupServices) AddClientService(user models.UserModel) error {
	var client models.ClientModel
	if err := setClientVless(&client, user); err != nil {
		return err
	}

	jsonReqData, err := json.Marshal(client)
	if err != nil {
		return &consts.CustomError{
			Message: consts.JSON_MARSHAL_ERROR.Message,
			Code:    consts.JSON_MARSHAL_ERROR.Code,
			Detail:  err.Error(),
		}
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
