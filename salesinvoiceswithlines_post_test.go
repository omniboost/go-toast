package asperion_test

import (
	"encoding/json"
	"fmt"
	"testing"

	asperion "github.com/omniboost/go-asperion"
)

func TestSalesInvoicesWithLinesPost(t *testing.T) {
	req := client.NewSalesInvoicesWithLinesPostRequest()
	req.RequestBody().Lines.Data = asperion.SalesInvoiceLines{
		{},
		{},
	}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
