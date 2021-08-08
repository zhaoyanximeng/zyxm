package course_price

import "time"

type PriceComment struct {
	Comment string `json:"comment" gorm:"column:comment"`
	Setdate time.Time `json:"setdate" gorm:"column:setdate"`
}

func NewPriceComment() *PriceComment {
	return &PriceComment{Setdate: time.Now()}
}

func WithComment(comment string) CoursePriceOption {
	return func(c *CoursePrice) {
		c.Comment.Comment = comment
	}
}

func WithSetdata(t time.Time) CoursePriceOption {
	return func(c *CoursePrice) {
		c.Comment.Setdate = t
	}
}
