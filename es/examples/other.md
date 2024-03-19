

```go

// checkPolicy 检查及更新索引声明周期策略
func checkPolicy(cfg *Config, cli *Client) (bool, error) {
	try := 0
	for _, indexName := range cfg.IndexNames {
		policyFileName := PolicyFilePrefix + indexName + PolicyFileSuffix
		policyCheckTextMd5 := iohelper.NewCheckTextMd5(policyFileName, "", cfg.ConfigPath, PolicyFilePrefix+indexName)
		policyChanged := policyCheckTextMd5.Change()

		for {
			exists, err := cli.ILMPolicyExists(context.Background(), policyFileName)
			fmt.Printf("ex: %v, err: %v\n", exists, err)
			try++
			if try > CreatePolicyRetry {
				panic("try to create policy over 50 times")
			}
			//if err != nil {
			//	continue
			//}

			// 更新 policy
			if !exists || policyChanged {
				// 存在，但是需要更新，所以先删除旧的
				if exists {
					err := cli.DeleteILMPolicy(context.TODO(), policyFileName)
					if err != nil {
						continue
					}
				}

				err := cli.PutILMPolicy(context.TODO(), policyFileName, cfg.ConfigPath)
				fmt.Printf("put err: %v\n", err)
				if err != nil {
					continue
				}

				err = policyCheckTextMd5.Write()
				exists = true
			}
			if exists {
				break
			}
			time.Sleep(2 * time.Second)
		}
	}
	return true, nil
}

// checkTemplate 检查及更新索引模版
func checkTemplate(cfg *Config, cli *Client) (bool, error) {
	try := 0
	for _, indexName := range cfg.IndexNames {
		templateFIleName := TemplateFilePrefix + indexName + TemplateFileSuffix
		templateCheckTextMd5 := iohelper.NewCheckTextMd5(templateFIleName, "", cfg.ConfigPath, TemplateFilePrefix+indexName)
		templateChanged := templateCheckTextMd5.Change()

		for {
			exists, err := cli.TemplateExists(context.Background(), templateFIleName)
			try++
			if try > CreateTemplateRetry {
				panic("try to create index template over 50 times")
			}

			if err != nil {
				fmt.Printf("template exists error: %v, exists: %v\n", err, exists)
			}
			fmt.Printf("template exists error: %v, exists: %v\n", err, exists)
			if !exists || templateChanged {
				// 存在，但是需要更新，所以先删除旧的
				if exists {
					err := cli.DeleteIndexTemplate(context.TODO(), templateFIleName)
					if err != nil {
						continue
					}
				}

				err := cli.PutIndexTemplate(context.TODO(), templateFIleName, cfg.ConfigPath)
				if err != nil {
					fmt.Printf("put index template error: %v\n", err)
					continue
				}

				err = templateCheckTextMd5.Write()
			}

			if exists {
				break
			}
			time.Sleep(2 * time.Second)
		}
	}
	return true, nil
}

// checkIndex 检查索引是否存在如果不存在则创建
func checkIndex(cfg *Config, cli *Client) (bool, error) {
	try := 0
	for _, indexName := range cfg.IndexNames {
		for {
			exists, err := cli.IndexExists(context.Background(), indexName)
			try++
			if try > CreateIndexRetry {
				panic("try to create index over 50 times")
			}
			if err != nil {
				fmt.Printf("index exists error: %v", err)
			}
			if !exists {
				newIndexName := fmt.Sprintf("<%s-{now/d}-00001>", indexName)
				indexBody := `{
					"aliases": {
						"%s": {"is_write_index": true}
					}
				}`

				indexBody = fmt.Sprintf(indexBody, indexName)

				res, err := cli.CreateIndex(context.TODO(), newIndexName, indexBody)
				if err != nil || res == false {
					fmt.Printf("create index failure: %v\n", err)
				}
				exists = true
			}
			if exists {
				break
			}
			time.Sleep(2 * time.Second)
		}
	}
	return true, nil
}
```