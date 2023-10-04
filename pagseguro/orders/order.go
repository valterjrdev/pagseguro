package orders

type (
	Amount struct {
		Value    int64   `json:"value,omitempty"`
		Currency string  `json:"currency,omitempty"`
		Summary  Summary `json:"summary,omitempty"`
	}

	Address struct {
		Street     string `json:"street,omitempty"`
		Number     string `json:"number,omitempty"`
		Locality   string `json:"locality,omitempty"`
		City       string `json:"city,omitempty"`
		Region     string `json:"region,omitempty"`
		RegionCode string `json:"region_code,omitempty"`
		Country    string `json:"country,omitempty"`
		PostalCode string `json:"postal_code,omitempty"`
	}

	Phone struct {
		Country string `json:"country,omitempty"`
		Area    string `json:"area,omitempty"`
		Number  string `json:"number,omitempty"`
		Type    string `json:"type,omitempty"`
	}

	Shipping struct {
		Address Address `json:"address,omitempty"`
	}

	Billing struct {
		Address Address `json:"address,omitempty"`
	}

	QrCodes struct {
		Amount Amount `json:"amount,omitempty"`
	}

	Customer struct {
		Name   string  `json:"name,omitempty"`
		Email  string  `json:"email,omitempty"`
		TaxID  string  `json:"tax_id,omitempty"`
		Phones []Phone `json:"phones,omitempty"`
	}

	Item struct {
		ReferenceID string `json:"reference_id,omitempty"`
		Name        string `json:"name,omitempty"`
		Quantity    int    `json:"quantity,omitempty"`
		UnitAmount  int    `json:"unit_amount,omitempty"`
	}

	Order struct {
		ReferenceID      string    `json:"reference_id,omitempty"`
		Customer         Customer  `json:"customer,omitempty"`
		Shipping         Shipping  `json:"shipping,omitempty"`
		Billing          Billing   `json:"billing,omitempty"`
		Charges          []Charge  `json:"charges,omitempty"`
		Items            []Item    `json:"items,omitempty"`
		QrCodes          []QrCodes `json:"qr_codes,omitempty"`
		NotificationUrls []string  `json:"notification_urls,omitempty"`
	}
)

func (o *Order) Endpoint() string {
	return "/orders"
}
