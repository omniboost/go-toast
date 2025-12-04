package toast

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/omniboost/go-toast/utils"
)

func (c *Client) NewPartnerRestaurantsGetRequest() PartnerRestaurantsGetRequest {
	r := PartnerRestaurantsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type PartnerRestaurantsGetRequest struct {
	client      *Client
	queryParams *PartnerRestaurantsGetRequestQueryParams
	pathParams  *PartnerRestaurantsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody PartnerRestaurantsGetRequestBody
}

func (r PartnerRestaurantsGetRequest) NewQueryParams() *PartnerRestaurantsGetRequestQueryParams {
	return &PartnerRestaurantsGetRequestQueryParams{}
}

type PartnerRestaurantsGetRequestQueryParams struct {
	StartDate    DateTime `schema:"startDate,omitempty"`
	EndDate      DateTime `schema:"endDate,omitempty"`
	PageSize     int      `schema:"pageSize,omitempty"`
	Page         int      `schema:"page,omitempty"`
	BusinessDate Date     `schema:"businessDate,omitempty"`
}

func (p PartnerRestaurantsGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *PartnerRestaurantsGetRequest) QueryParams() *PartnerRestaurantsGetRequestQueryParams {
	return r.queryParams
}

func (r PartnerRestaurantsGetRequest) NewPathParams() *PartnerRestaurantsGetRequestPathParams {
	return &PartnerRestaurantsGetRequestPathParams{}
}

type PartnerRestaurantsGetRequestPathParams struct{}

func (p *PartnerRestaurantsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *PartnerRestaurantsGetRequest) PathParams() *PartnerRestaurantsGetRequestPathParams {
	return r.pathParams
}

func (r *PartnerRestaurantsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *PartnerRestaurantsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *PartnerRestaurantsGetRequest) Method() string {
	return r.method
}

func (r PartnerRestaurantsGetRequest) NewRequestBody() PartnerRestaurantsGetRequestBody {
	return PartnerRestaurantsGetRequestBody{}
}

type PartnerRestaurantsGetRequestBody struct {
}

func (r *PartnerRestaurantsGetRequest) RequestBody() *PartnerRestaurantsGetRequestBody {
	return nil
}

func (r *PartnerRestaurantsGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *PartnerRestaurantsGetRequest) SetRequestBody(body PartnerRestaurantsGetRequestBody) {
	r.requestBody = body
}

func (r *PartnerRestaurantsGetRequest) NewResponseBody() *PartnerRestaurantsGetResponseBody {
	return &PartnerRestaurantsGetResponseBody{}
}

type PartnerRestaurantsGetResponseBody []struct {
	RestaurantGUID        string      `json:"restaurantGuid"`
	ManagementGroupGUID   string      `json:"managementGroupGuid"`
	RestaurantName        string      `json:"restaurantName"`
	LocationName          string      `json:"locationName"`
	CreatedByEmailAddress string      `json:"createdByEmailAddress"`
	ExternalGroupRef      interface{} `json:"externalGroupRef"`
	ExternalRestaurantRef interface{} `json:"externalRestaurantRef"`
	ModifiedDate          int64       `json:"modifiedDate"`
	CreatedDate           int64       `json:"createdDate"`
	IsoModifiedDate       time.Time   `json:"isoModifiedDate"`
	IsoCreatedDate        time.Time   `json:"isoCreatedDate"`
}

func (r *PartnerRestaurantsGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/partners/v1/restaurants", r.PathParams())
	return &u
}

func (r *PartnerRestaurantsGetRequest) Do(ctx context.Context) (PartnerRestaurantsGetResponseBody, error, *http.Response) {
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
