package asperion_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJournalEntriesWithLinesGet(t *testing.T) {
	req := client.NewJournalEntriesWithLinesGetRequest()
	req.QueryParams().BookYear = 2021
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
