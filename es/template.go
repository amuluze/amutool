// Package es
// Date: 2023/3/27 19:14
// Author: Amu
// Description: 索引模版
package es

import (
	"context"

	"github.com/olivere/elastic/v7"
)

func (c *Client) TemplateExists(ctx context.Context, templateName string) (bool, error) {
	res, err := c.IndexGetIndexTemplate(templateName).Human(true).Pretty(true).Do(ctx)
	if err != nil {
		return false, err
	}
	_, exists := res.IndexTemplates.ByName(templateName)
	return exists, nil
}

func (c *Client) GetIndexTemplate(ctx context.Context, templateName string) (*elastic.IndicesGetIndexTemplateData, error) {
	res, err := c.GetIndexTemplate(ctx, templateName)
	return res, err
}

func (c *Client) PutIndexTemplate(ctx context.Context, templateName string, templateBody string) (bool, error) {
	res, err := c.IndexPutIndexTemplate(templateName).Pretty(true).BodyString(templateBody).Do(ctx)
	return res.Acknowledged, err
}

func (c *Client) DeleteIndexTemplate(ctx context.Context, templateName string) (bool, error) {
	res, err := c.IndexDeleteIndexTemplate(templateName).Human(true).Do(ctx)
	return res.Acknowledged, err
}
