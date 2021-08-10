package errorstring

import "fmt"

const (
	DataErrorModelCreate = 1501 // 实体新增失败
	DataErrorModelUpdate = 1502	// 实体更新失败
)

type DataOperatorError struct {
	Code int
	Message string
}

func (doe *DataOperatorError) Error() string  {
	return doe.Message
}

func NewDataOperatorError(code int,message string) *DataOperatorError {
	return &DataOperatorError{
		Code:    code,
		Message: message,
	}
}

func NewModelCreateError(modelName,err string) *DataOperatorError {
	return &DataOperatorError{
		Code:    DataErrorModelCreate,
		Message: fmt.Sprintf("model %s create err:%s",modelName,err),
	}
}