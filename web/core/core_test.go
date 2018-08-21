package core

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"steve/common/constant"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_transmitHTTPRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(transmitHTTPRequest))
	defer server.Close()

	serverRequester = func(requestJSON *requestJSONData) (responseJSONData, error) {
		assert.Equal(t, "test-server", requestJSON.Server)
		assert.Equal(t, "test-cmd", requestJSON.Cmd)
		assert.Equal(t, []byte(`{"test-field": "test-data"}`), []byte(requestJSON.Data))
		return responseJSONData{
			Code: constant.HTTPOK,
			Msg:  "test-msg",
			Data: []byte(`{"test-response-field": "test-response-data"}`),
		}, nil
	}
	response, err := http.Post(server.URL, "application/json", bytes.NewReader([]byte(
		`{
			"server": "test-server", 
			"cmd": "test-cmd", 
			"Data": {"test-field": "test-data"}
		}`)))
	assert.Nil(t, err)
	responseData, err := ioutil.ReadAll(response.Body)

	responseJSON := responseJSONData{}
	assert.Nil(t, json.Unmarshal(responseData, &responseJSON))
	assert.Equal(t, int(constant.HTTPOK), responseJSON.Code)
	assert.Equal(t, "test-msg", responseJSON.Msg)

	rawData := map[string]interface{}{}

	assert.Nil(t, json.Unmarshal(responseJSON.Data, &rawData))
	assert.Equal(t, "test-response-data", rawData["test-response-field"])
}
