package toast

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-toast/utils"
)

func (c *Client) NewTimeEntriesGetRequest() TimeEntriesGetRequest {
	r := TimeEntriesGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type TimeEntriesGetRequest struct {
	client      *Client
	queryParams *TimeEntriesGetRequestQueryParams
	pathParams  *TimeEntriesGetRequestPathParams
	method      string
	headers     http.Header
	requestBody TimeEntriesGetRequestBody
}

func (r TimeEntriesGetRequest) NewQueryParams() *TimeEntriesGetRequestQueryParams {
	return &TimeEntriesGetRequestQueryParams{}
}

type TimeEntriesGetRequestQueryParams struct {
	BusinessDate        Date     `schema:"businessDate,omitempty"`
	EndDate             DateTime `schema:"endDate,omitempty"`
	IncludeArchived     bool     `schema:"includeArchived,omitempty"`
	IncludeMissedBreaks bool     `schema:"includeMissedBreaks,omitempty"`
	ModifiedEndDate     DateTime `schema:"modifiedEndDate,omitempty"`
	ModifiedStartDate   DateTime `schema:"modifiedStartDate,omitempty"`
	StartDate           DateTime `schema:"startDate,omitempty"`
	TimeEntryIds        []string `schema:"timeEntryIds,omitempty"`
}

func (p TimeEntriesGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *TimeEntriesGetRequest) QueryParams() *TimeEntriesGetRequestQueryParams {
	return r.queryParams
}

func (r TimeEntriesGetRequest) NewPathParams() *TimeEntriesGetRequestPathParams {
	return &TimeEntriesGetRequestPathParams{}
}

type TimeEntriesGetRequestPathParams struct {
}

func (p *TimeEntriesGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *TimeEntriesGetRequest) PathParams() *TimeEntriesGetRequestPathParams {
	return r.pathParams
}

func (r *TimeEntriesGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *TimeEntriesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *TimeEntriesGetRequest) Method() string {
	return r.method
}

func (r TimeEntriesGetRequest) NewRequestBody() TimeEntriesGetRequestBody {
	return TimeEntriesGetRequestBody{}
}

type TimeEntriesGetRequestBody struct {
}

func (r *TimeEntriesGetRequest) RequestBody() *TimeEntriesGetRequestBody {
	return nil
}

func (r *TimeEntriesGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *TimeEntriesGetRequest) SetRequestBody(body TimeEntriesGetRequestBody) {
	r.requestBody = body
}

func (r *TimeEntriesGetRequest) NewResponseBody() *TimeEntriesGetResponseBody {
	return &TimeEntriesGetResponseBody{}
}

type TimeEntriesGetResponseBody TimeEntries

func (r *TimeEntriesGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/labor/v1/timeEntries/", r.PathParams())
	return &u
}

func (r *TimeEntriesGetRequest) Do() (TimeEntriesGetResponseBody, error, *http.Response) {
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

func (r *TimeEntriesGetRequest) All() (TimeEntriesGetResponseBody, error) {
	body, err, _ := r.Do()
	if err != nil {
		return body, err
	}

	return body, nil
}
