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
		var daysToKeep int
		if game_mode == "ability_draft" {
			daysToKeep = 10
		} else {
			daysToKeep = 5
		}

		cutoffTime := time.Now().AddDate(0, 0, -daysToKeep).Unix()
		var total int64
		for {
			sql := fmt.Sprintf("DELETE FROM %s_matches WHERE start_time < ? LIMIT 1000", game_mode)
			result, err := db.Exec(sql, cutoffTime)
			if err != nil {
				return err
			}

			rowsAffected, err := result.RowsAffected()
			if err != nil {
				return err
			}
			total += rowsAffected

			fmt.Printf("Deleted %d rows from %s matches.\n", rowsAffected, game_mode)

			if rowsAffected < 1000 {
				break
			}
		}
		message := fmt.Sprintf("Deleted %d rows from %s matches.\n", total, game_mode)
		discord_logger.SendDiscordMessage(message)
	}
	return nil
}
