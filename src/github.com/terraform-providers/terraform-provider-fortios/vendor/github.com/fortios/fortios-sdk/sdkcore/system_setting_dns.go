package forticlient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	// "strconv"
)

// JSONSystemSettingDNS contains the parameters for Create and Update API function
type JSONSystemSettingDNS struct {
	Primary   string `json:"primary"`
	Secondary string `json:"secondary"`
}

// JSONCreateSystemSettingDNSOutput contains the output results for Create API function
type JSONCreateSystemSettingDNSOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// JSONUpdateSystemSettingDNSOutput contains the output results for Update API function
// Attention: Considering scalability, the previous structure and the current structure may change differently
type JSONUpdateSystemSettingDNSOutput struct {
	Vdom       string  `json:"vdom"`
	Mkey       string  `json:"mkey"`
	Status     string  `json:"status"`
	HTTPStatus float64 `json:"http_status"`
}

// CreateSystemSettingDNS API operation for FortiOS
func (c *FortiSDKClient) CreateSystemSettingDNS(params *JSONSystemSettingDNS) (output *JSONCreateSystemSettingDNSOutput, err error) {
	// HTTPMethod := "POST"
	// path := "/api/v2/cmdb/system/dns"
	// output = &JSONCreateSystemSettingDNSOutput{}
	// locJSON, err := json.Marshal(params)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// bytes := bytes.NewBuffer(locJSON)
	// req := c.NewRequest(HTTPMethod, path, nil, bytes)
	// err = req.Send()

	// body, err := ioutil.ReadAll(req.HTTPResponse.Body)

	// var result map[string]interface{}
	// json.Unmarshal([]byte(string(body)), &result)

	// req.HTTPResponse.Body.Close()

	// if result != nil {
	// 	if result["vdom"] != nil {
	// 		output.Vdom = result["vdom"].(string)
	// 	}
	// 	if result["mkey"] != nil {
	// 		output.Mkey = result["mkey"].(string)
	// 	}
	// 	if result["status"] != nil {
	// 		output.Status = result["status"].(string)
	// 	} else {
	// 		err = fmt.Errorf("cannot get the right response")
	// 		return
	// 	}
	// 	if result["http_status"] != nil {
	// 		output.HTTPStatus = result["http_status"].(float64)
	// 	}
	// } else {
	// 	err = fmt.Errorf("cannot get the right response")
	// 	return
	// }

	return
}

// UpdateSystemSettingDNS API operation for FortiOS set the dns server.
// Returns the execution result when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dns chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) UpdateSystemSettingDNS(params *JSONSystemSettingDNS, mkey string) (output *JSONUpdateSystemSettingDNSOutput, err error) {
	HTTPMethod := "PUT"
	path := "/api/v2/cmdb/system/dns"
	// path += "/" + mkey
	output = &JSONUpdateSystemSettingDNSOutput{}
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

// DeleteSystemSettingDNS API operation for FortiOS
func (c *FortiSDKClient) DeleteSystemSettingDNS(mkey string) (err error) {
	// HTTPMethod := "DELETE"
	// path := "/api/v2/cmdb/system/dns"
	// // path += "/" + mkey

	// req := c.NewRequest(HTTPMethod, path, nil, nil)
	// err = req.Send()

	// body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	// log.Printf("FOS-fortios response: %s", string(body))

	// var result map[string]interface{}
	// json.Unmarshal([]byte(string(body)), &result)

	// req.HTTPResponse.Body.Close()

	// if result != nil {
	// 	if result["status"] == nil {
	// 		err = fmt.Errorf("cannot get the right response")
	// 		return
	// 	}
	// } else {
	// 	err = fmt.Errorf("cannot get the right response")
	// 	return
	// }

	return
}

// ReadSystemSettingDNS API operation for FortiOS gets the dns server setting.
// Returns the requested dns server value when the request executes successfully.
// Returns error for service API and SDK errors.
// See the system - dns chapter in the FortiOS Handbook - CLI Reference.
func (c *FortiSDKClient) ReadSystemSettingDNS(mkey string) (output *JSONSystemSettingDNS, err error) {
	HTTPMethod := "GET"
	path := "/api/v2/cmdb/system/dns"
	// path += "/" + mkey

	req := c.NewRequest(HTTPMethod, path, nil, nil)
	err = req.Send()

	body, err := ioutil.ReadAll(req.HTTPResponse.Body)
	log.Printf("FOS-fortios reading response: %s", string(body))

	output = &JSONSystemSettingDNS{}
	var result map[string]interface{}
	json.Unmarshal([]byte(string(body)), &result)

	req.HTTPResponse.Body.Close()

	if result != nil {
		if result["status"] == nil {
			err = fmt.Errorf("cannot get the right response")
			return
		}

		mapTmp := (result["results"].(map[string]interface{}))

		if mapTmp == nil {
			return
		}

		if mapTmp["primary"] != nil {
			output.Primary = mapTmp["primary"].(string)
		}
		if mapTmp["secondary"] != nil {
			output.Secondary = mapTmp["secondary"].(string)
		}

	} else {
		err = fmt.Errorf("cannot get the right response")
		return
	}

	return
}
