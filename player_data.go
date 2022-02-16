package main

import "fmt"

type PlayerData struct {
	Global GlobalData `json:"global"`
}

type GlobalData struct {
	Name     string   `json:"name"`
	Uid      int64    `json:"uid"`
	Level    int      `json:"level"`
	Platform string   `json:"platform"`
	Rank     RankData `json:"rank"`
}

type RankData struct {
	RankScore int    `json:"rankScore"`
	RankName  string `json:"rankName"`
	RankDiv   int    `json:"rankDiv"`
}

func (g GlobalData) getRank() string {
	return fmt.Sprintf("%s - %d", g.Rank.RankName, g.Rank.RankDiv)
}
