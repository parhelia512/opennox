package opennox

import "unsafe"

func nox_swprintf(dst *byte, format string, args ...interface{}) {
	panic("TODO")
}

func cgoFuncEqual(a, b interface{}) bool {
	panic("TODO")
}

func cgoFuncAddr(fnc interface{}) uintptr {
	panic("TODO")
}

func cgoAsFunc[T any](fnc uint32, typ *T) T {
	panic("TODO")
}

type HANDLE = unsafe.Pointer

type int2 struct {
	field_0 int32
	field_4 int32
}

type int4 struct {
	field_0 int32
	field_4 int32
	field_8 int32
	field_C int32
}

type float2 struct {
	field_0 float32
	field_4 float32
}

type float4 struct {
	field_0 float32
	field_4 float32
	field_8 float32
	field_C float32
}
