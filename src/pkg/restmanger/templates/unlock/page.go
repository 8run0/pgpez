package unlock

type PageData struct {
	PageTitle  string
	CurrentDir string
	Passphrase string
}

func (pd *PageData) GetData() interface{} {
	return pd
}
