package entity

type File struct {
	Lines []string
}

func NewFile(data []string) *File {
	return &File{
		Lines: data,
	}
}
