// Code generated by Constructor annotation processor. DO NOT EDIT.
// versions:
//		go-annotation: 0.0.4-alpha
//		Constructor: 1.0.0

package internal

import (
	bytes "bytes"
	common "github.com/YReshetko/go-annotation/examples/constructor/internal/common"
	constraints "golang.org/x/exp/constraints"
	http "net/http"
)

func NewAnotherStructOverride(a SomeStructure, b *SomeStructure, buff bytes.Buffer, fn func(**SomeStructure) AnotherStruct) *AnotherStruct {
	returnValue := &AnotherStruct{
		a:    a,
		b:    b,
		buff: buff,
		fn:   fn,
	}

	return returnValue
}

type AnotherStructOption func(*AnotherStruct)

func NewAnotherStructOptional(opts ...AnotherStructOption) AnotherStruct {
	rt := &AnotherStruct{}
	for _, o := range opts {
		o(rt)
	}

	return *rt
}

func WithA(v SomeStructure) AnotherStructOption {
	return func(rt *AnotherStruct) {
		rt.a = v
	}
}

func WithB(v *SomeStructure) AnotherStructOption {
	return func(rt *AnotherStruct) {
		rt.b = v
	}
}

func WithBuff(v bytes.Buffer) AnotherStructOption {
	return func(rt *AnotherStruct) {
		rt.buff = v
	}
}

func WithFn(v func(**SomeStructure) AnotherStruct) AnotherStructOption {
	return func(rt *AnotherStruct) {
		rt.fn = v
	}
}

func NewSomeSomeStructureThisIsMyTemplate(anonimus struct {
	a int
	b float64
}, c *bool, d **complex128) SomeStructure {
	returnValue := SomeStructure{
		anonimus:   anonimus,
		c:          c,
		chanals:    make(chan []struct{ A http.Request }),
		chanalsCap: make(chan []struct{ A http.Request }, 5),
		d:          d,
		maps:       make(map[chan []int]struct{ A http.Request }, 5),
		slice:      make([]map[chan int]string, 5, 15),
		slice2:     []map[chan int]string{},
	}

	return returnValue
}

type SomeStructureOption func(*SomeStructure)

func NewSomeSomeStructureThisIsMyTemplateOptional(opts ...SomeStructureOption) SomeStructure {
	rt := &SomeStructure{
		chanals:    make(chan []struct{ A http.Request }),
		chanalsCap: make(chan []struct{ A http.Request }, 5),
		maps:       make(map[chan []int]struct{ A http.Request }, 5),
		slice:      make([]map[chan int]string, 5, 15),
		slice2:     []map[chan int]string{},
	}
	for _, o := range opts {
		o(rt)
	}

	return *rt
}

func WithAnonimus(v struct {
	a int
	b float64
}) SomeStructureOption {
	return func(rt *SomeStructure) {
		rt.anonimus = v
	}
}

func WithC(v *bool) SomeStructureOption {
	return func(rt *SomeStructure) {
		rt.c = v
	}
}

func WithChanals(v chan []struct{ A http.Request }) SomeStructureOption {
	return func(rt *SomeStructure) {
		rt.chanals = v
	}
}

func WithChanalsCap(v chan []struct{ A http.Request }) SomeStructureOption {
	return func(rt *SomeStructure) {
		rt.chanalsCap = v
	}
}

func WithD(v **complex128) SomeStructureOption {
	return func(rt *SomeStructure) {
		rt.d = v
	}
}

func WithMaps(v map[chan []int]struct{ A http.Request }) SomeStructureOption {
	return func(rt *SomeStructure) {
		rt.maps = v
	}
}

func WithSlice(v []map[chan int]string) SomeStructureOption {
	return func(rt *SomeStructure) {
		rt.slice = v
	}
}

func WithSlice2(v []map[chan int]string) SomeStructureOption {
	return func(rt *SomeStructure) {
		rt.slice2 = v
	}
}

type MySomeStructureBuilder struct {
	_anonimus_ struct {
		a int
		b float64
	}
	_c_          *bool
	_chanals_    chan []struct{ A http.Request }
	_chanalsCap_ chan []struct{ A http.Request }
	_d_          **complex128
	_maps_       map[chan []int]struct{ A http.Request }
	_slice_      []map[chan int]string
	_slice2_     []map[chan int]string
}

func NewSomeStructureBuilder() *MySomeStructureBuilder {
	return &MySomeStructureBuilder{}
}

func (b *MySomeStructureBuilder) BuildAnonimusField(v struct {
	a int
	b float64
}) *MySomeStructureBuilder {
	b._anonimus_ = v
	return b
}

func (b *MySomeStructureBuilder) BuildCField(v *bool) *MySomeStructureBuilder {
	b._c_ = v
	return b
}

func (b *MySomeStructureBuilder) BuildChanalsField(v chan []struct{ A http.Request }) *MySomeStructureBuilder {
	b._chanals_ = v
	return b
}

func (b *MySomeStructureBuilder) BuildChanalsCapField(v chan []struct{ A http.Request }) *MySomeStructureBuilder {
	b._chanalsCap_ = v
	return b
}

func (b *MySomeStructureBuilder) BuildDField(v **complex128) *MySomeStructureBuilder {
	b._d_ = v
	return b
}

func (b *MySomeStructureBuilder) BuildMapsField(v map[chan []int]struct{ A http.Request }) *MySomeStructureBuilder {
	b._maps_ = v
	return b
}

func (b *MySomeStructureBuilder) BuildSliceField(v []map[chan int]string) *MySomeStructureBuilder {
	b._slice_ = v
	return b
}

func (b *MySomeStructureBuilder) BuildSlice2Field(v []map[chan int]string) *MySomeStructureBuilder {
	b._slice2_ = v
	return b
}

func (b *MySomeStructureBuilder) Build() *SomeStructure {
	out := SomeStructure{}
	if b._chanals_ == nil {
		b._chanals_ = make(chan []struct{ A http.Request })
	}
	if b._chanalsCap_ == nil {
		b._chanalsCap_ = make(chan []struct{ A http.Request }, 5)
	}
	if b._maps_ == nil {
		b._maps_ = make(map[chan []int]struct{ A http.Request }, 5)
	}
	if b._slice_ == nil {
		b._slice_ = make([]map[chan int]string, 5, 15)
	}
	if b._slice2_ == nil {
		b._slice2_ = []map[chan int]string{}
	}

	out.anonimus = b._anonimus_
	out.c = b._c_
	out.chanals = b._chanals_
	out.chanalsCap = b._chanalsCap_
	out.d = b._d_
	out.maps = b._maps_
	out.slice = b._slice_
	out.slice2 = b._slice2_

	return &out
}

func NewStackQueueStruct[T comparable, V constraints.Integer](a stack[T], buff bytes.Buffer, fn func(**SomeStructure) AnotherStruct, q queue[V], simp T, vimp V) *StackQueueStruct[T, V] {
	returnValue := &StackQueueStruct[T, V]{
		a:    a,
		buff: buff,
		fn:   fn,
		q:    q,
		simp: simp,
		str:  make(chan map[T][]V),
		vimp: vimp,
	}
	returnValue.postConstruct1()
	returnValue.postConstruct2()
	returnValue.postConstruct3()

	return returnValue
}

type StackQueueStructOption[T comparable, V constraints.Integer] func(*StackQueueStruct[T, V])

func NewStackQueueStructOptional[T comparable, V constraints.Integer](opts ...StackQueueStructOption[T, V]) *StackQueueStruct[T, V] {
	rt := &StackQueueStruct[T, V]{
		str: make(chan map[T][]V),
	}
	for _, o := range opts {
		o(rt)
	}
	rt.postConstruct1()
	rt.postConstruct2()
	rt.postConstruct3()

	return rt
}

func WithSQSA[T comparable, V constraints.Integer](v stack[T]) StackQueueStructOption[T, V] {
	return func(rt *StackQueueStruct[T, V]) {
		rt.a = v
	}
}

func WithSQSBuff[T comparable, V constraints.Integer](v bytes.Buffer) StackQueueStructOption[T, V] {
	return func(rt *StackQueueStruct[T, V]) {
		rt.buff = v
	}
}

func WithSQSFn[T comparable, V constraints.Integer](v func(**SomeStructure) AnotherStruct) StackQueueStructOption[T, V] {
	return func(rt *StackQueueStruct[T, V]) {
		rt.fn = v
	}
}

func WithSQSQ[T comparable, V constraints.Integer](v queue[V]) StackQueueStructOption[T, V] {
	return func(rt *StackQueueStruct[T, V]) {
		rt.q = v
	}
}

func WithSQSSimp[T comparable, V constraints.Integer](v T) StackQueueStructOption[T, V] {
	return func(rt *StackQueueStruct[T, V]) {
		rt.simp = v
	}
}

func WithSQSStr[T comparable, V constraints.Integer](v chan map[T][]V) StackQueueStructOption[T, V] {
	return func(rt *StackQueueStruct[T, V]) {
		rt.str = v
	}
}

func WithSQSVimp[T comparable, V constraints.Integer](v V) StackQueueStructOption[T, V] {
	return func(rt *StackQueueStruct[T, V]) {
		rt.vimp = v
	}
}

type MyStackQueueStructBuilder[T comparable, V constraints.Integer] struct {
	_a_    stack[T]
	_buff_ bytes.Buffer
	_fn_   func(**SomeStructure) AnotherStruct
	_q_    queue[V]
	_simp_ T
	_str_  chan map[T][]V
	_vimp_ V
}

func NewStackQueueStructBuilder[T comparable, V constraints.Integer]() *MyStackQueueStructBuilder[T, V] {
	return &MyStackQueueStructBuilder[T, V]{}
}

func (b *MyStackQueueStructBuilder[T, V]) BuildAField(v stack[T]) *MyStackQueueStructBuilder[T, V] {
	b._a_ = v
	return b
}

func (b *MyStackQueueStructBuilder[T, V]) BuildBuffField(v bytes.Buffer) *MyStackQueueStructBuilder[T, V] {
	b._buff_ = v
	return b
}

func (b *MyStackQueueStructBuilder[T, V]) BuildFnField(v func(**SomeStructure) AnotherStruct) *MyStackQueueStructBuilder[T, V] {
	b._fn_ = v
	return b
}

func (b *MyStackQueueStructBuilder[T, V]) BuildQField(v queue[V]) *MyStackQueueStructBuilder[T, V] {
	b._q_ = v
	return b
}

func (b *MyStackQueueStructBuilder[T, V]) BuildSimpField(v T) *MyStackQueueStructBuilder[T, V] {
	b._simp_ = v
	return b
}

func (b *MyStackQueueStructBuilder[T, V]) BuildStrField(v chan map[T][]V) *MyStackQueueStructBuilder[T, V] {
	b._str_ = v
	return b
}

func (b *MyStackQueueStructBuilder[T, V]) BuildVimpField(v V) *MyStackQueueStructBuilder[T, V] {
	b._vimp_ = v
	return b
}

func (b *MyStackQueueStructBuilder[T, V]) Build() *StackQueueStruct[T, V] {
	out := StackQueueStruct[T, V]{}
	if b._str_ == nil {
		b._str_ = make(chan map[T][]V)
	}

	out.a = b._a_
	out.buff = b._buff_
	out.fn = b._fn_
	out.q = b._q_
	out.simp = b._simp_
	out.str = b._str_
	out.vimp = b._vimp_

	out.postConstruct1()
	out.postConstruct2()
	out.postConstruct3()

	return &out
}

func NewStackStruct[T stack[T]](a stack[T], fn func(**SomeStructure) AnotherStruct, q queue[stack[T]]) StackStruct[T] {
	returnValue := StackStruct[T]{
		a:  a,
		fn: fn,
		q:  q,
	}

	return returnValue
}

func NewEmbeddingConstructorExample[T comparable, V constraints.Integer](ExternalEmbeddedInterface common.ExternalEmbeddedInterface, SomeStruct *common.SomeStruct, SomeStructure SomeStructure, embeddedPrivateInterface embeddedPrivateInterface, privateEmbedded privateEmbedded) StructEmbedding[T, V] {
	returnValue := StructEmbedding[T, V]{
		ExternalEmbeddedInterface: ExternalEmbeddedInterface,
		SomeStruct:                SomeStruct,
		SomeStructure:             SomeStructure,
		embeddedPrivateInterface:  embeddedPrivateInterface,
		privateEmbedded:           privateEmbedded,
	}

	return returnValue
}

type StructEmbeddingOption[T comparable, V constraints.Integer] func(*StructEmbedding[T, V])

func NewStructEmbedding[T comparable, V constraints.Integer](opts ...StructEmbeddingOption[T, V]) StructEmbedding[T, V] {
	rt := &StructEmbedding[T, V]{}
	for _, o := range opts {
		o(rt)
	}

	return *rt
}

func WithExternalEmbeddedInterface[T comparable, V constraints.Integer](v common.ExternalEmbeddedInterface) StructEmbeddingOption[T, V] {
	return func(rt *StructEmbedding[T, V]) {
		rt.ExternalEmbeddedInterface = v
	}
}

func WithSomeStruct[T comparable, V constraints.Integer](v *common.SomeStruct) StructEmbeddingOption[T, V] {
	return func(rt *StructEmbedding[T, V]) {
		rt.SomeStruct = v
	}
}

func WithSomeStructure[T comparable, V constraints.Integer](v SomeStructure) StructEmbeddingOption[T, V] {
	return func(rt *StructEmbedding[T, V]) {
		rt.SomeStructure = v
	}
}

func WithEmbeddedPrivateInterface[T comparable, V constraints.Integer](v embeddedPrivateInterface) StructEmbeddingOption[T, V] {
	return func(rt *StructEmbedding[T, V]) {
		rt.embeddedPrivateInterface = v
	}
}

func WithPrivateEmbedded[T comparable, V constraints.Integer](v privateEmbedded) StructEmbeddingOption[T, V] {
	return func(rt *StructEmbedding[T, V]) {
		rt.privateEmbedded = v
	}
}

type StructEmbeddingBuilder[T comparable, V constraints.Integer] struct {
	_ExternalEmbeddedInterface_ common.ExternalEmbeddedInterface
	_SomeStruct_                *common.SomeStruct
	_SomeStructure_             SomeStructure
	_embeddedPrivateInterface_  embeddedPrivateInterface
	_privateEmbedded_           privateEmbedded
}

func NewStructEmbeddingBuilder[T comparable, V constraints.Integer]() *StructEmbeddingBuilder[T, V] {
	return &StructEmbeddingBuilder[T, V]{}
}

func (b *StructEmbeddingBuilder[T, V]) ExternalEmbeddedInterface(v common.ExternalEmbeddedInterface) *StructEmbeddingBuilder[T, V] {
	b._ExternalEmbeddedInterface_ = v
	return b
}

func (b *StructEmbeddingBuilder[T, V]) SomeStruct(v *common.SomeStruct) *StructEmbeddingBuilder[T, V] {
	b._SomeStruct_ = v
	return b
}

func (b *StructEmbeddingBuilder[T, V]) SomeStructure(v SomeStructure) *StructEmbeddingBuilder[T, V] {
	b._SomeStructure_ = v
	return b
}

func (b *StructEmbeddingBuilder[T, V]) EmbeddedPrivateInterface(v embeddedPrivateInterface) *StructEmbeddingBuilder[T, V] {
	b._embeddedPrivateInterface_ = v
	return b
}

func (b *StructEmbeddingBuilder[T, V]) PrivateEmbedded(v privateEmbedded) *StructEmbeddingBuilder[T, V] {
	b._privateEmbedded_ = v
	return b
}

func (b *StructEmbeddingBuilder[T, V]) Build() StructEmbedding[T, V] {
	out := StructEmbedding[T, V]{}

	out.ExternalEmbeddedInterface = b._ExternalEmbeddedInterface_
	out.SomeStruct = b._SomeStruct_
	out.SomeStructure = b._SomeStructure_
	out.embeddedPrivateInterface = b._embeddedPrivateInterface_
	out.privateEmbedded = b._privateEmbedded_

	return out
}

func NewTheThirdStruct(a SomeStructure, b *SomeStructure, c int, d int, fn func(**SomeStructure) AnotherStruct) TheThirdStruct {
	returnValue := TheThirdStruct{
		a:  a,
		b:  b,
		c:  c,
		d:  d,
		fn: fn,
	}

	return returnValue
}
