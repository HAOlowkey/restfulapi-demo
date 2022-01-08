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
}

func (q *QueryHostRequest) Offset() int {
	return q.PageSize * (q.PageNum - 1)
}

type DescribeHostRequest struct {
	Id int
}

type UpdateHostRequest struct {
	Id int
}

type DeleteHostRequest struct {
	Id int
}
