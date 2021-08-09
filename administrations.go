package asperion

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-asperion/utils"
)

func (c *Client) NewAdministrationsGetRequest() AdministrationsGetRequest {
	r := AdministrationsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type AdministrationsGetRequest struct {
	client      *Client
	queryParams *AdministrationsGetRequestQueryParams
	pathParams  *AdministrationsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody AdministrationsGetRequestBody
}

func (r AdministrationsGetRequest) NewQueryParams() *AdministrationsGetRequestQueryParams {
	return &AdministrationsGetRequestQueryParams{}
}

type AdministrationsGetRequestQueryParams struct {
	Page            int    `schema:"page,omitempty"`
	PageSize        int    `schema:"pageSize,omitempty"`
	Fields          string `schema:"fields,omitempty"`
	OrderBy         string `schema:"orderBy,omitempty"`
	MetaOnly        bool   `schema:"metaOnly,omitempty"`
	ID              int    `schema:"Id,omitempty"`
	Name            string `schema:"Name,omitempty"`
	TelephoneNumber string `schema:"TelephoneNumber,omitempty"`
	Email           string `schema:"Email,omitempty"`
}

func (p AdministrationsGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *AdministrationsGetRequest) QueryParams() *AdministrationsGetRequestQueryParams {
	return r.queryParams
}

func (r AdministrationsGetRequest) NewPathParams() *AdministrationsGetRequestPathParams {
	return &AdministrationsGetRequestPathParams{}
}

type AdministrationsGetRequestPathParams struct {
}

func (p *AdministrationsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *AdministrationsGetRequest) PathParams() *AdministrationsGetRequestPathParams {
	return r.pathParams
}

func (r *AdministrationsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *AdministrationsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *AdministrationsGetRequest) Method() string {
	return r.method
}

func (r AdministrationsGetRequest) NewRequestBody() AdministrationsGetRequestBody {
	return AdministrationsGetRequestBody{}
}

type AdministrationsGetRequestBody struct {
}

func (r *AdministrationsGetRequest) RequestBody() *AdministrationsGetRequestBody {
	return nil
}

func (r *AdministrationsGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *AdministrationsGetRequest) SetRequestBody(body AdministrationsGetRequestBody) {
	r.requestBody = body
}

func (r *AdministrationsGetRequest) NewResponseBody() *AdministrationsGetResponseBody {
	return &AdministrationsGetResponseBody{}
}

type AdministrationsGetResponseBody struct {
	Data Administrations `json:"_data"`
	Meta Meta            `json:"_meta"`
}

func (r *AdministrationsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/v1/administrations", r.PathParams())
	return &u
}

func (r *AdministrationsGetRequest) Do() (AdministrationsGetResponseBody, error) {
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
