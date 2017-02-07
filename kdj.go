package indicator

type Kdj struct {
	n1 int
	n2 int
	n3 int
}
// NewKdj(9, 3, 3)
func NewKdj(n1 int, n2 int, n3 int) *Kdj {
	return &Kdj{n1:n1, n2:n2, n3:n3}
}

func (this *Kdj)maxHigh(bids []Kline) (h float64) {
	h = bids[0].High
	for i := 0; i < len(bids); i++ {
		if bids[i].High > h {
			h = bids[i].High
		}
	}
	return
}

func (this *Kdj)minLow(bids []Kline) (l float64) {
	l = bids[0].Low
	for i := 0; i < len(bids); i++ {
		if (bids[i].Low < l) {
			l = bids[i].Low
		}
	}
	return
}

func (this *Kdj)sma(x []float64, n float64) (r []float64) {
	r = make([]float64, len(x))
	for i := 0; i < len(x); i++ {
		if i == 0 {
			r[i] = x[i]
		} else {
			r[i] = (1.0 * x[i] + (n - 1.0) * r[i - 1]) / n
		}
	}
	return
}



func (this *Kdj)Kdj(bids []Kline) (k, d, j []float64) {
	l := len(bids)
	rsv := make([]float64, l)
	j = make([]float64, l)
	rsv[0] = 50.0
	for i := 1; i < l; i++ {
		m := i + 1 - this.n1
		if m < 0 {
			m = 0
		}
		h := this.maxHigh(bids[m:i + 1])
		l := this.minLow(bids[m:i + 1])
		rsv[i] = (bids[i].Close - l) * 100.0 /( h - l)
		rsv[i] =rsv[i]
	}
	k = this.sma(rsv, float64(this.n2))
	d = this.sma(k, float64(this.n3))
	for i := 0; i < l; i++ {
		j[i] = 3.0 * k[i] - 2.0 * d[i]
	}
	return
}