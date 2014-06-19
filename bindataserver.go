// Bindataserver provides an easy way to serve assets loaded through http://github.com/jteeuwen/go-bindata.
//
// For example, given a directory like:
//
//    main.go
//    data/
//      secret.txt
//      assets/
//        website.html
//        website.css
//
// and the following command to generate:
//
//    $ go-bindata -prefix "/data" data
//
// the code to serve assets would be:
//
//     http.Handle("/assets/", http.StripPrefix("/", bindataserer.Bindata(_bindata)))
//
// The assets would be available under a url such as http://exmple.com/assets/website.html.
//
package bindataserver

import "net/http"

// BindataServer augments a map produced by github.com/jteeuwen/go-bindata into
// an http.Handler.
type Bindata map[string]func() ([]byte, error)

// ServeHTTP serves a request from the result of go-bindata. The request path must
// match the asset name exactly.
func (b Bindata) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f, ok := b[r.URL.Path]
	if !ok {
		http.NotFound(w, r)
		return
	}
	data, err := f()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}
