package go_thunder

import (
	"bytes"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type FwTcpResetOnError struct {
	Enable FwTcpResetOnErrorInstance `json:"reset-on-error-instance,omitempty"`
}

type FwTcpResetOnErrorInstance struct {
	Enable int    `json:"enable,omitempty"`
	UUID   string `json:"uuid,omitempty"`
}

func PostFwTcpResetOnError(id string, inst FwTcpResetOnError, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostFwTcpResetOnError")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/fw/tcp/reset-on-error", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwTcpResetOnError
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostFwTcpResetOnError REQ RES..........................", m)
			err := check_api_status("PostFwTcpResetOnError", data)
			if err != nil {
				return err
			}

		}
	}
	return err
}

func GetFwTcpResetOnError(id string, host string) (*FwTcpResetOnError, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetFwTcpResetOnError")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/fw/tcp/reset-on-error/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m FwTcpResetOnError
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetFwTcpResetOnError REQ RES..........................", m)
			err := check_api_status("GetFwTcpResetOnError", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
