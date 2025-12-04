package toast

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-toast/utils"
)

func (c *Client) NewPaymentGetRequest() PaymentGetRequest {
	r := PaymentGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type PaymentGetRequest struct {
	client      *Client
	queryParams *PaymentGetRequestQueryParams
	pathParams  *PaymentGetRequestPathParams
	method      string
	headers     http.Header
	requestBody PaymentGetRequestBody
}

func (r PaymentGetRequest) NewQueryParams() *PaymentGetRequestQueryParams {
	return &PaymentGetRequestQueryParams{}
}

type PaymentGetRequestQueryParams struct {}

func (p PaymentGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PaymentGetRequest) QueryParams() *PaymentGetRequestQueryParams {
	return r.queryParams
}

func (r PaymentGetRequest) NewPathParams() *PaymentGetRequestPathParams {
	return &PaymentGetRequestPathParams{}
}

type PaymentGetRequestPathParams struct {
	GUID string `schema:"GUID"`
}

func (p *PaymentGetRequestPathParams) Params() map[string]string {
	return map[string]string{
		"GUID": p.GUID,
	}
}

func (r *PaymentGetRequest) PathParams() *PaymentGetRequestPathParams {
	return r.pathParams
}

func (r *PaymentGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PaymentGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *PaymentGetRequest) Method() string {
	return r.method
}

func (r PaymentGetRequest) NewRequestBody() PaymentGetRequestBody {
	return PaymentGetRequestBody{}
}

type PaymentGetRequestBody struct {
}

func (r *PaymentGetRequest) RequestBody() *PaymentGetRequestBody {
	return nil
}

func (r *PaymentGetRequest) RequestBodyInterface() any {
	return nil
}

func (r *PaymentGetRequest) SetRequestBody(body PaymentGetRequestBody) {
	r.requestBody = body
}

func (r *PaymentGetRequest) NewResponseBody() *PaymentGetResponseBody {
	return &PaymentGetResponseBody{}
}

type PaymentGetResponseBody Payment

func (r *PaymentGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/orders/v2/payments/{{.GUID}}", r.PathParams())
	return &u
}

func (r *PaymentGetRequest) Do() (PaymentGetResponseBody, error, *http.Response) {
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


