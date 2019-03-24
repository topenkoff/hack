package main

type SendWarningRequest struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Image     string  `json:"image"`
	SectorID  int     `json:"sector_id"`
}

type GetSectorByCoordinatesRequest struct {
	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`
}

type GetWarningsRequest struct {
	SectorID int `json:"sector_id"`
}

type DeleteWarningRequest struct {
	WarningID int `json:"id"`
}

type GetNumbersRequest struct {
	Letter string `json:"letter"`
}

type CreateHashTagRequest struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	CountryID int     `json:"country_id"`
}

type CreateOrUpdateSessionRequest struct {
	Session   string  `json:"session"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
