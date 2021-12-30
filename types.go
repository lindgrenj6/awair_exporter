package main

type DeviceList struct {
	Devices []struct {
		DeviceID   int64  `json:"deviceId"`
		DeviceType string `json:"deviceType"`
		Name       string `json:"name"`
	} `json:"devices"`
}

type AirData struct {
	Timestamp   string
	Score       float64
	Humidity    float64
	Pm25        float64
	Temperature float64
	Co2         float64
	Voc         float64
}

type AirDataRaw struct {
	Data []struct {
		Score   float64 `json:"score"`
		Sensors []struct {
			Comp  string  `json:"comp"`
			Value float64 `json:"value"`
		} `json:"sensors"`
		Timestamp string `json:"timestamp"`
	} `json:"data"`
}

func (a *AirDataRaw) GetSensorValue(sensor string) float64 {
	for _, s := range a.Data[0].Sensors {
		if s.Comp == sensor {
			return s.Value
		}
	}

	return -1
}

func (a *AirDataRaw) getReading() *AirData {
	return &AirData{
		Timestamp:   a.Data[0].Timestamp,
		Score:       a.Data[0].Score,
		Humidity:    a.GetSensorValue("humid"),
		Pm25:        a.GetSensorValue("pm25"),
		Temperature: a.GetSensorValue("temp"),
		Co2:         a.GetSensorValue("co2"),
		Voc:         a.GetSensorValue("voc"),
	}
}
