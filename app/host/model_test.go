package host_test

import (
	"testing"

	"github.com/HAOlowkey/restfulapi-demo/app/host"
	"github.com/stretchr/testify/assert"
)

func TestHostUpdate(t *testing.T) {
	should := assert.New(t)

	h := host.NewDefaultHost()
	patch := host.NewDefaultHost()
	patch.Name = "patch01"

	err := h.PatchUpdate(patch.Resource, patch.Describe)
	if should.NoError(err) {
		should.Equal(patch.Name, h.Name)
	}
}
