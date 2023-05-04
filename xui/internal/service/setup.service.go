package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/google/uuid"

	"github.com/harleywinston/x-operators/xui/configs"
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
		return err
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
		return err
	}
	// log.Println(bytes.NewBuffer(jsonReqData))

	apiUrl := configs.BaseURL.ResolveReference(&url.URL{Path: "xui/API/inbounds/addClient"})
	req, err := http.NewRequest(http.MethodPost, apiUrl.String(), bytes.NewBuffer(jsonReqData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := configs.Clinet.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf(
			`xui api called successfully but an error happend!
			Status code: %v
			Response body: %v
			`,
			resp.StatusCode,
			resp.Body,
		)
	}

	return nil
}

func (s *SetupServices) DeleteClientService(user models.UserModel) error {
	return nil
}
