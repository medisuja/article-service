package requests

type (
	ParamArticle struct {
		Author string `json:"author" valid:"required~parameter is empty"`
		Title  string `json:"title"  valid:"required~parameter is empty"`
		Body   string `json:"body"   valid:"required~parameter is empty"`
	}
)
