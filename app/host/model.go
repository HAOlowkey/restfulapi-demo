package host

type Host struct {
	*Resource
	*Describe
}

type Resource struct {
}

type Describe struct {
}

type Set struct {
	Total int
	Items []*Host
}
