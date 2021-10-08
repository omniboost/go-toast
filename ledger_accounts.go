package asperion

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-asperion/utils"
)

func (c *Client) NewLedgerAccountsGetRequest() LedgerAccountsGetRequest {
	r := LedgerAccountsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type LedgerAccountsGetRequest struct {
	client      *Client
	queryParams *LedgerAccountsGetRequestQueryParams
	pathParams  *LedgerAccountsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody LedgerAccountsGetRequestBody
}

func (r LedgerAccountsGetRequest) NewQueryParams() *LedgerAccountsGetRequestQueryParams {
	return &LedgerAccountsGetRequestQueryParams{}
}

type LedgerAccountsGetRequestQueryParams struct {
	Page           int    `schema:"page,omitempty"`
	PageSize       int    `schema:"pageSize,omitempty"`
	Fields         string `schema:"fields,omitempty"`
	OrderBy        string `schema:"orderBy,omitempty"`
	MetaOnly       bool   `schema:"metaOnly,omitempty"`
	ID             int    `schema:"Id,omitempty"`
	Type           string `schema:"type,omitempty"`
	Description    string `schema:"description,omitempty"`
	ConversionCode string `schema:"conversionCode,omitempty"`
	CostCenter     string `schema:"costCenter,omitempty"`
}

func (p LedgerAccountsGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *LedgerAccountsGetRequest) QueryParams() *LedgerAccountsGetRequestQueryParams {
	return r.queryParams
}

func (r LedgerAccountsGetRequest) NewPathParams() *LedgerAccountsGetRequestPathParams {
	return &LedgerAccountsGetRequestPathParams{}
}

type LedgerAccountsGetRequestPathParams struct {
}

func (p *LedgerAccountsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *LedgerAccountsGetRequest) PathParams() *LedgerAccountsGetRequestPathParams {
	return r.pathParams
}

func (r *LedgerAccountsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *LedgerAccountsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *LedgerAccountsGetRequest) Method() string {
	return r.method
}

func (r LedgerAccountsGetRequest) NewRequestBody() LedgerAccountsGetRequestBody {
	return LedgerAccountsGetRequestBody{}
}

type LedgerAccountsGetRequestBody struct {
}

func (r *LedgerAccountsGetRequest) RequestBody() *LedgerAccountsGetRequestBody {
	return nil
}

func (r *LedgerAccountsGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *LedgerAccountsGetRequest) SetRequestBody(body LedgerAccountsGetRequestBody) {
	r.requestBody = body
}

func (r *LedgerAccountsGetRequest) NewResponseBody() *LedgerAccountsGetResponseBody {
	return &LedgerAccountsGetResponseBody{}
}

type LedgerAccountsGetResponseBody struct {
	Data LedgerAccounts `json:"_data"`
	Meta Meta           `json:"_meta"`
}

func (r *LedgerAccountsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/v1/ledgeraccounts", r.PathParams())
	return &u
}

func (r *LedgerAccountsGetRequest) Do() (LedgerAccountsGetResponseBody, error) {
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
