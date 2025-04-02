package toast

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-toast/utils"
)

func (c *Client) NewMenuV2GetRequest() MenuV2GetRequest {
	r := MenuV2GetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type MenuV2GetRequest struct {
	client      *Client
	queryParams *MenuV2GetRequestQueryParams
	pathParams  *MenuV2GetRequestPathParams
	method      string
	headers     http.Header
	requestBody MenuV2GetRequestBody
}

func (r MenuV2GetRequest) NewQueryParams() *MenuV2GetRequestQueryParams {
	return &MenuV2GetRequestQueryParams{}
}

type MenuV2GetRequestQueryParams struct {
}

func (p MenuV2GetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *MenuV2GetRequest) QueryParams() *MenuV2GetRequestQueryParams {
	return r.queryParams
}

func (r MenuV2GetRequest) NewPathParams() *MenuV2GetRequestPathParams {
	return &MenuV2GetRequestPathParams{}
}

type MenuV2GetRequestPathParams struct {
}

func (p *MenuV2GetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *MenuV2GetRequest) PathParams() *MenuV2GetRequestPathParams {
	return r.pathParams
}

func (r *MenuV2GetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *MenuV2GetRequest) SetMethod(method string) {
	r.method = method
}

func (r *MenuV2GetRequest) Method() string {
	return r.method
}

func (r MenuV2GetRequest) NewRequestBody() MenuV2GetRequestBody {
	return MenuV2GetRequestBody{}
}

type MenuV2GetRequestBody struct {
}

func (r *MenuV2GetRequest) RequestBody() *MenuV2GetRequestBody {
	return nil
}

func (r *MenuV2GetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *MenuV2GetRequest) SetRequestBody(body MenuV2GetRequestBody) {
	r.requestBody = body
}

func (r *MenuV2GetRequest) NewResponseBody() *MenuV2GetResponseBody {
	return &MenuV2GetResponseBody{}
}

type MenuV2GetResponseBody struct {
	RestaurantGUID     string `json:"restaurantGuid"`
	LastUpdated        string `json:"lastUpdated"`
	RestaurantTimeZone string `json:"restaurantTimeZone"`
	Menus              []struct {
		Name            string `json:"name"`
		GUID            string `json:"guid"`
		MultiLocationID string `json:"multiLocationId"`
		// MasterID            int      `json:"masterId"` -> deprecated
		Description         string   `json:"description"`
		POSName             string   `json:"posName"`
		POSButtonColorLight string   `json:"posButtonColorLight"`
		POSButtonColorDark  string   `json:"posButtonColorDark"`
		HighResImage        string   `json:"highResImage"`
		Image               string   `json:"image"`
		Visibility          []string `json:"visibility"`
		Availability        struct {
			AlwaysAvailable bool `json:"alwaysAvailable"`
			Schedule        []struct {
				Days       []string
				TimeRanges []struct {
					Start string `json:"start"`
					End   string `json:"end"`
				} `json:"timeRanges"`
			} `json:"schedule"`
		} `json:"availability"`
		MenuGroups []MenuGroupV2 `json:"menuGroups"`
	} `json:"menus"`
	ModifierGroupReferences    map[string]ModifierGroupV2    `json:"modifierGroupReferences"`
	ModifierOptionReferences   map[string]ModifierOptionV2   `json:"modifierOptionReferences"`
	PreModifierGroupReferences map[string]PreModifierGroupV2 `json:"preModifierGroupReferences"`
}

type PricingRulesV2 struct {
	TimeSpecificPricingRules []struct {
		TimeSpecificPrice float64 `json:"timeSpecificPrice"`
		BasePrice         float64 `json:"basePrice"`
		Schedule          []struct {
			Days       []string `json:"days"`
			TimeRanges []struct {
				Start string `json:"start"`
				End   string `json:"end"`
			} `json:"timeRanges"`
		} `json:"schedule"`
	} `json:"timeSpecificPricingRules"`
	SizeSpecificPricingGUID  string `json:"sizeSpecificPricingGuid"`
	SizeSequencePricingRules []struct {
		SizeName       string `json:"sizeName"`
		SizeGUID       string `json:"sizeGuid"`
		SequencePrices []struct {
			Sequence int     `json:"sequence"`
			Price    float64 `json:"price"`
		} `json:"sequencePrices"`
	} `json:"sizeSequencePricingRules"`
}

type MenuGroupV2 struct {
	Name            string `json:"name"`
	GUID            string `json:"guid"`
	MultiLocationID string `json:"multiLocationId"`
	// MasterID            int      `json:"masterId"` -> deprecated
	Description         string   `json:"description"`
	POSName             string   `json:"posName"`
	POSButtonColorLight string   `json:"posButtonColorLight"`
	POSButtonColorDark  string   `json:"posButtonColorDark"`
	Image               string   `json:"image"`
	Visibility          []string `json:"visibility"`
	ItemTags            []struct {
		Name string `json:"name"`
		GUID string `json:"guid"`
	} `json:"itemTags"`
	MenuGroups []MenuGroupV2 `json:"menuGroups"`
	MenuItems  []MenuItemV2  `json:"menuItems"`
}

type MenuItemV2 struct {
	Name            string `json:"name"`
	KitchenName     string `json:"kitchenName"`
	GUID            string `json:"guid"`
	MultiLocationID string `json:"multiLocationId"`
	// MasterID            int      `json:"masterId"` -> deprecated
	Description         string         `json:"description"`
	POSName             string         `json:"posName"`
	POSButtonColorLight string         `json:"posButtonColorLight"`
	POSButtonColorDark  string         `json:"posButtonColorDark"`
	Image               string         `json:"image"`
	Visibility          []string       `json:"visibility"`
	Price               float64        `json:"price"`
	PricingStrategy     string         `json:"pricingStrategy"`
	PricingRules        PricingRulesV2 `json:"pricingRules"`
	IsDeferred          bool           `json:"isDeferred"`
	IsDiscountable      bool           `json:"isDiscountable"`
	SalesCategory       struct {
		Name string `json:"name"`
		GUID string `json:"guid"`
	} `json:"salesCategory"`
	TaxInfo      []string `json:"taxInfo"`
	TaxInclusion string   `json:"taxInclusion"`
	ItemTags     []struct {
		Name string `json:"name"`
		GUID string `json:"guid"`
	} `json:"itemTags"`
	Plu               string `json:"plu"`
	Sku               string `json:"sku"`
	Calories          int    `json:"calories"`
	ContentAdvisories struct {
		Alcohol struct {
			ContainsAlcohol string `json:"containsAlcohol"`
		} `json:"alcohol"`
	} `json:"contentAdvisories"`
	UnitOfMeasure string `json:"unitOfMeasure"`
	Portions      []struct {
		Name                    string `json:"name"`
		GUID                    string `json:"guid"`
		ModifierGroupReferences []int  `json:"modifierGroupReferences"`
	} `json:"portions"`
	PrepTime                          int      `json:"prepTime"`
	PrepStations                      []string `json:"prepStations"`
	ModifierGroupReferences           []int    `json:"modifierGroupReferences"`
	EligiblePaymentAssistancePrograms []string `json:"eligiblePaymentAssistancePrograms"`
	// Length                            float64      `json:"length"`
	// Height                            float64      `json:"height"`
	// Width                             float64      `json:"width"`
	// DimensionUnitOfMeasure            string   `json:"dimensionUnitOfMeasure"`
	// Weight                            float64      `json:"weight"`
	// WeightUnitOfMeasure               string   `json:"weightUnitOfMeasure"`
	Images []string `json:"images"`
	// GuestCount                        float64      `json:"guestCount"`
}

type ModifierGroupV2 struct {
	Name            string `json:"name"`
	GUID            string `json:"guid"`
	ReferenceID     int    `json:"referenceId"`
	MultiLocationID string `json:"multiLocationId"`
	// MasterID            int      `json:"masterId"` -> deprecated
	POSName                           string         `json:"posName"`
	POSButtonColorLight               string         `json:"posButtonColorLight"`
	POSButtonColorDark                string         `json:"posButtonColorDark"`
	Visibility                        []string       `json:"visibility"`
	PricingStrategy                   string         `json:"pricingStrategy"`
	PricingRules                      PricingRulesV2 `json:"pricingRules"`
	DefaultOptionsChargePrice         string         `json:"defaultOptionsChargePrice"`
	DefaultOptionsSubstitutionPricing string         `json:"defaultOptionsSubstitutionPricing"`
	MinSelections                     int            `json:"minSelections"`
	MaxSelections                     int            `json:"maxSelections"`
	RequiredMode                      string         `json:"requiredMode"`
	IsMultiSelect                     bool           `json:"isMultiSelect"`
	PreModifierGroupReference         int            `json:"preModifierGroupReference"`
	ModifierOptionReferences          []int          `json:"modifierOptionReferences"`
}

type ModifierOptionV2 struct {
	ReferenceID     int    `json:"referenceId"`
	Name            string `json:"name"`
	KitchenName     string `json:"kitchenName"`
	GUID            string `json:"guid"`
	MultiLocationID string `json:"multiLocationId"`
	// MasterID            int      `json:"masterId"` -> deprecated
	Description         string         `json:"description"`
	POSName             string         `json:"posName"`
	POSButtonColorLight string         `json:"posButtonColorLight"`
	POSButtonColorDark  string         `json:"posButtonColorDark"`
	PrepStations        []string       `json:"prepStations"`
	Image               string         `json:"image"`
	Visibility          []string       `json:"visibility"`
	Price               int            `json:"price"`
	PricingStrategy     string         `json:"pricingStrategy"`
	PricingRules        PricingRulesV2 `json:"pricingRules"`
	SalesCategory       struct {
		Name string `json:"name"`
		GUID string `json:"guid"`
	} `json:"salesCategory"`
	TaxInfo               []string `json:"taxInfo"`
	ModifierOptionTaxInfo struct {
		TaxRateGuids         []string `json:"taxRateGuids"`
		OverrideItemTaxRates bool     `json:"overrideItemTaxRates"`
	} `json:"modifierOptionTaxInfo"`
	ItemTags []struct {
		Name string `json:"name"`
		GUID string `json:"guid"`
	} `json:"itemTags"`
	Plu               string `json:"plu"`
	Sku               string `json:"sku"`
	Calories          int    `json:"calories"`
	ContentAdvisories struct {
		Alcohol struct {
			ContainsAlcohol string `json:"containsAlcohol"`
		} `json:"alcohol"`
	} `json:"contentAdvisories"`
	UnitOfMeasure    string `json:"unitOfMeasure"`
	IsDefault        bool   `json:"isDefault"`
	AllowsDuplicates bool   `json:"allowsDuplicates"`
	Portions         []struct {
		Name                    string `json:"name"`
		GUID                    string `json:"guid"`
		ModifierGroupReferences []int  `json:"modifierGroupReferences"`
	} `json:"portions"`
	PrepTime                int      `json:"prepTime"`
	ModifierGroupReferences []int    `json:"modifierGroupReferences"`
	Length                  float64  `json:"length"`
	Height                  float64  `json:"height"`
	Width                   float64  `json:"width"`
	DimensionUnitOfMeasure  string   `json:"dimensionUnitOfMeasure"`
	Weight                  float64  `json:"weight"`
	WeightUnitOfMeasure     string   `json:"weightUnitOfMeasure"`
	Images                  []string `json:"images"`
	GuestCount              float64  `json:"guestCount"`
}

type PreModifierGroupV2 struct {
	Name            string `json:"name"`
	GUID            string `json:"guid"`
	MultiLocationID string `json:"multiLocationId"`
	PreModifiers    struct {
		Name                 string  `json:"name"`
		GUID                 string  `json:"guid"`
		MultiLocationID      string  `json:"multiLocationId"`
		FixedPrice           float64 `json:"fixedPrice"`
		MultiplicationFactor float64 `json:"multiplicationFactor"`
		DisplayMode          string  `json:"displayMode"`
		POSName              string  `json:"posName"`
		POSButtonColorLight  string  `json:"posButtonColorLight"`
		POSButtonColorDark   string  `json:"posButtonColorDark"`
	} `json:"preModifiers"`
}

func (r *MenuV2GetRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("/menus/v2/menus", r.PathParams())
	return &u
}

func (r *MenuV2GetRequest) Do() (MenuV2GetResponseBody, error, *http.Response) {
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
