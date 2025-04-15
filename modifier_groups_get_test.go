package toast_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestModifierGroupsGet(t *testing.T) {
	req := client.NewModifierGroupsGetRequest()
	resp, err, _ := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
