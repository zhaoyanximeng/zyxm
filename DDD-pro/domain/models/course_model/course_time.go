package course_model

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

type CourseTime struct {
	CourseSeconds int64 `gorm:"column:course_time"`
}

func NewCourseTime() *CourseTime {
	return &CourseTime{}
}

func WithCourseTime(seconds int64) CourseMainOption {
	return func(c *CourseMain) {
		c.CourseTime.CourseSeconds = seconds
	}
}

func (t CourseTime) MarshalJSON() ([]byte, error) {
	seconds := t.CourseSeconds
	hour := seconds / 3600
	minute := (seconds - hour*3600) / 60
	second := seconds - hour*3600 - minute*60
	return jsoniter.Marshal(map[string]interface{}{
		"course_seconds": seconds,
		"format":         fmt.Sprintf("%d时%d分%d秒", hour, minute, second),
	})
}

func (ct *CourseTime) AddSecond(sec int64) {
	ct.CourseSeconds += sec
}
