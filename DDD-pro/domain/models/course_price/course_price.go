package course_price

type CoursePrice struct {
	ID	int `json:"id" gorm:"column:id"`
	CourseID int `json:"course_id" gorm:"column:course_id"`
	Price *Price `json:"price" gorm:"embedded"`
	Comment *PriceComment `json:"comment" gorm:"embedded"`
	Iscurrent bool `json:"iscurrent" gorm:"column:iscurrent"`
}

func New(args ...CoursePriceOption) *CoursePrice {
	c := &CoursePrice{
		Price:NewPrice(),
		Comment: NewPriceComment(),
	}
	CoursePriceOptions(args).apply(c)
	return c
}

type CoursePriceOption func(c *CoursePrice)

type CoursePriceOptions []CoursePriceOption

func (cpos CoursePriceOptions) apply(u *CoursePrice) {
	for _,f := range cpos {
		f(u)
	}
}

func WithCourseID(cid int) CoursePriceOption {
	return func(c *CoursePrice) {
		c.ID = cid
	}
}

func WithIscurrent(b bool) CoursePriceOption {
	return func(c *CoursePrice) {
		c.Iscurrent = b
	}
}