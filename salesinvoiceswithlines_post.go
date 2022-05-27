package asperion

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-asperion/utils"
)

func (c *Client) NewSalesInvoicesWithLinesPostRequest() SalesInvoicesWithLinesPostRequest {
	r := SalesInvoicesWithLinesPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type SalesInvoicesWithLinesPostRequest struct {
	client      *Client
	queryParams *SalesInvoicesWithLinesPostRequestQueryParams
	pathParams  *SalesInvoicesWithLinesPostRequestPathParams
	method      string
	headers     http.Header
	requestBody SalesInvoicesWithLinesPostRequestBody
}

func (r SalesInvoicesWithLinesPostRequest) NewQueryParams() *SalesInvoicesWithLinesPostRequestQueryParams {
	return &SalesInvoicesWithLinesPostRequestQueryParams{}
}

type SalesInvoicesWithLinesPostRequestQueryParams struct {
	ValidateOnly bool `schema:"validateOnly,omitempty"`
}

func (p SalesInvoicesWithLinesPostRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *SalesInvoicesWithLinesPostRequest) QueryParams() *SalesInvoicesWithLinesPostRequestQueryParams {
	return r.queryParams
}

func (r SalesInvoicesWithLinesPostRequest) NewPathParams() *SalesInvoicesWithLinesPostRequestPathParams {
	return &SalesInvoicesWithLinesPostRequestPathParams{}
}

type SalesInvoicesWithLinesPostRequestPathParams struct {
}

func (p *SalesInvoicesWithLinesPostRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *SalesInvoicesWithLinesPostRequest) PathParams() *SalesInvoicesWithLinesPostRequestPathParams {
	return r.pathParams
}

func (r *SalesInvoicesWithLinesPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *SalesInvoicesWithLinesPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *SalesInvoicesWithLinesPostRequest) Method() string {
	return r.method
}

func (r SalesInvoicesWithLinesPostRequest) NewRequestBody() SalesInvoicesWithLinesPostRequestBody {
	return SalesInvoicesWithLinesPostRequestBody{}
}

type SalesInvoicesWithLinesPostRequestBody SalesInvoice

func (r *SalesInvoicesWithLinesPostRequest) RequestBody() *SalesInvoicesWithLinesPostRequestBody {
	return &r.requestBody
}

func (r *SalesInvoicesWithLinesPostRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *SalesInvoicesWithLinesPostRequest) SetRequestBody(body SalesInvoicesWithLinesPostRequestBody) {
	r.requestBody = body
}

func (r *SalesInvoicesWithLinesPostRequest) NewResponseBody() *SalesInvoicesWithLinesPostResponseBody {
	return &SalesInvoicesWithLinesPostResponseBody{}
}

type SalesInvoicesWithLinesPostResponseBody SalesInvoice

func (r *SalesInvoicesWithLinesPostRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/v1/salesinvoiceswithlines", r.PathParams())
	return &u
}

func (r *SalesInvoicesWithLinesPostRequest) Do() (SalesInvoicesWithLinesPostResponseBody, error) {
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
