package asperion

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-asperion/utils"
)

func (c *Client) NewJournalsGetRequest() JournalsGetRequest {
	r := JournalsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type JournalsGetRequest struct {
	client      *Client
	queryParams *JournalsGetRequestQueryParams
	pathParams  *JournalsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody JournalsGetRequestBody
}

func (r JournalsGetRequest) NewQueryParams() *JournalsGetRequestQueryParams {
	return &JournalsGetRequestQueryParams{}
}

type JournalsGetRequestQueryParams struct {
	Page           int    `schema:"page,omitempty"`
	PageSize       int    `schema:"pageSize,omitempty"`
	Fields         string `schema:"fields,omitempty"`
	OrderBy        string `schema:"orderBy,omitempty"`
	MetaOnly       bool   `schema:"metaOnly,omitempty"`
	ID             int    `schema:"Id,omitempty"`
	Description    string `schema:"description,omitempty"`
	ConversionCode string `schema:"conversionCode,omitempty"`
}

func (p JournalsGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *JournalsGetRequest) QueryParams() *JournalsGetRequestQueryParams {
	return r.queryParams
}

func (r JournalsGetRequest) NewPathParams() *JournalsGetRequestPathParams {
	return &JournalsGetRequestPathParams{}
}

type JournalsGetRequestPathParams struct {
}

func (p *JournalsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *JournalsGetRequest) PathParams() *JournalsGetRequestPathParams {
	return r.pathParams
}

func (r *JournalsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *JournalsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *JournalsGetRequest) Method() string {
	return r.method
}

func (r JournalsGetRequest) NewRequestBody() JournalsGetRequestBody {
	return JournalsGetRequestBody{}
}

type JournalsGetRequestBody struct {
}

func (r *JournalsGetRequest) RequestBody() *JournalsGetRequestBody {
	return nil
}

func (r *JournalsGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *JournalsGetRequest) SetRequestBody(body JournalsGetRequestBody) {
	r.requestBody = body
}

func (r *JournalsGetRequest) NewResponseBody() *JournalsGetResponseBody {
	return &JournalsGetResponseBody{}
}

type JournalsGetResponseBody struct {
	Data Journals `json:"_data"`
	Meta Meta     `json:"_meta"`
}

func (r *JournalsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/v1/journals", r.PathParams())
	return &u
}

func (r *JournalsGetRequest) Do() (JournalsGetResponseBody, error) {
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
