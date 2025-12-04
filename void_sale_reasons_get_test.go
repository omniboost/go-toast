package toast_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestVoidSaleReasonsGetRequest(t *testing.T) {
	req := client.NewVoidSaleReasonsGetRequest()
	resp, err, _ := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
