package client

// SchemaConfigLoader schema配置加载器转换
type SchemaConfigLoader struct {
}

func NewSchemaConfigLoader() *SchemaConfigLoader {
	return &SchemaConfigLoader{}
}

func (c *SchemaConfigLoader) LoaderConfig(schema string, url string) (err error) {
	return nil
}
