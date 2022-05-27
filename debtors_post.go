package asperion

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-asperion/utils"
)

func (c *Client) NewDebtorsPostRequest() DebtorsPostRequest {
	r := DebtorsPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type DebtorsPostRequest struct {
	client      *Client
	queryParams *DebtorsPostRequestQueryParams
	pathParams  *DebtorsPostRequestPathParams
	method      string
	headers     http.Header
	requestBody DebtorsPostRequestBody
}

func (r DebtorsPostRequest) NewQueryParams() *DebtorsPostRequestQueryParams {
	return &DebtorsPostRequestQueryParams{}
}

type DebtorsPostRequestQueryParams struct {
	ValidateOnly bool `schema:"validateOnly,omitempty"`
}

func (p DebtorsPostRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *DebtorsPostRequest) QueryParams() *DebtorsPostRequestQueryParams {
	return r.queryParams
}

func (r DebtorsPostRequest) NewPathParams() *DebtorsPostRequestPathParams {
	return &DebtorsPostRequestPathParams{}
}

type DebtorsPostRequestPathParams struct {
}

func (p *DebtorsPostRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *DebtorsPostRequest) PathParams() *DebtorsPostRequestPathParams {
	return r.pathParams
}

func (r *DebtorsPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *DebtorsPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *DebtorsPostRequest) Method() string {
	return r.method
}

func (r DebtorsPostRequest) NewRequestBody() DebtorsPostRequestBody {
	return DebtorsPostRequestBody{}
}

type DebtorsPostRequestBody Debtor

func (r *DebtorsPostRequest) RequestBody() *DebtorsPostRequestBody {
	return &r.requestBody
}

func (r *DebtorsPostRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *DebtorsPostRequest) SetRequestBody(body DebtorsPostRequestBody) {
	r.requestBody = body
}

func (r *DebtorsPostRequest) NewResponseBody() *DebtorsPostResponseBody {
	return &DebtorsPostResponseBody{}
}

type DebtorsPostResponseBody struct {
	Data Costunits `json:"_data"`
	Meta Meta      `json:"_meta"`
}

func (r *DebtorsPostRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/v1/debtors", r.PathParams())
	return &u
}

func (r *DebtorsPostRequest) Do() (DebtorsPostResponseBody, error) {
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
