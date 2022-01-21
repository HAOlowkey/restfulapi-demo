package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/HAOlowkey/restfulapi-demo/apps/host"
	"github.com/infraboard/mcube/http/context"
	"github.com/infraboard/mcube/http/request"
	"github.com/infraboard/mcube/http/response"
)

func (h *handler) CreateHost(w http.ResponseWriter, r *http.Request) {
	body, err := request.ReadBody(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	h.log.Debugf("receive body %s", string(body))
	req := host.NewDefaultHost()
	err = json.Unmarshal(body, req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	ins, err := h.host.CreateHost(r.Context(), req)

	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, ins)
}

func (h *handler) QueryHost(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query()
	psStr := qs.Get("page_size")
	pnStr := qs.Get("page_number")

	ps := 20
	pn := 1

	ps, _ = strconv.Atoi(psStr)
	pn, _ = strconv.Atoi(pnStr)

	req := &host.QueryHostRequest{
		PageSize:   int64(ps),
		PageNumber: int64(pn),
		Keywords:   qs.Get("keywords"),
	}

	set, err := h.host.QueryHost(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) DescribeHost(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)

	req := &host.DescribeHostRequest{
		Id: ctx.PS.ByName("id"),
	}

	host, err := h.host.DescribeHost(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, host)
}

func (h *handler) PutUpdateHost(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	body, err := request.ReadBody(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := host.NewPutUpdateHostRequest()
	h.log.Debugf("receive body %s", string(body))
	err = json.Unmarshal(body, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req.Host.Resource.Id = ctx.PS.ByName("id")

	ins, err := h.host.UpdateHost(r.Context(), req)

	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, ins)
}

func (h *handler) PatchUpdateHost(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	body, err := request.ReadBody(r)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req := host.NewPatchUpdateHostRequest()
	h.log.Debugf("receive body %s", string(body))
	err = json.Unmarshal(body, req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	req.Host.Resource.Id = ctx.PS.ByName("id")

	ins, err := h.host.UpdateHost(r.Context(), req)

	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, ins)
}

func (h *handler) DeleteHost(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	ins, err := h.host.DeleteHost(r.Context(), &host.DeleteHostRequest{Id: ctx.PS.ByName("id")})
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, ins)
}
