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

func (d *Document) getRequestBodyByModel(model interface{}) *openapi3.RequestBodyRef {
	body := &openapi3.RequestBodyRef{
		Value: openapi3.NewRequestBody(),
	}
	if model == nil {
		return body
	}

	if _, ok := d.OpenAPI.Components.Schemas[reflect.TypeOf(model).Elem().Name()]; ok {
		body.Value.Required = true
		body.Value.Content = openapi3.Content{
			"application/json": &openapi3.MediaType{Schema: &openapi3.SchemaRef{Ref: "#/components/schemas/" + d.OpenAPI.Components.Schemas[reflect.TypeOf(model).Elem().Name()].Value.Title}},
		}
	} else {
		schema := d.getRequestSchemaByModel(model)
		body.Value.Required = true
		body.Value.Content = openapi3.NewContentWithSchema(schema, []string{"application/json"})
	}
	return body
}

func (d *Document) getRequestSchemaByModel(model interface{}) *openapi3.Schema {
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
				embedSchema := d.getRequestSchemaByModel(value.Interface())
				for key, embedProperty := range embedSchema.Properties {
					schema.Properties[key] = embedProperty
				}
				for _, name := range embedSchema.Required {
					schema.Required = append(schema.Required, name)
				}
			}

			// json 表示 content-type 为 application/json
			// form 表示 content-type 为 application/x-www-form-urlencoded
			var fieldSchema *openapi3.Schema
			tag, err := tags.Get(constants.JSON)
			if err == nil {
				// 这里的 value 为空，通过 Interface() 方法获取其原值，用于判断其原值的类型，
				//fmt.Printf("===============> value interface: %+v value interface type: %+v\n", value.Interface(), reflect.TypeOf(value.Interface()))
				fieldSchema := d.getSchemaByType(value.Interface(), true)
				schema.Properties[tag.Name] = openapi3.NewSchemaRef("", fieldSchema)
			} else if tag, err = tags.Get(constants.FORM); err == nil {
				fieldSchema := d.getSchemaByType(value.Interface(), true)
				schema.Properties[tag.Name] = openapi3.NewSchemaRef("", fieldSchema)
			} else {
				continue
			}

			// validate 表示该索引字段的验证规则
			validateTag, err := tags.Get(constants.VALIDATE)
			if err == nil {
				if validateTag.Name == "required" {
					schema.Required = append(schema.Required, tag.Name)
				}
				options := validateTag.Options
				if len(options) > 0 {
					schema.Properties[tag.Name] = d.getValidateSchemaByOptions(value.Interface(), options)
					fieldSchema = schema.Properties[tag.Name].Value
				}
			}

			// 如果 fieldSchema 为空则直接返回
			if fieldSchema == nil {
				continue
			}

			// 解析字段默认值
			defaultTag, err := tags.Get(constants.DEFAULT)
			if err == nil {
				fieldSchema.Default = defaultTag.Name
			}
			// 解析字段示例值
			exampleTag, err := tags.Get(constants.EXAMPLE)
			if err == nil {
				fieldSchema.Example = exampleTag.Name
			}
			// 解析字段描述
			descriptionTag, err := tags.Get(constants.DESCRIPTION)
			if err == nil {
				fieldSchema.Description = descriptionTag.Name
			}
		}
	} else if type_.Kind() == reflect.Slice {
		schema = openapi3.NewArraySchema()
		schema.Items = &openapi3.SchemaRef{Value: d.getRequestSchemaByModel(reflect.New(type_.Elem()).Elem().Interface())}
	} else {
		schema = d.getSchemaByType(model, true)
	}
	return schema
}
