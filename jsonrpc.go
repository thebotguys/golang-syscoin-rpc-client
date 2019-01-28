package syscoinrpc

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/twinj/uuid"
)

// jsonRPCrequest represents a generic JSON RPC request.
type jsonRPCrequest struct {
	JSONRpcVersion string      `json:"jsonrpc,required"`
	Method         string      `json:"method,required"`
	Params         interface{} `json:"params,required"`
	ID             string      `json:"id,required"`
}

// jsonRPCrequest represents a generic JSON RPC response.
type jsonRPCresponse struct {
	Result json.RawMessage `json:"result,required"`
	Error  *errorMessage   `json:"error"`
	ID     string          `json:"id,required"`
}

type errorMessage struct {
	Code    int    `json:"code,required"`
	Message string `json:"message,required"`
}

func (err *errorMessage) Error() string {
	return err.Message
}

// do performs a JSON RPC Call.
//     url   : The endpoint url.
//     method: The name of the method which is going to be called.
//     params: The JSON object representing all the params.
func (c *Client) do(method string, params ...interface{} /*json.Marshaler*/) (json.RawMessage, error) {
	jsonReq := jsonRPCrequest{
		JSONRpcVersion: "1.0",
		Method:         method,
		Params:         params,
		ID:             uuid.NewV4().String(),
	}

	reqBody, err := json.Marshal(jsonReq)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.user, c.pass)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var jsonResp jsonRPCresponse
	err = json.Unmarshal(content, &jsonResp)
	if err != nil {
		return nil, err
	}
	if jsonResp.Error != nil {
		return nil, jsonResp.Error
	}

	return jsonResp.Result, nil
}
