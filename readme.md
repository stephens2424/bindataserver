# bindataserver
--
    import "github.com/stephens2424/bindataserver"

Bindataserver provides an easy way to serve assets loaded through
http://github.com/jteeuwen/go-bindata.

For example, given a directory like:

    main.go
    data/
      secret.txt
      assets/
        website.html
        website.css

and the following command to generate:

    $ go-bindata -prefix "/data" data/...

the code to serve assets would be:

    http.Handle("/assets/", http.StripPrefix("/", bindataserver.Bindata(_bindata)))

The assets would be available under a url such as
http://example.com/assets/website.html.

## Usage

#### type Bindata

```go
type Bindata map[string]func() ([]byte, error)
```

BindataServer augments a map produced by github.com/jteeuwen/go-bindata into an
http.Handler.

#### func (Bindata) ServeHTTP

```go
func (b Bindata) ServeHTTP(w http.ResponseWriter, r *http.Request)
```
ServeHTTP serves a request from the result of go-bindata. The request path must
match the asset name exactly.
