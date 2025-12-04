package toast

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-toast/utils"
)

func (c *Client) NewStockInventoryGetRequest() StockInventoryGetRequest {
	r := StockInventoryGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type StockInventoryGetRequest struct {
	client      *Client
	queryParams *StockInventoryGetRequestQueryParams
	pathParams  *StockInventoryGetRequestPathParams
	method      string
	headers     http.Header
	requestBody StockInventoryGetRequestBody
}

func (r StockInventoryGetRequest) NewQueryParams() *StockInventoryGetRequestQueryParams {
	return &StockInventoryGetRequestQueryParams{}
}

type StockInventoryGetRequestQueryParams struct {
	// PageSize     int      `schema:"pageSize,omitempty"`
	// Page         int      `schema:"page,omitempty"`
	Status string `schema:"status,omitempty"`
}

func (p StockInventoryGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *StockInventoryGetRequest) QueryParams() *StockInventoryGetRequestQueryParams {
	return r.queryParams
}

func (r StockInventoryGetRequest) NewPathParams() *StockInventoryGetRequestPathParams {
	return &StockInventoryGetRequestPathParams{}
}

type StockInventoryGetRequestPathParams struct {
}

func (p *StockInventoryGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *StockInventoryGetRequest) PathParams() *StockInventoryGetRequestPathParams {
	return r.pathParams
}

func (r *StockInventoryGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *StockInventoryGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *StockInventoryGetRequest) Method() string {
	return r.method
}

func (r StockInventoryGetRequest) NewRequestBody() StockInventoryGetRequestBody {
	return StockInventoryGetRequestBody{}
}

type StockInventoryGetRequestBody struct {
}

func (r *StockInventoryGetRequest) RequestBody() *StockInventoryGetRequestBody {
	return nil
}

func (r *StockInventoryGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *StockInventoryGetRequest) SetRequestBody(body StockInventoryGetRequestBody) {
	r.requestBody = body
}

func (r *StockInventoryGetRequest) NewResponseBody() *StockInventoryGetResponseBody {
	return &StockInventoryGetResponseBody{}
}

type StockInventoryGetResponseBody MenuItemInventories

func (r *StockInventoryGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/stock/v1/inventory", r.PathParams())
	return &u
}

func (r *StockInventoryGetRequest) Do(ctx context.Context) (StockInventoryGetResponseBody, error, *http.Response) {
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
