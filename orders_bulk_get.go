package toast

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-toast/utils"
)

func (c *Client) NewOrdersBulkGetRequest() OrdersBulkGetRequest {
	r := OrdersBulkGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type OrdersBulkGetRequest struct {
	client      *Client
	queryParams *OrdersBulkGetRequestQueryParams
	pathParams  *OrdersBulkGetRequestPathParams
	method      string
	headers     http.Header
	requestBody OrdersBulkGetRequestBody
}

func (r OrdersBulkGetRequest) NewQueryParams() *OrdersBulkGetRequestQueryParams {
	return &OrdersBulkGetRequestQueryParams{}
}

type OrdersBulkGetRequestQueryParams struct {
	StartDate    DateTime `schema:"startDate,omitempty"`
	EndDate      DateTime `schema:"endDate,omitempty"`
	PageSize     int      `schema:"pageSize,omitempty"`
	Page         int      `schema:"page,omitempty"`
	BusinessDate Date     `schema:"businessDate,omitempty"`
}

func (p OrdersBulkGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *OrdersBulkGetRequest) QueryParams() *OrdersBulkGetRequestQueryParams {
	return r.queryParams
}

func (r OrdersBulkGetRequest) NewPathParams() *OrdersBulkGetRequestPathParams {
	return &OrdersBulkGetRequestPathParams{}
}

type OrdersBulkGetRequestPathParams struct {
}

func (p *OrdersBulkGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *OrdersBulkGetRequest) PathParams() *OrdersBulkGetRequestPathParams {
	return r.pathParams
}

func (r *OrdersBulkGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *OrdersBulkGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *OrdersBulkGetRequest) Method() string {
	return r.method
}

func (r OrdersBulkGetRequest) NewRequestBody() OrdersBulkGetRequestBody {
	return OrdersBulkGetRequestBody{}
}

type OrdersBulkGetRequestBody struct {
}

func (r *OrdersBulkGetRequest) RequestBody() *OrdersBulkGetRequestBody {
	return nil
}

func (r *OrdersBulkGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *OrdersBulkGetRequest) SetRequestBody(body OrdersBulkGetRequestBody) {
	r.requestBody = body
}

func (r *OrdersBulkGetRequest) NewResponseBody() *OrdersBulkGetResponseBody {
	return &OrdersBulkGetResponseBody{}
}

type OrdersBulkGetResponseBody Orders

func (r *OrdersBulkGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/orders/v2/ordersBulk", r.PathParams())
	return &u
}

func (r *OrdersBulkGetRequest) Do() (OrdersBulkGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	err = r.client.InitToken(req)
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
