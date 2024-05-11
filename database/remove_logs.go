package database

func RemoveBinLogs() error {
	db, err := Connect()
	if err != nil {
		return err
	}

	defer db.Close()

	sql := "PURGE BINARY LOGS BEFORE NOW()"
	_, err = db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}
