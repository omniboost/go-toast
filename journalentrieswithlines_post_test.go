package asperion_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	asperion "github.com/omniboost/go-asperion"
)

func TestJournalEntriesWithLinesPost(t *testing.T) {
	req := client.NewJournalEntriesWithLinesPostRequest()
	req.RequestBody().Lines.Data = asperion.JournalEntryLines{
		{
			Amount:          100,
			CostCenterID:    30,
			BookingDate:     asperion.Date{time.Now()},
			Description:     "Debit",
			LedgerAccountID: 20,
		},
		{
			Amount:          -100,
			CostCenterID:    0,
			BookingDate:     asperion.Date{time.Now()},
			Description:     "Credit",
			LedgerAccountID: 20,
		},
	}
	req.RequestBody().JournalID = "90"
	req.RequestBody().Date = asperion.Date{time.Now()}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
