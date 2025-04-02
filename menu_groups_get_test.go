package toast_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/omniboost/go-toast"
)

func TestMenuGroupsGet(t *testing.T) {
	req := client.NewMenuGroupsGetRequest()
	req.QueryParams().LastModified = toast.DateTime{time.Date(2025, 3, 30, 0, 0, 0, 0, time.UTC)}
	resp, err, _ := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
