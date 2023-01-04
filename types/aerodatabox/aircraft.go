package aerodatabox

type Aircraft struct {
	Id               int           `json:"id"`
	Reg              string        `json:"reg"`
	Active           bool          `json:"active"`
	Serial           string        `json:"serial"`
	HexIcao          string        `json:"hexIcao"`
	AirlineName      string        `json:"airlineName"`
	IataCodeShort    string        `json:"iataCodeShort"`
	IcaoCode         string        `json:"icaoCode"`
	Model            string        `json:"model"`
	ModelCode        string        `json:"modelCode"`
	NumSeats         int           `json:"numSeats"`
	RolloutDate      string        `json:"rolloutDate"`
	FirstFlightDate  string        `json:"firstFlightDate"`
	DeliveryData     string        `json:"deliveryDate"`
	RegistrationDate string        `json:"registrationDate"`
	TypeName         string        `json:"typeName"`
	NumEngines       int           `json:"numEngines"`
	EngineType       string        `json:"engineType"`
	IsFreighter      bool          `json:"isFreighter"`
	ProductionLine   string        `json:"productionLine"`
	AgeYears         float32       `json:"ageYears"`
	Verified         bool          `json:"verified"`
	NumRegistration  int           `json:"numRegistrations"`
	Image            AircraftImage `json:"image"`
}
