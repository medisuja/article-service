package transformers

type (
	Transformer struct {
		Data interface{} `json:"data"`
	}

	CollectionTransformer struct {
		Data []interface{} `json:"data"`
	}
)
