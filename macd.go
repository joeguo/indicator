package indicator

type Macd struct {
	short  *Ema
	long   *Ema
	signal *Ema
	diff   float64
	dea    float64
	macd   float64
}

func NewMacd(short, long, signal int32) *Macd {
	return &Macd{short:NewEma(short), long:NewEma(long), signal:NewEma(signal)}
}

func (this *Macd)Update(price float64) (float64, float64, float64) {
	s := this.short.Update(price)
	l := this.long.Update(price)
	this.diff = s - l
	this.dea = this.signal.Update(this.diff)
	this.macd = 2.0 * (this.diff - this.dea)

	return this.diff, this.dea, this.macd
}

func (this *Macd)Clone() *Macd {
	return &Macd{short:this.short.Clone(), long:this.long.Clone(), signal:this.signal.Clone(), diff:this.diff, dea:this.dea, macd:this.macd}

}