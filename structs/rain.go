package structs

// RainStation holds current and recent rain data for one province.
type RainStation struct {
	ProvinceID   int       `json:"province_id"`
	ProvinceName string    `json:"province_name"`
	ProvinceEn   string    `json:"province_en"`
	Region       string    `json:"region"`
	Lat          float64   `json:"lat"`
	Lng          float64   `json:"lng"`
	Amount24h    float64   `json:"amount_24h"`  // mm in last 24 h
	Amount7d     float64   `json:"amount_7d"`   // mm in last 7 days
	Intensity    string    `json:"intensity"`   // None | Light | Moderate | Heavy | VeryHeavy
	HourlyData   []float64 `json:"hourly_data"` // last 24 hours, index 0 = 23 h ago
}

// DailyForecast represents a single day rain forecast.
type DailyForecast struct {
	Date      string  `json:"date"`
	Amount    float64 `json:"amount"`
	Intensity string  `json:"intensity"`
}

// RainForecast holds a 7-day rain forecast for one province.
type RainForecast struct {
	ProvinceID   int             `json:"province_id"`
	ProvinceName string          `json:"province_name"`
	Daily        []DailyForecast `json:"daily"`
}

// RainOverview aggregates current rain stations.
type RainOverview struct {
	UpdateTimestamp int64         `json:"update_timestamp"`
	Stations        []RainStation `json:"stations"`
}

// RainForecastOverview aggregates province forecasts.
type RainForecastOverview struct {
	UpdateTimestamp int64          `json:"update_timestamp"`
	Forecasts       []RainForecast `json:"forecasts"`
}
