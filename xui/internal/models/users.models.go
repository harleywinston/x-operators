package models

type ClientSettingsModel struct {
	ID         string `json:"id"`
	Flow       string `json:"flow"`
	Email      string `json:"email"`
	LimitIP    int    `json:"limitIp"`
	TotalGB    int    `json:"totalGB"`
	ExpiryTime int    `json:"expiryTime"`
	Enable     bool   `json:"enable"`
	TgID       string `json:"tgId"`
	SubID      string `json:"subId"`
}

type ClientModel struct {
	ID       int                 `json:"id"`
	Settings ClientSettingsModel `json:"settings"`
}

type UserModel struct {
	Email    string `json:"email"    gorm:"unique"`
	Username string `json:"username"`
	Passwd   string `json:"password"`
	GroupsID int    `json:"group_id"`
}
