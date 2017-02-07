package indicator

import "time"

type Kline struct {
	Time  time.Time
	Open  float64
	Close float64
	High  float64
	Low   float64
	Vol   float64
	Money float64
}
