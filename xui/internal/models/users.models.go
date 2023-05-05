package models

import (
	"encoding/json"

	"github.com/harleywinston/x-operators/xui/consts"
)

type ClientSettings interface {
	GetSettingsString() (string, error)
}

type VlessClientSettings struct {
	ID         string `json:"id"`
	Flow       string `json:"flow"`
	Email      string `json:"email"`
	LimitIP    int    `json:"limitIp"`
	TotalGB    int    `json:"totalGB"`
	ExpiryTime int64  `json:"expiryTime"`
	Enable     bool   `json:"enable"`
	TgID       string `json:"tgId"`
	SubID      string `json:"subId"`
}

func (c *VlessClientSettings) GetSettingsString() (string, error) {
	jsonData, err := json.Marshal(map[string][]interface{}{"clients": {c}})
	if err != nil {
		return "", &consts.CustomError{
			Message: consts.JSON_MARSHAL_ERROR.Message,
			Code:    consts.JSON_MARSHAL_ERROR.Code,
			Detail:  err.Error(),
		}
	}

	return string(jsonData), nil
}

type TrojanClientSettings struct {
	Password   string `json:"Password"`
	Flow       string `json:"flow"`
	Email      string `json:"email"`
	LimitIP    int    `json:"limitIp"`
	TotalGB    int    `json:"totalGB"`
	ExpiryTime int64  `json:"expiryTime"`
	Enable     bool   `json:"enable"`
	TgID       string `json:"tgId"`
	SubID      string `json:"subId"`
}

func (c *TrojanClientSettings) GetSettingsString() (string, error) {
	jsonData, err := json.Marshal(map[string][]interface{}{"clients": {c}})
	if err != nil {
		return "", &consts.CustomError{
			Message: consts.JSON_MARSHAL_ERROR.Message,
			Code:    consts.JSON_MARSHAL_ERROR.Code,
			Detail:  err.Error(),
		}
	}

	return string(jsonData), nil
}

type ClientModel struct {
	ID       int            `json:"id"`
	Settings ClientSettings `json:"settings"`
}

type UserModel struct {
	Email      string `json:"email"`
	Username   string `json:"username"`
	Passwd     string `json:"password"`
	ExpiryTime int64  `json:"expiryTime"`
	GroupsID   int    `json:"group_id"`
}
