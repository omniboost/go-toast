package toast_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestRevenueCentersGet(t *testing.T) {
	req := client.NewRevenueCentersGetRequest()
	resp, err, _ := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
