package asperion_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestDebtorsGet(t *testing.T) {
	req := client.NewDebtorsGetRequest()
	req.QueryParams().Name = "ABC Automotive"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
