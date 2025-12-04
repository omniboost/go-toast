package toast_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/omniboost/go-toast"
)

func TestMenuItemsGet(t *testing.T) {
	req := client.NewMenuItemsGetRequest()
	req.QueryParams().LastModified = toast.DateTime{time.Date(2025, 3, 30, 0, 0, 0, 0, time.UTC)}
	resp, err, _ := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
