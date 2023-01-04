package aerodatabox

type AircraftImage struct {
	Url         string `json:"url"`
	WebUrl      string `json:"webUrl"`
	Author      string `json:"author"`
	Title       string `json:"title"`
	Description string `json:"description"`
	License     string `json:"license"`
}
