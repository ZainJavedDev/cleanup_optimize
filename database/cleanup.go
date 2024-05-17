package database

import (
	"fmt"
	"time"

	"github.com/ZainJavedDev/cleanup_optimize/discord_logger"
)

var GAME_MODES = []string{"ranked", "unranked", "turbo", "ability_draft"}

func RemoveOldMatches() error {
	db, err := Connect()
	if err != nil {
		return err
	}

	defer db.Close()

	for _, game_mode := range GAME_MODES {

		sevenDaysAgo := time.Now().AddDate(0, 0, -5).Unix()
		sql := fmt.Sprintf("DELETE FROM %s_matches WHERE start_time < $1", game_mode)
		result, err := db.Exec(sql, sevenDaysAgo)
		if err != nil {
			return err
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}
		fmt.Printf("Deleted %d rows.\n", rowsAffected)

		message := fmt.Sprintf("Deleted %d rows.\n", rowsAffected)
		discord_logger.SendDiscordMessage(message)
	}
	return nil
}
