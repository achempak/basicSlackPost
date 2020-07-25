package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	url := "https://hooks.slack.com/services/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"
	test, _ := json.Marshal(struct {
		Text string `json:"text"`
	}{"Hello world!"})
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(test))
	if err != nil {
		return
	}
	req.Header.Set("Content-type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	if resp.StatusCode != http.StatusCreated {
		resp.Body.Close()
		fmt.Println("Post failed: %s", resp.Status)
		return
	}
	resp.Body.Close()
}
