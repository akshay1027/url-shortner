package types

import "time"

// type Response struct {
// 	Status  int        `json:"status"`
// 	Success bool       `json:"success"`
// 	Message string     `json:"message"`
// 	Data    *fiber.Map `json:"data"`
// }

type Request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type Response struct {
	Success         bool          `json:"success"`
	URL             string        `json:"url"`
	CustomShort     string        `json:"short"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"rate_limit"`
	XRateLimitReset time.Duration `json:"rate_limit_reset"`
}
