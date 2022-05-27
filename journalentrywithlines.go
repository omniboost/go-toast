package asperion

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-asperion/utils"
)

func (c *Client) NewJournalEntryWithLinesGetRequest() JournalEntryWithLinesGetRequest {
	r := JournalEntryWithLinesGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type JournalEntryWithLinesGetRequest struct {
	client      *Client
	queryParams *JournalEntryWithLinesGetRequestQueryParams
	pathParams  *JournalEntryWithLinesGetRequestPathParams
	method      string
	headers     http.Header
	requestBody JournalEntryWithLinesGetRequestBody
}

func (r JournalEntryWithLinesGetRequest) NewQueryParams() *JournalEntryWithLinesGetRequestQueryParams {
	return &JournalEntryWithLinesGetRequestQueryParams{}
}

type JournalEntryWithLinesGetRequestQueryParams struct {
	Page            int    `schema:"page,omitempty"`
	PageSize        int    `schema:"pageSize,omitempty"`
	Fields          string `schema:"fields,omitempty"`
	OrderBy         string `schema:"orderBy,omitempty"`
	MetaOnly        bool   `schema:"metaOnly,omitempty"`
	ID              int    `schema:"Id,omitempty"`
	Name            string `schema:"Name,omitempty"`
	TelephoneNumber string `schema:"TelephoneNumber,omitempty"`
	Email           string `schema:"Email,omitempty"`
}

func (p JournalEntryWithLinesGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *JournalEntryWithLinesGetRequest) QueryParams() *JournalEntryWithLinesGetRequestQueryParams {
	return r.queryParams
}

func (r JournalEntryWithLinesGetRequest) NewPathParams() *JournalEntryWithLinesGetRequestPathParams {
	return &JournalEntryWithLinesGetRequestPathParams{}
}

type JournalEntryWithLinesGetRequestPathParams struct {
	ID int `schema:"id"`
}

func (p *JournalEntryWithLinesGetRequestPathParams) Params() map[string]string {
	return map[string]string{
		"id": strconv.Itoa(p.ID),
	}
}

func (r *JournalEntryWithLinesGetRequest) PathParams() *JournalEntryWithLinesGetRequestPathParams {
	return r.pathParams
}

func (r *JournalEntryWithLinesGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *JournalEntryWithLinesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *JournalEntryWithLinesGetRequest) Method() string {
	return r.method
}

func (r JournalEntryWithLinesGetRequest) NewRequestBody() JournalEntryWithLinesGetRequestBody {
	return JournalEntryWithLinesGetRequestBody{}
}

type JournalEntryWithLinesGetRequestBody struct {
}

func (r *JournalEntryWithLinesGetRequest) RequestBody() *JournalEntryWithLinesGetRequestBody {
	return nil
}

func (r *JournalEntryWithLinesGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *JournalEntryWithLinesGetRequest) SetRequestBody(body JournalEntryWithLinesGetRequestBody) {
	r.requestBody = body
}

func (r *JournalEntryWithLinesGetRequest) NewResponseBody() *JournalEntryWithLinesGetResponseBody {
	return &JournalEntryWithLinesGetResponseBody{}
}

type JournalEntryWithLinesGetResponseBody JournalEntryWithLines

func (r *JournalEntryWithLinesGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/v1/journalentrieswithlines/{{.id}}", r.PathParams())
	return &u
}

func (r *JournalEntryWithLinesGetRequest) Do() (JournalEntryWithLinesGetResponseBody, error) {
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
