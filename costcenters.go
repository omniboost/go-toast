package asperion

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-asperion/utils"
)

func (c *Client) NewCostcentersGetRequest() CostcentersGetRequest {
	r := CostcentersGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CostcentersGetRequest struct {
	client      *Client
	queryParams *CostcentersGetRequestQueryParams
	pathParams  *CostcentersGetRequestPathParams
	method      string
	headers     http.Header
	requestBody CostcentersGetRequestBody
}

func (r CostcentersGetRequest) NewQueryParams() *CostcentersGetRequestQueryParams {
	return &CostcentersGetRequestQueryParams{}
}

type CostcentersGetRequestQueryParams struct {
	Page           int    `schema:"page,omitempty"`
	PageSize       int    `schema:"pageSize,omitempty"`
	Fields         string `schema:"fields,omitempty"`
	OrderBy        string `schema:"orderBy,omitempty"`
	MetaOnly       bool   `schema:"metaOnly,omitempty"`
	ID             int    `schema:"Id,omitempty"`
	Description    string `schema:"description,omitempty"`
	ConversionCode string `schema:"conversionCode,omitempty"`
}

func (p CostcentersGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CostcentersGetRequest) QueryParams() *CostcentersGetRequestQueryParams {
	return r.queryParams
}

func (r CostcentersGetRequest) NewPathParams() *CostcentersGetRequestPathParams {
	return &CostcentersGetRequestPathParams{}
}

type CostcentersGetRequestPathParams struct {
}

func (p *CostcentersGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CostcentersGetRequest) PathParams() *CostcentersGetRequestPathParams {
	return r.pathParams
}

func (r *CostcentersGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CostcentersGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *CostcentersGetRequest) Method() string {
	return r.method
}

func (r CostcentersGetRequest) NewRequestBody() CostcentersGetRequestBody {
	return CostcentersGetRequestBody{}
}

type CostcentersGetRequestBody struct {
}

func (r *CostcentersGetRequest) RequestBody() *CostcentersGetRequestBody {
	return nil
}

func (r *CostcentersGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *CostcentersGetRequest) SetRequestBody(body CostcentersGetRequestBody) {
	r.requestBody = body
}

func (r *CostcentersGetRequest) NewResponseBody() *CostcentersGetResponseBody {
	return &CostcentersGetResponseBody{}
}

type CostcentersGetResponseBody struct {
	Data Costcenters `json:"_data"`
	Meta Meta        `json:"_meta"`
}

func (r *CostcentersGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/v1/costcenters", r.PathParams())
	return &u
}

func (r *CostcentersGetRequest) Do() (CostcentersGetResponseBody, error) {
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
