package asperion

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-asperion/utils"
)

func (c *Client) NewSalesInvoicesWithLinesFilePostRequest() SalesInvoicesWithLinesFilePostRequest {
	r := SalesInvoicesWithLinesFilePostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.formParams = r.NewSalesInvoicesWithLinesFilePostFormParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type SalesInvoicesWithLinesFilePostRequest struct {
	client      *Client
	queryParams *SalesInvoicesWithLinesFilePostRequestQueryParams
	pathParams  *SalesInvoicesWithLinesFilePostRequestPathParams
	formParams  *SalesInvoicesWithLinesFilePostFormParams
	method      string
	headers     http.Header
	requestBody SalesInvoicesWithLinesFilePostRequestBody
}

func (r SalesInvoicesWithLinesFilePostRequest) NewQueryParams() *SalesInvoicesWithLinesFilePostRequestQueryParams {
	return &SalesInvoicesWithLinesFilePostRequestQueryParams{}
}

type SalesInvoicesWithLinesFilePostRequestQueryParams struct {
	ValidateOnly bool `schema:"validateOnly,omitempty"`
}

func (p SalesInvoicesWithLinesFilePostRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *SalesInvoicesWithLinesFilePostRequest) QueryParams() *SalesInvoicesWithLinesFilePostRequestQueryParams {
	return r.queryParams
}

func (r *SalesInvoicesWithLinesFilePostRequest) FormParams() *SalesInvoicesWithLinesFilePostFormParams {
	return r.formParams
}

func (r SalesInvoicesWithLinesFilePostRequest) NewPathParams() *SalesInvoicesWithLinesFilePostRequestPathParams {
	return &SalesInvoicesWithLinesFilePostRequestPathParams{}
}

func (r SalesInvoicesWithLinesFilePostRequest) NewSalesInvoicesWithLinesFilePostFormParams() *SalesInvoicesWithLinesFilePostFormParams {
	return &SalesInvoicesWithLinesFilePostFormParams{}
}

type SalesInvoicesWithLinesFilePostFormParams struct {
	File FormFile
}

func (p SalesInvoicesWithLinesFilePostFormParams) Values() url.Values {
	return url.Values{}
}

func (p SalesInvoicesWithLinesFilePostFormParams) Files() map[string]FormFile {
	return map[string]FormFile{
		"File": p.File,
	}
}

type SalesInvoicesWithLinesFilePostRequestPathParams struct {
	ID string
}

func (p *SalesInvoicesWithLinesFilePostRequestPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ID,
	}
}

func (r *SalesInvoicesWithLinesFilePostRequest) PathParams() *SalesInvoicesWithLinesFilePostRequestPathParams {
	return r.pathParams
}

func (r *SalesInvoicesWithLinesFilePostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *SalesInvoicesWithLinesFilePostRequest) SetMethod(method string) {
	r.method = method
}

func (r *SalesInvoicesWithLinesFilePostRequest) Method() string {
	return r.method
}

func (r SalesInvoicesWithLinesFilePostRequest) NewRequestBody() SalesInvoicesWithLinesFilePostRequestBody {
	return bytes.NewReader([]byte{})
}

type SalesInvoicesWithLinesFilePostRequestBody io.Reader

func (r *SalesInvoicesWithLinesFilePostRequest) RequestBody() *SalesInvoicesWithLinesFilePostRequestBody {
	return &r.requestBody
}

func (r *SalesInvoicesWithLinesFilePostRequest) RequestBodyInterface() interface{} {
	return r.requestBody
}

func (r *SalesInvoicesWithLinesFilePostRequest) SetRequestBody(body SalesInvoicesWithLinesFilePostRequestBody) {
	r.requestBody = body
}

func (r *SalesInvoicesWithLinesFilePostRequest) NewResponseBody() *SalesInvoicesWithLinesFilePostResponseBody {
	return &SalesInvoicesWithLinesFilePostResponseBody{}
}

type SalesInvoicesWithLinesFilePostResponseBody File

func (r *SalesInvoicesWithLinesFilePostRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/v1/salesinvoiceswithlines/{{.id}}/file", r.PathParams())
	return &u
}

func (r *SalesInvoicesWithLinesFilePostRequest) Do() (SalesInvoicesWithLinesFilePostResponseBody, error) {
	// Create http request
	req, err := r.client.NewFormRequest(nil, r.Method(), *r.URL(), r.FormParams())
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
