package client_test

import (
	"context"
	"testing"

	"github.com/HAOlowkey/restfulapi-demo/apps/host"
	"github.com/HAOlowkey/restfulapi-demo/client"
	"github.com/stretchr/testify/assert"
)

func TestHost(t *testing.T) {
	should := assert.New(t)
	client, err := client.NewClient(client.NewDefaultClientConfig())
	should.NoError(err)
	set, err := client.Host().QueryHost(context.Background(), &host.QueryHostRequest{
		PageSize:   20,
		PageNumber: 1,
	})
	should.NoError(err)

	t.Log(set.Items)

}
