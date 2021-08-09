package asperion_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCostunitsGet(t *testing.T) {
	req := client.NewCostunitsGetRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
