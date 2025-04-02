package toast_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestStockInventoryGet(t *testing.T) {
	req := client.NewStockInventoryGetRequest()
	req.QueryParams().Status = "QUANTITY"
	resp, err, _ := req.Do()
	if err != nil {
		t.Error(err)
	}
	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
