package impl

import (
	"context"

	"github.com/HAOlowkey/restfulapi-demo/app/host"
)

func (ins *impl) CreateHost(context.Context, *host.Host) (*host.Host, error) {
	return nil, nil
}

func (ins *impl) QueryHost(context.Context, *host.QueryHostRequest) (*host.Set, error) {
	return nil, nil
}

func (ins *impl) DescribeHost(context.Context, *host.DescribeHostRequest) (*host.Host, error) {
	return nil, nil
}

func (ins *impl) UpdateHost(context.Context, *host.UpdateHostRequest) (*host.Host, error) {
	return nil, nil
}

func (ins *impl) DeleteHost(context.Context, *host.DeleteHostRequest) (*host.Host, error) {
	return nil, nil
}
