package api

import (
	"GOproject/GIT/memory_note/serializer"
	"encoding/json"
	"fmt"
)

func ErrorResponse(err error) serializer.Response {
	_, ok := err.(*json.UnmarshalTypeError)
	if ok {
		return serializer.Response{
			Status:  40001,
			Message: "json不匹配",
			Error:   fmt.Sprint(err),
		}
	}
	return serializer.Response{
		Status:  40001,
		Message: "参数错误",
		Error:   fmt.Sprint(err),
	}
}
