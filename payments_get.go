package toast

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-toast/utils"
)

func (c *Client) NewPaymentsGetRequest() PaymentsGetRequest {
	r := PaymentsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type PaymentsGetRequest struct {
	client      *Client
	queryParams *PaymentsGetRequestQueryParams
	pathParams  *PaymentsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody PaymentsGetRequestBody
}

func (r PaymentsGetRequest) NewQueryParams() *PaymentsGetRequestQueryParams {
	return &PaymentsGetRequestQueryParams{}
}

type PaymentsGetRequestQueryParams struct {
	PaidBusinessDate   Date `schema:"paidBusinessDate,omitempty"`
	RefundBusinessDate Date `schema:"refundBusinessDate,omitempty"`
	VoidBusinessDate   Date `schema:"voidBusinessDate,omitempty"`
	PageSize           int  `schema:"pageSize,omitempty"`
	Page               int  `schema:"page,omitempty"`
}

func (p PaymentsGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PaymentsGetRequest) QueryParams() *PaymentsGetRequestQueryParams {
	return r.queryParams
}

func (r PaymentsGetRequest) NewPathParams() *PaymentsGetRequestPathParams {
	return &PaymentsGetRequestPathParams{}
}

type PaymentsGetRequestPathParams struct {
}

func (p *PaymentsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PaymentsGetRequest) PathParams() *PaymentsGetRequestPathParams {
	return r.pathParams
}

func (r *PaymentsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PaymentsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *PaymentsGetRequest) Method() string {
	return r.method
}

func (r PaymentsGetRequest) NewRequestBody() PaymentsGetRequestBody {
	return PaymentsGetRequestBody{}
}

type PaymentsGetRequestBody struct {
}

func (r *PaymentsGetRequest) RequestBody() *PaymentsGetRequestBody {
	return nil
}

func (r *PaymentsGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *PaymentsGetRequest) SetRequestBody(body PaymentsGetRequestBody) {
	r.requestBody = body
}

func (r *PaymentsGetRequest) NewResponseBody() *PaymentsGetResponseBody {
	return &PaymentsGetResponseBody{}
}

type PaymentsGetResponseBody []string

func (r *PaymentsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/orders/v2/payments", r.PathParams())
	return &u
}

func (r *PaymentsGetRequest) Do() (PaymentsGetResponseBody, error, *http.Response) {
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

func (r *PaymentsGetRequest) All() (PaymentsGetResponseBody, error) {
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
