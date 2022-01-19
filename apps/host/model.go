package host

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/imdario/mergo"
)

func NewDefaultHost() *Host {
	return &Host{
		Resource: &Resource{
			CreateAt: time.Now().UnixMilli(),
		},
		Describe: &Describe{},
	}
}

var validate = validator.New()

func (h *Host) Validate() error {
	return validate.Struct(h)
}

func (s *Set) Add(h *Host) {
	s.Items = append(s.Items, h)
	s.Total += 1
}

func (h *Host) PutUpdate(res *Resource, desc *Describe) {
	h.Resource.UpdateAt = time.Now().UnixNano() / 1000000
	if res != nil {
		h.Resource = res
	}

	if desc != nil {
		h.Describe = desc
	}
}

func (h *Host) PatchUpdate(res *Resource, desc *Describe) error {
	h.Resource.UpdateAt = time.Now().UnixNano() / 1000000

	if res != nil {
		err := mergo.MergeWithOverwrite(h.Resource, res)
		if err != nil {
			return err
		}
	}

	if desc != nil {
		err := mergo.MergeWithOverwrite(h.Describe, desc)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewDefaultSet() *Set {
	return &Set{
		Items: []*Host{},
	}
}
