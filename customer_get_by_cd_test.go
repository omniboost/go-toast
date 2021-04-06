package vismanet_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestCustomerGetByCD(t *testing.T) {
	req := client.NewCustomerGetByCD()
	req.PathParams().CustomerCD = "202103"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
