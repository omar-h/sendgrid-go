package contactdb

import (
	"encoding/json"
	"errors"

	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
)

// ErrorResponse is the response when you get a non-200 code.
type ErrorResponse struct {
	ErrorCode int        `json:"error_code,omitempty"`
	Errors    []APIError `json:"errors,omitempty"`
}

// APIError is an error which has been returned from the API.
type APIError struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

// Error will fulfill the error interface.
func (e APIError) Error() string {
	return e.Message
}

// sendRequest will send the request to the API.
func sendRequest(apiKey, path, payload string, method rest.Method) ([]byte, error) {
	request := sendgrid.GetRequest(apiKey, path, "https://api.sendgrid.com")
	request.Method = method
	if payload != "" {
		request.Body = []byte(payload)
	}

	response, err := sendgrid.API(request)
	if err != nil {
		return nil, err
	}

	switch response.StatusCode {
	case 200:
		break
	case 201:
		break
	case 202:
		break
	case 204:
		break
	default:
		errRes := &ErrorResponse{}
		err := json.Unmarshal([]byte(response.Body), errRes)
		if err != nil {
			return nil, err
		}

		errRes.ErrorCode = response.StatusCode

		if len(errRes.Errors) <= 0 {
			return nil, errors.New("an unknown API error has occured")
		}

		return nil, errRes.Errors[0]
	}

	return []byte(response.Body), nil
}

// SendGETRequest will send a GET request to the API.
func SendGETRequest(apiKey, path string) ([]byte, error) {
	return sendRequest(apiKey, path, "", rest.Get)
}

// SendPOSTRequest will send a POST request to the API.
func SendPOSTRequest(apiKey, path, payload string) ([]byte, error) {
	return sendRequest(apiKey, path, payload, rest.Post)
}

// SendPATCHRequest will send a PATCH request to the API.
func SendPATCHRequest(apiKey, path, payload string) ([]byte, error) {
	return sendRequest(apiKey, path, payload, rest.Patch)
}

// SendDELETERequest will send a GET request to the API.
func SendDELETERequest(apiKey, path, payload string) ([]byte, error) {
	return sendRequest(apiKey, path, payload, rest.Delete)
}
