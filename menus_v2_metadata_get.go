package toast

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-toast/utils"
)

func (c *Client) NewMenuV2MetadataGetRequest() MenuV2MetadataGetRequest {
	r := MenuV2MetadataGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type MenuV2MetadataGetRequest struct {
	client      *Client
	queryParams *MenuV2MetadataGetRequestQueryParams
	pathParams  *MenuV2MetadataGetRequestPathParams
	method      string
	headers     http.Header
	requestBody MenuV2MetadataGetRequestBody
}

func (r MenuV2MetadataGetRequest) NewQueryParams() *MenuV2MetadataGetRequestQueryParams {
	return &MenuV2MetadataGetRequestQueryParams{}
}

type MenuV2MetadataGetRequestQueryParams struct {
}

func (p MenuV2MetadataGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *MenuV2MetadataGetRequest) QueryParams() *MenuV2MetadataGetRequestQueryParams {
	return r.queryParams
}

func (r MenuV2MetadataGetRequest) NewPathParams() *MenuV2MetadataGetRequestPathParams {
	return &MenuV2MetadataGetRequestPathParams{}
}

type MenuV2MetadataGetRequestPathParams struct {
}

func (p *MenuV2MetadataGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *MenuV2MetadataGetRequest) PathParams() *MenuV2MetadataGetRequestPathParams {
	return r.pathParams
}

func (r *MenuV2MetadataGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *MenuV2MetadataGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *MenuV2MetadataGetRequest) Method() string {
	return r.method
}

func (r MenuV2MetadataGetRequest) NewRequestBody() MenuV2MetadataGetRequestBody {
	return MenuV2MetadataGetRequestBody{}
}

type MenuV2MetadataGetRequestBody struct {
}

func (r *MenuV2MetadataGetRequest) RequestBody() *MenuV2MetadataGetRequestBody {
	return nil
}

func (r *MenuV2MetadataGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *MenuV2MetadataGetRequest) SetRequestBody(body MenuV2MetadataGetRequestBody) {
	r.requestBody = body
}

func (r *MenuV2MetadataGetRequest) NewResponseBody() *MenuV2MetadataGetResponseBody {
	return &MenuV2MetadataGetResponseBody{}
}

type MenuV2MetadataGetResponseBody struct {
	RestaurantGUID string `json:"restaurantGuid"`
	LastUpdated    string `json:"lastUpdated"`
}

func (r *MenuV2MetadataGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/menus/v2/metadata", r.PathParams())
	return &u
}

func (r *MenuV2MetadataGetRequest) Do(ctx context.Context) (MenuV2MetadataGetResponseBody, error, *http.Response) {
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
