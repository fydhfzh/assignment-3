package entity

import "time"

type Weather struct {
	Water       int    `json:"water"`
	Wind        int    `json:"wind"`
	WaterStatus string `json:"water_status"`
	WindStatus  string `json:"wind_status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (w *Weather) DetermineStatus(){
	if w.Water <= 5 {
		w.WaterStatus = "Aman"
	} else if w.Water >= 6 && w.Water <= 8 {
		w.WaterStatus = "Siaga"
	} else if w.Water > 8 {
		w.WaterStatus = "Bahaya"
	}

	if w.Wind <= 6 {
		w.WindStatus = "Aman"
	} else if w.Wind >= 7 && w.Wind <= 15 {
		w.WindStatus = "Siaga"
	} else if w.Wind > 15 {
		w.WindStatus = "Bahaya"
	}
}