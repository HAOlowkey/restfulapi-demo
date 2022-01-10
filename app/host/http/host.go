package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/HAOlowkey/restfulapi-demo/app/host"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
	"github.com/julienschmidt/httprouter"
)

func (h *handler) CreateHost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body, err := request.ReadBody(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	h.log.Debugf("receive body %s", string(body))
	req := host.NewDefaultHost()
	json.Unmarshal(body, req)
	ins, err := h.host.CreateHost(r.Context(), req)

	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, ins)
}

func (h *handler) QueryHost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	qs := r.URL.Query()
	psStr := qs.Get("page_size")
	pnStr := qs.Get("page_number")

	ps := 20
	pn := 1

	ps, _ = strconv.Atoi(psStr)
	pn, _ = strconv.Atoi(pnStr)

	req := &host.QueryHostRequest{
		PageSize: ps,
		PageNum:  pn,
		KeyWords: qs.Get("keywords"),
	}

	set, err := h.host.QueryHost(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) DescribeHost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	req := &host.DescribeHostRequest{
		Id: p.ByName("id"),
	}

	host, err := h.host.DescribeHost(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, host)
}
