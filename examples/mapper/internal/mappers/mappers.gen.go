// Code generated by Mapper annotation processor. DO NOT EDIT.
// versions:
//		go: go1.18.3
//		go-annotation: 0.0.4-alpha
//		Mapper: 0.0.1-alpha

package mappers

import (
	common "github.com/YReshetko/go-annotation/examples/mapper/internal/models/common"
	input "github.com/YReshetko/go-annotation/examples/mapper/internal/models/input"
	output "github.com/YReshetko/go-annotation/examples/mapper/internal/models/output"
	"strconv"
	"strings"
	time "time"
)

var _ PrimitivesMapper = (*PrimitivesMapperImpl)(nil)

type PrimitivesMapperImpl struct{}

func (_this_ PrimitivesMapperImpl) Primitives(in *input.Primitives) output.Primitives {
	out := output.Primitives{}
	if in != nil {
		out.Bool = in.Bool
		out.Byte = in.Byte
		out.Int = in.Int
		out.Int8 = in.Int8
		out.Int16 = in.Int16
		out.Int32 = in.Int32
		out.Int64 = in.Int64
		out.Float32 = in.Float32
		out.Float64 = in.Float64
		out.Complex64 = in.Complex64
		out.Complex128 = in.Complex128
		out.String = in.String
		out.Uintptr = in.Uintptr
		out.Rune = in.Rune
		out.PtrBool = in.PtrBool
		out.PtrUnt = in.PtrUnt
		out.PtrUnt8 = in.PtrUnt8
		out.PtrUnt16 = in.PtrUnt16
		out.PtrUnt32 = in.PtrUnt32
		out.PtrUnt64 = in.PtrUnt64
		out.PtrByte = in.PtrByte
		out.PtrInt = in.PtrInt
		out.PtrInt8 = in.PtrInt8
		out.PtrInt16 = in.PtrInt16
		out.PtrInt32 = in.PtrInt32
		out.PtrInt64 = in.PtrInt64
		out.PtrFloat32 = in.PtrFloat32
		out.PtrFloat64 = in.PtrFloat64
		out.PtrComplex64 = in.PtrComplex64
		out.PtrComplex128 = in.PtrComplex128
		out.PtrString = in.PtrString
		out.PtrUintptr = in.PtrUintptr
		out.PtrRune = in.PtrRune
		out.CreateAt = &in.CreateAt
		if in.MAP != nil {
			out.MAP = *in.MAP
		}
	}

	return out
}

func (_this_ PrimitivesMapperImpl) ConvertMethod(in0 *Inner, in1 ThisFileStruct) output.Primitives {
	_var_0 := output.Primitives{}
	_var_0.Bool = func(v string) bool {
		res := v == "1" || strings.ToLower(v) == "true"
		return res
	}(in1.Bool)
	_var_0.String = in1.String
	_var_0.PtrUnt64 = func(v string) *uint64 {
		res := func() uint64 { o, _ := strconv.Atoi(v); return uint64(o) }()
		return &res
	}(in1.PtrUnt64)
	_var_0.PtrInt = &in1.PtrInt
	_var_0.PtrComplex64 = func(v string) *complex64 {
		res := func() complex64 { o, _ := strconv.ParseComplex(v, 64); return complex64(o) }()
		return &res
	}(in1.PtrComplex64)
	if in1.PtrComplex128 != nil {
		_var_0.PtrComplex128 = func(v string) *complex128 {
			res := func() complex128 { o, _ := strconv.ParseComplex(v, 128); return o }()
			return &res
		}(*in1.PtrComplex128)
	}
	if in1.Int != nil {
		_var_0.Int = *in1.Int
	}
	if in1.Float32 != nil {
		_var_0.Float32 = func(v int) float32 {
			res := float32(v)
			return res
		}(*in1.Float32)
	}
	if in0 != nil {
		_var_0.PtrBool = in0.PtrBool
	}
	if in1.PtrFloat64 != nil {
		_var_0.PtrFloat64 = func(v string) *float64 {
			res := func() float64 { o, _ := strconv.ParseFloat(v, 64); return o }()
			return &res
		}(*in1.PtrFloat64)
	}

	return _var_0
}

func (_this_ PrimitivesMapperImpl) ConvertMethod2(in0 Inner, in1 *ThisFileStruct) output.Primitives {
	_var_0 := output.Primitives{}
	_var_0.PtrBool = in0.PtrBool
	if in1 != nil {
		_var_0.Bool = func(v string) bool {
			res := v == "1" || strings.ToLower(v) == "true"
			return res
		}(in1.Bool)
		_var_0.String = in1.String
		_var_0.PtrUnt64 = func(v string) *uint64 {
			res := func() uint64 { o, _ := strconv.Atoi(v); return uint64(o) }()
			return &res
		}(in1.PtrUnt64)
		_var_0.PtrInt = &in1.PtrInt
		_var_0.PtrComplex64 = func(v string) *complex64 {
			res := func() complex64 { o, _ := strconv.ParseComplex(v, 64); return complex64(o) }()
			return &res
		}(in1.PtrComplex64)
		if in1.Int != nil {
			_var_0.Int = *in1.Int
		}
		if in1.Float32 != nil {
			_var_0.Float32 = func(v int) float32 {
				res := float32(v)
				return res
			}(*in1.Float32)
		}
		if in1.PtrFloat64 != nil {
			_var_0.PtrFloat64 = func(v string) *float64 {
				res := func() float64 { o, _ := strconv.ParseFloat(v, 64); return o }()
				return &res
			}(*in1.PtrFloat64)
		}
		if in1.PtrComplex128 != nil {
			_var_0.PtrComplex128 = func(v string) *complex128 {
				res := func() complex128 { o, _ := strconv.ParseComplex(v, 128); return o }()
				return &res
			}(*in1.PtrComplex128)
		}
	}

	return _var_0
}

var _ ConstantMapper = (*ConstantMapperImpl)(nil)

type ConstantMapperImpl struct{}

func (_this_ ConstantMapperImpl) Primitives() output.Primitives {
	_var_0 := output.Primitives{}
	_var_0.Float32 = func(v string) float32 {
		res := func() float32 { o, _ := strconv.ParseFloat(v, 32); return float32(o) }()
		return res
	}("3.14")
	_var_0.String = "Hello"
	_var_0.PtrFloat32 = func(v string) *float32 {
		res := func() float32 { o, _ := strconv.ParseFloat(v, 32); return float32(o) }()
		return &res
	}("3.14")
	_var_1 := "World"
	_var_0.PtrString = &_var_1

	return _var_0
}

var _ BaseStructuresMapper = (*BaseStructuresMapperImpl)(nil)

type BaseStructuresMapperImpl struct{}

func (_this_ BaseStructuresMapperImpl) Structures1(in *input.StructuresMapping) output.StructuresMapping {
	_var_0 := output.StructuresMapping{}
	if in != nil {
		_var_0.Field1 = in.Field1
		_var_0.Field2 = &in.Field2
		_var_0.Field4 = in.Field4
		if in.Field3 != nil {
			_var_0.Field3 = *in.Field3
		}
	}

	return _var_0
}

func (_this_ BaseStructuresMapperImpl) Structures2(in *input.StructuresMapping2) output.StructuresMapping2 {
	_var_0 := output.StructuresMapping2{}
	if in != nil {
		_var_0.AnotherField2 = &in.Field2.Field1.Field1
		_var_0.AnotherField4 = &in.Field2.Field1.Field1
		if in.Field1.Field1.Field2 != nil {
			_var_0.AnotherField1 = *in.Field1.Field1.Field2
		}
		if in.Field1.Field2 != nil && in.Field1.Field2.Field2 != nil {
			_var_0.AnotherField3 = *in.Field1.Field2.Field2
		}
	}

	return _var_0
}

func (_this_ BaseStructuresMapperImpl) Structures3(in1 *common.Common, in2 common.Common) output.StructuresMapping2 {
	_var_0 := output.StructuresMapping2{}
	_var_0.AnotherField4 = &in2
	if in1 != nil {
		_var_0.AnotherField3 = *in1
	}

	return _var_0
}

func (_this_ BaseStructuresMapperImpl) Structures4(in1 string, in2 *complex128, in3 *uint64) output.Primitives {
	_var_0 := output.Primitives{}
	_var_0.String = in1
	_var_0.PtrString = &in1
	if in3 != nil {
		_var_0.Uint64 = *in3
	}
	if in2 != nil {
		_var_0.Complex128 = *in2
		_var_0.PtrComplex128 = in2
	}

	return _var_0
}

func (_this_ BaseStructuresMapperImpl) TimeMapping(t time.Time) time.Time {
	_var_0 := time.Time{}

	return _var_0
}

var _ FunctionMapper = (*FunctionMapperImpl)(nil)

type FunctionMapperImpl struct{}

func (_this_ FunctionMapperImpl) Function1(in *input.StructuresMapping2) output.StructuresMapping2 {
	_var_0 := output.StructuresMapping2{}
	if in != nil && in.Field1.Field1.Field2 != nil && in.Field1.Field2 != nil && in.Field1.Field2.Field2 != nil {
		_var_0.AnotherField1 = _this_.fieldToField(
			in.Field1.Field1.Field1,
			in.Field1.Field1.Field2,
			in.Field1.Field2.Field1,
			in.Field1.Field2.Field2)
	}

	return _var_0
}

func (_this_ FunctionMapperImpl) Function2(in *input.StructuresMapping2) output.StructuresMapping2 {
	_var_0 := output.StructuresMapping2{}
	if in != nil && in.Field1.Field1.Field2 != nil && in.Field1.Field2 != nil && in.Field1.Field2.Field2 != nil {
		_var_0.AnotherField1 = fieldToField(
			in.Field1.Field1.Field1,
			in.Field1.Field1.Field2,
			in.Field1.Field2.Field1,
			in.Field1.Field2.Field2)
	}

	return _var_0
}

func (_this_ FunctionMapperImpl) fieldToField(in0 common.Common, in1 *common.Common, in2 common.Common, in3 *common.Common) common.Common {
	_var_0 := common.Common{}
	_var_0.Field1 = in0.Field1
	_var_0.Field2 = in0.Field2
	_var_0.Field3 = in0.Field3
	_var_0.Slice = in0.Slice

	return _var_0
}

var _ SliceMapping = (*SliceMappingImpl)(nil)

type SliceMappingImpl struct{}

func (_this_ SliceMappingImpl) Function1(in *input.SliceStruct) output.SliceStruct {
	_var_0 := output.SliceStruct{}
	if in != nil && in.Slice != nil {
		_var_0.Slice = sliceInOut(in.Slice)
	}

	return _var_0
}

func (_this_ SliceMappingImpl) Function2(in *input.SliceStruct) output.SliceStruct {
	_var_0 := output.SliceStruct{}
	if in != nil && in.Slice != nil {

		_var_1 := *in.Slice
		_var_2 := make([]output.StructuresMapping2, len(_var_1), len(_var_1))
		for _var_3, _var_4 := range _var_1 {
			_var_2[_var_3] = _this_.genMapper(_var_4)
		}
		_var_0.Slice = _var_2

	}

	return _var_0
}

func (_this_ SliceMappingImpl) genMapper(in input.Local2) output.StructuresMapping2 {
	_var_0 := output.StructuresMapping2{}

	return _var_0
}

func (_this_ SliceMappingImpl) Function3(in *input.SliceStruct2) output.SliceStruct2 {
	_var_0 := output.SliceStruct2{}
	if in != nil {

		_var_1 := in.Slice
		_var_2 := make([]common.Common2, len(_var_1), len(_var_1))
		for _var_3, _var_4 := range _var_1 {
			_var_2[_var_3] = _this_.genMapper2(_var_4)
		}
		_var_0.Slice = _var_2

	}

	return _var_0
}

func (_this_ SliceMappingImpl) genMapper2(in common.Common) common.Common2 {
	_var_0 := common.Common2{}
	_var_0.Field1 = in.Field1
	_var_0.Field2 = in.Field2
	_var_0.Field3 = in.Field3
	_var_0.Slice = in.Slice

	return _var_0
}

var _ MapMapping = (*MapMappingImpl)(nil)

type MapMappingImpl struct{}

func (_this_ MapMappingImpl) Function1(in input.MapStruct) output.MapStruct {
	_var_0 := output.MapStruct{}

	_var_1 := in.Map
	_var_2 := make(map[common.MapKey]common.Common2, len(_var_1))
	for _var_3, _var_4 := range _var_1 {
		_var_5, _var_6 := _this_.genMapper(_var_3, _var_4)
		_var_2[_var_5] = _var_6
	}
	_var_0.Map = _var_2

	return _var_0
}

func (_this_ MapMappingImpl) genMapper(k common.MapKey, v common.Common2) (common.MapKey, common.Common2) {
	_var_0 := common.MapKey{}
	_var_0.Field1 = k.Field1
	_var_0.Field2 = k.Field2
	_var_0.Field3 = k.Field3
	_var_1 := common.Common2{}
	_var_1.Field1 = k.Field1
	_var_1.Field2 = k.Field2
	_var_1.Field3 = k.Field3
	_var_1.Slice = v.Slice

	return _var_0, _var_1
}

func (_this_ MapMappingImpl) Function2(mapStruct input.MapStruct2) output.MapStruct2 {
	_var_0 := output.MapStruct2{}
	_var_0.Map = mapStruct.Map

	return _var_0
}
