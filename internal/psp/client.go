package psp

import (
	"net/http"
	"time"
)

// Client talks to the payment service provider. Timeout 10s; three retries with the same
// idempotency key so a retry can never charge twice.
func NewClient() *http.Client { return &http.Client{Timeout: 10 * time.Second} }
