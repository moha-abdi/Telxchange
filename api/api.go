package api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strconv"

	"github.com/moha-abdi/telxchange/api/types/requests"
	"github.com/moha-abdi/telxchange/api/types/responses"
	"github.com/moha-abdi/telxchange/internal/exchange"
	"github.com/moha-abdi/telxchange/internal/exchange/network"
)

type APIClient struct {
	BaseURL    string
	httpClient *http.Client
	username   string
	sessionId  string
	deviceId   string
}

func NewApiClient(baseURL, username, deviceId string) *APIClient {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile(filepath.Join(homeDir, "golog.log"), os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	return &APIClient{
		BaseURL:  baseURL,
		username: username,
		deviceId: deviceId,
		httpClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					KeyLogWriter: f,
				},
			},
		},
	}
}

// Notice that this method currently does not have any usage
// but it will be used to set common attributes as the
// code logic increases.
func (c *APIClient) setCommonAttributes(req interface{}) {
	v := reflect.ValueOf(req).Elem()
	if v.Kind() == reflect.Struct {
		serviceInfo := v.FieldByName("ServiceInfo")
		if serviceInfo.IsValid() {
			requestAttributes := serviceInfo.FieldByName("RequestAttributes")
			if requestAttributes.IsValid() {
				mobileNo := requestAttributes.FieldByName("MobileNo")
				if mobileNo.IsValid() && mobileNo.CanSet() {
					mobileNo.SetString(c.username)
				}
				deviceId := requestAttributes.FieldByName("DeviceId")
				if deviceId.IsValid() && deviceId.CanSet() {
					deviceId.SetString(c.deviceId)
				}
				sessionId := requestAttributes.FieldByName("SessionId")
				if sessionId.IsValid() && sessionId.CanSet() {
					sessionId.SetString(c.sessionId)
				}
			}
		}
	}
}

func (c *APIClient) doRequest(endpoint string, req interface{}, resp interface{}) error {
	c.setCommonAttributes(req)

	data, err := json.Marshal(req)
	if err != nil {
		return err
	}

	url := c.BaseURL + endpoint
	response, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if strResp, ok := resp.(*string); ok {
		*strResp = string(body)
		return nil
	}

	return json.Unmarshal(body, resp)
}

func (c *APIClient) CreateToken() (string, error) {
	req := requests.NewTokenRequest()
	req.ServiceInfo.RequestAttributes.Username = c.username

	var resp responses.TokenResponse
	err := c.doRequest("/api/createToken", req, &resp)
	if err != nil {
		return "", err
	}

	return resp.ServiceInfo.ResponseAttributes.Token, nil
}

func (c *APIClient) Login(password string) (responses.LoginResponse, error) {
	req := requests.NewLoginRequest()
	req.ServiceInfo.RequestAttributes.Username = c.username
	req.ServiceInfo.RequestAttributes.UserPassword = password
	req.ServiceInfo.RequestAttributes.DeviceId = c.deviceId

	var resp responses.LoginResponse
	err := c.doRequest("/api/login", req, &resp)
	if err == nil {
		c.sessionId = resp.ServiceInfo.ResponseAttributes.SessionId
	}
	return resp, err
}

func (c *APIClient) Request2FA() (responses.VerificationCodeResponse, error) {
	req := requests.NewVerificationCodeRequest()

	var resp responses.VerificationCodeResponse
	err := c.doRequest("/api/requestVerificationCode", req, &resp)
	return resp, err
}

func (c *APIClient) AuthenticateDevice(otpCode string) (responses.AuthDeviceResponse, error) {
	req := requests.NewAuthDeviceRequest()
	req.ServiceInfo.RequestAttributes.OtpCode = otpCode

	var resp responses.AuthDeviceResponse
	err := c.doRequest("/api/authenticateDevice", req, &resp)
	return resp, err
}

func (c *APIClient) GetBalance(accountId string) (responses.BalanceQuertResponse, error) {
	req := requests.NewBalanceQueryRequest()
	req.ServiceInfo.RequestAttributes.AccountId = accountId

	var resp responses.BalanceQuertResponse
	err := c.doRequest("/api/balanceQuery", req, &resp)
	return resp, err
}

func (c *APIClient) GetExchangeRate(partnerID string) (responses.ExchangeRateResponse, error) {
	req := requests.NewExchangeRateRequest()
	req.ServiceInfo.RequestAttributes.PartnerUID = partnerID
	req.ServiceInfo.RequestAttributes.TargetCurrencyCode = exchange.USD.Itoa()
	req.ServiceInfo.RequestAttributes.SourceCurrencyCode = exchange.SLSH.Itoa()

	var resp responses.ExchangeRateResponse
	err := c.doRequest("/api/getMerchantExchangeRateByCurrencyCode", req, &resp)
	return resp, err
}

func (c *APIClient) GetPartnerInfo(
	network network.Network,
	partnerID string,
) (responses.PartnerInfo, *APIError) {
	req := requests.NewPartnerInfoRequest()
	req.ServiceInfo.RequestAttributes.SubPartnerUID = network.Code.Itoa() + partnerID

	var resp responses.PartnerInfoResponse
	err := c.doRequest("/api/getPartnerInfoByUID", req, &resp)
	if err != nil {
		// "-1" is used for runtime internal errors
		return responses.PartnerInfo{}, NewAPIError("-1", err.Error())
	}

	if resp.ServiceInfo.ResponseAttributes.ResultCode != "2001" {
		return responses.PartnerInfo{}, NewAPIError(
			resp.ServiceInfo.ResponseAttributes.ResultCode,
			resp.ServiceInfo.ResponseAttributes.ReplyMessage,
		)
	}
	resp.ServiceInfo.ResponseAttributes.PartnerInfo.ID, _ = strconv.Atoi(partnerID)

	return resp.ServiceInfo.ResponseAttributes.PartnerInfo, nil
}

func (c *APIClient) TestEndpoint(endpoint string) (string, error) {
	req := requests.NewExchangeRateRequest()
	req.ServiceInfo.RequestAttributes.PartnerUID = ""
	req.ServiceInfo.RequestAttributes.TargetCurrencyCode = exchange.USD.Itoa()
	req.ServiceInfo.RequestAttributes.SourceCurrencyCode = exchange.SLSH.Itoa()

	var resp string
	err := c.doRequest(endpoint, req, &resp)
	return resp, err
}
