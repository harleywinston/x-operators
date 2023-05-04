package models

type ClientSettingsModel struct {
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

type ClientModel struct {
	ID       int    `json:"id"`
	Settings string `json:"settings"`
}

type UserModel struct {
	Email      string `json:"email"`
	Username   string `json:"username"`
	Passwd     string `json:"password"`
	ExpiryTime int64  `json:"expiryTime"`
	GroupsID   int    `json:"group_id"`
}
