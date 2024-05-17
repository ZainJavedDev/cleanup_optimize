package discord_logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type DiscordMessage struct {
	Content string `json:"content"`
}

func SendDiscordMessage(message string) error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	var token, webhookId string
	token = os.Getenv("WEBHOOK_ID")
	webhookId = os.Getenv("WEBHOOK_TOKEN")
	webhookURL := fmt.Sprintf("https://discord.com/api/webhooks/%s/%s", token, webhookId)

	msg := DiscordMessage{
		Content: "pgsql: " + message,
	}
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(msgBytes))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		fmt.Println(resp.Body)
	}
	return nil
}
