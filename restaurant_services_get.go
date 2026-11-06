package toast

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-toast/utils"
)

func (c *Client) NewRestaurantServicesGetRequest() RestaurantServicesGetRequest {
	r := RestaurantServicesGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type RestaurantServicesGetRequest struct {
	client      *Client
	queryParams *RestaurantServicesGetRequestQueryParams
	pathParams  *RestaurantServicesGetRequestPathParams
	method      string
	headers     http.Header
	requestBody RestaurantServicesGetRequestBody
}

func (r RestaurantServicesGetRequest) NewQueryParams() *RestaurantServicesGetRequestQueryParams {
	return &RestaurantServicesGetRequestQueryParams{}
}

type RestaurantServicesGetRequestQueryParams struct {
	PageToken    string   `schema:"pageToken,omitempty"`
	LastModified DateTime `schema:"lastModified,omitempty"`
}

func (p RestaurantServicesGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *RestaurantServicesGetRequest) QueryParams() *RestaurantServicesGetRequestQueryParams {
	return r.queryParams
}

func (r RestaurantServicesGetRequest) NewPathParams() *RestaurantServicesGetRequestPathParams {
	return &RestaurantServicesGetRequestPathParams{}
}

type RestaurantServicesGetRequestPathParams struct {
}

func (p *RestaurantServicesGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *RestaurantServicesGetRequest) PathParams() *RestaurantServicesGetRequestPathParams {
	return r.pathParams
}

func (r *RestaurantServicesGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *RestaurantServicesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *RestaurantServicesGetRequest) Method() string {
	return r.method
}

func (r RestaurantServicesGetRequest) NewRequestBody() RestaurantServicesGetRequestBody {
	return RestaurantServicesGetRequestBody{}
}

type RestaurantServicesGetRequestBody struct {
}

func (r *RestaurantServicesGetRequest) RequestBody() *RestaurantServicesGetRequestBody {
	return nil
}

func (r *RestaurantServicesGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *RestaurantServicesGetRequest) SetRequestBody(body RestaurantServicesGetRequestBody) {
	r.requestBody = body
}

func (r *RestaurantServicesGetRequest) NewResponseBody() *RestaurantServicesGetResponseBody {
	return &RestaurantServicesGetResponseBody{}
}

type RestaurantServicesGetResponseBody RestaurantServices

func (r *RestaurantServicesGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/config/v2/restaurantServices/", r.PathParams())
	return &u
}

func (r *RestaurantServicesGetRequest) Do(ctx context.Context) (RestaurantServicesGetResponseBody, error, *http.Response) {
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

func (r *RestaurantServicesGetRequest) All(ctx context.Context) (RestaurantServicesGetResponseBody, error) {
	body, err, resp := r.Do(ctx)
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
		body, err, resp = r.Do(ctx)
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
