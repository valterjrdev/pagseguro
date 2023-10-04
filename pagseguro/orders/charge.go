package orders

import (
	"time"
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

func (r *Charge) Endpoint() string {
	return "/charges"
}
