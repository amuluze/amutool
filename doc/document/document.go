// Package document
// Date: 2022/9/23 13:29
// Author: Amu
// Description:
package document

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"

	"gitee.com/amuluze/amutool/doc/router"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/ghodss/yaml"
)

type DocInfoLicense struct {
	Name string
	URL  string
}

type DocInfoContact struct {
	Name  string
	URL   string
	Email string
}

type DocInfo struct {
	Title       string          // 标题
	Description string          // api 描述
	Version     string          // 版本
	License     *DocInfoLicense // 许可证信息
	Contact     *DocInfoContact // 联系人信息
}

type DocSecurityScheme struct {
	Name string
	In   string
	Type string
}

type DocServer struct {
	URL         string // 服务器url
	Description string // 描述
}

type Document struct {
	OpenAPI         *openapi3.T                          // openapi 实例
	OpenAPIJsonFile string                               // json 格式的 api 文档存放路径
	OpenAPIYamlFile string                               // yaml 格式的 api 文档存放路径
	Paths           map[string]map[string]*router.Router // 存储 AddPath 添加的 Path 信息
	Components      []interface{}                        // 存储 AddComponent 添加的 Component 信息
	Servers         openapi3.Servers
}

func New(title, description, version string, options ...Option) *Document {
	info := &openapi3.Info{
		Title:       title,
		Description: description,
		Version:     version,
	}

	openapi := &openapi3.T{
		Info:       info,
		OpenAPI:    "3.0.0",
		Components: openapi3.NewComponents(),
		Tags:       openapi3.Tags{},
		Paths:      map[string]*openapi3.PathItem{},
		Security:   openapi3.SecurityRequirements{map[string][]string{"http": {}}},
	}

	doc := &Document{
		OpenAPI:         openapi,
		OpenAPIJsonFile: "./docs/openapi.json",
		OpenAPIYamlFile: "./docs/openapi.yaml",
	}

	for _, option := range options {
		option(doc)
	}

	return doc
}

func (d *Document) Build() {
	d.buildComponents()
	d.buildPaths()
	_ = d.WriteToJson()
	_ = d.WriteToYaml()
}

func (d *Document) buildComponents() {
	schemas := make(map[string]*openapi3.SchemaRef)
	for _, component := range d.Components {
		if component == nil {
			continue
		}

		name, schema := d.getComponentsByModel(component)
		schemas[name] = schema
	}
	d.OpenAPI.Components.Schemas = schemas
}

func (d *Document) buildPaths() {
	paths := make(openapi3.Paths)
	for path, m := range d.Paths {
		pathItem := &openapi3.PathItem{}
		for method, r := range m {
			operation := &openapi3.Operation{
				Tags:      r.Tags,
				Summary:   r.Description,
				Responses: d.getResponses(r.Responses),
			}
			requestBody := d.getRequestBodyByModel(r.Request)
			parameters := d.getParametersByModel(r.Request, r.Header, r.Cookie)
			switch method {
			case http.MethodGet:
				pathItem.Get = operation
				operation.Parameters = parameters
			case http.MethodPost:
				pathItem.Post = operation
			case http.MethodPut:
				pathItem.Put = operation
			case http.MethodDelete:
				pathItem.Delete = operation
			}
			if method != http.MethodGet && requestBody.Value.Content != nil {
				operation.RequestBody = requestBody
			}

			if !r.HasSecurity {
				operation.Security = &openapi3.SecurityRequirements{}
			}
		}
		paths[d.fixPath(path)] = pathItem
	}
	d.OpenAPI.Paths = paths
}

func (d *Document) AddPath(path, method, description, tag string, request interface{}, responses interface{}, hasSecurity bool) {
	tags := []string{tag}
	rt := router.New(
		path,
		method,
		description,
		tags,
		router.Requests(request),
		router.Responses(responses),
		router.HasSecurity(hasSecurity),
	)
	if d.Paths == nil {
		d.Paths = make(map[string]map[string]*router.Router)
	}
	if d.Paths[path] == nil {
		d.Paths[path] = make(map[string]*router.Router)
	}
	d.Paths[path][method] = rt
}

func (d *Document) AddComponent(component interface{}) {
	d.Components = append(d.Components, component)
}

// /:id -> /{id}
func (d *Document) fixPath(path string) string {
	reg := regexp.MustCompile("/:([0-9a-zA-Z]+)")
	return reg.ReplaceAllString(path, "/{${1}}")
}

func (d *Document) MarshalJSON() ([]byte, error) {
	return d.OpenAPI.MarshalJSON()
}

func (d *Document) MarshalYAML() ([]byte, error) {
	res, _ := d.MarshalJSON()
	return yaml.JSONToYAML(res)
}

func (d *Document) WriteToJson() error {
	abs, err := filepath.Abs(filepath.Dir(d.OpenAPIJsonFile))
	if err != nil {
		return err
	}

	_, err = os.Stat(abs)
	if err != nil {
		err := os.Mkdir(abs, 0766)
		if err != nil {
			return err
		}
	}
	fileObj, err := os.OpenFile(d.OpenAPIJsonFile, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
	if err != nil {
		fmt.Printf("Failed to open json file %s: %v", d.OpenAPIJsonFile, err)
		return err
	}
	content, err := d.MarshalJSON()

	if err != nil {
		fmt.Printf("marshal swagger json error: %v", err)
		return err
	}

	if _, err := io.WriteString(fileObj, string(content)); err == nil {
		fmt.Println("success generated json file")
		return nil
	}
	return err
}

func (d *Document) WriteToYaml() error {
	abs, err := filepath.Abs(filepath.Dir(d.OpenAPIYamlFile))
	if err != nil {
		return err
	}

	_, err = os.Stat(abs)
	if err != nil {
		err := os.Mkdir(abs, 0766)
		if err != nil {
			return err
		}
	}
	fileObj, err := os.OpenFile(d.OpenAPIYamlFile, os.O_RDWR|os.O_TRUNC|os.O_CREATE, 0766)
	if err != nil {
		fmt.Printf("Failed to open yaml file %s: %v", d.OpenAPIYamlFile, err)
		return err
	}
	content, err := d.MarshalYAML()

	if err != nil {
		fmt.Printf("marshal swagger yaml error: %v", err)
		return err
	}

	if _, err := io.WriteString(fileObj, string(content)); err == nil {
		fmt.Println("success generated yaml file.")
		return nil
	}
	return err
}
