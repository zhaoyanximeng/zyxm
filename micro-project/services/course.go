package services

import (
	"context"
	. "micro-project/proto/course"
)

type CourseServiceImpl struct{}

func NewCourseServiceImpl() *CourseServiceImpl {
	return &CourseServiceImpl{}
}

func (cs *CourseServiceImpl) ListForTop(ctx context.Context, in *ListRequest, out *ListResponse) error {
	ret := make([]*CourseModel, 0)
	ret = append(ret, NewCourseModel(1, "golang"), NewCourseModel(2, "rust"))
	out.Result = ret
	return nil
}

func NewCourseModel(id int32, name string) *CourseModel {
	return &CourseModel{
		CourseId:   id,
		CourseName: name,
	}
}
