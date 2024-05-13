package mux

import (
	"io"
	"net/http"
	"saviusz/melon-server/src/errors"
)

type Context struct {
	Req    *http.Request
	w      http.ResponseWriter
	Params map[string]string
}

func (ctx *Context) createResponse(output interface{}, code int) (err error) {
	req_mime := getRequestedMime(*ctx.Req)
	output_data, res_mime, err := marshal(output, req_mime)

	if err != nil {
		return err
	}

	ctx.w.Header().Add("Content-Type", res_mime)
	ctx.w.WriteHeader(code)
	ctx.w.Write(output_data)

	return nil
}

func (ctx *Context) getProvidedMime() string {
	return ctx.Req.Header.Get("Content-Type")
}

func (ctx *Context) Body(out interface{}) *errors.ApiError {

	mime := ctx.getProvidedMime()
	body, err := io.ReadAll(ctx.Req.Body)
	if err != nil {
		return errors.InternalServerError(err)
	}

	err2 := unmarshal(&out, body, mime)
	if err2 != nil {
		return err2
	}

	return nil
}

func (ctx *Context) ShapelessBody() (interface{}, *errors.ApiError) {
	mime := ctx.getProvidedMime()
	body, err := io.ReadAll(ctx.Req.Body)
	if err != nil {
		return nil, errors.InternalServerError(err)
	}

	out, err2 := unmarshalRaw(body, mime)
	if err2 != nil {
		return nil, err2
	}

	return out, nil
}
