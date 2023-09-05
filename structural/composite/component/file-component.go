package component

type FileComponent interface {
	GetSize() int
	GetName() string
	SetSize(size int)
	SetName(name string)
	AddFileComponent(fileComponent FileComponent)
	PrintSubComponents()
	IsLeafNode() bool
}
