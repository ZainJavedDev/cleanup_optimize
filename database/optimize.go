package database

import (
	"fmt"

	"github.com/ZainJavedDev/cleanup_optimize/discord_logger"
)

func VacuumFullTables() error {
	db, err := Connect()
	if err != nil {
		return err
	}
	var sql string
	defer db.Close()
	for _, game_mode := range GAME_MODES {
		sql = fmt.Sprintf("VACUUM FULL %s_matches", game_mode)
		_, err = db.Exec(sql)
		if err != nil {
			return err
		}

		sql = fmt.Sprintf("VACUUM FULL %s_players", game_mode)
		_, err = db.Exec(sql)
		if err != nil {
			return err
		}
		message := fmt.Sprintf("Vacuumed full on %s tables", game_mode)
		err = discord_logger.SendDiscordMessage(message)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(message)
	}
	return nil
}
