package dao

type Model interface {
	GetID() uint64
	GetTable() string
	Validate() error
}
