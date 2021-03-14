package masscan

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
)

type (
	XMLRun struct {
		Scanner       string   `xml:"scanner,attr" json:"scanner"`
		Start         string   `xml:"start,attr" json:"start"`
		Version       string   `xml:"version,attr" json:"version"`
		OutputVersion string   `xml:"xmloutputversion,attr" json:"outputversion"`
		ScanInfo      ScanInfo `xml:"scaninfo" json:"scaninfo"`
		Host          []Host   `xml:"host" json:"host"`
	}
	ScanInfo struct {
		Type  string `xml:"type,attr" json:"type"`
		Proto string `xml:"protocol,attr" json:"proto"`
	}
	Address struct {
		Addr     string `xml:"addr,attr" json:"addr"`
		AddrType string `xml:"addrtype,attr" json:"addrtype"`
	}
	State struct {
		State     string `xml:"state,attr" json:"state"`
		Reason    string `xml:"reason,attr" json:"reason"`
		ReasonTTL string `xml:"reason_ttl,attr" json:"reasonttl"`
	}
	Host struct {
		Endtime string  `xml:"endtime,attr" json:"endtime"`
		Address Address `xml:"address" json:"address"`
		Ports   []Port  `xml:"ports>port" json:"ports"`
	}
	Port struct {
		Protocol string  `xml:"protocol,attr" json:"proto"`
		PortId   string  `xml:"portid,attr" json:"port"`
		State    State   `xml:"state" json:"state"`
		Service  Service `xml:"service" json:"service"`
	}
	Service struct {
		Name   string `xml:"name,attr" json:"name"`
		Banner string `xml:"banner,attr" json:"banner"`
	}
	JsonRun struct {
		IP        string   `json:"ip"`
		Timestamp string   `json:"timestamp"`
		Port      JsonPort `json:"ports"`
	}
	JsonPort struct {
		PortId   uint16 `json:"port"`
		Protocol string `json:"proto"`
		Status   string `json:"status"`
		Reason   string `json:"reason"`
		TTL      uint8  `json:"ttl"`
	}
)

func ParseXML(filePath string) XMLRun {
	r := XMLRun{}
	dat, _ := ioutil.ReadFile(filePath)
	xml.Unmarshal(dat, &r)
	return r
}

func ParseJSON(filePath string) []JsonRun {
	r := []JsonRun{}
	dat, _ := ioutil.ReadFile(filePath)
	json.Unmarshal(dat, &r)
	return r
}
