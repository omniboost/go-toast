package asperion

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-asperion/utils"
)

func (c *Client) NewVATCodesGetRequest() VATCodesGetRequest {
	r := VATCodesGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type VATCodesGetRequest struct {
	client      *Client
	queryParams *VATCodesGetRequestQueryParams
	pathParams  *VATCodesGetRequestPathParams
	method      string
	headers     http.Header
	requestBody VATCodesGetRequestBody
}

func (r VATCodesGetRequest) NewQueryParams() *VATCodesGetRequestQueryParams {
	return &VATCodesGetRequestQueryParams{}
}

type VATCodesGetRequestQueryParams struct {
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

func (p VATCodesGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *VATCodesGetRequest) QueryParams() *VATCodesGetRequestQueryParams {
	return r.queryParams
}

func (r VATCodesGetRequest) NewPathParams() *VATCodesGetRequestPathParams {
	return &VATCodesGetRequestPathParams{}
}

type VATCodesGetRequestPathParams struct {
}

func (p *VATCodesGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *VATCodesGetRequest) PathParams() *VATCodesGetRequestPathParams {
	return r.pathParams
}

func (r *VATCodesGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *VATCodesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *VATCodesGetRequest) Method() string {
	return r.method
}

func (r VATCodesGetRequest) NewRequestBody() VATCodesGetRequestBody {
	return VATCodesGetRequestBody{}
}

type VATCodesGetRequestBody struct {
}

func (r *VATCodesGetRequest) RequestBody() *VATCodesGetRequestBody {
	return nil
}

func (r *VATCodesGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *VATCodesGetRequest) SetRequestBody(body VATCodesGetRequestBody) {
	r.requestBody = body
}

func (r *VATCodesGetRequest) NewResponseBody() *VATCodesGetResponseBody {
	return &VATCodesGetResponseBody{}
}

type VATCodesGetResponseBody struct {
	Data Administrations `json:"_data"`
}

func (r *VATCodesGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/v1/vatcodes", r.PathParams())
	return &u
}

func (r *VATCodesGetRequest) Do() (VATCodesGetResponseBody, error) {
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
