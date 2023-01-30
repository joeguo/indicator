package indicator

import (
	"math"
)

type Boll struct {
	n int
	k float64
}

//NewBoll(20, 2)
func NewBoll(n int, k int32) *Boll {
	return &Boll{n: n, k: float64(k)}
}
func (this *Boll) sma(lines []Kline) float64 {
	s := len(lines)
	var sum float64 = 0
	for i := 0; i < s; i++ {
		sum += float64(lines[i].Close)
	}
	return sum / float64(s)
}
func (this *Boll) dma(lines []Kline, ma float64) float64 {
	s := len(lines)
	//log.Println(s)
	var sum float64 = 0
	for i := 0; i < s; i++ {
		sum += (lines[i].Close - ma) * (lines[i].Close - ma)
	}
	return math.Sqrt(sum / float64(this.n))
}

func (this *Boll) Boll(lines []Kline) (mid []float64, up []float64, low []float64) {
	l := len(lines)

	mid = make([]float64, l)
	up = make([]float64, l)
	low = make([]float64, l)
	if l < this.n {
		return
	}
	for i := l - 1; i > this.n-1; i-- {
		ps := lines[(i - this.n + 1): i+1 ]
		mid[i] = this.sma(ps)
		dm := this.k * this.dma(ps, mid[i])
		up[i] = mid[i] + dm
		low[i] = mid[i] - dm
	}

	return
}
