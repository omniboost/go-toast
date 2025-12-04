package toast_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestOrdersGet(t *testing.T) {
	req := client.NewOrdersGetRequest()
	req.PathParams().OrderGUID = "71213cdd-b21a-43f2-ab40-04a824034f37"
	resp, err, _ := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
