// Package es
// Date: 2023/4/20 17:03
// Author: Amu
// Description:
package es

import "github.com/olivere/elastic/v7"

const (
	CreateIndexRetry    = 50
	CreatePolicyRetry   = 50
	CreateTemplateRetry = 50
	PolicyFilePrefix    = "policy_"
	PolicyFileSuffix    = ".json"
	TemplateFilePrefix  = "template_"
	TemplateFileSuffix  = ".json"
)

var (
	NewMatchQuery             = elastic.NewMatchQuery
	NewTermQuery              = elastic.NewTermQuery
	NewTermsQuery             = elastic.NewTermsQuery
	NewTermsQueryFrom         = elastic.NewTermsQueryFromStrings
	NewBoolQuery              = elastic.NewBoolQuery
	NewRangeQuery             = elastic.NewRangeQuery
	NewNestedQuery            = elastic.NewNestedQuery
	NewMatchAllQuery          = elastic.NewMatchAllQuery
	NewMatchPhraseQuery       = elastic.NewMatchPhraseQuery
	NewMatchPhrasePrefixQuery = elastic.NewMatchPhrasePrefixQuery
	NewRegexpQuery            = elastic.NewRegexpQuery

	NewTermsAggregation     = elastic.NewTermsAggregation
	NewAvgAggregation       = elastic.NewAvgAggregation
	NewDateRangeAggregation = elastic.NewDateRangeAggregation
	NewFilterAggregation    = elastic.NewFilterAggregation
	NewFiltersAggregation   = elastic.NewFiltersAggregation
	NewSumAggregation       = elastic.NewSumAggregation
	NewMaxAggregation       = elastic.NewMaxAggregation
	NewMinAggregation       = elastic.NewMinAggregation
)
