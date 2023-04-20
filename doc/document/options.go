// Package document
// Date: 2022/9/23 13:30
// Author: Amu
// Description:
package document

import (
	"gitee.com/amuluze/amutool/doc/security"
	"github.com/getkin/kin-openapi/openapi3"
)

type Option func(document *Document)

func SetContact(contact DocInfoContact) Option {
	c := &openapi3.Contact{
		Name:  contact.Name,
		URL:   contact.URL,
		Email: contact.Email,
	}
	return func(doc *Document) {
		doc.OpenAPI.Info.Contact = c
	}
}
func SetLicense(license DocInfoLicense) Option {
	l := &openapi3.License{
		Name: license.Name,
		URL:  license.URL,
	}
	return func(doc *Document) {
		doc.OpenAPI.Info.License = l
	}
}

func SetSecuritySchema(schema security.ISecurity) Option {
	return func(doc *Document) {
		doc.OpenAPI.Components.SecuritySchemes = map[string]*openapi3.SecuritySchemeRef{
			"httpx": {Value: schema.Scheme()},
		}
	}
}

//func SetSecuritySchema(schema DocSecurityScheme) Option {
//	return func(doc *Document) {
//		doc.OpenAPI.Components.SecuritySchemes = map[string]*openapi3.SecuritySchemeRef{
//			"ApiAuth": {Value: &openapi3.SecurityScheme{Type: schema.Type, Name: schema.Name, In: schema.In}},
//		}
//	}
//}
