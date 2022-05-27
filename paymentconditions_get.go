package asperion

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-asperion/utils"
)

func (c *Client) NewPaymentConditionsGetRequest() PaymentConditionsGetRequest {
	r := PaymentConditionsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type PaymentConditionsGetRequest struct {
	client      *Client
	queryParams *PaymentConditionsGetRequestQueryParams
	pathParams  *PaymentConditionsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody PaymentConditionsGetRequestBody
}

func (r PaymentConditionsGetRequest) NewQueryParams() *PaymentConditionsGetRequestQueryParams {
	return &PaymentConditionsGetRequestQueryParams{}
}

type PaymentConditionsGetRequestQueryParams struct {
	Page        int    `schema:"page,omitempty"`
	PageSize    int    `schema:"pageSize,omitempty"`
	Fields      string `schema:"fields,omitempty"`
	OrderBy     string `schema:"orderBy,omitempty"`
	MetaOnly    bool   `schema:"metaOnly,omitempty"`
	ID          int    `schema:"Id,omitempty"`
	Description string `schema:"Description,omitempty"`
}

func (p PaymentConditionsGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PaymentConditionsGetRequest) QueryParams() *PaymentConditionsGetRequestQueryParams {
	return r.queryParams
}

func (r PaymentConditionsGetRequest) NewPathParams() *PaymentConditionsGetRequestPathParams {
	return &PaymentConditionsGetRequestPathParams{}
}

type PaymentConditionsGetRequestPathParams struct {
}

func (p *PaymentConditionsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PaymentConditionsGetRequest) PathParams() *PaymentConditionsGetRequestPathParams {
	return r.pathParams
}

func (r *PaymentConditionsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PaymentConditionsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *PaymentConditionsGetRequest) Method() string {
	return r.method
}

func (r PaymentConditionsGetRequest) NewRequestBody() PaymentConditionsGetRequestBody {
	return PaymentConditionsGetRequestBody{}
}

type PaymentConditionsGetRequestBody struct {
}

func (r *PaymentConditionsGetRequest) RequestBody() *PaymentConditionsGetRequestBody {
	return nil
}

func (r *PaymentConditionsGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *PaymentConditionsGetRequest) SetRequestBody(body PaymentConditionsGetRequestBody) {
	r.requestBody = body
}

func (r *PaymentConditionsGetRequest) NewResponseBody() *PaymentConditionsGetResponseBody {
	return &PaymentConditionsGetResponseBody{}
}

type PaymentConditionsGetResponseBody struct {
	Data PaymentConditions `json:"_data"`
	Meta Meta              `json:"_meta"`
}

func (r *PaymentConditionsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/v1/paymentconditions", r.PathParams())
	return &u
}

func (r *PaymentConditionsGetRequest) Do() (PaymentConditionsGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
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
