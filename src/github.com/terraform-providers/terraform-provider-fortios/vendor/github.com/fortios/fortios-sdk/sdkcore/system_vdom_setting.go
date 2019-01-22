package forticlient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"fmt"
	"strconv"
)

// JSONSystemVdomSetting contains the parameters for Create and Update API function
type JSONSystemVdomSetting struct {
	Name       string `json:"name"`
	ShortName       string `json:"short-name"`
	Temporary       string `json:"temporary"`
}

// JSONCreateSystemVdomSettingOutput contains the output results for Create API function
type JSONCreateSystemVdomSettingOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// JSONUpdateSystemVdomSettingOutput contains the output results for Update API function
// Attention: Considering scalability, the previous structure and the current structure may change differently
type JSONUpdateSystemVdomSettingOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// CreateSystemVdomSetting API operation for FortiOS creates a new vdom.
// Returns the index value of the vdom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) CreateSystemVdomSetting(params *JSONSystemVdomSetting) (output *JSONCreateSystemVdomSettingOutput, err error) {
	HTTPMethod := "POST"
	path := "/api/v2/cmdb/system/vdom"
	output = &JSONCreateSystemVdomSettingOutput{}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)
	req := c.NewRequest(HTTPMethod, path, nil, bytes)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["vdom"] != nil {
			output.Vdom = result["vdom"].(string)
		}
		if result["mkey"] != nil {
			output.Mkey = result["mkey"].(string)
		}
		if result["status"] != nil {
			output.Status = result["status"].(string)
		} else {
			err = fmt.Errorf("cannot get the right response")
			return
		}
		if result["http_status"] != nil {
			output.HTTPStatus = result["http_status"].(float64)
		}
	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}

// UpdateSystemVdomSetting API operation for FortiOS updates the specified vdom.
// Returns the index value of the vdom and execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemVdomSetting(params *JSONSystemVdomSetting, mkey string) (output *JSONUpdateSystemVdomSettingOutput, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/vdom"
	path += "/" + mkey
	output = &JSONUpdateSystemVdomSettingOutput{}
	locJSON, err := json.Marshal(params)
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes := bytes.NewBuffer(locJSON)
	req := c.NewRequest(HTTPMethod, path, nil, bytes)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	log.Printf("FOS-fortios response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["vdom"] != nil {
			output.Vdom = result["vdom"].(string)
		}
		if result["mkey"] != nil {
			output.Mkey = result["mkey"].(string)
		}
		if result["status"] != nil {
			output.Status = result["status"].(string)
		} else {
			err = fmt.Errorf("cannot get the right response")
			return
		}
		if result["http_status"] != nil {
			output.HTTPStatus = result["http_status"].(float64)
		}
	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}

// DeleteSystemVdomSetting API operation for FortiOS deletes the specified vdom.
// Returns error for service API and SDK errors.
// See the system - vdom chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) DeleteSystemVdomSetting(mkey string) (err error) {
	HTTPMethod := "DELETE"
	path := "/api/v2/cmdb/system/vdom"
	path += "/" + mkey

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	log.Printf("FOS-fortios response: %s", string(body))

	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["status"] == nil {
			err = fmt.Errorf("cannot get the right response")
			return
		}
	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}

// ReadSystemVdomSetting API operation for FortiOS gets the vdom
// with the specified index value.
// Returns the requested vdom value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - vdom chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemVdomSetting(mkey string) (output *JSONSystemVdomSetting, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/vdom"
	path += "/" + mkey

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	log.Printf("FOS-fortios reading response: %s", string(body))

	output = &JSONSystemVdomSetting{}
	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["status"] == nil {
			err = fmt.Errorf("cannot get the right response")
			return
		}

		mapTmp := (result["results"].([]interface {}))[0].(map[string]interface {})

		if mapTmp == nil {
			return
		}

		if mapTmp["name"] != nil {
			output.Name = mapTmp["name"].(string)
		}
		if mapTmp["short-name"] != nil {
			output.ShortName = mapTmp["short-name"].(string)
		}
		if mapTmp["temporary"] != nil {
			output.Temporary = strconv.Itoa(int(mapTmp["temporary"].(float64)))
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}
