package asperion_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestPaymentConditionsGet(t *testing.T) {
	req := client.NewPaymentConditionsGetRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
