package services

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/aouyuu/thai-flood-radar/data"
	"github.com/aouyuu/thai-flood-radar/structs"
)

// regionBaseRain returns the expected daily rainfall (mm) for a region in February.
func regionBaseRain(region string) float64 {
	switch region {
	case "South":
		return 18.0 // Gulf-coast south gets more rain year-round
	case "Northeast":
		return 2.5
	case "North":
		return 1.5
	case "Central":
		return 2.0
	case "East":
		return 5.0
	case "West":
		return 2.0
	default:
		return 2.0
	}
}

// floodRisk returns a 0–1 weighting for how flood-prone a province is.
func floodRisk(provinceID int) float64 {
	risks := map[int]float64{
		10: 0.45, // Bangkok
		14: 0.85, // Ayutthaya – historically flood-prone
		15: 0.70, // Ang Thong
		16: 0.60, // Lop Buri
		17: 0.70, // Sing Buri
		18: 0.60, // Chai Nat
		24: 0.65, // Chachoengsao
		34: 0.72, // Ubon Ratchathani
		35: 0.68, // Yasothon
		40: 0.62, // Khon Kaen
		44: 0.65, // Maha Sarakham
		45: 0.70, // Roi Et
		60: 0.72, // Nakhon Sawan
		80: 0.82, // Nakhon Si Thammarat
		84: 0.65, // Surat Thani
		90: 0.75, // Songkhla
		92: 0.68, // Trang
		94: 0.65, // Pattani
		96: 0.78, // Narathiwat
	}
	if r, ok := risks[provinceID]; ok {
		return r
	}
	return 0.30
}

// rainIntensity categorises a 24 h rainfall amount.
func rainIntensity(mm float64) string {
	switch {
	case mm <= 0:
		return "None"
	case mm < 10:
		return "Light"
	case mm < 35:
		return "Moderate"
	case mm < 70:
		return "Heavy"
	default:
		return "VeryHeavy"
	}
}

// dailySeed returns a deterministic per-province seed for a given calendar day.
func dailySeed(t time.Time, provinceID int) int64 {
	y, m, d := t.Date()
	return int64(y*10000+int(m)*100+d)*1000 + int64(provinceID)
}

// GetAffectedOverview returns the flood-affected overview for a given date.
func GetAffectedOverview(date time.Time) structs.AffectedOverview {
	var areas []structs.AreaOverview

	for _, p := range data.Provinces {
		r := rand.New(rand.NewSource(dailySeed(date, p.ID)))
		base := regionBaseRain(p.Region)
		risk := floodRisk(p.ID)

		// rain: base ±50 % random
		rain := base * (0.5 + r.Float64()*1.5)

		// only include province if flooding threshold reached
		if rain < 8 || risk < 0.55 {
			// small chance of residual flooding even in dry provinces
			if r.Float64() > 0.92 {
				val := int64(float64(r.Intn(500)+50) * risk)
				areas = append(areas, structs.AreaOverview{
					ID: int64(p.ID), Name: p.Name, Affected: &val,
				})
			}
			continue
		}

		val := int64(rain * risk * float64(r.Intn(800)+200))
		areas = append(areas, structs.AreaOverview{
			ID: int64(p.ID), Name: p.Name, Affected: &val,
		})
	}

	return structs.AffectedOverview{
		UpdateTimestamp: time.Now().Unix(),
		Date:            date.Unix(),
		AffectedAreas:   areas,
	}
}

// GetAffectedProvince returns detailed district/sub-district flood data.
func GetAffectedProvince(provinceID int, from, to time.Time) (structs.Area, bool) {
	p, ok := data.ProvinceByID(provinceID)
	if !ok {
		return structs.Area{}, false
	}

	r := rand.New(rand.NewSource(dailySeed(to, provinceID)))
	risk := floodRisk(provinceID)

	districts := make([]structs.District, 0, 5)
	numDistricts := r.Intn(4) + 2

	for d := 1; d <= numDistricts; d++ {
		subs := make([]structs.SubDistrict, 0, 4)
		numSubs := r.Intn(4) + 1
		for s := 1; s <= numSubs; s++ {
			affected := int64(float64(r.Intn(2000)+100) * risk)
			subs = append(subs, structs.SubDistrict{
				ID:              int64(s),
				UpdateTimestamp: time.Now().Unix(),
				Name:            fmt.Sprintf("ตำบลที่ %d", s),
				Affected:        affected,
			})
		}
		districts = append(districts, structs.District{
			ID:              int64(d),
			UpdateTimestamp: time.Now().Unix(),
			Name:            fmt.Sprintf("อำเภอที่ %d", d),
			Subdistrict:     subs,
		})
	}

	return structs.Area{
		ID:              int64(p.ID),
		UpdateTimestamp: time.Now().Unix(),
		Name:            p.Name,
		District:        districts,
	}, true
}

// GetRainOverview returns current 24 h rain data for all provinces.
func GetRainOverview() structs.RainOverview {
	now := time.Now()
	stations := make([]structs.RainStation, 0, len(data.Provinces))

	for _, p := range data.Provinces {
		base := regionBaseRain(p.Region)

		// Build 24-hour series (adds some intra-day variation).
		hourly := make([]float64, 24)
		dayR := rand.New(rand.NewSource(dailySeed(now, p.ID)))
		total24 := 0.0
		for h := 0; h < 24; h++ {
			// More rain during afternoon / evening hours (13-20).
			peakFactor := 1.0
			hour := (now.Hour() - 23 + h + 24) % 24
			if hour >= 13 && hour <= 20 {
				peakFactor = 1.8
			}
			mm := math.Max(0, base/24*peakFactor*(0.2+dayR.Float64()*1.8))
			// Round to 1 decimal.
			mm = math.Round(mm*10) / 10
			hourly[h] = mm
			total24 += mm
		}
		total24 = math.Round(total24*10) / 10

		// 7-day total: sum of daily amounts.
		total7d := 0.0
		for day := 0; day < 7; day++ {
			t7 := now.AddDate(0, 0, -day)
			dr := rand.New(rand.NewSource(dailySeed(t7, p.ID)))
			daily := base * (0.3 + dr.Float64()*2.0)
			total7d += daily
		}
		total7d = math.Round(total7d*10) / 10

		stations = append(stations, structs.RainStation{
			ProvinceID:   p.ID,
			ProvinceName: p.Name,
			ProvinceEn:   p.NameEn,
			Region:       p.Region,
			Lat:          p.Lat,
			Lng:          p.Lng,
			Amount24h:    total24,
			Amount7d:     total7d,
			Intensity:    rainIntensity(total24),
			HourlyData:   hourly,
		})
	}

	return structs.RainOverview{
		UpdateTimestamp: now.Unix(),
		Stations:        stations,
	}
}

// GetRainForecast returns a 7-day daily rain forecast for all provinces.
func GetRainForecast() structs.RainForecastOverview {
	now := time.Now()
	forecasts := make([]structs.RainForecast, 0, len(data.Provinces))

	for _, p := range data.Provinces {
		base := regionBaseRain(p.Region)
		daily := make([]structs.DailyForecast, 7)

		for d := 0; d < 7; d++ {
			t := now.AddDate(0, 0, d)
			dr := rand.New(rand.NewSource(dailySeed(t, p.ID+500))) // offset seed for forecast
			amount := math.Max(0, base*(0.2+dr.Float64()*2.0))
			amount = math.Round(amount*10) / 10
			daily[d] = structs.DailyForecast{
				Date:      t.Format("2006-01-02"),
				Amount:    amount,
				Intensity: rainIntensity(amount),
			}
		}

		forecasts = append(forecasts, structs.RainForecast{
			ProvinceID:   p.ID,
			ProvinceName: p.Name,
			Daily:        daily,
		})
	}

	return structs.RainForecastOverview{
		UpdateTimestamp: now.Unix(),
		Forecasts:       forecasts,
	}
}

// GetAlerts returns currently active flood / rain alerts.
func GetAlerts() structs.AlertsResponse {
	now := time.Now()

	// Realistic alerts for February (southern Thailand monsoon aftermath).
	rawAlerts := []struct {
		pid         int
		alertType   string
		level       string
		description string
		hoursAgo    int
		duration    int // hours
	}{
		{96, "HeavyRain", "Warning",
			"ฝนตกหนักต่อเนื่อง คาดมีน้ำท่วมขังในพื้นที่ลุ่มต่ำ",
			6, 48},
		{96, "RiverFlood", "Watch",
			"ระดับน้ำในแม่น้ำสูงกว่าตลิ่ง ประชาชนควรระวัง",
			12, 72},
		{94, "HeavyRain", "Watch",
			"ฝนตกหนักบางพื้นที่ เฝ้าระวังน้ำท่วมฉับพลัน",
			4, 36},
		{90, "RiverFlood", "Watch",
			"ทะเลสาบสงขลาระดับน้ำสูง อาจส่งผลต่อพื้นที่โดยรอบ",
			18, 96},
		{80, "HeavyRain", "Watch",
			"ฝนตกหนักในพื้นที่สูง เฝ้าระวังดินโคลนถล่มและน้ำป่า",
			8, 48},
		{14, "RiverFlood", "Watch",
			"ระดับน้ำในแม่น้ำเจ้าพระยาสูงขึ้น เฝ้าระวังพื้นที่ริมแม่น้ำ",
			24, 120},
		{92, "HeavyRain", "Warning",
			"ฝนตกหนักต่อเนื่อง น้ำท่วมขังในพื้นที่ชุมชน",
			3, 24},
	}

	alerts := make([]structs.Alert, 0, len(rawAlerts))
	for i, a := range rawAlerts {
		p, ok := data.ProvinceByID(a.pid)
		if !ok {
			continue
		}
		issued := now.Add(-time.Duration(a.hoursAgo) * time.Hour)
		expires := issued.Add(time.Duration(a.duration) * time.Hour)
		if now.After(expires) {
			continue
		}
		alerts = append(alerts, structs.Alert{
			ID:           fmt.Sprintf("ALT-%04d-%02d", now.Year()*100+int(now.Month()), i+1),
			ProvinceID:   p.ID,
			ProvinceName: p.Name,
			Type:         a.alertType,
			Level:        a.level,
			Description:  a.description,
			IssuedAt:     issued.Unix(),
			ExpiresAt:    expires.Unix(),
		})
	}

	return structs.AlertsResponse{
		UpdateTimestamp: now.Unix(),
		TotalAlerts:     len(alerts),
		Alerts:          alerts,
	}
}
