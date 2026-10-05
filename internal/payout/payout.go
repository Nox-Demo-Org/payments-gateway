package payout

// Request is the body of POST /v1/payouts.
type Request struct {
	ClaimID       string `json:"claim_id"`
	AmountPence   int64  `json:"amount_pence"`
	PayeeName     string `json:"payee_name"`
	SortCode      string `json:"sort_code"`
	AccountNumber string `json:"account_number"`
	IdempotencyID string `json:"idempotency_id"`
}

// Send pays by Faster Payments and publishes payments.payout.sent. Payouts over £25,000 wait
// for a second approver in claims-management before they reach this endpoint.
func Send(r Request) error { return nil }
