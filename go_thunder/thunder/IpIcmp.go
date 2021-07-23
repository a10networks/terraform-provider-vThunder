package go_thunder

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"util"
)

type Icmp struct {
	UUID IcmpInstance `json:"icmp,omitempty"`
}

type IcmpInstance struct {
	Redirect    int    `json:"redirect,omitempty"`
	Unreachable int    `json:"unreachable,omitempty"`
	UUID        string `json:"uuid,omitempty"`
}

func PostIpIcmp(id string, inst Icmp, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside PostIpIcmp")
	payloadBytes, err := json.Marshal(inst)
	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))
	if err != nil {
		logger.Println("[INFO] Marshalling failed with error ", err)
	}

	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/ip/icmp", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Icmp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)

		} else {
			logger.Println("[INFO] PostIpIcmp REQ RES..........................", m)
			err := check_api_status("PostIpIcmp", data)
			if err != nil {
				return err
			}
		}
	}
return err
}

func GetIpIcmp(id string, host string) (*Icmp, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id
	logger.Println("[INFO] Inside GetIpIcmp")

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/ip/icmp/", nil, headers)

	if err != nil {
		logger.Println("The HTTP request failed with error ", err)
		return nil, err

	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m Icmp
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			logger.Println("Unmarshal error ", err)
			return nil, err
		} else {
			logger.Println("[INFO] GetIpIcmp REQ RES..........................", m)
			err := check_api_status("GetIpIcmp", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}

}
