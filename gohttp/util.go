package gohttp

import (
	"encoding/json"
	"encoding/xml"
	"net"
)

//Json转XML
func Json2Xml(jsonString string, value interface{}) (string, error) {
	if err := json.Unmarshal([]byte(jsonString), value); err != nil {
		return "", err
	}
	xml, err := xml.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(xml), nil
}

//XML转Json
func Xml2Json(xmlString string, value interface{}) (string, error) {
	if err := xml.Unmarshal([]byte(xmlString), value); err != nil {
		return "", err
	}
	js, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(js), nil
}

//获取设备IP地址
func GetInterfaceIps() ([]net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}
	ips := make([]net.IP, 0)
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP)
			}
		}
	}
	return ips, nil
}
