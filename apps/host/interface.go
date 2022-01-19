package host

func (q *QueryHostRequest) Offset() int64 {
	return q.PageSize * (q.PageNumber - 1)
}

func NewPutUpdateHostRequest() *UpdateHostRequest {
	return &UpdateHostRequest{
		UpdateMode: UpdateMode_PUT,
		Host:       NewDefaultHost(),
	}
}

func NewPatchUpdateHostRequest() *UpdateHostRequest {
	return &UpdateHostRequest{
		UpdateMode: UpdateMode_PATCH,
		Host:       NewDefaultHost(),
	}
}
