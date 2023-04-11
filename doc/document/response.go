// Package document
// Date: 2022/9/23 13:30
// Author: Amu
// Description:
package document

import (
	"reflect"

	"gitee.com/amuluze/amutool/doc/constants"
	"github.com/fatih/structtag"
	"github.com/getkin/kin-openapi/openapi3"
)

type FailedResponse struct {
	Err string `json:"err"` // 响应错误，来自 service 层的错误信息
	Msg string `json:"msg"` // 错误消息，来自 controller 层的错误信息
}

func (d *Document) getResponses(response interface{}) openapi3.Responses {
	ret := openapi3.NewResponses()
	var schema *openapi3.Schema
	var content openapi3.Content

	ret["204"] = &openapi3.ResponseRef{Value: &openapi3.Response{Content: nil, Description: &constants.NOCONTENT}}
	ret["403"] = &openapi3.ResponseRef{Value: &openapi3.Response{Content: nil, Description: &constants.FORBIDDEN}}
	ret["404"] = &openapi3.ResponseRef{Value: &openapi3.Response{Content: nil, Description: &constants.NOTFOUND}}

	if _, ok := d.OpenAPI.Components.Schemas[reflect.TypeOf(&FailedResponse{}).Elem().Name()]; ok {
		content = openapi3.Content{"application/json": &openapi3.MediaType{Schema: &openapi3.SchemaRef{Ref: "#/components/schemas/" + d.OpenAPI.Components.Schemas[reflect.TypeOf(&FailedResponse{}).Elem().Name()].Value.Title}}}
	} else {
		content = openapi3.NewContentWithJSONSchema(d.getResponseSchemaByModel(&FailedResponse{}))
	}
	ret["400"] = &openapi3.ResponseRef{
		Value: &openapi3.Response{
			Content:     content,
			Description: &constants.FAILED,
		},
	}
	ret["500"] = &openapi3.ResponseRef{
		Value: &openapi3.Response{
			Content:     content,
			Description: &constants.FAILED,
		},
	}

	if response == nil {
		return ret
	}
	if _, ok := d.OpenAPI.Components.Schemas[reflect.TypeOf(response).Elem().Name()]; ok {
		ret["200"] = &openapi3.ResponseRef{
			Value: &openapi3.Response{
				Content:     openapi3.Content{"application/json": &openapi3.MediaType{Schema: &openapi3.SchemaRef{Ref: "#/components/schemas/" + d.OpenAPI.Components.Schemas[reflect.TypeOf(response).Elem().Name()].Value.Title}}},
				Description: &constants.SUCCESS,
			},
		}
	} else {
		schema = d.getResponseSchemaByModel(response)
		content = openapi3.NewContentWithJSONSchema(schema)

		ret["200"] = &openapi3.ResponseRef{
			Value: &openapi3.Response{
				Content:     content,
				Description: &constants.SUCCESS,
			},
		}
	}
	return ret
}

func (d *Document) getResponseSchemaByModel(model interface{}) *openapi3.Schema {
	schema := openapi3.NewObjectSchema()
	if model == nil {
		return schema
	}

	// 获取 model 的类型对象和值对象
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

			// 如果 field 在 components 中，则映射到 components
			if _, ok := d.OpenAPI.Components.Schemas[field.Name]; ok {
				schema.Properties[field.Name] = &openapi3.SchemaRef{Ref: "#/components/schemas/" + d.OpenAPI.Components.Schemas[field.Name].Value.Title}
				continue
			}

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
				embedSchema := d.getResponseSchemaByModel(value.Interface())
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
				fieldSchema.Default = defaultTag.Name
			}
			exampleTag, err := tags.Get(constants.EXAMPLE)
			if err == nil {
				fieldSchema.Example = exampleTag.Name
			}
		}
	} else if type_.Kind() == reflect.Slice {
		schema = openapi3.NewArraySchema()
		schema.Items = &openapi3.SchemaRef{Value: d.getResponseSchemaByModel(reflect.New(type_.Elem()).Elem().Interface())}
	} else {
		schema = d.getSchemaByType(model, false)
	}
	return schema
}
