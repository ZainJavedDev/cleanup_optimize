package database

import (
	"fmt"
	"log"
	"time"

	"github.com/ZainJavedDev/cleanup_optimize/discord_logger"
)

var GAME_MODES = []string{"ranked", "unranked", "turbo"}

func RemoveOldMatches() {
	db, err := Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	for _, game_mode := range GAME_MODES {

		sevenDaysAgo := time.Now().AddDate(0, 0, -5).Unix()
		var total int64
		for {
			sql := fmt.Sprintf("DELETE FROM %s_matches WHERE start_time < ? LIMIT 1000", game_mode)
			result, err := db.Exec(sql, sevenDaysAgo)
			if err != nil {
				log.Fatal(err)
			}

			rowsAffected, err := result.RowsAffected()
			if err != nil {
				log.Fatal(err)
			}
			total += rowsAffected

			fmt.Printf("Deleted %d rows.\n", rowsAffected)

			if rowsAffected < 1000 {
				break
			}
		}
		message := fmt.Sprintf("Deleted %d rows.\n", total)
		discord_logger.SendDiscordMessage(message)
	}
}
