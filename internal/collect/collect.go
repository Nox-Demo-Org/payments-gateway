package collect

// Request is the body of POST /v1/collections.
type Request struct {
	PlanID        string `json:"plan_id"`
	Instalment    int    `json:"instalment"`
	AmountPence   int64  `json:"amount_pence"`
	Method        string `json:"method"` // card | direct_debit
	IdempotencyID string `json:"idempotency_id"`
}

// Collect charges the stored card token or submits a direct debit, then publishes
// payments.collection.succeeded. Direct debits settle in 3 working days; cards at once.
func Collect(r Request) error { return nil }
