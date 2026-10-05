# payments-gateway

Moves money for Tidewell Mutual: collects premiums by card and direct debit, and pays out settled claims by bank transfer. Talks to our payment service provider (PSP). Holds no card numbers: the PSP returns a token.

Owned by **Billing & Payments**. On-call: `#tw-billing`.

| Contract | Kind | Consumers |
| --- | --- | --- |
| `POST /v1/collections` | REST | billing-service |
| `POST /v1/payouts` | REST | claims-management |
| `payments.collection.succeeded` | Event | billing-service |
| `payments.payout.sent` | Event | claims-management, notifications-hub |

Every request carries an `idempotency_id`. A repeat with the same id returns the first result and never moves money twice.
