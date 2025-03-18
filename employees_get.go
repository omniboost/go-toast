package toast

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-toast/utils"
)

func (c *Client) NewEmployeesGetRequest() EmployeesGetRequest {
	r := EmployeesGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type EmployeesGetRequest struct {
	client      *Client
	queryParams *EmployeesGetRequestQueryParams
	pathParams  *EmployeesGetRequestPathParams
	method      string
	headers     http.Header
	requestBody EmployeesGetRequestBody
}

func (r EmployeesGetRequest) NewQueryParams() *EmployeesGetRequestQueryParams {
	return &EmployeesGetRequestQueryParams{}
}

type EmployeesGetRequestQueryParams struct {
	EmployeeIDs []string `schema:"employeeIds,omitempty"`
}

func (p EmployeesGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *EmployeesGetRequest) QueryParams() *EmployeesGetRequestQueryParams {
	return r.queryParams
}

func (r EmployeesGetRequest) NewPathParams() *EmployeesGetRequestPathParams {
	return &EmployeesGetRequestPathParams{}
}

type EmployeesGetRequestPathParams struct {
}

func (p *EmployeesGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *EmployeesGetRequest) PathParams() *EmployeesGetRequestPathParams {
	return r.pathParams
}

func (r *EmployeesGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *EmployeesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *EmployeesGetRequest) Method() string {
	return r.method
}

func (r EmployeesGetRequest) NewRequestBody() EmployeesGetRequestBody {
	return EmployeesGetRequestBody{}
}

type EmployeesGetRequestBody struct {
}

func (r *EmployeesGetRequest) RequestBody() *EmployeesGetRequestBody {
	return nil
}

func (r *EmployeesGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *EmployeesGetRequest) SetRequestBody(body EmployeesGetRequestBody) {
	r.requestBody = body
}

func (r *EmployeesGetRequest) NewResponseBody() *EmployeesGetResponseBody {
	return &EmployeesGetResponseBody{}
}

type EmployeesGetResponseBody Employees

func (r *EmployeesGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/labor/v1/employees/", r.PathParams())
	return &u
}

func (r *EmployeesGetRequest) Do() (EmployeesGetResponseBody, error, *http.Response) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
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

func (r *EmployeesGetRequest) All() (EmployeesGetResponseBody, error) {
	body, err, resp := r.Do()
	if err != nil {
		return body, err
	}

	concat := body
	token, err := r.client.GetPageToken(resp)
	if err != nil {
		return concat, err
	}

	for token != "" {
		r.QueryParams().PageToken = token
		body, err, resp = r.Do()
		if err != nil {
			return concat, err
		}

		concat = append(concat, body...)
		token, err = r.client.GetPageToken(resp)
		if err != nil {
			return concat, err
		}
	}

	return concat, nil
}
