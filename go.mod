module github.com/kllpw/pgpez

go 1.14

replace golang.org/x/crypto => github.com/ProtonMail/crypto v0.0.0-20200416114516-1fa7f403fb9c

require (
	github.com/ProtonMail/gopenpgp/v2 v2.0.1
	github.com/gorilla/mux v1.7.4
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/webview/webview v0.0.0-20200724072439-e0c01595b361
	github.com/zserge/lorca v0.1.9
	golang.org/x/crypto v0.0.0-20191011191535-87dc89f01550 // indirect
)
