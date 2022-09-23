// Package document
// Date: 2022/9/23 13:29
// Author: Amu
// Description:
package document

import (
	"encoding/json"
	"reflect"
	"strconv"

	"gitee.com/amuluze/amutool/doc/constants"
	"gitee.com/amuluze/amutool/log"
	"github.com/fatih/structtag"
	"github.com/getkin/kin-openapi/openapi3"
)

func (d *Document) getComponentsByModel(model interface{}) (string, *openapi3.SchemaRef) {
	if model == nil {
		return "", nil
	}

	body := &openapi3.SchemaRef{
		Value: openapi3.NewAllOfSchema(),
	}
	componentTitle := reflect.TypeOf(model).Elem().Name()
	body.Value = d.getComponentSchemaByModel(model)
	body.Value.Title = componentTitle

	return componentTitle, body
}

func (d *Document) getComponentSchemaByModel(model interface{}) *openapi3.Schema {
	schema := openapi3.NewObjectSchema()
	if model == nil {
		return schema
	}

	type_ := reflect.TypeOf(model)
	value_ := reflect.ValueOf(model)
	// 如果类型对象和值对象的种类是 指针，则获取指针指向的元素类型
	if type_.Kind() == reflect.Ptr {
		type_ = type_.Elem()
	}
	if value_.Kind() == reflect.Ptr {
		value_ = value_.Elem()
	}

	if type_.Kind() == reflect.Struct {
		// 遍历结构体的成员
		for i := 0; i < type_.NumField(); i++ {
			if type_.Kind() == reflect.Struct && value_.Kind() == reflect.Invalid {
				value_ = reflect.New(type_).Elem()
			}
			// 根据索引，获取索引对应结构体字段的类型信息（包括字段名，字段类型，字段索引，字段tag等信息）
			field := type_.Field(i)
			// 根据索引，获取索引对应结构体字段的值信息（因为这里是一个没有初始化的实例，因此值信息为空）
			value := value_.Field(i)

			// 根据索引，获取该索引字段的 tags
			tags, err := structtag.Parse(string(field.Tag))
			if err != nil {
				panic(err)
			}

			// exclude 表示放弃解析该字段
			_, err = tags.Get(constants.EXCLUDE)
			if err == nil {
				continue
			}

			// embed 表示嵌套结构
			_, err = tags.Get(constants.EMBED)
			if err == nil {
				embedSchema := d.getComponentSchemaByModel(value.Interface())
				for key, embedProperty := range embedSchema.Properties {
					schema.Properties[key] = embedProperty
				}
				for _, name := range embedSchema.Required {
					schema.Required = append(schema.Required, name)
				}
			}

			var fieldSchema *openapi3.Schema
			tag, err := tags.Get(constants.JSON)
			if err == nil {
				fieldSchema = d.getSchemaByType(value.Interface(), false)
				schema.Properties[tag.Name] = openapi3.NewSchemaRef("", fieldSchema)
			} else {
				continue
			}

			validateTag, err := tags.Get(constants.VALIDATE)
			if err == nil && validateTag.Name == "required" {
				schema.Required = append(schema.Required, tag.Name)
			}
			descriptionTag, err := tags.Get(constants.DESCRIPTION)
			if err == nil {
				fieldSchema.Description = descriptionTag.Name
			}
			defaultTag, err := tags.Get(constants.DEFAULT)
			if err == nil {
				log.Info(">>>>>>>>", field.Type.String(), "    ", defaultTag.Value())
				switch field.Type.String() {
				case "int":
					defaultValue, _ := strconv.ParseInt(defaultTag.Value(), 10, 64)
					log.Info(defaultValue)
					fieldSchema.Default = defaultValue
				default:
					fieldSchema.Default = defaultTag.Value()
				}

			}
			exampleTag, err := tags.Get(constants.EXAMPLE)
			if err == nil {
				fieldSchema.Example = exampleTag.Name
			}
			titleTag, err := tags.Get(constants.TITLE)
			if err == nil {
				fieldSchema.Title = titleTag.Name
			}
			enumTag, err := tags.Get(constants.ENUM)
			if err == nil {
				var ee []interface{}
				err := json.Unmarshal([]byte(enumTag.Value()), &ee)
				if err != nil {
					log.Error(enumTag.String(), "设置错误")
				}
				fieldSchema.Enum = ee
			}
		}
	} else if type_.Kind() == reflect.Slice {
		schema = openapi3.NewArraySchema()
		schema.Items = &openapi3.SchemaRef{Value: d.getResponseSchemaByModel(reflect.New(type_.Elem()).Elem().Interface())}
	} else {
		schema = d.getSchemaByType(model, false)
	}
	return schema

	return nil
}
