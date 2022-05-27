package asperion_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSalesInvoicesWithLinesGet(t *testing.T) {
	req := client.NewSalesInvoicesWithLinesGetRequest()
	req.QueryParams().InvoiceNumber = "20221214"
	req.QueryParams().InvoiceNumberComparison = "Contains"
	req.QueryParams().JournalID = "70"
	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
