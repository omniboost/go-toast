package toast_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestRestaurantGet(t *testing.T) {
	req := client.NewRestaurantGetRequest()
	req.PathParams().GUID = client.ToastRestaurantExternalID()
	resp, err, _ := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
