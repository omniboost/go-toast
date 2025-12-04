package toast_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestPartnerRestaurantsGet(t *testing.T) {
	req := client.NewPartnerRestaurantsGetRequest()
	resp, err, _ := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
