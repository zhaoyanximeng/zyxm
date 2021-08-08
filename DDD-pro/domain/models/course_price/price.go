package course_price

import "github.com/shopspring/decimal"

type Price struct {
	MarketPrice decimal.Decimal	`json:"market_price" gorm:"column:market_price"`
	SalePrice decimal.Decimal `json:"sale_price" gorm:"column:sale_price"`
	Discount decimal.Decimal `json:"discount" gorm:"-"`
}

func NewPrice() *Price {
	return &Price{}
}

func WithMarketPrice(price float64) CoursePriceOption {
	return func(c *CoursePrice) {
		c.Price.MarketPrice = decimal.NewFromFloat(price)
		c.Price.setDiscount()
	}
}

func WithSalePrice(price float64) CoursePriceOption {
	return func(c *CoursePrice) {
		c.Price.SalePrice = decimal.NewFromFloat(price)
		c.Price.setDiscount()
	}
}

func (p *Price) setDiscount() {
	zero := decimal.NewFromInt(0)
	if p.MarketPrice.GreaterThan(zero) && p.SalePrice.GreaterThan(zero) {
		p.Discount = p.MarketPrice.Sub(p.SalePrice).
			DivRound(p.MarketPrice,2).Sub(decimal.NewFromInt(1)).Abs()
	}
}
