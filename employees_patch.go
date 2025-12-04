package toast

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-toast/utils"
)

func (c *Client) NewEmployeesPatchRequest() EmployeesPatchRequest {
	r := EmployeesPatchRequest{
		client:  c,
		method:  http.MethodPatch,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type EmployeesPatchRequest struct {
	client      *Client
	queryParams *EmployeesPatchRequestQueryParams
	pathParams  *EmployeesPatchRequestPathParams
	method      string
	headers     http.Header
	requestBody EmployeesPatchRequestBody
}

func (r EmployeesPatchRequest) NewQueryParams() *EmployeesPatchRequestQueryParams {
	return &EmployeesPatchRequestQueryParams{}
}

type EmployeesPatchRequestQueryParams struct{}

func (p EmployeesPatchRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *EmployeesPatchRequest) QueryParams() *EmployeesPatchRequestQueryParams {
	return r.queryParams
}

func (r EmployeesPatchRequest) NewPathParams() *EmployeesPatchRequestPathParams {
	return &EmployeesPatchRequestPathParams{}
}

type EmployeesPatchRequestPathParams struct {
	EmployeeID string `schema:"employee_id"`
}

func (p *EmployeesPatchRequestPathParams) Params() map[string]string {
	return map[string]string{
		"employee_id": p.EmployeeID,
	}
}

func (r *EmployeesPatchRequest) PathParams() *EmployeesPatchRequestPathParams {
	return r.pathParams
}

func (r *EmployeesPatchRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *EmployeesPatchRequest) SetMethod(method string) {
	r.method = method
}

func (r *EmployeesPatchRequest) Method() string {
	return r.method
}

func (r EmployeesPatchRequest) NewRequestBody() EmployeesPatchRequestBody {
	return EmployeesPatchRequestBody{}
}

type EmployeesPatchRequestBody PatchEmployee

func (r *EmployeesPatchRequest) RequestBody() *EmployeesPatchRequestBody {
	return &r.requestBody
}

func (r *EmployeesPatchRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *EmployeesPatchRequest) SetRequestBody(body EmployeesPatchRequestBody) {
	r.requestBody = body
}

func (r *EmployeesPatchRequest) NewResponseBody() *EmployeesPatchResponseBody {
	return &EmployeesPatchResponseBody{}
}

type EmployeesPatchResponseBody Employees

func (r *EmployeesPatchRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/labor/v1/employees/{{.employee_id}}", r.PathParams())
	return &u
}

func (r *EmployeesPatchRequest) Do(ctx context.Context) (EmployeesPatchResponseBody, error, *http.Response) {
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
