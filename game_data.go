package main

import (
	"fmt"
	"strings"
)

type GameData struct {
	PlayerUid         string     `json:"uid"`
	PlayerName        string     `json:"name"`
	GameMode          string     `json:"gameMode"`
	LegendPlayed      string     `json:"legendPlayed"`
	StartTimestamp    int64      `json:"gameStartTimestamp"`
	EndTimestamp      int64      `json:"gameEndTimestamp"`
	GameDuration      int        `json:"gameLengthSecs"`
	GameStats         []GameStat `json:"gameData"`
	BRScoreChange     int        `json:"BRScoreChange"`
	ArenasScoreChange int        `json:"ArenasScoreChange"`
}

type GameStat struct {
	Key   string `json:"key"`
	Value int    `json:"value"`
	Name  string `json:"name"`
}

func (g GameData) getGameId() string {
	return fmt.Sprintf("%s-%d", g.PlayerUid, g.StartTimestamp)
}

func (g GameData) getDamageDone() int {
	return g.getStatValueByKey("damage")
}

func (g GameData) getKills() int {
	return g.getStatValueByKey("kills")
}

func (g GameData) getScoreChange() int {
	if g.GameMode == "BATTLE_ROYALE" {
		return g.BRScoreChange
	}

	return g.ArenasScoreChange
}

func (g GameData) getStatValueByKey(key string) int {
	for _, stat := range g.GameStats {
		if strings.Contains(stat.Key, key) {
			return stat.Value
		}
	}

	return 0
}
