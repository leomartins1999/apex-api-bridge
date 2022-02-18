package deserializers

import (
	"apex-api-sync/models"
	"encoding/json"
)

func DeserializeGames(data []byte) ([]models.GameData, error) {
	games := make([]models.GameData, 0)

	err := json.Unmarshal(data, &games)

	return games, err
}
