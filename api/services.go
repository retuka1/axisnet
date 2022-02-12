package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/adipatiarya/apis/utils"
)

var headers = make(map[string]string)

func init() {

	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Connection"] = "keep-alive"
}

func RequestOtp(msisdn string) (Response, error) {
	var resp Response
	const BASE_URL = "https://otp.api.axis.co.id/otp"

	var content = func(msisdn string) string {
		b, err := json.Marshal(&RequestOtpParams{Msisdn: utils.Encrypt(msisdn, PRIVATE_KEY)})
		if err != nil {
			return ""
		}
		return utils.Encrypt(string(b), PRIVATE_KEY)
	}(msisdn)
	headers["X-API-KEY"] = "ipRsTxOLbgogxEiXj1P73JVWgQk9CEXMSuIsaFO8t6I="
	headers["x-app-version"] = "7.12.1"

	resp, err := fetchData("POST", headers, BASE_URL, content)
	if err != nil {
		return resp, err
	}
	return resp, nil

}

type TokenResult struct {
	Token string `json:"token"`
}

func LoginOtp(msisdn, otp string) (Response, error) {

	var err error

	const BASE_URL = "https://otp.api.axis.co.id/otplogin"

	var content = func(msisdn string) string {
		b, err := json.Marshal(&LoginOtpParams{
			Brand:      "samsung",
			DeviceOs:   "30",
			DeviceType: "android",
			FbId:       "a08b868811f7b868bac73b1516e9cba2",
			FcmToken:   "dM9O5-lOSg-hO_G2HyDTm-:APA91bErnFea4EJEXBV1WX7xGNLoF0xIMJ0psqE7rF4gesjkcR3z_SvMKgqRlwBBZUiNr524aW1r79L6O4mGYryFgCgTMsydx9lshcK6dyqOEcPs-vXesz2wogmv8hdFkat37oWPkj7G",
			Model:      "SM-A505F",
			Msisdn:     utils.Encrypt(msisdn, PRIVATE_KEY),
			OtpCode:    otp,
		})
		if err != nil {
			return ""
		}
		return utils.Encrypt(string(b), PRIVATE_KEY)
	}(msisdn)
	headers["X-API-KEY"] = "ipRsTxOLbgogxEiXj1P73JVWgQk9CEXMSuIsaFO8t6I="
	headers["x-app-version"] = "7.12.1"

	resp, err := fetchData("POST", headers, BASE_URL, content)

	if err != nil {
		return resp, err
	}

	if !resp.Status {
		return resp, nil
	}

	s := utils.Decrypt(resp.Data, PRIVATE_KEY)

	var result map[string]string

	json.Unmarshal([]byte(s), &result)

	resp.Data = result["token"]

	return resp, nil
}

type BuyPackageParams struct {
	ServiceId string `json:"service_id"`
	Type      string `json:"type"`
}

func BuyPackage(token, service_id string) (Response, error) {
	BASE_URL := "https://trxpackages.api.axis.co.id/package/buy/v2"

	payload, _ := json.Marshal(&BuyPackageParams{
		ServiceId: service_id,
		Type:      "PACKAGE",
	})
	headers["Authorization"] = token
	headers["x-app-version"] = "7.12.1"
	content := base64.StdEncoding.EncodeToString(payload)
	resp, err := fetchData("POST", headers, BASE_URL, string(content))
	if err != nil {
		return resp, err
	}
	return resp, nil
}

type UserClaimParam struct {
	ServiceType string `json:"service_type"`
	ServiceId   string `json:"service_id"`
}

func ClaimPackage(token, serviceType, serviceId string) (Response, error) {
	BASE_URL := "https://trxpackages.api.axis.co.id/package/user/claim"

	payload, _ := json.Marshal(&UserClaimParam{
		ServiceType: serviceType,
		ServiceId:   serviceId,
	})
	headers["Authorization"] = token
	headers["x-app-version"] = "7.12.1"
	content := base64.StdEncoding.EncodeToString(payload)

	resp, err := fetchData("POST", headers, BASE_URL, content)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func fetchData(method string, headers map[string]string, baseUrl string, payload string) (Response, error) {
	var err error

	var client = &http.Client{}
	var resp Response

	var param = url.Values{}
	param.Set("content", payload)
	var body = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest(method, baseUrl, body)
	if err != nil {
		return resp, err
	}
	for key, value := range headers {
		request.Header.Set(key, value)
	}

	response, err := client.Do(request)
	if err != nil {
		return resp, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
