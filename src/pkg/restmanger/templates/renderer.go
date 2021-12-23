package templates

import "C"
import (
	"fmt"
	"html/template"
	"io"

	"github.com/kllpw/pgpez/src/pkg/restmanger/templates/header"
)

type Page = struct {
	Path string
	Key  string
}

var (
	Header   = Page{Path: "pkg/restmanger/templates/header/header.html", Key: "header.html"}
	Footer   = Page{Path: "pkg/restmanger/templates/footer/footer.html", Key: "footer.html"}
	Index    = Page{Path: "pkg/restmanger/templates/index/index.html", Key: "index.html"}
	Home     = Page{Path: "pkg/restmanger/templates/home/home.html", Key: "home.html"}
	Keys     = Page{Path: "pkg/restmanger/templates/keys/keys.html", Key: "keys.html"}
	Key      = Page{Path: "pkg/restmanger/templates/keys/key.html", Key: "key.html"}
	Error    = Page{Path: "pkg/restmanger/templates/error/error.html", Key: "error.html"}
	Messages = Page{Path: "pkg/restmanger/templates/messages/messages.html", Key: "messages.html"}
	Contacts = Page{Path: "pkg/restmanger/templates/contacts/contacts.html", Key: "contacts.html"}
	Unlock   = Page{Path: "pkg/restmanger/templates/unlock/unlock.html", Key: "unlock.html"}
)

type Renderer interface {
	RenderTemplate(w io.Writer, page Page, data PageData)
	ToggleDarkMode()
}

type PageData interface {
	GetData() interface{}
}

type rendererImpl struct {
	templates *template.Template
	darkmode  bool
}

var DefaultRenderer = &rendererImpl{
	templates: createTemplates(),
}

func createTemplates() *template.Template {
	t := template.New(Index.Key)
	t.ParseFiles(Index.Path)
	t.New(Header.Key).ParseFiles(Header.Path)
	t.New(Footer.Key).ParseFiles(Footer.Path)
	t.New(Home.Key).ParseFiles(Home.Path)
	t.New(Keys.Key).ParseFiles(Keys.Path)
	t.New(Key.Key).ParseFiles(Key.Path)
	t.New(Error.Key).ParseFiles(Error.Path)
	t.New(Messages.Key).ParseFiles(Messages.Path)
	t.New(Contacts.Key).ParseFiles(Contacts.Path)
	t.New(Unlock.Key).ParseFiles(Unlock.Path)
	return t
}

func (r *rendererImpl) RenderTemplate(w io.Writer, page Page, data PageData) {
	err := r.renderHeader(w, page)
	t := r.templates.Lookup(page.Key)
	err = t.Execute(w, data.GetData())
	if err != nil {
		fmt.Println(err)
	}
}

func (r *rendererImpl) ToggleDarkMode() {
	r.darkmode = !r.darkmode
}

func (r *rendererImpl) renderHeader(w io.Writer, page Page) error {

	hdr := r.templates.Lookup(Header.Key)
	hd := header.PageData{}
	menuActive := true
	if page == Unlock {
		menuActive = false
	}
	if page == Index {
		hd = header.PageData{
			PageTitle:      "pgpez",
			MenuActive:     menuActive,
			KeysActive:     false,
			MessagesActive: false,
			ContactsActive: false,
			DarkMode:       r.darkmode,
		}
	} else if page == Unlock {
		hd = header.PageData{
			PageTitle:      "pgpez - unlock",
			MenuActive:     menuActive,
			KeysActive:     false,
			MessagesActive: false,
			ContactsActive: false,
			DarkMode:       r.darkmode,
		}
	} else if page == Messages {
		hd = header.PageData{
			PageTitle:      "pgpez - messages",
			MenuActive:     menuActive,
			KeysActive:     false,
			MessagesActive: true,
			ContactsActive: false,
			DarkMode:       r.darkmode,
		}
	} else if page == Keys || page == Key {
		hd = header.PageData{
			PageTitle:      "pgpez - keys",
			MenuActive:     menuActive,
			KeysActive:     true,
			MessagesActive: false,
			ContactsActive: false,
			DarkMode:       r.darkmode,
		}
	} else if page == Contacts {
		hd = header.PageData{
			PageTitle:      "pgpez - contacts",
			MenuActive:     menuActive,
			KeysActive:     false,
			MessagesActive: false,
			ContactsActive: true,
			DarkMode:       r.darkmode,
		}
	}
	err := hdr.ExecuteTemplate(w, Header.Key, hd)
	if err != nil {

	}
	return err
}
