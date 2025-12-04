package toast

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-toast/utils"
)

func (c *Client) NewEmployeesPostRequest() EmployeesPostRequest {
	r := EmployeesPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type EmployeesPostRequest struct {
	client      *Client
	queryParams *EmployeesPostRequestQueryParams
	pathParams  *EmployeesPostRequestPathParams
	method      string
	headers     http.Header
	requestBody EmployeesPostRequestBody
}

func (r EmployeesPostRequest) NewQueryParams() *EmployeesPostRequestQueryParams {
	return &EmployeesPostRequestQueryParams{}
}

type EmployeesPostRequestQueryParams struct{}

func (p EmployeesPostRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *EmployeesPostRequest) QueryParams() *EmployeesPostRequestQueryParams {
	return r.queryParams
}

func (r EmployeesPostRequest) NewPathParams() *EmployeesPostRequestPathParams {
	return &EmployeesPostRequestPathParams{}
}

type EmployeesPostRequestPathParams struct {
}

func (p *EmployeesPostRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *EmployeesPostRequest) PathParams() *EmployeesPostRequestPathParams {
	return r.pathParams
}

func (r *EmployeesPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *EmployeesPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *EmployeesPostRequest) Method() string {
	return r.method
}

func (r EmployeesPostRequest) NewRequestBody() EmployeesPostRequestBody {
	return EmployeesPostRequestBody{}
}

type EmployeesPostRequestBody PostEmployee

func (r *EmployeesPostRequest) RequestBody() *EmployeesPostRequestBody {
	return &r.requestBody
}

func (r *EmployeesPostRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *EmployeesPostRequest) SetRequestBody(body EmployeesPostRequestBody) {
	r.requestBody = body
}

func (r *EmployeesPostRequest) NewResponseBody() *EmployeesPostResponseBody {
	return &EmployeesPostResponseBody{}
}

type EmployeesPostResponseBody Employees

func (r *EmployeesPostRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/labor/v1/employees/", r.PathParams())
	return &u
}

func (r *EmployeesPostRequest) Do(ctx context.Context) (EmployeesPostResponseBody, error, *http.Response) {
	// Create http request
	req, err := r.client.NewRequest(ctx, r)
	if err != nil {
		return *r.NewResponseBody(), err, nil
	}

	err = r.client.InitToken(req)
	if err != nil {
		return *r.NewResponseBody(), err, nil
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err, nil
	}

	responseBody := r.NewResponseBody()
	resp, err := r.client.Do(req, responseBody)
	return *responseBody, err, resp
}
