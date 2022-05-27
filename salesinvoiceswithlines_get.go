package asperion

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-asperion/utils"
)

func (c *Client) NewSalesInvoicesWithLinesGetRequest() SalesInvoicesWithLinesGetRequest {
	r := SalesInvoicesWithLinesGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type SalesInvoicesWithLinesGetRequest struct {
	client      *Client
	queryParams *SalesInvoicesWithLinesGetRequestQueryParams
	pathParams  *SalesInvoicesWithLinesGetRequestPathParams
	method      string
	headers     http.Header
	requestBody SalesInvoicesWithLinesGetRequestBody
}

func (r SalesInvoicesWithLinesGetRequest) NewQueryParams() *SalesInvoicesWithLinesGetRequestQueryParams {
	return &SalesInvoicesWithLinesGetRequestQueryParams{}
}

type SalesInvoicesWithLinesGetRequestQueryParams struct {
	Page                 int     `schema:"page,omitempty"`
	PageSize             int     `schema:"pageSize,omitempty"`
	Fields               string  `schema:"fields,omitempty"`
	OrderBy              string  `schema:"orderBy,omitempty"`
	MetaOnly             bool    `schema:"metaOnly,omitempty"`
	ID                   int     `schema:"Id,omitempty"`
	AmountExclVat        float64 `schema:"AmountExclVat,omitempty"`
	AmountInclVat        float64 `schema:"AmountInclVat,omitempty"`
	Bookyear             int     `schema:"Bookyear,omitempty"`
	Bookperiod           int     `schema:"Bookperiod,omitempty"`
	CostCenterId         int     `schema:"CostCenterId,omitempty"`
	DebtorCustomerNumber string  `schema:"DebtorCustomerNumber,omitempty"`
	Date                 Date    `schema:"Date,omitempty"`
	DebtorID             int     `schema:"DebtorId,omitempty"`
	Description          string  `schema:"Description,omitempty"`
	InvoiceNumber        string  `schema:"InvoiceNumber,omitempty"`
	// Sets the comparison type for the InvoiceNumber filter:
	// - Contains
	// - Equals
	// - NotEquals
	// - StartsWith
	// - EndsWith
	// when not provided defaults to : Contains
	InvoiceNumberComparison string `schema:"InvoiceNumber_comparison,omitempty"`
	JournalID               string `schema:"JournalId,omitempty"`
	// Filter on Status
	// - 1 = Concept
	// - 2 = InJournalentry
	// - 3 = InLedger
	// - 6 = Payable
	// - 7 = Paid
	// - 8 = Hold
	// Available values : 1, 2, 3, 6, 7, 8
	Status             int     `schema:"Status,omitempty"`
	PaymentConditionID int     `schema:"PaymentConditionId,omitempty"`
	VATAmount          float64 `schema:"VatAmount,omitempty"`
}

func (p SalesInvoicesWithLinesGetRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *SalesInvoicesWithLinesGetRequest) QueryParams() *SalesInvoicesWithLinesGetRequestQueryParams {
	return r.queryParams
}

func (r SalesInvoicesWithLinesGetRequest) NewPathParams() *SalesInvoicesWithLinesGetRequestPathParams {
	return &SalesInvoicesWithLinesGetRequestPathParams{}
}

type SalesInvoicesWithLinesGetRequestPathParams struct {
}

func (p *SalesInvoicesWithLinesGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *SalesInvoicesWithLinesGetRequest) PathParams() *SalesInvoicesWithLinesGetRequestPathParams {
	return r.pathParams
}

func (r *SalesInvoicesWithLinesGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *SalesInvoicesWithLinesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *SalesInvoicesWithLinesGetRequest) Method() string {
	return r.method
}

func (r SalesInvoicesWithLinesGetRequest) NewRequestBody() SalesInvoicesWithLinesGetRequestBody {
	return SalesInvoicesWithLinesGetRequestBody{}
}

type SalesInvoicesWithLinesGetRequestBody struct {
}

func (r *SalesInvoicesWithLinesGetRequest) RequestBody() *SalesInvoicesWithLinesGetRequestBody {
	return nil
}

func (r *SalesInvoicesWithLinesGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *SalesInvoicesWithLinesGetRequest) SetRequestBody(body SalesInvoicesWithLinesGetRequestBody) {
	r.requestBody = body
}

func (r *SalesInvoicesWithLinesGetRequest) NewResponseBody() *SalesInvoicesWithLinesGetResponseBody {
	return &SalesInvoicesWithLinesGetResponseBody{}
}

type SalesInvoicesWithLinesGetResponseBody struct {
	Data SalesInvoices `json:"_data"`
	Meta Meta          `json:"_meta"`
}

func (r *SalesInvoicesWithLinesGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/v1/salesinvoiceswithlines", r.PathParams())
	return &u
}

func (r *SalesInvoicesWithLinesGetRequest) Do() (SalesInvoicesWithLinesGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}

func (r *SalesInvoicesWithLinesGetRequest) All() (SalesInvoicesWithLinesGetResponseBody, error) {
	resp, err := r.Do()
	if err != nil {
		return resp, err
	}

	concat := resp.Data

	for resp.Meta.NextPage != "" {
		u, err := url.Parse(resp.Meta.NextPage)
		if err != nil {
			return resp, err
		}

		page, err := strconv.Atoi(u.Query().Get("page"))
		if err != nil {
			return resp, err
		}

		r.QueryParams().Page = page
		resp, err = r.Do()
		if err != nil {
			return resp, err
		}

		concat = append(concat, resp.Data...)
	}

	resp.Data = concat
	return resp, nil
}
