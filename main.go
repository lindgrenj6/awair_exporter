package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	score       = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "score", Help: "awair score"}, []string{"device"})
	humidity    = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "humidity", Help: "humidity at current sensor"}, []string{"device"})
	pm25        = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "pm25", Help: "pm25 at current sensor"}, []string{"device"})
	temperature = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "temperature", Help: "temperature at current sensor"}, []string{"device"})
	co2         = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "co2", Help: "co2 at current sensor"}, []string{"device"})
	voc         = promauto.NewGaugeVec(prometheus.GaugeOpts{Name: "voc", Help: "voc at current sensor"}, []string{"device"})
)

func main() {
	go func() {
		for {
			devices, err := GetDeviceList()
			if err != nil {
				panic(err)
			}

			for _, dev := range devices.Devices {
				rawData, err := GetAirDataForDevice(dev.DeviceType, dev.DeviceID)
				if err != nil {
					panic(err)
				}

				reading := rawData.getReading()
				label := prometheus.Labels{"device": dev.Name}

				score.With(label).Set(reading.Score)
				humidity.With(label).Set(reading.Humidity)
				pm25.With(label).Set(reading.Pm25)
				temperature.With(label).Set(reading.Temperature)
				co2.With(label).Set(reading.Co2)
				voc.With(label).Set(reading.Voc)

				fmt.Printf("fetched data for %v\n", reading.Timestamp)
			}

			time.Sleep(5 * time.Minute)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
