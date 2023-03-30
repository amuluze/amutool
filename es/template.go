// Package es
// Date: 2023/3/27 19:14
// Author: Amu
// Description: 索引模版
package es

import (
	"context"
	"fmt"
)

const (
	CreateTemplateRetry = 50
	TemplateFilePrefix  = "template_"
	TemplateFileSuffix  = ".json"
)

func (c *Client) TemplateExists(ctx context.Context, templateName string) (bool, error) {
	exists, err := c.GetIndexTemplate(ctx, templateName)
	if err != nil {
		return false, err
	}
	fmt.Printf("exists: %#v\n", exists)
	return true, nil
}

func (c *Client) GetIndexTemplate(ctx context.Context, templateName string) (map[string]interface{}, error) {
	do, err := c.IndexGetIndexTemplate(templateName).Human(true).Pretty(true).Do(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Printf("do: %#v\n", do)
	return nil, nil
}

func (c *Client) PutIndexTemplate(ctx context.Context, templateName string, templatePath string) error {
	fileName := fmt.Sprint(templatePath, "/", TemplateFilePrefix, templateName, TemplateFileSuffix)
	bodyString, err := c.ReadFile(fileName)
	if err != nil {
		return err
	}

	res, err := c.IndexPutIndexTemplate(templateName).Pretty(true).BodyString(bodyString).Do(context.TODO())
	if err != nil || !res.Acknowledged {
		fmt.Printf("template body string: %v\n", err)
		return err
	}
	return nil
}

func (c *Client) DeleteIndexTemplate(ctx context.Context, templateName string) error {
	res, err := c.IndexDeleteIndexTemplate(templateName).Human(true).Do(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("res: %#v\n", res)
	return nil
}
