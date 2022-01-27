package storage

type Resource struct {
	Storage Storage
}

func NewResource() *Resource {
	r := &Resource{
		Storage: NewKvStorage(),
	}
	return r
}
