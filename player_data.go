package main

type PlayerData struct {
	Global GlobalData `json:"global"`
}

type GlobalData struct {
	Name     string `json:"name"`
	Uid      int64  `json:"uid"`
	Level    int    `json:"level"`
	Platform string `json:"platform"`
}
