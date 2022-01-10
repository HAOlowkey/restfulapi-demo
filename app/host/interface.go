package host

import "context"

type Service interface {
	CreateHost(context.Context, *Host) (*Host, error)
	QueryHost(context.Context, *QueryHostRequest) (*Set, error)
	DescribeHost(context.Context, *DescribeHostRequest) (*Host, error)
	UpdateHost(context.Context, *UpdateHostRequest) (*Host, error)
	DeleteHost(context.Context, *DeleteHostRequest) (*Host, error)
}

type QueryHostRequest struct {
	PageSize int
	PageNum  int
	KeyWords string
}

func (q *QueryHostRequest) Offset() int {
	return q.PageSize * (q.PageNum - 1)
}

type DescribeHostRequest struct {
	Id string
}

type UpdateMode int

const (
	PATCH UpdateMode = iota
	PUT
)

func NewPutUpdateHostRequest() *UpdateHostRequest {
	return &UpdateHostRequest{
		UpdateMode: PUT,
		Host:       NewDefaultHost(),
	}
}

func NewPatchUpdateHostRequest() *UpdateHostRequest {
	return &UpdateHostRequest{
		UpdateMode: PATCH,
		Host:       NewDefaultHost(),
	}
}

type UpdateHostRequest struct {
	UpdateMode
	*Host
}

type DeleteHostRequest struct {
	Id string
}
