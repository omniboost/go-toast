package ikentoo_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestBusinessesGet(t *testing.T) {
	req := client.NewBusinessesGetRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
