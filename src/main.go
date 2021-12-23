package main

import (
	"github.com/kllpw/pgpez/src/pkg/restmanger"
	_ "github.com/mattn/go-sqlite3"
	"github.com/webview/webview"
)

/*
func main() {
	ui, err := lorca.New("localhost:80/unlock", "", 1024, 768)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()
	mngr := restmanger.NewManager()
	go mngr.Start()
	ui.Load("http://localhost:80/unlock")
	<-ui.Done()
}
*/

func main() {
	mngr := restmanger.NewManager()
	go mngr.Start()
	w := webview.New(true)
	defer w.Destroy()
	w.SetTitle("pgpez")
	w.SetSize(1024, 768, webview.HintNone)
	w.Navigate("http://localhost:80")
	w.Run()
}
