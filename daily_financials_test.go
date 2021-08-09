package asperion_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	asperion "github.com/omniboost/go-asperion"
)

func TestDailyFinancialsGet(t *testing.T) {
	req := client.NewDailyFinancialsGetRequest()
	req.PathParams().BusinessID = businessID
	req.QueryParams().Date = asperion.Date{time.Now().AddDate(0, 0, -1)}
	req.QueryParams().Include = append(req.QueryParams().Include, "payments")
	req.QueryParams().Include = append(req.QueryParams().Include, "payment")
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
