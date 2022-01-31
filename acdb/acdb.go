package acdb

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

var acReg []acInfo

func Import(dbUrl string) {

}
