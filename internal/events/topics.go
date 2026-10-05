package events

const (
	TopicCollectionSucceeded = "payments.collection.succeeded" // {plan_id, instalment, amount_pence, collected_at}
	TopicPayoutSent          = "payments.payout.sent"          // {claim_id, amount_pence, sent_at}
)
