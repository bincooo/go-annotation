// Code generated by Constructor annotation processor. DO NOT EDIT.
// versions:
//		go-annotation: 0.0.3-alpha
//		Constructor: 0.0.3

package generators

import (
	utils "github.com/YReshetko/go-annotation/annotations/mapper/utils"
	annotation "github.com/YReshetko/go-annotation/pkg"
	ast "go/ast"
)

type structureTypeGeneratorBuilder struct {
	value structureTypeGenerator
}

func newStructureTypeGeneratorBuilder() *structureTypeGeneratorBuilder {
	return &structureTypeGeneratorBuilder{}
}

func (b *structureTypeGeneratorBuilder) setNode(v annotation.Node) *structureTypeGeneratorBuilder {
	b.value.node = v
	return b
}

func (b *structureTypeGeneratorBuilder) setName(v string) *structureTypeGeneratorBuilder {
	b.value.name = v
	return b
}

func (b *structureTypeGeneratorBuilder) setParentAlias(v string) *structureTypeGeneratorBuilder {
	b.value.parentAlias = v
	return b
}

func (b *structureTypeGeneratorBuilder) setFields(v []*ast.Field) *structureTypeGeneratorBuilder {
	b.value.fields = v
	return b
}

func (b *structureTypeGeneratorBuilder) setFieldGenerators(v []*fieldGenerator) *structureTypeGeneratorBuilder {
	b.value.fieldGenerators = v
	return b
}

func (b *structureTypeGeneratorBuilder) build() *structureTypeGenerator {
	if b.value.fieldGenerators == nil {
		b.value.fieldGenerators = []*fieldGenerator{}
	}
	b.value.buildFields()

	return &b.value
}

type sliceTypeGeneratorBuilder struct {
	value sliceTypeGenerator
}

func newSliceTypeGeneratorBuilder() *sliceTypeGeneratorBuilder {
	return &sliceTypeGeneratorBuilder{}
}

func (b *sliceTypeGeneratorBuilder) setNode(v annotation.Node) *sliceTypeGeneratorBuilder {
	b.value.node = v
	return b
}

func (b *sliceTypeGeneratorBuilder) setParentAlias(v string) *sliceTypeGeneratorBuilder {
	b.value.parentAlias = v
	return b
}

func (b *sliceTypeGeneratorBuilder) setSliceType(v ast.Node) *sliceTypeGeneratorBuilder {
	b.value.sliceType = v
	return b
}

func (b *sliceTypeGeneratorBuilder) build() *sliceTypeGenerator {
	b.value.buildFields()

	return &b.value
}

type cacheBuilder struct {
	value cache
}

func newCache() *cacheBuilder {
	return &cacheBuilder{}
}

func (b *cacheBuilder) setVarPrefix(v string) *cacheBuilder {
	b.value.varPrefix = v
	return b
}

func (b *cacheBuilder) setImports(v map[Import]struct{}) *cacheBuilder {
	b.value.imports = v
	return b
}

func (b *cacheBuilder) setNode(v *utils.Node[string, string]) *cacheBuilder {
	b.value.node = v
	return b
}

func (b *cacheBuilder) build() *cache {

	return &b.value
}

type fieldGeneratorBuilder struct {
	value fieldGenerator
}

func newFieldGeneratorBuilder() *fieldGeneratorBuilder {
	return &fieldGeneratorBuilder{}
}

func (b *fieldGeneratorBuilder) setNode(v annotation.Node) *fieldGeneratorBuilder {
	b.value.node = v
	return b
}

func (b *fieldGeneratorBuilder) setAst(v ast.Node) *fieldGeneratorBuilder {
	b.value.ast = v
	return b
}

func (b *fieldGeneratorBuilder) setName(v string) *fieldGeneratorBuilder {
	b.value.name = v
	return b
}

func (b *fieldGeneratorBuilder) setAlias(v string) *fieldGeneratorBuilder {
	b.value.alias = v
	return b
}

func (b *fieldGeneratorBuilder) setParentAlias(v string) *fieldGeneratorBuilder {
	b.value.parentAlias = v
	return b
}

func (b *fieldGeneratorBuilder) setImportPkg(v string) *fieldGeneratorBuilder {
	b.value.importPkg = v
	return b
}

func (b *fieldGeneratorBuilder) setIsPointer(v bool) *fieldGeneratorBuilder {
	b.value.isPointer = v
	return b
}

func (b *fieldGeneratorBuilder) build() *fieldGenerator {
	b.value.buildFields()

	return &b.value
}

type MapperGeneratorBuilder struct {
	value MapperGenerator
}

func NewMapperGeneratorBuilder() *MapperGeneratorBuilder {
	return &MapperGeneratorBuilder{}
}

func (b *MapperGeneratorBuilder) IntName(v string) *MapperGeneratorBuilder {
	b.value.intName = v
	return b
}

func (b *MapperGeneratorBuilder) StructName(v string) *MapperGeneratorBuilder {
	b.value.structName = v
	return b
}

func (b *MapperGeneratorBuilder) Mgs(v []*methodGenerator) *MapperGeneratorBuilder {
	b.value.mgs = v
	return b
}

func (b *MapperGeneratorBuilder) Node(v annotation.Node) *MapperGeneratorBuilder {
	b.value.node = v
	return b
}

func (b *MapperGeneratorBuilder) IntType(v *ast.InterfaceType) *MapperGeneratorBuilder {
	b.value.intType = v
	return b
}

func (b *MapperGeneratorBuilder) Build() *MapperGenerator {
	if b.value.mgs == nil {
		b.value.mgs = []*methodGenerator{}
	}
	b.value.buildMethodGenerators()

	return &b.value
}

type methodGeneratorBuilder struct {
	value methodGenerator
}

func newMethodGeneratorBuilder() *methodGeneratorBuilder {
	return &methodGeneratorBuilder{}
}

func (b *methodGeneratorBuilder) setNode(v annotation.Node) *methodGeneratorBuilder {
	b.value.node = v
	return b
}

func (b *methodGeneratorBuilder) setName(v string) *methodGeneratorBuilder {
	b.value.name = v
	return b
}

func (b *methodGeneratorBuilder) setOverloading(v *overloading) *methodGeneratorBuilder {
	b.value.overloading = v
	return b
}

func (b *methodGeneratorBuilder) setInput(v []*ast.Field) *methodGeneratorBuilder {
	b.value.input = v
	return b
}

func (b *methodGeneratorBuilder) setOutput(v []*ast.Field) *methodGeneratorBuilder {
	b.value.output = v
	return b
}

func (b *methodGeneratorBuilder) setInputGenerators(v []*fieldGenerator) *methodGeneratorBuilder {
	b.value.inputGenerators = v
	return b
}

func (b *methodGeneratorBuilder) setOutputGenerator(v []*fieldGenerator) *methodGeneratorBuilder {
	b.value.outputGenerator = v
	return b
}

func (b *methodGeneratorBuilder) build() *methodGenerator {
	if b.value.inputGenerators == nil {
		b.value.inputGenerators = []*fieldGenerator{}
	}
	if b.value.outputGenerator == nil {
		b.value.outputGenerator = []*fieldGenerator{}
	}
	b.value.buildOutput()
	b.value.buildInput()

	return &b.value
}

func newOverloading(isIgnoreDefault bool) *overloading {
	returnValue := &overloading{
		isIgnoreDefault: isIgnoreDefault,
		mappings:        make(map[string]mapping),
	}

	return returnValue
}

type primitiveTypeGeneratorBuilder struct {
	value primitiveTypeGenerator
}

func newPrimitiveTypeGeneratorBuilder() *primitiveTypeGeneratorBuilder {
	return &primitiveTypeGeneratorBuilder{}
}

func (b *primitiveTypeGeneratorBuilder) setName(v string) *primitiveTypeGeneratorBuilder {
	b.value.name = v
	return b
}

func (b *primitiveTypeGeneratorBuilder) build() *primitiveTypeGenerator {
	b.value.postConstruct()

	return &b.value
}
