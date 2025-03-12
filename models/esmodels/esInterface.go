package esmodels

type ESIndexInterFace interface {
	Index() string
	Mapping() string
	IndexExists() bool
	CreateIndex() error
	RemoveIndex() error
}
