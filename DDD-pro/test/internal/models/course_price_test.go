package models

import (
	"eventbus/domain/models/course_price"
	"eventbus/test/internal/conf"
	"testing"
)

func CoursePriceAdd(cid int,mPrice,sPrice float64,comment string,current bool) error {
	cp := course_price.New(
		course_price.WithCourseID(cid),
		course_price.WithMarketPrice(mPrice),
		course_price.WithSalePrice(sPrice),
		course_price.WithComment(comment),
		course_price.WithIscurrent(current),
	)
	return conf.DB.Table("course_prices").Create(cp).Error
}

func TestCoursePriceAdd(t *testing.T) {
	data := []struct{
		cid int
		mPrice float64
		sPrice float64
		comment string
		current bool
	} {
		{cid:201,sPrice:30.02,mPrice:21,comment:"双11活动",current:false},
		{cid:202,sPrice:0,mPrice:21,comment:"活动2",current:false},
		{cid:204,sPrice:55,mPrice:0,comment:"活动3",current:false},
		{cid:203,sPrice:109,mPrice:200,comment:"活动4",current:false},
	}
	for i := range data {
		err := CoursePriceAdd(data[i].cid,data[i].mPrice,data[i].sPrice,data[i].comment,data[i].current)
		if err != nil {
			t.Errorf("课程价格id=%d创建出错，原因是:%s",data[i].cid,err.Error())
		}
	}
}