package models

type TadaOutput struct {
	Products []struct {
		ApplicationFeeWithPaymentMethods struct {
			Corporate  float64 `json:"corporate,omitempty"`
			Creditcard float64 `json:"creditcard"`
			Netsqr     float64 `json:"netsqr"`
			OcbcPao    float64 `json:"ocbc_pao"`
		} `json:"applicationFeeWithPaymentMethods"`
		AvailableOptions []struct {
			Currency      string      `json:"currency"`
			Name          string      `json:"name"`
			OriginalPrice interface{} `json:"originalPrice"`
			Price         int         `json:"price"`
			Type          int         `json:"type"`
		} `json:"availableOptions"`
		AvailablePaymentMethods []string `json:"availablePaymentMethods"`
		CarGroup                int      `json:"carGroup"`
		CarType                 int      `json:"carType"`
		CcApplicationFee        struct {
			Aba                   float64 `json:"aba"`
			Adyen                 float64 `json:"adyen"`
			AdyenAmericanexpress  float64 `json:"adyen_americanexpress"`
			Moovpay               float64 `json:"moovpay"`
			Stripe                float64 `json:"stripe"`
			StripeAmericanexpress float64 `json:"stripe_americanexpress"`
		} `json:"ccApplicationFee"`
		Code                   string `json:"code"`
		Currency               string `json:"currency"`
		CustomScalePointPrices struct {
			ApplicationFeeWithPaymentMethods int `json:"applicationFeeWithPaymentMethods"`
			CcApplicationFee                 int `json:"ccApplicationFee"`
			NetPrice                         int `json:"netPrice"`
			Price                            int `json:"price"`
		} `json:"customScalePointPrices"`
		CustomerPayFee bool `json:"customerPayFee"`
		Description    struct {
			En string `json:"en"`
		} `json:"description"`
		DescriptionLocal     string  `json:"descriptionLocal"`
		Discount             int     `json:"discount"`
		DisplayOrder         int     `json:"displayOrder"`
		DisplayPrice         *string `json:"displayPrice"`
		FareType             int     `json:"fareType"`
		FlexiblePriceEnabled bool    `json:"flexiblePriceEnabled"`
		Meta                 struct {
			AvailablePaymentMethods   string  `json:"available_payment_methods"`
			CustomerPayTransactionFee bool    `json:"customer_pay_transaction_fee"`
			MaxPriceWeight            float64 `json:"max_price_weight,omitempty"`
			MinPriceWeight            float64 `json:"min_price_weight,omitempty"`
		} `json:"meta"`
		MinimumRiderSuggestionPrices struct {
			ABA            int `json:"ABA,omitempty"`
			ALIPAYP        int `json:"ALIPAY_P,omitempty"`
			CASH           int `json:"CASH,omitempty"`
			CMCB           int `json:"CMCB,omitempty"`
			CORPORATE      int `json:"CORPORATE,omitempty"`
			CREDITCARD     int `json:"CREDITCARD,omitempty"`
			CURRENTBALANCE int `json:"CURRENT_BALANCE,omitempty"`
			DELIVERY       int `json:"DELIVERY,omitempty"`
			MOMOP          int `json:"MOMO_P,omitempty"`
			NETSQR         int `json:"NETSQR,omitempty"`
			OCBCPAO        int `json:"OCBC_PAO,omitempty"`
			PAO            int `json:"PAO,omitempty"`
			PIPAYP         int `json:"PIPAY_P,omitempty"`
			TADAPAY        int `json:"TADAPAY,omitempty"`
			WING           int `json:"WING,omitempty"`
		} `json:"minimumRiderSuggestionPrices"`
		Na               bool        `json:"na"`
		NetPrice         float64     `json:"netPrice"`
		Price            float64     `json:"price"`
		ProductType      int         `json:"productType"`
		Region           string      `json:"region"`
		RegionId         int         `json:"regionId"`
		ReservationFee   interface{} `json:"reservationFee"`
		SeatCount        int         `json:"seatCount"`
		VoucherAvailable bool        `json:"voucherAvailable"`
	} `json:"products"`
	Routes struct {
		Field1 struct {
			Distance         int    `json:"distance"`
			DistanceMeter    int    `json:"distanceMeter"`
			Duration         int    `json:"duration"`
			DurationSeconds  int    `json:"durationSeconds"`
			EncodedPolyline  string `json:"encodedPolyline"`
			OverviewPolyline string `json:"overviewPolyline"`
			TollFee          int    `json:"tollFee"`
		} `json:"0"`
	} `json:"routes"`
}
