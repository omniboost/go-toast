package asperion

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-asperion/utils"
)

func (c *Client) NewDebtorsGetRequest() DebtorsGetRequest {
	r := DebtorsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type DebtorsGetRequest struct {
	client      *Client
	queryParams *DebtorsGetRequestQueryParams
	pathParams  *DebtorsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody DebtorsGetRequestBody
}

func (r DebtorsGetRequest) NewQueryParams() *DebtorsGetRequestQueryParams {
	return &DebtorsGetRequestQueryParams{}
}

type DebtorsGetRequestQueryParams struct {
	Page               int    `schema:"page,omitempty"`
	PageSize           int    `schema:"pageSize,omitempty"`
	Fields             string `schema:"fields,omitempty"`
	OrderBy            string `schema:"orderBy,omitempty"`
	MetaOnly           bool   `schema:"metaOnly,omitempty"`
	CustomerNumber     string `schema:"CustomerNumber,omitempty"`
	Bban               string `schema:"Bban,omitempty"`
	DirectDebit        bool   `schema:"DirectDebit,omitempty"`
	IBAN               string `schema:"Iban,omitempty"`
	BIC                string `schema:"Bic,omitempty"`
	COCNumber          string `schema:"CocNumber,omitempty"`
	VATNumber          string `schema:"VatNumber,omitempty"`
	City               string `schema:"City,omitempty"`
	Address            string `schema:"Address,omitempty"`
	AddressLine2       string `schema:"AddressLine2,omitempty"`
	AddressLine3       string `schema:"AddressLine3,omitempty"`
	AttentionTo        string `schema:"AttentionTo,omitempty"`
	CountryCode        string `schema:"CountryCode,omitempty"`
	PostCode           string `schema:"PostCode,omitempty"`
	DebtorKindId       string `schema:"DebtorKindId,omitempty"`
	PaymentConditionId int    `schema:"PaymentConditionId,omitempty"`
	PhoneNumber        string `schema:"PhoneNumber,omitempty"`
	FaxNumber          string `schema:"FaxNumber,omitempty"`
	MobileNumber       string `schema:"MobileNumber,omitempty"`
	EmailAddress       string `schema:"EmailAddress,omitempty"`
	Comments           string `schema:"Comments,omitempty"`
	Characteristic4    string `schema:"Characteristic4,omitempty"`
	Characteristic5    string `schema:"Characteristic5,omitempty"`
	PaymentMethodID    string `schema:"PaymentMethodId,omitempty"`
	DateAddedFrom      Date   `schema:"DateAdded_from,omitempty"`
	DateAddedTo        Date   `schema:"DateAdded_to,omitempty"`
	Blocked            bool   `schema:"Blocked,omitempty"`
	NoReminder         bool   `schema:"NoReminder,omitempty"`
	ID                 int    `schema:"Id,omitempty"`
	Name               string `schema:"Name,omitempty"`
}

func (p DebtorsGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *DebtorsGetRequest) QueryParams() *DebtorsGetRequestQueryParams {
	return r.queryParams
}

func (r DebtorsGetRequest) NewPathParams() *DebtorsGetRequestPathParams {
	return &DebtorsGetRequestPathParams{}
}

type DebtorsGetRequestPathParams struct {
}

func (p *DebtorsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *DebtorsGetRequest) PathParams() *DebtorsGetRequestPathParams {
	return r.pathParams
}

func (r *DebtorsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *DebtorsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *DebtorsGetRequest) Method() string {
	return r.method
}

func (r DebtorsGetRequest) NewRequestBody() DebtorsGetRequestBody {
	return DebtorsGetRequestBody{}
}

type DebtorsGetRequestBody struct {
}

func (r *DebtorsGetRequest) RequestBody() *DebtorsGetRequestBody {
	return nil
}

func (r *DebtorsGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *DebtorsGetRequest) SetRequestBody(body DebtorsGetRequestBody) {
	r.requestBody = body
}

func (r *DebtorsGetRequest) NewResponseBody() *DebtorsGetResponseBody {
	return &DebtorsGetResponseBody{}
}

type DebtorsGetResponseBody struct {
	Data Debtors `json:"_data"`
	Meta Meta    `json:"_meta"`
}

func (r *DebtorsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/v1/debtors", r.PathParams())
	return &u
}

func (r *DebtorsGetRequest) Do() (DebtorsGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	req.Header.Add("X-Tenant-Id", strconv.Itoa(r.client.TenantID()))

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
