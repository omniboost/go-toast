package toast_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestLoginPost(t *testing.T) {
	req := client.NewLoginPostRequest()
	req.RequestBody().ClientID = client.ClientID()
	req.RequestBody().ClientSecret = client.ClientSecret()
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
