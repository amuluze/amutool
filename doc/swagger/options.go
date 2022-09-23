// Package swagger
// Date: 2022/9/23 13:28
// Author: Amu
// Description:
package swagger

type Option func(*Swagger)

func SetFilePath(filePath string) Option {
	return func(swagger *Swagger) {
		swagger.FilePath = filePath
	}
}
