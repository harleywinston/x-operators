package models

import (
	"encoding/json"

	"github.com/harleywinston/x-operators/pkg/consts"
)

type StreamSettingsInterface interface{}

type SniffingInterface interface{}

type ClientSettings interface {
	GetSettingsString() (string, error)
}

type ClientStatsModel struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Enable    bool   `json:"enable"`
	InboundID int    `json:"inboundId"`
}

type InboundStatsModel struct {
	ID             int                     `json:"id"`
	Port           int                     `json:"port"`
	Protocol       string                  `json:"protocol"`
	StreamSettings StreamSettingsInterface `json:"streamSettings"`
	Sniffing       SniffingInterface       `json:"sniffing"`
	ClientStats    []ClientStatsModel      `json:"clientStats"`
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
	Password   string `json:"password"`
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
