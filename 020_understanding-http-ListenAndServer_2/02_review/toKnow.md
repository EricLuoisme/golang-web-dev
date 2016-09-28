# ListenAndServe

- [http.ListenAndServe](https://godoc.org/net/http#ListenAndServe)
``` Go
func ListenAndServe(addr string, handler Handler) error
```

********************
# HANDLER

- [http.Handler](https://godoc.org/net/http#Handler)
``` Go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```
# HANDLE

- [http.Handle](https://godoc.org/net/http#Handle)
``` Go
func Handle(pattern string, handler Handler)
```
********************
# SERVEMUX

- [http.NewServeMux](https://godoc.org/net/http#NewServeMux)
``` Go
func NewServeMux() *ServeMux
```

- [http.ServeHTTP](https://godoc.org/net/http#ServeMux.ServeHTTP)
``` Go
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request)
```

********************
# HANDLERFUNC & func(ResponseWriter, *Request)

- [http.HandlerFunc](https://godoc.org/net/http#HandlerFunc)
``` Go
type HandlerFunc func(ResponseWriter, *Request)
```

- [http.ServeHTTP](https://godoc.org/net/http#HandlerFunc.ServeHTTP)
``` Go
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)
```

# HANDLEFUNC

- [http.HandleFunc](https://godoc.org/net/http#HandleFunc)
``` Go
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
```

********************