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
	GUID          string      `json:"guid"`
	EntityType    string      `json:"entityType"`
	ExternalID    interface{} `json:"externalId"`
	RevenueCenter struct {
		GUID       string      `json:"guid"`
		EntityType string      `json:"entityType"`
		ExternalID interface{} `json:"externalId"`
	} `json:"revenueCenter"`
	Server struct {
		GUID       string      `json:"guid"`
		EntityType string      `json:"entityType"`
		ExternalID interface{} `json:"externalId"`
	} `json:"server"`
	LastModifiedDevice struct {
		ID string `json:"id"`
	} `json:"lastModifiedDevice"`
	Source            string   `json:"source"`
	VoidDate          DateTime `json:"voidDate"`
	Duration          int      `json:"duration"`
	BusinessDate      Date     `json:"businessDate"`
	PaidDate          DateTime `json:"paidDate"`
	RestaurantService struct {
		GUID       string      `json:"guid"`
		EntityType string      `json:"entityType"`
		ExternalID interface{} `json:"externalId"`
	} `json:"restaurantService"`
	Voided                   bool        `json:"voided"`
	EstimatedFulfillmentDate DateTime    `json:"estimatedFulfillmentDate"`
	Table                    interface{} `json:"table"`
	RequiredPrepTime         string      `json:"requiredPrepTime"`
	ApprovalStatus           string      `json:"approvalStatus"`
	DeliveryInfo             interface{} `json:"deliveryInfo"`
	ServiceArea              interface{} `json:"serviceArea"`
	CurbsidePickupInfo       interface{} `json:"curbsidePickupInfo"`
	NumberOfGuests           int         `json:"numberOfGuests"`
	DiningOption             interface{} `json:"diningOption"`
	OpenedDate               DateTime    `json:"openedDate"`
	VoidBusinessDate         Date        `json:"voidBusinessDate"`
	Checks                   []struct {
		GUID          string      `json:"guid"`
		EntityType    string      `json:"entityType"`
		ExternalID    interface{} `json:"externalId"`
		DisplayNumber string      `json:"displayNumber"`
		Payments      []struct {
			GUID                  string      `json:"guid"`
			EntityType            string      `json:"entityType"`
			ExternalID            interface{} `json:"externalId"`
			OriginalProcessingFee interface{} `json:"originalProcessingFee"`
			Server                struct {
				GUID       string      `json:"guid"`
				EntityType string      `json:"entityType"`
				ExternalID interface{} `json:"externalId"`
			} `json:"server"`
			CashDrawer         interface{} `json:"cashDrawer"`
			LastModifiedDevice struct {
				ID string `json:"id"`
			} `json:"lastModifiedDevice"`
			RefundStatus       string      `json:"refundStatus"`
			IsProcessedOffline interface{} `json:"isProcessedOffline"`
			Type               string      `json:"type"`
			VoidInfo           interface{} `json:"voidInfo"`
			CheckGUID          string      `json:"checkGuid"`
			OtherPayment       struct {
				GUID       string      `json:"guid"`
				EntityType string      `json:"entityType"`
				ExternalID interface{} `json:"externalId"`
			} `json:"otherPayment"`
			PaidDate           DateTime    `json:"paidDate"`
			OrderGUID          string      `json:"orderGuid"`
			CardEntryMode      interface{} `json:"cardEntryMode"`
			PaymentStatus      string      `json:"paymentStatus"`
			Amount             float64     `json:"amount"`
			TipAmount          float64     `json:"tipAmount"`
			AmountTendered     float64     `json:"amountTendered"`
			CardType           interface{} `json:"cardType"`
			HouseAccount       interface{} `json:"houseAccount"`
			McaRepaymentAmount interface{} `json:"mcaRepaymentAmount"`
			CreatedDevice      struct {
				ID string `json:"id"`
			} `json:"createdDevice"`
			PaidBusinessDate Date        `json:"paidBusinessDate"`
			Last4Digits      interface{} `json:"last4Digits"`
			Refund           interface{} `json:"refund"`
		} `json:"payments"`
		AppliedDiscounts   []interface{} `json:"appliedDiscounts"`
		LastModifiedDevice struct {
			ID string `json:"id"`
		} `json:"lastModifiedDevice"`
		VoidDate           DateTime    `json:"voidDate"`
		PaidDate           DateTime    `json:"paidDate"`
		AppliedLoyaltyInfo interface{} `json:"appliedLoyaltyInfo"`
		Voided             bool        `json:"voided"`
		PaymentStatus      string      `json:"paymentStatus"`
		Amount             float64     `json:"amount"`
		TabName            string      `json:"tabName"`
		TaxExempt          bool        `json:"taxExempt"`
		OpenedDate         DateTime    `json:"openedDate"`
		TotalAmount        float64     `json:"totalAmount"`
		Selections         []struct {
			GUID                   string        `json:"guid"`
			EntityType             string        `json:"entityType"`
			ExternalID             interface{}   `json:"externalId"`
			Deferred               bool          `json:"deferred"`
			PreDiscountPrice       float64       `json:"preDiscountPrice"`
			VoidReason             interface{}   `json:"voidReason"`
			OptionGroup            interface{}   `json:"optionGroup"`
			DisplayName            string        `json:"displayName"`
			AppliedDiscounts       []interface{} `json:"appliedDiscounts"`
			Modifiers              []interface{} `json:"modifiers"`
			SeatNumber             int           `json:"seatNumber"`
			VoidDate               DateTime      `json:"voidDate"`
			FulfillmentStatus      string        `json:"fulfillmentStatus"`
			OptionGroupPricingMode interface{}   `json:"optionGroupPricingMode"`
			SalesCategory          struct {
				GUID       string      `json:"guid"`
				EntityType string      `json:"entityType"`
				ExternalID interface{} `json:"externalId"`
			} `json:"salesCategory"`
			SelectionType string  `json:"selectionType"`
			Price         float64 `json:"price"`
			Voided        bool    `json:"voided"`
			AppliedTaxes  []struct {
				GUID       string `json:"guid"`
				EntityType string `json:"entityType"`
				TaxRate    struct {
					GUID       string `json:"guid"`
					EntityType string `json:"entityType"`
				} `json:"taxRate"`
				Rate                          float64 `json:"rate"`
				Name                          string  `json:"name"`
				TaxAmount                     float64 `json:"taxAmount"`
				Type                          string  `json:"type"`
				FacilitatorCollectAndRemitTax bool    `json:"facilitatorCollectAndRemitTax"`
			} `json:"appliedTaxes"`
			StoredValueTransactionID interface{} `json:"storedValueTransactionId"`
			ItemGroup                struct {
				GUID            string      `json:"guid"`
				EntityType      string      `json:"entityType"`
				ExternalID      interface{} `json:"externalId"`
				MultiLocationID string      `json:"multiLocationId"`
			} `json:"itemGroup"`
			Item struct {
				GUID            string      `json:"guid"`
				EntityType      string      `json:"entityType"`
				ExternalID      interface{} `json:"externalId"`
				MultiLocationID string      `json:"multiLocationId"`
			} `json:"item"`
			TaxInclusion     string      `json:"taxInclusion"`
			Quantity         float64     `json:"quantity"`
			ReceiptLinePrice float64     `json:"receiptLinePrice"`
			UnitOfMeasure    string      `json:"unitOfMeasure"`
			RefundDetails    interface{} `json:"refundDetails"`
			ToastGiftCard    interface{} `json:"toastGiftCard"`
			Tax              float64     `json:"tax"`
			DiningOption     interface{} `json:"diningOption"`
			VoidBusinessDate Date        `json:"voidBusinessDate"`
			CreatedDate      DateTime    `json:"createdDate"`
			PreModifier      interface{} `json:"preModifier"`
			ModifiedDate     DateTime    `json:"modifiedDate"`
		} `json:"selections"`
		VoidBusinessDate Date     `json:"voidBusinessDate"`
		CreatedDate      DateTime `json:"createdDate"`
		Deleted          bool     `json:"deleted"`
		CreatedDevice    struct {
			ID string `json:"id"`
		} `json:"createdDevice"`
		ClosedDate            DateTime      `json:"closedDate"`
		DeletedDate           DateTime      `json:"deletedDate"`
		ModifiedDate          DateTime      `json:"modifiedDate"`
		TaxAmount             float64       `json:"taxAmount"`
		AppliedServiceCharges []interface{} `json:"appliedServiceCharges"`
		Customer              interface{}   `json:"customer"`
	} `json:"checks"`
	Deleted       bool `json:"deleted"`
	CreatedDevice struct {
		ID string `json:"id"`
	} `json:"createdDevice"`
	CreatedDate     DateTime    `json:"createdDate"`
	ClosedDate      DateTime    `json:"closedDate"`
	DeletedDate     DateTime    `json:"deletedDate"`
	ModifiedDate    DateTime    `json:"modifiedDate"`
	PromisedDate    Date        `json:"promisedDate"`
	ChannelGUID     interface{} `json:"channelGuid"`
	PricingFeatures []string    `json:"pricingFeatures"`
}

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
		Address1  string  `json:"address1"`
		Address2  string  `json:"address2"`
		City      string  `json:"city"`
		StateCode string  `json:"stateCode"`
		ZipCode   string  `json:"zipCode"`
		Country   string  `json:"country"`
		Phone     string  `json:"phone"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"location"`
	Schedules struct {
		DaySchedules struct {
			Num600000000314630485 struct {
				ScheduleName string `json:"scheduleName"`
				Services     []struct {
					Name  string `json:"name"`
					Hours struct {
						StartTime string `json:"startTime"`
						EndTime   string `json:"endTime"`
					} `json:"hours"`
					Overnight bool `json:"overnight"`
				} `json:"services"`
				OpenTime  string `json:"openTime"`
				CloseTime string `json:"closeTime"`
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
