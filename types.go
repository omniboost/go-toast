package toast

import "time"

type Token struct {
	TokenType    string `json:"tokenType"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expiresIn"`
	AccessToken  string `json:"accessToken"`
	IDToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
	expiry       time.Time
}

func (t Token) IsExpired() bool {
	return time.Now().After(t.expiry)
}

type Orders []Order

type Order struct {
	GUID          string        `json:"guid"`
	EntityType    string        `json:"entityType"`
	ExternalID    string        `json:"externalId"`
	RevenueCenter RevenueCenter `json:"revenueCenter"`
	Server        struct {
		GUID       string `json:"guid"`
		EntityType string `json:"entityType"`
		ExternalID string `json:"externalId"`
	} `json:"server"`
	LastModifiedDevice struct {
		ID string `json:"id"`
	} `json:"lastModifiedDevice"`
	Source            string   `json:"source"`
	VoidDate          DateTime `json:"voidDate"`
	Duration          *int     `json:"duration"` // switched to pointer
	DisplayNumber     string   `json:"displayNumber"`
	BusinessDate      Date     `json:"businessDate"`
	PaidDate          DateTime `json:"paidDate"`
	RestaurantService struct {
		GUID       string `json:"guid"`
		EntityType string `json:"entityType"`
		ExternalID string `json:"externalId"`
	} `json:"restaurantService"`
	Voided                   *bool    `json:"voided"` // switched to pointer
	EstimatedFulfillmentDate DateTime `json:"estimatedFulfillmentDate"`
	Table                    struct {
		GUID       string `json:"guid"`
		EntityType string `json:"entityType"`
		ExternalID string `json:"externalId"`
	} `json:"table"`
	RequiredPrepTime string `json:"requiredPrepTime"`
	ApprovalStatus   string `json:"approvalStatus"`
	DeliveryInfo     struct {
		Address1           string   `json:"address1"`
		Address2           string   `json:"address2"`
		City               string   `json:"city"`
		AdministrativeArea string   `json:"administrativeArea"`
		State              string   `json:"state"`
		ZipCode            string   `json:"zipCode"`
		Country            string   `json:"country"`
		Latitude           *float64 `json:"latitude"`  // switched to pointer
		Longitude          *float64 `json:"longitude"` // switched to pointer
		Notes              string   `json:"notes"`
		DeliveredDate      DateTime `json:"deliveredDate"`
		DispatchedDate     DateTime `json:"dispatchedDate"`
		DeliveryEmployee   struct {
			GUID       string `json:"guid"`
			EntityType string `json:"entityType"`
			ExternalID string `json:"externalId"`
		} `json:"deliveryEmployee"`
		DeliveryState string `json:"deliveryState"`
	} `json:"deliveryInfo"`
	ServiceArea struct {
		GUID       string `json:"guid"`
		EntityType string `json:"entityType"`
		ExternalID string `json:"externalId"`
	} `json:"serviceArea"`
	CurbsidePickupInfo struct {
		GUID                 string `json:"guid"`
		EntityType           string `json:"entityType"`
		TransportColor       string `json:"transportColor"`
		TransportDescription string `json:"transportDescription"`
		Notes                string `json:"notes"`
	} `json:"curbsidePickupInfo"`
	NumberOfGuests      *int `json:"numberOfGuests"` // switched to pointer
	DeliveryServiceInfo struct {
		DriverName                 string   `json:"driverName"`
		DriverPhoneNumber          string   `json:"driverPhoneNumber"`
		EntityType                 string   `json:"entityType"`
		GUID                       string   `json:"guid"`
		OriginalQuotedDeliveryDate DateTime `json:"originalQuotedDeliveryDate"`
		ProviderID                 string   `json:"providerId"`
		ProviderInfo               string   `json:"providerInfo"`
		ProviderName               string   `json:"providerName"`
	} `json:"deliveryServiceInfo"`
	DiningOption     DiningOption `json:"diningOption"`
	OpenedDate       DateTime     `json:"openedDate"`
	VoidBusinessDate Date         `json:"voidBusinessDate"`
	Checks           []struct {
		GUID          string `json:"guid"`
		EntityType    string `json:"entityType"`
		ExternalID    string `json:"externalId"`
		DisplayNumber string `json:"displayNumber"`
		Duration      *int   `json:"duration"` // switched to pointer
		Payments      []struct {
			GUID                  string   `json:"guid"`
			EntityType            string   `json:"entityType"`
			ExternalID            string   `json:"externalId"`
			OriginalProcessingFee *float64 `json:"originalProcessingFee"` // switched to pointer
			Server                struct {
				GUID       string `json:"guid"`
				EntityType string `json:"entityType"`
				ExternalID string `json:"externalId"`
			} `json:"server"`
			CashDrawer struct {
				GUID       string `json:"guid"`
				EntityType string `json:"entityType"`
				ExternalID string `json:"externalId"`
			} `json:"cashDrawer"`
			LastModifiedDevice struct {
				ID string `json:"id"`
			} `json:"lastModifiedDevice"`
			RefundStatus       string      `json:"refundStatus"`
			IsProcessedOffline interface{} `json:"isProcessedOffline"` // not in the API documentation
			Type               string      `json:"type"`
			VoidInfo           struct {
				VoidUser struct {
					GUID       string `json:"guid"`
					EntityType string `json:"entityType"`
					ExternalID string `json:"externalId"`
				} `json:"voidUser"`
				VoidApprover struct {
					GUID       string `json:"guid"`
					EntityType string `json:"entityType"`
					ExternalID string `json:"externalId"`
				} `json:"voidApprover"`
				VoidDate         DateTime `json:"voidDate"`
				VoidBusinessDate Date     `json:"voidBusinessDate"`
				VoidReason       struct {
					GUID       string `json:"guid"`
					EntityType string `json:"entityType"`
					ExternalID string `json:"externalId"`
				} `json:"voidReason"`
			} `json:"voidInfo"`
			CheckGUID    string `json:"checkGuid"`
			OtherPayment struct {
				GUID       string `json:"guid"`
				EntityType string `json:"entityType"`
				ExternalID string `json:"externalId"`
			} `json:"otherPayment"`
			PaidDate       DateTime `json:"paidDate"`
			OrderGUID      string   `json:"orderGuid"`
			CardEntryMode  string   `json:"cardEntryMode"`
			PaymentStatus  string   `json:"paymentStatus"`
			Amount         *float64 `json:"amount"`         // switched to pointer
			TipAmount      *float64 `json:"tipAmount"`      // switched to pointer
			AmountTendered *float64 `json:"amountTendered"` // switched to pointer
			CardType       string   `json:"cardType"`
			HouseAccount   struct {
				GUID       string `json:"guid"`
				EntityType string `json:"entityType"`
				ExternalID string `json:"externalId"`
			} `json:"houseAccount"`
			McaRepaymentAmount *float64 `json:"mcaRepaymentAmount"` // switched to pointer
			CreatedDevice      struct {
				ID string `json:"id"`
			} `json:"createdDevice"`
			PaidBusinessDate      Date   `json:"paidBusinessDate"`
			Last4Digits           string `json:"last4Digits"`
			Refund                Refund `json:"refund"` // used to be RefundDetails
			CardPaymentID         string `json:"cardPaymentId"`
			TenderTransactionGUID string `json:"tenderTransactionGuid"`
		} `json:"payments"`
		AppliedDiscounts AppliedDiscounts `json:"appliedDiscounts"`

		LastModifiedDevice struct {
			ID string `json:"id"`
		} `json:"lastModifiedDevice"`
		VoidDate           DateTime `json:"voidDate"`
		PaidDate           DateTime `json:"paidDate"`
		AppliedLoyaltyInfo struct {
			GUID                    string `json:"guid"`
			EntityType              string `json:"entityType"`
			LoyaltyIdentifier       string `json:"loyaltyIdentifier"`
			MaskedLoyaltyIdentifier string `json:"maskedLoyaltyIdentifier"`
			Vendor                  string `json:"vendor"`
			AccrualFamilyGuid       string `json:"accrualFamilyGuid"`
			AccrualText             string `json:"accrualText"`
		} `json:"appliedLoyaltyInfo"`
		Voided        *bool    `json:"voided"` // switched to pointer
		PaymentStatus string   `json:"paymentStatus"`
		Amount        *float64 `json:"amount"` // switched to pointer
		TabName       string   `json:"tabName"`
		TaxExempt     *bool    `json:"taxExempt"` // switched to pointer
		OpenedBy      struct {
			GUID       string `json:"guid"`
			EntityType string `json:"entityType"`
			ExternalID string `json:"externalId"`
		} `json:"openedBy"`
		OpenedDate  DateTime `json:"openedDate"`
		TotalAmount *float64 `json:"totalAmount"` // switched to pointer
		Selections  []struct {
			Selection
			Modifiers []Selection `json:"modifiers"`
		} `json:"selections"`
		VoidBusinessDate Date     `json:"voidBusinessDate"`
		CreatedDate      DateTime `json:"createdDate"`
		Deleted          *bool    `json:"deleted"` // switched to pointer
		CreatedDevice    struct {
			ID string `json:"id"`
		} `json:"createdDevice"`
		ClosedDate            DateTime              `json:"closedDate"`
		DeletedDate           DateTime              `json:"deletedDate"`
		ModifiedDate          DateTime              `json:"modifiedDate"`
		TaxAmount             *float64              `json:"taxAmount"` // switched to pointer
		AppliedServiceCharges AppliedServiceCharges `json:"appliedServiceCharges"`
		Customer              struct {
			GUID             string `json:"guid"`
			EntityType       string `json:"entityType"`
			FirstName        string `json:"firstName"`
			LastName         string `json:"lastName"`
			Phone            string `json:"phone"`
			PhoneCountryCode string `json:"phoneCountryCode"`
			Email            string `json:"email"`
		} `json:"customer"`
	} `json:"checks"`
	Deleted       *bool `json:"deleted"` // switched to pointer
	CreatedDevice struct {
		ID string `json:"id"`
	} `json:"createdDevice"`
	CreatedDate          DateTime `json:"createdDate"`
	ClosedDate           DateTime `json:"closedDate"`
	DeletedDate          DateTime `json:"deletedDate"`
	ModifiedDate         DateTime `json:"modifiedDate"`
	PromisedDate         DateTime `json:"promisedDate"`
	ChannelGUID          string   `json:"channelGuid"`
	PricingFeatures      []string `json:"pricingFeatures"`
	AppliedPackagingInfo struct {
		AppliedPackagingItems []struct {
			EntityType       string   `json:"entityType"`
			GUID             string   `json:"guid"`
			GuestDisplayName string   `json:"guestDisplayName"`
			Inclusion        string   `json:"inclusion"`
			ItemConfigId     string   `json:"itemConfigId"`
			ItemTypes        []string `json:"itemTypes"`
		} `json:"appliedPackagingItems"`
		EntityType string `json:"entityType"`
		GUID       string `json:"guid"`
	} `json:"appliedPackagingInfo"`
	ExcessFood        *bool `json:"excessFood"`        // switched to pointer
	CreatedInTestMode *bool `json:"createdInTestMode"` // switched to pointer
}

type Restaurants []Restaurant

type Restaurant struct {
	GUID    string `json:"guid"`
	General struct {
		Name                string      `json:"name"`
		LocationName        string      `json:"locationName"`
		LocationCode        interface{} `json:"locationCode"`
		Description         string      `json:"description"`
		TimeZone            string      `json:"timeZone"`
		CloseoutHour        int         `json:"closeoutHour"`
		ManagementGroupGUID string      `json:"managementGroupGuid"`
	} `json:"general"`
	Urls struct {
		Website          interface{} `json:"website"`
		Facebook         interface{} `json:"facebook"`
		Twitter          string      `json:"twitter"`
		OrderOnline      string      `json:"orderOnline"`
		PurchaseGiftCard string      `json:"purchaseGiftCard"`
		CheckGiftCard    string      `json:"checkGiftCard"`
	} `json:"urls"`
	Location struct {
		Address1  string `json:"address1"`
		Address2  string `json:"address2"`
		City      string `json:"city"`
		StateCode string `json:"stateCode"`
		ZipCode   string `json:"zipCode"`
		Country   string `json:"country"`
		Phone     string `json:"phone"`
		Latitude  int    `json:"latitude"`
		Longitude int    `json:"longitude"`
	} `json:"location"`
	Schedules struct {
		DaySchedules struct {
			Num600000000314630485 struct {
				ScheduleName string      `json:"scheduleName"`
				Services     interface{} `json:"services"`
				OpenTime     string      `json:"openTime"`
				CloseTime    string      `json:"closeTime"`
			} `json:"600000000314630485"`
		} `json:"daySchedules"`
		WeekSchedule struct {
			Monday    string `json:"monday"`
			Tuesday   string `json:"tuesday"`
			Wednesday string `json:"wednesday"`
			Thursday  string `json:"thursday"`
			Friday    string `json:"friday"`
			Saturday  string `json:"saturday"`
			Sunday    string `json:"sunday"`
		} `json:"weekSchedule"`
	} `json:"schedules"`
	Delivery struct {
		Enabled bool        `json:"enabled"`
		Minimum interface{} `json:"minimum"`
		Area    interface{} `json:"area"`
	} `json:"delivery"`
	OnlineOrdering struct {
		Enabled                bool   `json:"enabled"`
		Scheduling             bool   `json:"scheduling"`
		SpecialRequests        bool   `json:"specialRequests"`
		SpecialRequestsMessage string `json:"specialRequestsMessage"`
		PaymentOptions         struct {
			Delivery struct {
				Cash      bool `json:"cash"`
				CcSameDay bool `json:"ccSameDay"`
				CcFuture  bool `json:"ccFuture"`
			} `json:"delivery"`
			Takeout struct {
				Cash      bool `json:"cash"`
				CcSameDay bool `json:"ccSameDay"`
				CcFuture  bool `json:"ccFuture"`
				CcInStore bool `json:"ccInStore"`
			} `json:"takeout"`
			CcTip bool `json:"ccTip"`
		} `json:"paymentOptions"`
	} `json:"onlineOrdering"`
	PrepTimes struct {
		DeliveryPrepTime        int `json:"deliveryPrepTime"`
		DeliveryTimeAfterOpen   int `json:"deliveryTimeAfterOpen"`
		DeliveryTimeBeforeClose int `json:"deliveryTimeBeforeClose"`
		TakeoutPrepTime         int `json:"takeoutPrepTime"`
		TakeoutTimeAfterOpen    int `json:"takeoutTimeAfterOpen"`
		TakeoutTimeBeforeClose  int `json:"takeoutTimeBeforeClose"`
		TakeoutThrottlingTime   int `json:"takeoutThrottlingTime"`
		DeliveryThrottlingTime  int `json:"deliveryThrottlingTime"`
	} `json:"prepTimes"`
}

type MenuGroups []MenuGroup

type MenuGroup struct {
	GUID       string      `json:"guid"`
	EntityType string      `json:"entityType"`
	ExternalID interface{} `json:"externalId"`
	Parent     struct {
		GUID       string `json:"guid"`
		EntityType string `json:"entityType"`
		ExternalID string `json:"externalId"`
	} `json:"parent"`
	Images []struct {
		URL string `json:"url"`
	} `json:"images"`
	Visibility    string `json:"visibility"`
	UnitOfMeasure string `json:"unitOfMeasure"`
	OptionGroups  []struct {
		GUID       string `json:"guid"`
		EntityType string `json:"entityType"`
		ExternalID string `json:"externalId"`
	} `json:"optionGroups"`
	Menu struct {
		GUID       string `json:"guid"`
		EntityType string `json:"entityType"`
		ExternalID string `json:"externalId"`
	} `json:"menu"`
	InheritUnitOfMeasure bool `json:"inheritUnitOfMeasure"`
	Subgroups            []struct {
		GUID       string `json:"guid"`
		EntityType string `json:"entityType"`
		ExternalID string `json:"externalId"`
	} `json:"subgroups"`
	InheritOptionGroups bool   `json:"inheritOptionGroups"`
	OrderableOnline     string `json:"orderableOnline"`
	Name                string `json:"name"`
	Items               []struct {
		GUID       string `json:"guid"`
		EntityType string `json:"entityType"`
		ExternalID string `json:"externalId"`
	} `json:"items"`
}

type ModifierGroups []ModifierGroup
type ModifierGroup struct {
	GUID       string `json:"guid"`
	EntityType string `json:"entityType"`
	ExternalID string `json:"externalId"`
	Name       string `json:"name"`
	Options    []struct {
		GUID       string `json:"guid"`
		EntityType string `json:"entityType"`
		ExternalID string `json:"externalId"`
	} `json:"options"`
	MinSelections int `json:"minSelections"`
	MaxSelections int `json:"maxSelections"`
}

type MenuItems []MenuItem

type MenuItem struct {
	GUID       string `json:"guid"`
	EntityType string `json:"entityType"`
	ExternalID string `json:"externalId"`
	Images     []struct {
		URL string `json:"url"`
	} `json:"images"`
	Visibility    string `json:"visibility"`
	UnitOfMeasure string `json:"unitOfMeasure"`
	OptionGroups  []struct {
		GUID       string `json:"guid"`
		EntityType string `json:"entityType"`
		ExternalID string `json:"externalId"`
	} `json:"optionGroups"`
	Calories             int    `json:"calories"`
	Type                 string `json:"type"`
	InheritUnitOfMeasure bool   `json:"inheritUnitOfMeasure"`
	InheritOptionGroups  bool   `json:"inheritOptionGroups"`
	OrderableOnline      string `json:"orderableOnline"`
	Name                 string `json:"name"`
	Plu                  string `json:"plu"`
	Sku                  string `json:"sku"`
}

type Discounts []Discount

type Discount struct {
	GUID                string      `json:"guid"`
	EntityType          string      `json:"entityType"`
	Amount              float64     `json:"amount"`
	SelectionType       string      `json:"selectionType"`
	NonExclusive        bool        `json:"nonExclusive"`
	Percentage          float64     `json:"percentage"`
	Name                string      `json:"name"`
	Active              bool        `json:"active"`
	ItemPickingPriority string      `json:"itemPickingPriority"`
	Type                string      `json:"type"`
	FixedTotal          interface{} `json:"fixedTotal"`
}

type SalesCategories []SalesCategory

type SalesCategory struct {
	GUID            string `json:"guid"`
	EntityType      string `json:"entityType"`
	Name            string `json:"name"`
	ExternalID      string `json:"externalId"`
	MultiLocationID string `json:"multiLocationId"`
}

type DiningOptions []DiningOption

type DiningOption struct {
	GUID       string `json:"guid"`
	EntityType string `json:"entityType"`
	ExternalID string `json:"externalId"`
	Name       string `json:"name"`     // not in the API documentation
	Curbside   bool   `json:"curbside"` // not in the API documentation
	Behavior   string `json:"behavior"` // not in the API documentation
}

type Employees []Employee
type Employee struct {
	GUID           string      `json:"guid"`
	EntityType     string      `json:"entityType"`
	ExternalID     interface{} `json:"externalId"`
	V2EmployeeGUID string      `json:"v2EmployeeGuid"`
	LastName       string      `json:"lastName"`
	WageOverrides  []struct {
		Wage         float64 `json:"wage"`
		JobReference struct {
			GUID       string      `json:"guid"`
			EntityType string      `json:"entityType"`
			ExternalID interface{} `json:"externalId"`
		} `json:"jobReference"`
	} `json:"wageOverrides"`
	PhoneNumberCountryCode string      `json:"phoneNumberCountryCode"`
	FirstName              string      `json:"firstName"`
	ChosenName             string      `json:"chosenName"`
	CreatedDate            string      `json:"createdDate"`
	PhoneNumber            interface{} `json:"phoneNumber"`
	Deleted                bool        `json:"deleted"`
	DeletedDate            string      `json:"deletedDate"`
	JobReferences          []struct {
		GUID       string      `json:"guid"`
		EntityType string      `json:"entityType"`
		ExternalID interface{} `json:"externalId"`
	} `json:"jobReferences"`
	ModifiedDate       string `json:"modifiedDate"`
	ExternalEmployeeID string `json:"externalEmployeeId"`
	Email              string `json:"email"`
}

type PostEmployee struct {
	EntityType     string      `json:"entityType"`
	ExternalID     interface{} `json:"externalId,omitempty"`
	V2EmployeeGUID string      `json:"v2EmployeeGuid,omitempty"`
	LastName       string      `json:"lastName,omitempty"`
	WageOverrides  []struct {
		Wage         float64 `json:"wage,omitempty"`
		JobReference struct {
			GUID       string      `json:"guid,omitempty"`
			EntityType string      `json:"entityType,omitempty"`
			ExternalID interface{} `json:"externalId,omitempty"`
		} `json:"jobReference,omitempty"`
	} `json:"wageOverrides,omitempty"`
	PhoneNumberCountryCode string      `json:"phoneNumberCountryCode,omitempty"`
	FirstName              string      `json:"firstName,omitempty"`
	ChosenName             string      `json:"chosenName,omitempty"`
	CreatedDate            string      `json:"createdDate,omitempty"`
	PhoneNumber            interface{} `json:"phoneNumber,omitempty"`
	JobReferences          []struct {
		GUID       string      `json:"guid,omitempty"`
		EntityType string      `json:"entityType,omitempty"`
		ExternalID interface{} `json:"externalId,omitempty"`
	} `json:"jobReferences,omitempty"`
	ModifiedDate       string `json:"modifiedDate,omitempty"`
	ExternalEmployeeID string `json:"externalEmployeeId,omitempty"`
	Email              string `json:"email"`
}

type PatchEmployee struct {
	FirstName          string `json:"firstName,omitempty"`
	ChosenName         string `json:"chosenName,omitempty"`
	LastName           string `json:"lastName,omitempty"`
	ExternalEmployeeID string `json:"externalEmployeeId,omitempty"`
	PassCode           string `json:"passCode,omitempty"`
}

type RevenueCenters []RevenueCenter

type RevenueCenter struct {
	GUID        string `json:"guid"`
	EntityType  string `json:"entityType"`
	Name        string `json:"name"`        // not in the API documentation
	Description string `json:"description"` // not in the API documentation
	ExternalID  string `json:"externalId"`
}

type AlternativePaymentTypes []AlternativePaymentType

type AlternativePaymentType struct {
	GUID       string      `json:"guid"`
	EntityType string      `json:"entityType"`
	ExternalID interface{} `json:"externalId"`
	Name       string      `json:"name"`
}

type TaxRates []TaxRate

type TaxRate struct {
	GUID                string        `json:"guid"`
	EntityType          string        `json:"entityType"`
	IsDefault           bool          `json:"isDefault"`
	ConditionalTaxRates []interface{} `json:"conditionalTaxRates"`
	TaxTable            []interface{} `json:"taxTable"`
	Rate                float64       `json:"rate"`
	RoundingType        string        `json:"roundingType"`
	Name                string        `json:"name"`
	Type                string        `json:"type"`
}

type Selection struct {
	GUID             string   `json:"guid"`
	EntityType       string   `json:"entityType"`
	ExternalID       string   `json:"externalId"`
	Deferred         *bool    `json:"deferred"`         // switched to pointer
	PreDiscountPrice *float64 `json:"preDiscountPrice"` // switched to pointer
	VoidReason       struct {
		GUID       string `json:"guid"`
		EntityType string `json:"entityType"`
		ExternalID string `json:"externalId"`
	} `json:"voidReason"`
	OptionGroup struct {
		GUID            string `json:"guid"`
		EntityType      string `json:"entityType"`
		ExternalID      string `json:"externalId"`
		MultiLocationID string `json:"multiLocationId"`
	} `json:"optionGroup"`
	DisplayName            string           `json:"displayName"`
	AppliedDiscounts       AppliedDiscounts `json:"appliedDiscounts"`
	SeatNumber             *int             `json:"seatNumber"` // switched to pointer
	VoidDate               DateTime         `json:"voidDate"`
	FulfillmentStatus      string           `json:"fulfillmentStatus"`
	OptionGroupPricingMode string           `json:"optionGroupPricingMode"`
	SalesCategory          struct {
		GUID            string `json:"guid"`
		EntityType      string `json:"entityType"`
		ExternalID      string `json:"externalId"`
		MultiLocationID string `json:"multiLocationId"`
	} `json:"salesCategory"` // used to be SalesCategory but is actually a reference to SalesCategory (only name is missing)
	SelectionType            string       `json:"selectionType"`
	Price                    *float64     `json:"price"`  // switched to pointer
	Voided                   *bool        `json:"voided"` // switched to pointer
	AppliedTaxes             AppliedTaxes `json:"appliedTaxes"`
	StoredValueTransactionID interface{}  `json:"storedValueTransactionId"` // not in the API documentation
	ItemGroup                struct {
		GUID            string `json:"guid"`
		EntityType      string `json:"entityType"`
		ExternalID      string `json:"externalId"`
		MultiLocationID string `json:"multiLocationId"`
	} `json:"itemGroup"`
	Item struct {
		GUID            string `json:"guid"`
		EntityType      string `json:"entityType"`
		ExternalID      string `json:"externalId"`
		MultiLocationID string `json:"multiLocationId"`
	} `json:"item"`
	TaxInclusion     string        `json:"taxInclusion"`
	Quantity         *float64      `json:"quantity"`         // switched to pointer
	ReceiptLinePrice *float64      `json:"receiptLinePrice"` // switched to pointer
	UnitOfMeasure    string        `json:"unitOfMeasure"`
	RefundDetails    RefundDetails `json:"refundDetails"`
	ToastGiftCard    interface{}   `json:"toastGiftCard"` // not in the API documentation
	Tax              *float64      `json:"tax"`           // switched to pointer
	DiningOption     struct {
		GUID       string `json:"guid"`
		EntityType string `json:"entityType"`
		ExternalID string `json:"externalId"`
	} `json:"diningOption"`
	VoidBusinessDate Date     `json:"voidBusinessDate"`
	CreatedDate      DateTime `json:"createdDate"`
	PreModifier      struct {
		GUID            string `json:"guid"`
		EntityType      string `json:"entityType"`
		ExternalID      string `json:"externalId"`
		MultiLocationID string `json:"multiLocationId"`
	} `json:"preModifier"`
	ModifiedDate        DateTime `json:"modifiedDate"`
	ExternalPriceAmount *float64 `json:"externalPriceAmount"` // switched to pointer
	SplitOrigin         struct {
		GUID       string `json:"guid"`
		EntityType string `json:"entityType"`
	} `json:"splitOrigin"`
	OpenPriceAmount *float64 `json:"openPriceAmount"` // switched to pointer
}

type AppliedDiscounts []AppliedDiscount

type AppliedDiscount struct {
	AppliedDiscountReason struct {
		Comment        string `json:"comment"`
		Description    string `json:"description"`
		Name           string `json:"name"`
		DiscountReason struct {
			EntityType string `json:"entityType"`
			GUID       string `json:"guid"`
		} `json:"discountReason"`
	} `json:"appliedDiscountReason"`
	AppliedPromoCode string `json:"appliedPromoCode"`
	Approver         struct {
		EntityType string `json:"entityType"`
		ExternalID string `json:"externalId"`
		GUID       string `json:"guid"`
	} `json:"approver"`
	ComboItems []struct {
		EntityType string `json:"entityType"`
		ExternalID string `json:"externalId"`
		GUID       string `json:"guid"`
	} `json:"comboItems"`
	Discount struct {
		EntityType string `json:"entityType"`
		GUID       string `json:"guid"`
	} `json:"discount"`
	DiscountAmount  *float64 `json:"discountAmount"`  // switched to pointer
	DiscountPercent *float64 `json:"discountPercent"` // switched to pointer
	DiscountType    string   `json:"discountType"`
	EntityType      string   `json:"entityType"`
	ExternalID      string   `json:"externalId"`
	GUID            string   `json:"guid"`
	LoyaltyDetails  struct {
		Vendor      string `json:"vendor"`
		ReferenceID string `json:"referenceId"`
	} `json:"loyaltyDetails"`
	Name                 string   `json:"name"`
	NonTaxDiscountAmount *float64 `json:"nonTaxDiscountAmount"` // switched to pointer
	ProcessingState      string   `json:"processingState"`
	Triggers             []struct {
		Quantity  *float64 `json:"quantity"` // switched to pointer
		Selection struct {
			EntityType string `json:"entityType"`
			ExternalID string `json:"externalId"`
			GUID       string `json:"guid"`
		} `json:"selection"`
	} `json:"triggers"`
}

type AppliedTaxes []AppliedTax

type AppliedTax struct {
	GUID       string `json:"guid"`
	EntityType string `json:"entityType"`
	TaxRate    struct {
		GUID       string `json:"guid"`
		EntityType string `json:"entityType"`
	} `json:"taxRate"`
	Rate                          *float64 `json:"rate"` // switched to pointer
	Name                          string   `json:"name"`
	TaxAmount                     *float64 `json:"taxAmount"` // switched to pointer
	Type                          string   `json:"type"`
	FacilitatorCollectAndRemitTax *bool    `json:"facilitatorCollectAndRemitTax"` // switched to pointer
	Jurisdiction                  string   `json:"jurisdiction"`
	JurisdictionType              string   `json:"jurisdictionType"`
	DisplayName                   string   `json:"displayName"`
}

type RefundDetails struct { // used to be share for both RefundDetails and Refund
	RefundAmount *float64 `json:"refundAmount"` // switched to pointer
	// RefundBusinessDate Date     `json:"refundBusinessDate"`
	// RefundDate         DateTime `json:"refundDate"`
	// RefundStrategy     string   `json:"refundStrategy"`
	RefundTransaction struct {
		EntityType string `json:"entityType"`
		GUID       string `json:"guid"`
	} `json:"refundTransaction"`
	// TipRefundAmount float64 `json:"tipRefundAmount"`
	TaxRefundAmount *float64 `json:"taxRefundAmount"` // switched to pointer
}

type Refund struct {
	RefundAmount       *float64 `json:"refundAmount"` // switched to pointer
	RefundBusinessDate Date     `json:"refundBusinessDate"`
	RefundDate         DateTime `json:"refundDate"`
	RefundTransaction  struct {
		EntityType string `json:"entityType"`
		GUID       string `json:"guid"`
	} `json:"refundTransaction"`
	TipRefundAmount *float64 `json:"tipRefundAmount"` // switched to pointer
}

type AppliedServiceCharges []AppliedServiceCharge

type AppliedServiceCharge struct {
	AppliedTaxes  AppliedTaxes  `json:"appliedTaxes"`
	ChargeAmount  *float64      `json:"chargeAmount"` // switched to pointer
	ChargeType    string        `json:"chargeType"`
	Delivery      *bool         `json:"delivery"`    // switched to pointer
	Destination   string        `json:"destination"` // not in the API documentation
	DineIn        *bool         `json:"dineIn"`      // switched to pointer
	EntityType    string        `json:"entityType"`
	ExternalID    interface{}   `json:"externalId"`
	Gratuity      *bool         `json:"gratuity"` // switched to pointer
	GUID          string        `json:"guid"`
	Name          string        `json:"name"`
	PaymentGUID   interface{}   `json:"paymentGuid"`
	RefundDetails RefundDetails `json:"refundDetails"`
	ServiceCharge struct {
		EntityType string      `json:"entityType"`
		ExternalID interface{} `json:"externalId"`
		GUID       string      `json:"guid"`
	} `json:"serviceCharge"`
	ServiceChargeCalculation string `json:"serviceChargeCalculation"`
	ServiceChargeCategory    string `json:"serviceChargeCategory"`
	Takeout                  *bool  `json:"takeout"` // switched to pointer
	Taxable                  *bool  `json:"taxable"` // switched to pointer
}

type TimeEntries []TimeEntry
type TimeEntry struct {
	GUID         string   `json:"guid"`
	EntityType   string   `json:"entityType"`
	ExternalID   string   `json:"externalId"`
	CreatedDate  DateTime `json:"createdDate"`
	ModifiedDate DateTime `json:"modifiedDate"`
	DeletedDate  DateTime `json:"deletedDate"`
	Deleted      bool     `json:"deleted"`
	JobReference struct {
		GUID       string `json:"guid"`
		EntityType string `json:"entityType"`
		ExternalID string `json:"externalId"`
	} `json:"jobReference"`
	EmployeeReference struct {
		GUID       string `json:"guid"`
		EntityType string `json:"entityType"`
		ExternalID string `json:"externalId"`
	} `json:"employeeReference"`
	ShiftReference struct {
		GUID       string `json:"guid"`
		EntityType string `json:"entityType"`
		ExternalID string `json:"externalId"`
	} `json:"shiftReference"`
	InDate         DateTime `json:"inDate"`
	OutDate        DateTime `json:"outDate"`
	AutoClockedOut bool     `json:"autoClockedOut"`
	BusinessDate   string   `json:"businessDate"`
	RegularHours   float64  `json:"regularHours"`
	OvertimeHours  float64  `json:"overtimeHours"`
	HourlyWage     float64  `json:"hourlyWage"`
	Breaks         []struct {
		GUID      string `json:"guid"`
		BreakType struct {
			GUID       string `json:"guid"`
			EntityType string `json:"entityType"`
		} `json:"breakType"`
		Paid          bool     `json:"paid"`
		InDate        DateTime `json:"inDate"`
		OutDate       DateTime `json:"outDate"`
		Missed        bool     `json:"missed"`
		AuditResponse bool     `json:"auditResponse"`
	} `json:"breaks"`
	DeclaredCashTips              float64 `json:"declaredCashTips"`
	NonCashTips                   float64 `json:"nonCashTips"`
	CashGratuityServiceCharges    float64 `json:"cashGratuityServiceCharges"`
	NonCashGratuityServiceCharges float64 `json:"nonCashGratuityServiceCharges"`
	TipsWithheld                  float64 `json:"tipsWithheld"`
	NonCashSales                  float64 `json:"nonCashSales"`
	CashSales                     float64 `json:"cashSales"`
}

type MenuItemInventories []MenuItemInventory

type MenuItemInventory struct {
	GUID             string   `json:"guid"`
	ItemGUIDValidity string   `json:"itemGuidValidity"`
	MultiLocationID  string   `json:"multiLocationId"`
	Quantity         *float64 `json:"quantity"`
	Status           string   `json:"status"`
	VersionID        string   `json:"versionId"`
}

type MenuOptionGroups []MenuOptionGroup

type MenuOptionGroup struct {
	GUID          string `json:"guid"`
	EntityType    string `json:"entityType"`
	ExternalID    string `json:"externalId"`
	MaxSelections int    `json:"maxSelections"`
	MinSelections int    `json:"minSelections"`
	Name          string `json:"name"`
	Options       []struct {
		GUID       string `json:"guid"`
		EntityType string `json:"entityType"`
		ExternalID string `json:"externalId"`
	} `json:"options"`
}
