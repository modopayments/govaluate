package govaluate

import (
	"errors"
	"fmt"
)

/*
	Struct used to test "parameter calls".
*/
type dummyParameter struct {
	String    string
	Int       int64
	BoolFalse bool
	Nil       interface{}
	Nested    dummyNestedParameter
}

func (this dummyParameter) Func() string {
	return "funk"
}

func (this dummyParameter) Func2() (string, error) {
	return "frink", nil
}

func (this *dummyParameter) Func3() string {
	return "fronk"
}

func (this dummyParameter) FuncArgStr(arg1 string) string {
	return arg1
}

func (this dummyParameter) TestArgs(str string, ui uint, ui8 uint8, ui16 uint16, ui32 uint32, ui64 uint64, i int, i8 int8, i16 int16, i32 int32, i64 int64, f32 float32, f64 int64, b bool) string {

	var sum int64

	sum = int64(ui) + int64(ui8) + int64(ui16) + int64(ui32) + int64(ui64)
	sum += int64(i) + int64(i8) + int64(i16) + int64(i32) + int64(i64)
	sum += int64(f32)

	if b {
		sum += f64
	}

	return fmt.Sprintf("%v: %v", str, sum)
}

func (this dummyParameter) AlwaysFail() (interface{}, error) {
	return nil, errors.New("function should always fail")
}

type dummyNestedParameter struct {
	Funk string
}

func (this dummyNestedParameter) Dunk(arg1 string) string {
	return arg1 + "dunk"
}

var dummyParameterInstance = dummyParameter{
	String:    "string!",
	Int:       int64(101),
	BoolFalse: false,
	Nil:       nil,
	Nested: dummyNestedParameter{
		Funk: "funkalicious",
	},
}

var fooParameter = EvaluationParameter{
	Name:  "foo",
	Value: dummyParameterInstance,
}

var fooPtrParameter = EvaluationParameter{
	Name:  "fooptr",
	Value: &dummyParameterInstance,
}

var fooFailureParameters = map[string]interface{}{
	"foo":    fooParameter.Value,
	"fooptr": &fooPtrParameter.Value,
}
