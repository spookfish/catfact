package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Fact []struct {
	Status struct {
		Verified  bool `json:"verified"`
		SentCount int  `json:"sentCount"`
	} `json:"status"`
	ID        string    `json:"_id"`
	User      string    `json:"user"`
	Text      string    `json:"text"`
	V         int       `json:"__v"`
	Source    string    `json:"source"`
	UpdatedAt time.Time `json:"updatedAt"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
	Deleted   bool      `json:"deleted"`
	Used      bool      `json:"used"`
}

func main() {
	resp, err := http.Get("https://cat-fact.herokuapp.com/facts")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var fact Fact
	json.Unmarshal(body, &fact)
	fmt.Println("Daily cat fact:", fact[0].Text)
}
