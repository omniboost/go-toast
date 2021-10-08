package asperion

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-asperion/utils"
)

func (c *Client) NewJournalEntriesWithLinesGetRequest() JournalEntriesWithLinesGetRequest {
	r := JournalEntriesWithLinesGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type JournalEntriesWithLinesGetRequest struct {
	client      *Client
	queryParams *JournalEntriesWithLinesGetRequestQueryParams
	pathParams  *JournalEntriesWithLinesGetRequestPathParams
	method      string
	headers     http.Header
	requestBody JournalEntriesWithLinesGetRequestBody
}

func (r JournalEntriesWithLinesGetRequest) NewQueryParams() *JournalEntriesWithLinesGetRequestQueryParams {
	return &JournalEntriesWithLinesGetRequestQueryParams{}
}

type JournalEntriesWithLinesGetRequestQueryParams struct {
	Page     int    `schema:"page,omitempty"`
	PageSize int    `schema:"pageSize,omitempty"`
	Fields   string `schema:"fields,omitempty"`
	OrderBy  string `schema:"orderBy,omitempty"`
	MetaOnly bool   `schema:"metaOnly,omitempty"`

	ID                   int    `schema:"Id,omitempty"`
	BookPeriod           int    `schema:"BookPeriod,omitempty"`
	BookYear             int    `schema:"BookYear"`
	Date                 string `schema:"Date,omitempty"`
	JournalID            string `schema:"JournalId,omitempty"`
	Description          string `schema:"Description,omitempty"`
	SourceDocumentNumber string `schema:"SourceDocumentNumber,omitempty"`
	Status               string `schema:"Status,omitempty"`
	State                string `schema:"State,omitempty"`
}

func (p JournalEntriesWithLinesGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *JournalEntriesWithLinesGetRequest) QueryParams() *JournalEntriesWithLinesGetRequestQueryParams {
	return r.queryParams
}

func (r JournalEntriesWithLinesGetRequest) NewPathParams() *JournalEntriesWithLinesGetRequestPathParams {
	return &JournalEntriesWithLinesGetRequestPathParams{}
}

type JournalEntriesWithLinesGetRequestPathParams struct {
}

func (p *JournalEntriesWithLinesGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *JournalEntriesWithLinesGetRequest) PathParams() *JournalEntriesWithLinesGetRequestPathParams {
	return r.pathParams
}

func (r *JournalEntriesWithLinesGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *JournalEntriesWithLinesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *JournalEntriesWithLinesGetRequest) Method() string {
	return r.method
}

func (r JournalEntriesWithLinesGetRequest) NewRequestBody() JournalEntriesWithLinesGetRequestBody {
	return JournalEntriesWithLinesGetRequestBody{}
}

type JournalEntriesWithLinesGetRequestBody struct {
}

func (r *JournalEntriesWithLinesGetRequest) RequestBody() *JournalEntriesWithLinesGetRequestBody {
	return nil
}

func (r *JournalEntriesWithLinesGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *JournalEntriesWithLinesGetRequest) SetRequestBody(body JournalEntriesWithLinesGetRequestBody) {
	r.requestBody = body
}

func (r *JournalEntriesWithLinesGetRequest) NewResponseBody() *JournalEntriesWithLinesGetResponseBody {
	return &JournalEntriesWithLinesGetResponseBody{}
}

type JournalEntriesWithLinesGetResponseBody struct {
	Data JournalEntriesWithLines `json:"_data"`
	Meta Meta                    `json:"_meta"`
}

func (r *JournalEntriesWithLinesGetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/v1/journalentrieswithlines", r.PathParams())
	return &u
}

func (r *JournalEntriesWithLinesGetRequest) Do() (JournalEntriesWithLinesGetResponseBody, error) {
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
