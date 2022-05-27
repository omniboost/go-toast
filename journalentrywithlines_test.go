package asperion_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJournalEntryWithLinesGet(t *testing.T) {
	req := client.NewJournalEntryWithLinesGetRequest()
	req.PathParams().ID = 35
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
