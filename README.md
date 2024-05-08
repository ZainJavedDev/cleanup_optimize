# Cleanup and Optimize

This is a Go project that includes functionalities for database cleanup and optimization, as well as a Discord logger for monitoring purpose. The application runs as a cron job daily for cleaning up and optimizing the matches explorer database tables.

## Purpose

This project was needed as the matches explorer is able to get around 1.2 Million matches and over 10 Million players data in a single day, but only needs to keep the data for 8 days. So as older match replays expire, their data is also deleted from the database. The optimization of the database tables is needed as MariaDB Storage Engine does not free disk space even if some rows are deleted. 

## Project Structure

- `main.go`: The entry point of the application.
- `database/`: Contains the Go files for database operations.
  - `cleanup.go`: Contains the cleanup operations for the database.
  - `connect.go`: Contains the connection setup for the database.
  - `optimize.go`: Contains the optimization operations for the database.
- `discord_logger/`: Contains the Go files for logging to Discord.
  - `discord_logger.go`: Contains the implementation for sending messages to Discord.


## How to Run

1. Install the dependencies:

```sh
go mod download
```

2. Run the application 
```sh
go run main.go
```

Please make sure to update your .env file with the necessary environment variables before running the application.

You would need the following values:

- `MYSQL_DB`: The name of the database.
- `MYSQL_USER`: The user for the database.
- `MYSQL_PASSWORD`: The password for the database.
- `WEBHOOK_ID`: The ID of the Discord webhook.
- `WEBHOOK_TOKEN`: The Token of the Discord webhook.

## Future Work

Even after optimizing the tables, not all disk space is freed by MySQL as binary logs and other table logs are created as a backup. I would be implementing their cleanup as well. 