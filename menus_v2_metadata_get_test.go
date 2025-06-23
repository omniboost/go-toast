package toast_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMenuV2MetadataGet(t *testing.T) {
	req := client.NewMenuV2MetadataGetRequest()
	resp, err, _ := req.Do()
	if err != nil {
		t.Error(err)
	}
	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
