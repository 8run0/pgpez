package index

type PageData struct {
	PageTitle string
	Error     string
}

func (pd *PageData) GetData() interface{} {
	return pd
}
