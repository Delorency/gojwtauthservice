package tools

import (
	"auth/internal/schemes"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func SendToWebhook(ip, wburl string) {
	webbody := schemes.WebhookPayload{Message: fmt.Sprintf("Попытка входа с IP: %s", ip)}

	jsonbytes, err := json.Marshal(webbody)
	if err != nil {
		log.Printf("refresh service | Convert ot byte error\n")
	}

	resp, err := http.Post(wburl, "application/json", bytes.NewReader(jsonbytes))

	if err != nil {
		log.Printf("refresh service | send to webhook error\n")
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		log.Printf("Webhook bad status: %d\n", resp.StatusCode)
	}
}
