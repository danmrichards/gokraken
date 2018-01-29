package gokraken

import (
	"context"
	"fmt"
	"testing"
)

func TestUserData_Balance(t *testing.T) {
	// TODO: Table test this with a mock api.

	k := NewWithAuth(
		"YcYrrxEJvXNk/mwAhPXkaXykMFpqaf7psfdh5lPLPEXTMpHWu8el7tYt",
		"D/dLihFVqHctD4QJJUU/B2jUNjvEk4ZDoRz27h1rBncjJDoyvQplawH3U4suyP08KXV4zphOYzrpNvjYCdGbfw==",
	)

	resp, err := k.UserData.Balance(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("balance: %+v\n", resp)
}
