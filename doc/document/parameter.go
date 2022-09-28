// Package document
// Date: 2022/9/23 13:30
// Author: Amu
// Description:
package document

import (
	"reflect"

	v1 "gitee.com/amuluze/amutool/log"

	"gitee.com/amuluze/amutool/doc/constants"
	"github.com/fatih/structtag"
	"github.com/getkin/kin-openapi/openapi3"
)

func (d *Document) getParametersByModel(models ...interface{}) openapi3.Parameters {
	v1.Info("parameters models: ", models)
	parameters := openapi3.NewParameters()
	for _, model := range models {
		if model == nil {
			continue
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
			// 变量结构体的成员
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

				// embed 表示嵌套结构
				_, err = tags.Get(constants.EMBED)
				if err == nil {
					embedParameters := d.getParametersByModel(value.Interface())
					for _, embedParameter := range embedParameters {
						parameters = append(parameters, embedParameter)
					}
				}
				v1.Info("value interface: ", value)
				parameter := &openapi3.Parameter{
					Schema: openapi3.NewSchemaRef("", d.getSchemaByType(value.Interface(), true)),
				}
				jsonTag, err := tags.Get(constants.JSON)
				if err == nil {
					parameter.In = "json"
					parameter.Name = jsonTag.Name
				}
				queryTag, err := tags.Get(constants.QUERY)
				if err == nil {
					parameter.In = openapi3.ParameterInQuery
					parameter.Name = queryTag.Name
				}
				uriTag, err := tags.Get(constants.URI)
				if err == nil {
					parameter.In = openapi3.ParameterInPath
					parameter.Name = uriTag.Name
				}
				headerTag, err := tags.Get(constants.HEADER)
				if err == nil {
					parameter.In = openapi3.ParameterInHeader
					parameter.Name = headerTag.Name
				}
				cookieTag, err := tags.Get(constants.COOKIE)
				if err == nil {
					parameter.In = openapi3.ParameterInCookie
					parameter.Name = cookieTag.Name
				}
				if parameter.In == "" {
					continue
				}
				descriptionTag, err := tags.Get(constants.DESCRIPTION)
				if err == nil {
					parameter.WithDescription(descriptionTag.Name)
				}
				validateTag, err := tags.Get(constants.VALIDATE)
				if err == nil {
					parameter.WithRequired(validateTag.Name == "required")
					options := validateTag.Options
					if len(options) > 0 {
						parameter.Schema = d.getValidateSchemaByOptions(value.Interface(), options)
					}
				}
				defaultTag, err := tags.Get(constants.DEFAULT)
				if err == nil {
					parameter.Schema.Value.WithDefault(defaultTag.Name)
				}
				exampleTag, err := tags.Get(constants.EXAMPLE)
				if err == nil {
					parameter.Schema.Value.Example = exampleTag.Name
				}
				parameters = append(parameters, &openapi3.ParameterRef{
					Value: parameter,
				})
			}
		}
	}

	v1.Info("parameters: ", parameters)
	return parameters
}
