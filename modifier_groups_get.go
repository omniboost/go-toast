package toast

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-toast/utils"
)

func (c *Client) NewModifierGroupsGetRequest() ModifierGroupsGetRequest {
	r := ModifierGroupsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type ModifierGroupsGetRequest struct {
	client      *Client
	queryParams *ModifierGroupsGetRequestQueryParams
	pathParams  *ModifierGroupsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody ModifierGroupsGetRequestBody
}

func (r ModifierGroupsGetRequest) NewQueryParams() *ModifierGroupsGetRequestQueryParams {
	return &ModifierGroupsGetRequestQueryParams{}
}

type ModifierGroupsGetRequestQueryParams struct {
	StartDate    DateTime `schema:"startDate,omitempty"`
	EndDate      DateTime `schema:"endDate,omitempty"`
	PageSize     int      `schema:"pageSize,omitempty"`
	Page         int      `schema:"page,omitempty"`
	BusinessDate Date     `schema:"businessDate,omitempty"`
}

func (p ModifierGroupsGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *ModifierGroupsGetRequest) QueryParams() *ModifierGroupsGetRequestQueryParams {
	return r.queryParams
}

func (r ModifierGroupsGetRequest) NewPathParams() *ModifierGroupsGetRequestPathParams {
	return &ModifierGroupsGetRequestPathParams{}
}

type ModifierGroupsGetRequestPathParams struct {
}

func (p *ModifierGroupsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *ModifierGroupsGetRequest) PathParams() *ModifierGroupsGetRequestPathParams {
	return r.pathParams
}

func (r *ModifierGroupsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *ModifierGroupsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *ModifierGroupsGetRequest) Method() string {
	return r.method
}

func (r ModifierGroupsGetRequest) NewRequestBody() ModifierGroupsGetRequestBody {
	return ModifierGroupsGetRequestBody{}
}

type ModifierGroupsGetRequestBody struct {
}

func (r *ModifierGroupsGetRequest) RequestBody() *ModifierGroupsGetRequestBody {
	return nil
}

func (r *ModifierGroupsGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *ModifierGroupsGetRequest) SetRequestBody(body ModifierGroupsGetRequestBody) {
	r.requestBody = body
}

func (r *ModifierGroupsGetRequest) NewResponseBody() *ModifierGroupsGetResponseBody {
	return &ModifierGroupsGetResponseBody{}
}

type ModifierGroupsGetResponseBody ModifierGroups

func (r *ModifierGroupsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/config/v2/menuOptionGroups", r.PathParams())
	return &u
}

func (r *ModifierGroupsGetRequest) Do() (ModifierGroupsGetResponseBody, error, *http.Response) {
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

func (r *ModifierGroupsGetRequest) All() (ModifierGroupsGetResponseBody, error) {
	body, err, resp := r.Do()
	if err != nil {
		return body, err
	}

	concat := body
	next, err := r.client.GetNextPage(resp)
	if err != nil {
		return concat, err
	}

	for next != 0 {
		r.QueryParams().Page = next
		body, err, resp = r.Do()
		if err != nil {
			return concat, err
		}

		concat = append(concat, body...)
		next, err = r.client.GetNextPage(resp)
		if err != nil {
			return concat, err
		}
	}

	return concat, nil
}
