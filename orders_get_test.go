package toast_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestOrdersGet(t *testing.T) {
	req := client.NewOrdersGetRequest()
	req.PathParams().OrderGUID = "5b70e749-a365-4d13-a78b-94a55ea7e8e7"
	resp, err, _ := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}

