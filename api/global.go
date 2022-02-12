package api

type RequestOtpParams struct {
	Msisdn string `json:"msisdn"`
}
type LoginOtpParams struct {
	Brand      string `json:"brand"`
	DeviceOs   string `json:"device_os"`
	DeviceType string `json:"device_type"`
	FbId       string `json:"fb_id"`
	FcmToken   string `json:"fcm_token"`
	Model      string `json:"model"`
	Msisdn     string `json:"msisdn"`
	OtpCode    string `json:"otp_code"`
}

const PRIVATE_KEY = "ae432353a4eb4eb9"

type Response struct {
	Status       bool   `json:"status"`
	StatusCode   int    `json:"status_code"`
	ErrorMessage string `json:"error_message"`
	Data         string `json:"data"`
}
