package mux

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
	"saviusz/melon-server/src/errors"
)

type Mux struct {
	tree Tree
}

func New() Mux {
	return Mux{
		tree: *NewTree(),
	}
}

func (mux Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := Context{req: r, w: w}

	code, obj, err := mux.findHandler(r)(ctx)

	var err2 error

	if err != nil {
		code, obj, err2 = handleError(ctx, err)
	}

	if err2 != nil {
		w.Write([]byte("500 Internal server error"))
		return
	}

	if err := ctx.createResponse(obj, code); err != nil {
		w.Write([]byte("500 Internal server error"))
		return
	}

}

func handleError(ctx Context, e error) (code int, obj interface{}, err error) {
	switch fe := e.(type) {
	case errors.ApiError:
		code = fe.Code
		obj = fe
	default:
		code = 500
		obj = errors.ApiError{
			Type:     "/#",
			Title:    "Internal server error",
			Instance: "0",
			Detail:   e.Error(),
		}
	}

	return
}

func (mux *Mux) findHandler(r *http.Request) EndpointHandler {
	handler, params, err := mux.tree.Find(r.URL.Path, r.Method)
	if err == nil {
		return func(ctx Context) (int, interface{}, error) {
			newCtx := ctx
			newCtx.Params = params
			return handler(newCtx)
		}
	}
	return NotFoundHandler
}

func (mux *Mux) GET(path string, handler EndpointHandler) {
	err := mux.tree.Add(path, "GET", handler)
	if err != nil {
		log.Fatalf("Error adding path %s %s: %v", "GET", path, err)
	}
}

func (mux *Mux) POST(path string, handler EndpointHandler) {
	err := mux.tree.Add(path, "POST", handler)
	if err != nil {
		log.Fatalf("Error adding path %s %s: %v", "POST", path, err)
	}
}

func NotFoundHandler(ctx Context) (int, interface{}, error) {
	return 0, nil, errors.ApiError{
		Title:  "Not found",
		Code:   404,
		Detail: "Couldn't find resource for current URI",
		Type:   "about:blank",
	}
}

type EndpointHandler func(Context) (int, interface{}, error)

func getRequestedMime(req http.Request) string {
	acceptHeader := req.Header.Get("Accept")
	contentHeader := req.Header.Get("Content-Type")

	if acceptHeader != "" {
		return acceptHeader
	} else if contentHeader != "" {
		return contentHeader
	} else {
		return "application/json"
	}
}

func marshal(i interface{}, mime string) (output_data []byte, mime_o string, err error) {
	switch mime {
	default:
	case "application/json":
		output_data, err = json.Marshal(i)
		mime_o = "application/json; charset=utf-8"
	case "application/xml":
		output_data, err = xml.MarshalIndent(i, "", " ")
		mime_o = "application/xml"
	}
	return
}
