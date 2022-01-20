package client

import (
	"github.com/HAOlowkey/restfulapi-demo/apps/host"
	"google.golang.org/grpc"
)

func NewClient(conf *ClientConfig) (*Client, error) {
	conn, err := grpc.Dial(conf.Addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &Client{
		conf: conf,
		conn: conn,
	}, nil
}

type Client struct {
	conf *ClientConfig
	conn *grpc.ClientConn
}

func (c *Client) Host() host.ServiceClient {
	return host.NewServiceClient(c.conn)
}
