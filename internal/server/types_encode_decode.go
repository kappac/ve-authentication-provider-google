package server

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func EncodeVEValidateTokenRequest(r *http.Request, req interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(req); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func DecodeVEValidateTokenRequest(r *http.Request) (interface{}, error) {
	var req VEValidateTokenRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func EncodeVEValidateTokenResponse(w http.ResponseWriter, req interface{}) error {
	return json.NewEncoder(w).Encode(req)
}

func DecodeVEValidateTokenResponse(res *http.Response) (interface{}, error) {
	var resp VEValidateTokenResponse
	err := json.NewDecoder(res.Body).Decode(&resp)
	return resp, err
}
