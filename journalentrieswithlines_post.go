package asperion

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-asperion/utils"
)

func (c *Client) NewJournalEntriesWithLinesPostRequest() JournalEntriesWithLinesPostRequest {
	r := JournalEntriesWithLinesPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type JournalEntriesWithLinesPostRequest struct {
	client      *Client
	queryParams *JournalEntriesWithLinesPostRequestQueryParams
	pathParams  *JournalEntriesWithLinesPostRequestPathParams
	method      string
	headers     http.Header
	requestBody JournalEntriesWithLinesPostRequestBody
}

func (r JournalEntriesWithLinesPostRequest) NewQueryParams() *JournalEntriesWithLinesPostRequestQueryParams {
	return &JournalEntriesWithLinesPostRequestQueryParams{}
}

type JournalEntriesWithLinesPostRequestQueryParams struct {
	ValidateOnly bool `schema:"validateOnly,omitempty"`
}

func (p JournalEntriesWithLinesPostRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *JournalEntriesWithLinesPostRequest) QueryParams() *JournalEntriesWithLinesPostRequestQueryParams {
	return r.queryParams
}

func (r JournalEntriesWithLinesPostRequest) NewPathParams() *JournalEntriesWithLinesPostRequestPathParams {
	return &JournalEntriesWithLinesPostRequestPathParams{}
}

type JournalEntriesWithLinesPostRequestPathParams struct {
}

func (p *JournalEntriesWithLinesPostRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *JournalEntriesWithLinesPostRequest) PathParams() *JournalEntriesWithLinesPostRequestPathParams {
	return r.pathParams
}

func (r *JournalEntriesWithLinesPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *JournalEntriesWithLinesPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *JournalEntriesWithLinesPostRequest) Method() string {
	return r.method
}

func (r JournalEntriesWithLinesPostRequest) NewRequestBody() JournalEntriesWithLinesPostRequestBody {
	return JournalEntriesWithLinesPostRequestBody{}
}

type JournalEntriesWithLinesPostRequestBody JournalEntryWithLines

func (r *JournalEntriesWithLinesPostRequest) RequestBody() *JournalEntriesWithLinesPostRequestBody {
	return &r.requestBody
}

func (r *JournalEntriesWithLinesPostRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *JournalEntriesWithLinesPostRequest) SetRequestBody(body JournalEntriesWithLinesPostRequestBody) {
	r.requestBody = body
}

func (r *JournalEntriesWithLinesPostRequest) NewResponseBody() *JournalEntriesWithLinesPostResponseBody {
	return &JournalEntriesWithLinesPostResponseBody{}
}

type JournalEntriesWithLinesPostResponseBody JournalEntryWithLines

func (r *JournalEntriesWithLinesPostRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/v1/journalentrieswithlines", r.PathParams())
	return &u
}

func (r *JournalEntriesWithLinesPostRequest) Do() (JournalEntriesWithLinesPostResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	req.Header.Add("X-Tenant-Id", strconv.Itoa(r.client.TenantID()))

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
