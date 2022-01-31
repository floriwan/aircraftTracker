package acdb

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

type acInfo struct {
	Icao         string `json:"icao"`
	Reg          string `json:"reg"`
	IcaoType     string `json:"icaotype"`
	Year         string `json:"year"`
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	Owneop       string `json:"ownop"`
	FaaPia       bool   `json:"faa_pia"`
	FaaLadd      bool   `json:"faa_ladd"`
	ShortType    string `json:"shor_type"`
	Mil          bool   `json:"mil"`
}

var acReg map[string]acInfo

func GetAcInfo(reg string) (acInfo, error) {

	if ac, ok := acReg[reg]; ok {
		return ac, nil
	}

	return acInfo{}, fmt.Errorf("no aircraft with registration %v found", reg)

}

func Import(r io.Reader) {
	log.Printf("import new aircraft data ...")
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)

	acReg = make(map[string]acInfo)

	for scanner.Scan() {
		var data acInfo
		dataText := strings.Replace(scanner.Text(), "\\\\\"", "", -1)
		if err := json.Unmarshal([]byte(dataText), &data); err != nil {
			log.Fatalf("%v\n%v", err, dataText)
		}
		acReg[data.Reg] = data
	}

	log.Printf("%v aircrafts imported", len(acReg))

}
