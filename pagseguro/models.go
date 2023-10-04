package pagseguro

import (
	"net/http"
	"time"
)

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

		http.Client
	}
)

type (
	Summary struct {
		Total    int64 `json:"total,omitempty"`
		Paid     int64 `json:"paid,omitempty"`
		Refunded int64 `json:"refunded,omitempty"`
	}

	Holder struct {
		Name    string  `json:"name,omitempty"`
		TaxID   string  `json:"tax_id,omitempty"`
		Email   string  `json:"email,omitempty"`
		Address Address `json:"address,omitempty"`
	}

	Card struct {
		Holder       Holder `json:"holder,omitempty"`
		Number       string `json:"number,omitempty"`
		ExpMonth     string `json:"exp_month,omitempty"`
		ExpYear      string `json:"exp_year,omitempty"`
		SecurityCode string `json:"security_code,omitempty"`
	}

	InstructionLines struct {
		Line1 string `json:"line_1,omitempty"`
		Line2 string `json:"line_2,omitempty"`
	}

	Boleto struct {
		DueDate          string           `json:"due_date,omitempty"`
		InstructionLines InstructionLines `json:"instruction_lines,omitempty"`
		Holder           Holder           `json:"holder,omitempty"`
	}

	PaymentMethod struct {
		Type           string `json:"type,omitempty"`
		Installments   int64  `json:"installments,omitempty"`
		Capture        bool   `json:"capture,omitempty"`
		SoftDescriptor string `json:"soft_descriptor,omitempty"`
		Card           Card   `json:"card,omitempty"`
		Boleto         Boleto `json:"boleto,omitempty"`
	}

	Links struct {
		Rel   string `json:"rel,omitempty"`
		Href  string `json:"href,omitempty"`
		Media string `json:"media,omitempty"`
		Type  string `json:"type,omitempty"`
	}

	PaymentResponse struct {
		Code    string `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
	}

	Charge struct {
		ID               string            `json:"id,omitempty"`
		ReferenceID      string            `json:"reference_id,omitempty"`
		Description      string            `json:"description,omitempty"`
		Amount           Amount            `json:"amount,omitempty"`
		Status           string            `json:"status,omitempty"`
		PaymentMethod    PaymentMethod     `json:"payment_method,omitempty"`
		PaymentResponse  PaymentResponse   `json:"payment_response,omitempty"`
		Metadata         map[string]string `json:"metadata,omitempty"`
		NotificationUrls []string          `json:"notification_urls,omitempty"`
		Links            []Links           `json:"links,omitempty"`
		NotificationURL  []string          `json:"notification_url,omitempty"`
		CreatedAt        time.Time         `json:"created_at,omitempty"`
		PaidAt           time.Time         `json:"paid_at,omitempty"`
	}
)
