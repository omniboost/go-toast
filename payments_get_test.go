package toast_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/omniboost/go-toast"
)

func TestPaymentsGet(t *testing.T) {
	req := client.NewPaymentsGetRequest()
	req.QueryParams().PaidBusinessDate = toast.Date{time.Date(2025, 9, 26, 0, 0, 0, 0, time.UTC)}
	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	for _, paymentID := range resp {
		paymentReq := client.NewPaymentGetRequest()
		paymentReq.PathParams().GUID = paymentID
		paymentResp, err, _ := paymentReq.Do()
		if err != nil {
			t.Error(err)
		}

		b, _ := json.MarshalIndent(paymentResp, "", "  ")
		fmt.Println(string(b))
	}
}
