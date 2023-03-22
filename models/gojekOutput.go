package models

type T struct {
	Data struct {
		BaseToken              string      `json:"base_token"`
		BusinessBooking        interface{} `json:"business_booking"`
		Currency               string      `json:"currency"`
		DestinationServiceArea struct {
			CountryCode string `json:"country_code"`
			Currency    string `json:"currency"`
			Id          int    `json:"id"`
			TzName      string `json:"tz_name"`
		} `json:"destination_service_area"`
		Disclaimer struct {
			Message string `json:"message"`
		} `json:"disclaimer"`
		OriginServiceArea struct {
			CountryCode string `json:"country_code"`
			Currency    string `json:"currency"`
			Id          int    `json:"id"`
			TzName      string `json:"tz_name"`
		} `json:"origin_service_area"`
		ServiceTypes []struct {
			AllocationStrategies []struct {
				Type string `json:"type"`
			} `json:"allocation_strategies"`
			AllocationStrategy struct {
				Type string `json:"type"`
			} `json:"allocation_strategy"`
			Errors         interface{} `json:"errors"`
			Id             int         `json:"id"`
			PaymentMethods []struct {
				BaseFareWithVoucher    int  `json:"base_fare_with_voucher"`
				BaseFareWithoutVoucher int  `json:"base_fare_without_voucher"`
				Id                     int  `json:"id"`
				MaxPriceWithVoucher    int  `json:"max_price_with_voucher"`
				MaxPriceWithoutVoucher int  `json:"max_price_without_voucher"`
				MinPriceWithVoucher    int  `json:"min_price_with_voucher"`
				MinPriceWithoutVoucher int  `json:"min_price_without_voucher"`
				PlatformFee            int  `json:"platform_fee"`
				PriceWithVoucher       int  `json:"price_with_voucher"`
				PriceWithoutVoucher    int  `json:"price_without_voucher"`
				Surge                  bool `json:"surge"`
				PriceRange             struct {
					SubtractWithVoucher    int `json:"subtract_with_voucher"`
					SubtractWithoutVoucher int `json:"subtract_without_voucher"`
				} `json:"price_range,omitempty"`
			} `json:"payment_methods"`
			PickupEta    string `json:"pickup_eta"`
			PricingToken string `json:"pricing_token"`
			Route        struct {
				DistanceInMeters int    `json:"distance_in_meters"`
				Eta              int    `json:"eta"`
				Polyline         string `json:"polyline"`
				RouteType        string `json:"route_type"`
			} `json:"route"`
			ServiceProvider    string `json:"service_provider"`
			ServiceTypeDetails struct {
				Description        string `json:"description"`
				ShouldShowBenefits bool   `json:"should_show_benefits"`
			} `json:"service_type_details"`
			Success     string `json:"success"`
			FareWeblink string `json:"fare_weblink,omitempty"`
		} `json:"service_types"`
		TermsAndConditions interface{} `json:"terms_and_conditions"`
		VoucherFlow        string      `json:"voucher_flow"`
	} `json:"data"`
	Errors  interface{} `json:"errors"`
	Success bool        `json:"success"`
}
