package indicator

type Ema struct {
	Weight float64
	result float64
	age    uint32
}

func NewEma(weight int32) *Ema {
	return &Ema{Weight:float64(weight)}
}

func (this *Ema)Update(price float64) float64 {
	if this.age == 0 {
		this.result = price
	} else {
		this.result = (2.0 * price + (this.Weight - 1.0) * this.result) / (this.Weight + 1.0)
	}
	this.age += 1
	return this.result
}

func (this *Ema)Clone() *Ema {
	return &Ema{Weight:this.Weight, result:this.result, age:this.age}

}