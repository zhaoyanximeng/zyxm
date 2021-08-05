package course

import (
	"encoding/json"
	"time"
)

type CourseAddDate struct {
	Date time.Time `gorm:"column:course_adddate"`
}

func NewCourseAddDate() *CourseAddDate {
	return &CourseAddDate{Date: time.Now()}
}

func (cad *CourseAddDate) GetDefaultFormat() string {
	return cad.Date.Format("2006:01:02 15:04:05")
}

func (cad *CourseAddDate) GetUnixFormat() int64 {
	return cad.Date.Unix()
}

func (cad *CourseAddDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Date          time.Time `json:"date"`
		DefaultFormat string    `json:"default_format"`
		UnixFormat    int64     `json:"unix_format"`
	}{
		Date:          cad.Date,
		DefaultFormat: cad.GetDefaultFormat(),
		UnixFormat:    cad.GetUnixFormat(),
	})
}
