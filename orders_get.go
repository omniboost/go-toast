package toast

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-toast/utils"
)

func (c *Client) NewOrdersGetRequest() OrdersGetRequest {
	r := OrdersGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type OrdersGetRequest struct {
	client      *Client
	queryParams *OrdersGetRequestQueryParams
	pathParams  *OrdersGetRequestPathParams
	method      string
	headers     http.Header
	requestBody OrdersGetRequestBody
}

func (r OrdersGetRequest) NewQueryParams() *OrdersGetRequestQueryParams {
	return &OrdersGetRequestQueryParams{}
}

type OrdersGetRequestQueryParams struct {
}

func (p OrdersGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *OrdersGetRequest) QueryParams() *OrdersGetRequestQueryParams {
	return r.queryParams
}

func (r OrdersGetRequest) NewPathParams() *OrdersGetRequestPathParams {
	return &OrdersGetRequestPathParams{}
}

type OrdersGetRequestPathParams struct {
	OrderGUID string `schema:"order_guid"`
}

func (p *OrdersGetRequestPathParams) Params() map[string]string {
	return map[string]string{
		"order_guid": p.OrderGUID,
	}
}

func (r *OrdersGetRequest) PathParams() *OrdersGetRequestPathParams {
	return r.pathParams
}

func (r *OrdersGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *OrdersGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *OrdersGetRequest) Method() string {
	return r.method
}

func (r OrdersGetRequest) NewRequestBody() OrdersGetRequestBody {
	return OrdersGetRequestBody{}
}

type OrdersGetRequestBody struct {
}

func (r *OrdersGetRequest) RequestBody() *OrdersGetRequestBody {
	return nil
}

func (r *OrdersGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *OrdersGetRequest) SetRequestBody(body OrdersGetRequestBody) {
	r.requestBody = body
}

func (r *OrdersGetRequest) NewResponseBody() *OrdersGetResponseBody {
	return &OrdersGetResponseBody{}
}

type OrdersGetResponseBody Orders

func (r *OrdersGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/orders/v2/orders/{{.order_guid}}", r.PathParams())
	return &u
}

func (r *OrdersGetRequest) Do() (OrdersGetResponseBody, error, *http.Response) {
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
