package asperion

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-asperion/utils"
)

func (c *Client) NewSalesInvoicesGetRequest() SalesInvoicesGetRequest {
	r := SalesInvoicesGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type SalesInvoicesGetRequest struct {
	client      *Client
	queryParams *SalesInvoicesGetRequestQueryParams
	pathParams  *SalesInvoicesGetRequestPathParams
	method      string
	headers     http.Header
	requestBody SalesInvoicesGetRequestBody
}

func (r SalesInvoicesGetRequest) NewQueryParams() *SalesInvoicesGetRequestQueryParams {
	return &SalesInvoicesGetRequestQueryParams{}
}

type SalesInvoicesGetRequestQueryParams struct {
	// Page            int    `schema:"page,omitempty"`
	// PageSize        int    `schema:"pageSize,omitempty"`
	// Fields          string `schema:"fields,omitempty"`
	// OrderBy         string `schema:"orderBy,omitempty"`
	// MetaOnly        bool   `schema:"metaOnly,omitempty"`
	// ID              int    `schema:"Id,omitempty"`
	// Name            string `schema:"Name,omitempty"`
	// TelephoneNumber string `schema:"TelephoneNumber,omitempty"`
	// Email           string `schema:"Email,omitempty"`
}

func (p SalesInvoicesGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *SalesInvoicesGetRequest) QueryParams() *SalesInvoicesGetRequestQueryParams {
	return r.queryParams
}

func (r SalesInvoicesGetRequest) NewPathParams() *SalesInvoicesGetRequestPathParams {
	return &SalesInvoicesGetRequestPathParams{}
}

type SalesInvoicesGetRequestPathParams struct {
}

func (p *SalesInvoicesGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *SalesInvoicesGetRequest) PathParams() *SalesInvoicesGetRequestPathParams {
	return r.pathParams
}

func (r *SalesInvoicesGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *SalesInvoicesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *SalesInvoicesGetRequest) Method() string {
	return r.method
}

func (r SalesInvoicesGetRequest) NewRequestBody() SalesInvoicesGetRequestBody {
	return SalesInvoicesGetRequestBody{}
}

type SalesInvoicesGetRequestBody struct {
}

func (r *SalesInvoicesGetRequest) RequestBody() *SalesInvoicesGetRequestBody {
	return nil
}

func (r *SalesInvoicesGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *SalesInvoicesGetRequest) SetRequestBody(body SalesInvoicesGetRequestBody) {
	r.requestBody = body
}

func (r *SalesInvoicesGetRequest) NewResponseBody() *SalesInvoicesGetResponseBody {
	return &SalesInvoicesGetResponseBody{}
}

type SalesInvoicesGetResponseBody struct {
	Data SalesInvoices `json:"_data"`
	Meta Meta          `json:"_meta"`
}

func (r *SalesInvoicesGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/v1/salesinvoices", r.PathParams())
	return &u
}

func (r *SalesInvoicesGetRequest) Do() (SalesInvoicesGetResponseBody, error) {
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
