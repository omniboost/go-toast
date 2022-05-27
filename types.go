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
	AdministrationID   int    `json:"administrationId"`
	AdministrationName string `json:"administrationName"`
	ProfileName        string `json:"profileName"`
	PreviousPage       string `json:"previousPage"`
	NextPage           string `json:"nextPage"`
	TotalCount         int    `json:"totalCount"`
	PageCount          int    `json:"pageCount"`
	Self               string `json:"self"`
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

type Journals []Journal

type Journal struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Meta        struct {
		Self string `json:"self"`
	} `json:"_meta"`
}

type LedgerAccounts []LedgerAccount

type LedgerAccount struct {
	ID          int         `json:"id"`
	Type        int         `json:"type"`
	Description string      `json:"description"`
	VATCode     interface{} `json:"vatCode"`
	Disabled    bool        `json:"disabled"`
	Rcsfi       interface{} `json:"rcsfi"`
	Meta        struct {
		Self string `json:"self"`
	} `json:"_meta"`
}

type JournalEntriesWithLines []JournalEntryWithLines

type JournalEntryWithLines struct {
	Lines struct {
		Data JournalEntryLines `json:"_data"`
	} `json:"lines"`
	ID                   int    `json:"id,omitempty"`
	Bookperiod           int    `json:"bookperiod"`
	Bookyear             int    `json:"bookyear"`
	Date                 Date   `json:"date"`
	JournalID            string `json:"journalId"`
	OpeningBalance       int    `json:"openingBalance,omitempty"`
	ClosingBalance       int    `json:"closingBalance,omitempty"`
	Description          string `json:"description"`
	SourceDocumentNumber string `json:"sourceDocumentNumber,omitempty"`
	State                int    `json:"state,omitempty"`
	// Meta                 struct {
	// 	AdministrationID   int    `json:"administrationId"`
	// 	AdministrationName string `json:"administrationName"`
	// 	ProfileName        string `json:"profileName"`
	// 	Self               string `json:"self"`
	// 	AdditionalProp1    struct {
	// 	} `json:"additionalProp1"`
	// 	AdditionalProp2 struct {
	// 	} `json:"additionalProp2"`
	// 	AdditionalProp3 struct {
	// 	} `json:"additionalProp3"`
	// } `json:"_meta"`
}

type JournalEntryLines []JournalEntryLine

type JournalEntryLine struct {
	ID                int     `json:"id,omitempty"`
	Amount            float64 `json:"amount"`
	BookingDate       Date    `json:"bookingDate"`
	CreditorID        int     `json:"creditorId,omitempty"`
	DebtorID          int     `json:"debtorId,omitempty"`
	PurchaseInvoiceID int     `json:"purchaseInvoiceId,omitempty"`
	SalesInvoiceID    int     `json:"salesInvoiceId,omitempty"`
	Description       string  `json:"description"`
	LedgerAccountID   int     `json:"ledgerAccountId,omitempty"`
	VATAmount         float64 `json:"vatAmount"`
	VATCode           string  `json:"vatCode"`
	CostCenterID      int     `json:"costCenterId"`
	// Meta              struct {
	// 	AdministrationID   int    `json:"administrationId"`
	// 	AdministrationName string `json:"administrationName"`
	// 	ProfileName        string `json:"profileName"`
	// 	Self               string `json:"self"`
	// 	AdditionalProp1    struct {
	// 	} `json:"additionalProp1"`
	// 	AdditionalProp2 struct {
	// 	} `json:"additionalProp2"`
	// 	AdditionalProp3 struct {
	// 	} `json:"additionalProp3"`
	// } `json:"_meta"`
}

type SalesInvoices []SalesInvoice

type SalesInvoice struct {
	ID                   int     `json:"id,omitempty"`
	AmountExclVAT        float64 `json:"amountExclVat,omitempty"`
	AmountInclVAT        float64 `json:"amountInclVat,omitempty"`
	Bookyear             int     `json:"bookyear"`
	Bookperiod           int     `json:"bookperiod"`
	DebtorCustomerNumber string  `json:"debtorCustomerNumber"`
	Date                 Date    `json:"date"`
	Description          string  `json:"description"`
	InvoiceNumber        string  `json:"invoiceNumber"`
	JournalID            string  `json:"journalId"`
	Status               int     `json:"status,omitempty"`
	PaymentConditionID   int     `json:"paymentConditionId,omitempty"`
	PaymentDueDate       Date    `json:"paymentDueDate,omitempty"`
	VATAmount            float64 `json:"vatAmount,omitempty"`
	Lines                struct {
		Data SalesInvoiceLines `json:"_data"`
		Meta struct {
			TotalCount int `json:"totalCount"`
		} `json:"_meta"`
	} `json:"lines"`
	Meta struct {
		Self string `json:"self"`
	} `json:"_meta"`
}

type SalesInvoiceLines []SalesInvoiceLine

type SalesInvoiceLine struct {
	Amount          float64 `json:"amount"`
	CostCenterID    int     `json:"costCenterId,omitempty"`
	CostUnitID      int     `json:"costUnitId,omitempty"`
	Description     string  `json:"description"`
	ID              int     `json:"id,omitempty"`
	LedgerAccountID int     `json:"ledgerAccountId"`
	VATAmount       float64 `json:"vatAmount,omitempty"`
	VATCode         string  `json:"vatCode"`
}

type VATCodes []VATCode

type VATCode struct {
	ID          string      `json:"id"`
	Description string      `json:"description"`
	SalesCode   bool        `json:"salesCode"`
	Percentage  float64     `json:"percentage"`
	IsDefault   bool        `json:"isDefault"`
	ValidFrom   interface{} `json:"validFrom"`
	ValidUntil  interface{} `json:"validUntil"`
	CountryCode string      `json:"countryCode"`
	Meta        struct {
		Self string `json:"self"`
	} `json:"_meta"`
}

type Debtors []Debtor

type Debtor struct {
	CustomerNumber     string      `json:"customerNumber,omitempty"`
	Bban               string      `json:"bban,omitempty"`
	DirectDebit        bool        `json:"directDebit,omitempty"`
	Iban               string      `json:"iban,omitempty"`
	Bic                string      `json:"bic,omitempty"`
	CocNumber          string      `json:"cocNumber,omitempty"`
	VatNumber          string      `json:"vatNumber,omitempty"`
	City               string      `json:"city,omitempty"`
	Address            string      `json:"address,omitempty"`
	AddressLine2       interface{} `json:"addressLine2,omitempty"`
	AddressLine3       interface{} `json:"addressLine3,omitempty"`
	AttentionTo        string      `json:"attentionTo,omitempty"`
	CountryCode        string      `json:"countryCode,omitempty"`
	PostCode           string      `json:"postCode,omitempty"`
	DebtorKindID       string      `json:"debtorKindId,omitempty"`
	PaymentConditionID int         `json:"paymentConditionId,omitempty"`
	PhoneNumber        string      `json:"phoneNumber,omitempty"`
	FaxNumber          string      `json:"faxNumber,omitempty"`
	MobileNumber       string      `json:"mobileNumber,omitempty"`
	EmailAddress       string      `json:"emailAddress,omitempty"`
	Comments           interface{} `json:"comments,omitempty"`
	Characteristic4    interface{} `json:"characteristic4,omitempty"`
	Characteristic5    interface{} `json:"characteristic5,omitempty"`
	PaymentMethodID    string      `json:"paymentMethodId,omitempty"`
	DateAdded          string      `json:"dateAdded,omitempty"`
	DateLastChange     interface{} `json:"dateLastChange,omitempty"`
	Blocked            bool        `json:"blocked,omitempty"`
	NoReminder         bool        `json:"noReminder,omitempty"`
	ID                 int         `json:"id,omitempty"`
	Name               string      `json:"name,omitempty"`
	Meta               struct {
		Self string `json:"self"`
	} `json:"_meta"`
}

type PaymentConditions []PaymentCondition

type PaymentCondition struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Meta        struct {
		Self string `json:"self"`
	} `json:"_meta"`
}

type File struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	Extension     string `json:"extension"`
	Size          int    `json:"size"`
	CreatedOn     string `json:"createdOn"`
	FileURL       string `json:"fileUrl"`
	FullImageURL  string `json:"fullImageUrl"`
	Thumbnail1URL string `json:"thumbnail1Url"`
	Thumbnail2URL string `json:"thumbnail2Url"`
	Thumbnail3URL string `json:"thumbnail3Url"`
	Meta          struct {
		AdministrationID   int    `json:"administrationId"`
		AdministrationName string `json:"administrationName"`
		ProfileName        string `json:"profileName"`
		Self               string `json:"self"`
		AdditionalProp1    string `json:"additionalProp1"`
		AdditionalProp2    string `json:"additionalProp2"`
		AdditionalProp3    string `json:"additionalProp3"`
	} `json:"_meta"`
}
