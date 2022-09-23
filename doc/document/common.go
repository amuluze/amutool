// Package document
// Date: 2022/9/23 13:29
// Author: Amu
// Description:
package document

import (
	"mime/multipart"
	"strconv"
	"strings"
	"time"

	"github.com/getkin/kin-openapi/openapi3"
)

func (d *Document) getSchemaByType(t interface{}, request bool) *openapi3.Schema {
	var schema *openapi3.Schema
	var m float64
	m = float64(0)
	switch t.(type) {
	case int, int8, int16, *int, *int8, *int16:
		schema = openapi3.NewIntegerSchema()
	case uint, uint8, uint16, *uint, *uint8, *uint16:
		schema = openapi3.NewIntegerSchema()
		schema.Min = &m
	case int32, *int32:
		schema = openapi3.NewInt32Schema()
	case uint32, *uint32:
		schema = openapi3.NewInt32Schema()
		schema.Min = &m
	case int64, *int64:
		schema = openapi3.NewInt64Schema()
	case uint64, *uint64:
		schema = openapi3.NewInt64Schema()
		schema.Min = &m
	case string, *string:
		schema = openapi3.NewStringSchema()
	case time.Time, *time.Time:
		schema = openapi3.NewDateTimeSchema()
	case float32, float64, *float32, *float64:
		schema = openapi3.NewFloat64Schema()
	case bool, *bool:
		schema = openapi3.NewBoolSchema()
	case []byte, *[]byte:
		schema = openapi3.NewBytesSchema()
	case *multipart.FileHeader:
		schema = openapi3.NewStringSchema()
		schema.Format = "binary"
	case []*multipart.FileHeader:
		schema = openapi3.NewArraySchema()
		schema.Items = &openapi3.SchemaRef{
			Value: &openapi3.Schema{
				Type:   "string",
				Format: "binary",
			},
		}
	default:
		if request {
			schema = d.getRequestSchemaByModel(t)
		} else {
			schema = d.getResponseSchemaByModel(t)
		}
	}
	return schema
}

func (d *Document) getValidateSchemaByOptions(value interface{}, options []string) *openapi3.SchemaRef {
	// 字段参数校验 tag 解析，支持 枚举-oneof 最大值-max 最小值-min 长度-len
	schema := openapi3.NewSchemaRef("", d.getSchemaByType(value, true))
	for _, option := range options {
		if strings.HasPrefix(option, "oneof=") {
			optionItems := strings.Split(option[6:], " ")
			enums := make([]interface{}, len(optionItems))
			for i, optionItem := range optionItems {
				enums[i] = optionItem
			}
			schema.Value.WithEnum(enums...)
		}
		if strings.HasPrefix(option, "max=") {
			value, err := strconv.ParseFloat(option[4:], 64)
			if err != nil {
				panic(err)
			}
			schema.Value.WithMax(value)
		}
		if strings.HasPrefix(option, "min=") {
			value, err := strconv.ParseFloat(option[4:], 64)
			if err != nil {
				panic(err)
			}
			schema.Value.WithMin(value)
		}
		if strings.HasPrefix(option, "len=") {
			value, err := strconv.ParseInt(option[4:], 10, 64)
			if err != nil {
				panic(err)
			}
			schema.Value.WithLength(value)
		}
	}
	return schema
}
