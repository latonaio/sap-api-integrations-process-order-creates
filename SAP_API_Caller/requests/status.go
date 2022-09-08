package requests

type Status struct {
	ManufacturingOrder string  `json:"ManufacturingOrder"`
	StatusCode         string  `json:"StatusCode"`
	IsUserStatus       bool    `json:"IsUserStatus"`
	StatusShortName    *string `json:"StatusShortName"`
	StatusName         *string `json:"StatusName"`
}
