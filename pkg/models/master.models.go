package models

type MasterUsers struct {
	Email      string `json:"email"       gorm:"unique"`
	Username   string `json:"username"`
	Passwd     string `json:"password"`
	ExpiryTime int64  `json:"expiryTime"`
	GroupsID   int    `json:"group_id"`
	FuckedUser bool   `json:"fucked_user"`
}

type MasterState struct {
	Users []MasterUsers
}
