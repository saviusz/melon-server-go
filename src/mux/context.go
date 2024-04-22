package mux

import "net/http"

type Context struct {
	req    *http.Request
	w      http.ResponseWriter
	Params map[string]string
}

func (ctx *Context) createResponse(output interface{}, code int) (err error) {
	req_mime := getRequestedMime(*ctx.req)
	output_data, res_mime, err := marshal(output, req_mime)

	if err != nil {
		return err
	}

	ctx.w.Header().Add("Content-Type", res_mime)
	ctx.w.WriteHeader(code)
	ctx.w.Write(output_data)

	return nil
}
