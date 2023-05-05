package models

type UserModel struct {
	Email      string `json:"email"`
	Username   string `json:"username"`
	Passwd     string `json:"password"`
	ExpiryTime int64  `json:"expiryTime"`
	GroupsID   int    `json:"group_id"`
}

type MasterState struct {
	Users []UserModel
}
