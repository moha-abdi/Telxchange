package requests

type LocationInformation struct {
	CellID    string `json:"cellId"`
	LACId     string `json:"LACId"`
	MNC       string `json:"MNC"`
	MCC       string `json:"MCC"`
	IP        string `json:"IP"`
	MAC       string `json:"MAC"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}
