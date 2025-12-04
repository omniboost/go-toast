package toast

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-toast/utils"
)

func (c *Client) NewVoidSaleReasonsGetRequest() VoidSaleReasonsGetRequest {
	r := VoidSaleReasonsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type VoidSaleReasonsGetRequest struct {
	client      *Client
	queryParams *VoidSaleReasonsGetRequestQueryParams
	pathParams  *VoidSaleReasonsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody VoidSaleReasonsGetRequestBody
}

func (r VoidSaleReasonsGetRequest) NewQueryParams() *VoidSaleReasonsGetRequestQueryParams {
	return &VoidSaleReasonsGetRequestQueryParams{}
}

type VoidSaleReasonsGetRequestQueryParams struct{}

func (p VoidSaleReasonsGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *VoidSaleReasonsGetRequest) QueryParams() *VoidSaleReasonsGetRequestQueryParams {
	return r.queryParams
}

func (r VoidSaleReasonsGetRequest) NewPathParams() *VoidSaleReasonsGetRequestPathParams {
	return &VoidSaleReasonsGetRequestPathParams{}
}

type VoidSaleReasonsGetRequestPathParams struct {
}

func (p *VoidSaleReasonsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *VoidSaleReasonsGetRequest) PathParams() *VoidSaleReasonsGetRequestPathParams {
	return r.pathParams
}

func (r *VoidSaleReasonsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *VoidSaleReasonsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *VoidSaleReasonsGetRequest) Method() string {
	return r.method
}

func (r VoidSaleReasonsGetRequest) NewRequestBody() VoidSaleReasonsGetRequestBody {
	return VoidSaleReasonsGetRequestBody{}
}

type VoidSaleReasonsGetRequestBody struct {
}

func (r *VoidSaleReasonsGetRequest) RequestBody() *VoidSaleReasonsGetRequestBody {
	return nil
}

func (r *VoidSaleReasonsGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *VoidSaleReasonsGetRequest) SetRequestBody(body VoidSaleReasonsGetRequestBody) {
	r.requestBody = body
}

func (r *VoidSaleReasonsGetRequest) NewResponseBody() *VoidSaleReasonsGetResponseBody {
	return &VoidSaleReasonsGetResponseBody{}
}

type VoidSaleReasonsGetResponseBody VoidSaleReasons

func (r *VoidSaleReasonsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/config/v2/voidReasons", r.PathParams())
	return &u
}

func (r *VoidSaleReasonsGetRequest) Do(ctx context.Context) (VoidSaleReasonsGetResponseBody, error, *http.Response) {
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

func (r *VoidSaleReasonsGetRequest) All(ctx context.Context) (VoidSaleReasonsGetResponseBody, error) {
	body, err, _ := r.Do(ctx)
	if err != nil {
		return body, err
	}

	return body, nil
}
