package toast

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-toast/utils"
)

func (c *Client) NewMenuOptionGroupsGetRequest() MenuOptionGroupsGetRequest {
	r := MenuOptionGroupsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type MenuOptionGroupsGetRequest struct {
	client      *Client
	queryParams *MenuOptionGroupsGetRequestQueryParams
	pathParams  *MenuOptionGroupsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody MenuOptionGroupsGetRequestBody
}

func (r MenuOptionGroupsGetRequest) NewQueryParams() *MenuOptionGroupsGetRequestQueryParams {
	return &MenuOptionGroupsGetRequestQueryParams{}
}

type MenuOptionGroupsGetRequestQueryParams struct {
	PageToken    string   `schema:"pageToken,omitempty"`
	LastModified DateTime `schema:"lastModified,omitempty"`
}

func (p MenuOptionGroupsGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *MenuOptionGroupsGetRequest) QueryParams() *MenuOptionGroupsGetRequestQueryParams {
	return r.queryParams
}

func (r MenuOptionGroupsGetRequest) NewPathParams() *MenuOptionGroupsGetRequestPathParams {
	return &MenuOptionGroupsGetRequestPathParams{}
}

type MenuOptionGroupsGetRequestPathParams struct {
}

func (p *MenuOptionGroupsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *MenuOptionGroupsGetRequest) PathParams() *MenuOptionGroupsGetRequestPathParams {
	return r.pathParams
}

func (r *MenuOptionGroupsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *MenuOptionGroupsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *MenuOptionGroupsGetRequest) Method() string {
	return r.method
}

func (r MenuOptionGroupsGetRequest) NewRequestBody() MenuOptionGroupsGetRequestBody {
	return MenuOptionGroupsGetRequestBody{}
}

type MenuOptionGroupsGetRequestBody struct {
}

func (r *MenuOptionGroupsGetRequest) RequestBody() *MenuOptionGroupsGetRequestBody {
	return nil
}

func (r *MenuOptionGroupsGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *MenuOptionGroupsGetRequest) SetRequestBody(body MenuOptionGroupsGetRequestBody) {
	r.requestBody = body
}

func (r *MenuOptionGroupsGetRequest) NewResponseBody() *MenuOptionGroupsGetResponseBody {
	return &MenuOptionGroupsGetResponseBody{}
}

type MenuOptionGroupsGetResponseBody MenuOptionGroups

func (r *MenuOptionGroupsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/config/v2/menuOptionGroups", r.PathParams())
	return &u
}

func (r *MenuOptionGroupsGetRequest) Do() (MenuOptionGroupsGetResponseBody, error, *http.Response) {
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

func (r *MenuOptionGroupsGetRequest) All() (MenuOptionGroupsGetResponseBody, error) {
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
