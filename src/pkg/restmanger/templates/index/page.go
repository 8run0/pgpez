package index

type PageData struct {
	PageTitle      string
	WelcomeMessage string
}

func (pd *PageData) GetData() interface{} {
	return pd
}
