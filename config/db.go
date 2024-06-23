package config

import (
	"fmt"

	supa "github.com/supabase-community/supabase-go"
)

const (
	supabaseURL = "https://fcmbzfgubcylcmzsprwc.supabase.co"
	apiKey      = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImZjbWJ6Zmd1YmN5bGNtenNwcndjIiwicm9sZSI6ImFub24iLCJpYXQiOjE3MTExOTgwMjksImV4cCI6MjAyNjc3NDAyOX0.E1QFqkxZlLnRaNEv6zvviTx7w-jFiobYu9DjpR_P0a0"
)

func CreateClient() *supa.Client {
	client, err := supa.NewClient(supabaseURL, apiKey, nil)
	if err != nil {
		fmt.Println("cannot initalize client", err)
	}
	return client

}
