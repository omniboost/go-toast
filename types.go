package asperion

type Administrations []Administration

type Administration struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Meta struct {
		Self string `json:"self"`
	} `json:"_meta"`
}

type Meta struct {
	AdministrationID   int         `json:"administrationId"`
	AdministrationName string      `json:"administrationName"`
	ProfileName        string      `json:"profileName"`
	PreviousPage       interface{} `json:"previousPage"`
	NextPage           interface{} `json:"nextPage"`
	TotalCount         int         `json:"totalCount"`
	PageCount          int         `json:"pageCount"`
	Self               string      `json:"self"`
}

type Costcenters []Costcenter

type Costcenter struct {
	ID             int         `json:"id"`
	Description    string      `json:"description"`
	ConversionCode interface{} `json:"conversionCode"`
	Meta           struct {
		Self string `json:"self"`
	} `json:"_meta"`
}

type Costunits []Costunit

type Costunit struct {
}
