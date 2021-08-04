package ikentoo

import (
	"strings"
	"time"
)

type Includes []string

func (i Includes) MarshalSchema() string {
	return strings.Join(i, ",")
}

type Links struct {
	Self struct {
		Href      string `json:"href"`
		Templated bool   `json:"templated"`
	} `json:"self"`
}

type BusinessList []Business

type Business struct {
	BusinessName      string `json:"businessName"`
	BusinessID        int    `json:"businessId"`
	CurrencyCode      string `json:"currencyCode"`
	BusinessLocations []struct {
		BlName   string `json:"blName"`
		BlID     int64  `json:"blId"`
		Country  string `json:"country"`
		Timezone string `json:"timezone"`
	} `json:"businessLocations"`
}

type Sales []Sale

type Sale struct {
	AccountReference string `json:"accountReference"`
	AccountFiscID    string `json:"accountFiscId"`
	Source           struct {
		InitialAccountID  string `json:"initialAccountId"`
		PreviousAccountID string `json:"previousAccountId"`
	} `json:"source,omitempty"`
	SalesLines []struct {
		ID                       string `json:"id"`
		TotalNetAmountWithTax    string `json:"totalNetAmountWithTax"`
		TotalNetAmountWithoutTax string `json:"totalNetAmountWithoutTax"`
		MenuListPrice            string `json:"menuListPrice"`
		UnitCostPrice            string `json:"unitCostPrice"`
		ServiceCharge            string `json:"serviceCharge"`
		DiscountAmount           string `json:"discountAmount"`
		TaxAmount                string `json:"taxAmount"`
		DiscountType             string `json:"discountType"`
		DiscountCode             string `json:"discountCode"`
		DiscountName             string `json:"discountName"`
		AccountDiscountAmount    string `json:"accountDiscountAmount"`
		TotalDiscountAmount      string `json:"totalDiscountAmount"`
		AccountDiscountType      string `json:"accountDiscountType"`
		AccountDiscountCode      string `json:"accountDiscountCode"`
		AccountDiscountName      string `json:"accountDiscountName"`
		Sku                      string `json:"sku"`
		Name                     string `json:"name"`
		StatisticGroup           string `json:"statisticGroup"`
		Quantity                 string `json:"quantity"`
		TaxRatePercentage        string `json:"taxRatePercentage"`
		AccountingGroup          struct {
			AccountingGroupID int64  `json:"accountingGroupId"`
			Name              string `json:"name"`
			Code              string `json:"code"`
		} `json:"accountingGroup"`
		Currency   string        `json:"currency"`
		Tags       []interface{} `json:"tags"`
		Categories []struct {
			Category string `json:"category"`
			Value    string `json:"value"`
		} `json:"categories"`
		TimeOfSale   time.Time `json:"timeOfSale"`
		DeviceID     int       `json:"deviceId"`
		DeviceName   string    `json:"deviceName"`
		VoidReason   string    `json:"voidReason"`
		ParentLineID string    `json:"parentLineId,omitempty"`
	} `json:"salesLines"`
	Payments []struct {
		Code             string `json:"code"`
		Description      string `json:"description"`
		PaymentMethodID  int    `json:"paymentMethodId"`
		NetAmountWithTax string `json:"netAmountWithTax"`
		Currency         string `json:"currency"`
		Tip              string `json:"tip"`
		Type             string `json:"type"`
		DeviceID         int    `json:"deviceId"`
		DeviceName       string `json:"deviceName"`
	} `json:"payments"`
	TimeOfOpening      time.Time `json:"timeOfOpening"`
	TimeOfCloseAndPaid time.Time `json:"timeOfCloseAndPaid"`
	TableName          string    `json:"tableName"`
	Type               string    `json:"type"`
	ExternalReferences []string  `json:"externalReferences,omitempty"`
	NbCovers           float64   `json:"nbCovers"`
	DineIn             bool      `json:"dineIn"`
	DeviceID           int       `json:"deviceId"`
	DeviceName         string    `json:"deviceName"`
	VoidReason         string    `json:"voidReason"`
	ReceiptID          string    `json:"receiptId,omitempty"`
}
