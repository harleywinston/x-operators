package models

type StreamSettingsInterface interface{}

type SniffingInterface interface{}

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
	ClientStats    ClientStatsModel        `json:"clientStats"`
}
