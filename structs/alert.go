package structs

// Alert represents a flood or rain alert for a province.
type Alert struct {
	ID           string `json:"id"`
	ProvinceID   int    `json:"province_id"`
	ProvinceName string `json:"province_name"`
	Type         string `json:"type"`        // FlashFlood | RiverFlood | HeavyRain | StormSurge | Landslide
	Level        string `json:"level"`       // Watch | Warning | Emergency
	Description  string `json:"description"`
	IssuedAt     int64  `json:"issued_at"`
	ExpiresAt    int64  `json:"expires_at"`
}

// AlertsResponse is the top-level response for the alerts endpoint.
type AlertsResponse struct {
	UpdateTimestamp int64   `json:"update_timestamp"`
	TotalAlerts     int     `json:"total_alerts"`
	Alerts          []Alert `json:"alerts"`
}
