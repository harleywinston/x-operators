package service

import (
	"encoding/json"

	"github.com/google/uuid"

	"github.com/harleywinston/x-operators/xui/internal/models"
)

type SetupServices struct{}

func getFlow() string {
	return "xtls-rprx-vision-udp443"
}

func getVlessId() string {
	return uuid.New().String()
}

// func getExpiryTime() int64 {
// 	now := time.Now()
// 	future := now.AddDate(0, 0, 30)
// 	return future.UnixNano() / int64(time.Millisecond)
// }

func (s *SetupServices) AddClientService(user models.UserModel) error {
	var client models.ClientModel
	var clientSetting models.ClientSettingsModel
	client.ID = 1

	clientSetting.ID = getVlessId()
	clientSetting.Flow = getFlow()
	clientSetting.Email = user.Email
	clientSetting.LimitIP = 2
	clientSetting.TotalGB = 100
	clientSetting.ExpiryTime = user.ExpiryTime
	clientSetting.TgID = ""
	clientSetting.SubID = ""
	clientSetting.Enable = true

	settingData, err := json.Marshal(clientSetting)
	if err != nil {
		return err
	}
	client.Settings = string(settingData)

	return nil
}

func (s *SetupServices) DeleteClientService(user models.UserModel) error {
	return nil
}
