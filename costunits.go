package asperion

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-asperion/utils"
)

func (c *Client) NewCostunitsGetRequest() CostunitsGetRequest {
	r := CostunitsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CostunitsGetRequest struct {
	client      *Client
	queryParams *CostunitsGetRequestQueryParams
	pathParams  *CostunitsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody CostunitsGetRequestBody
}

func (r CostunitsGetRequest) NewQueryParams() *CostunitsGetRequestQueryParams {
	return &CostunitsGetRequestQueryParams{}
}

type CostunitsGetRequestQueryParams struct {
	Page           int    `schema:"page,omitempty"`
	PageSize       int    `schema:"pageSize,omitempty"`
	Fields         string `schema:"fields,omitempty"`
	OrderBy        string `schema:"orderBy,omitempty"`
	MetaOnly       bool   `schema:"metaOnly,omitempty"`
	ID             int    `schema:"Id,omitempty"`
	Type           string `schema:"type,omitempty"`
	Description    string `schema:"description,omitempty"`
	ConversionCode string `schema:"conversionCode,omitempty"`
	CostCenter     string `schema:"costCenter,omitempty"`
}

func (p CostunitsGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CostunitsGetRequest) QueryParams() *CostunitsGetRequestQueryParams {
	return r.queryParams
}

func (r CostunitsGetRequest) NewPathParams() *CostunitsGetRequestPathParams {
	return &CostunitsGetRequestPathParams{}
}

type CostunitsGetRequestPathParams struct {
}

func (p *CostunitsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *CostunitsGetRequest) PathParams() *CostunitsGetRequestPathParams {
	return r.pathParams
}

func (r *CostunitsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CostunitsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *CostunitsGetRequest) Method() string {
	return r.method
}

func (r CostunitsGetRequest) NewRequestBody() CostunitsGetRequestBody {
	return CostunitsGetRequestBody{}
}

type CostunitsGetRequestBody struct {
}

func (r *CostunitsGetRequest) RequestBody() *CostunitsGetRequestBody {
	return nil
}

func (r *CostunitsGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *CostunitsGetRequest) SetRequestBody(body CostunitsGetRequestBody) {
	r.requestBody = body
}

func (r *CostunitsGetRequest) NewResponseBody() *CostunitsGetResponseBody {
	return &CostunitsGetResponseBody{}
}

type CostunitsGetResponseBody struct {
	Data Costunits `json:"_data"`
	Meta Meta      `json:"_meta"`
}

func (r *CostunitsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/v1/costunits", r.PathParams())
	return &u
}

func (r *CostunitsGetRequest) Do() (CostunitsGetResponseBody, error) {
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
