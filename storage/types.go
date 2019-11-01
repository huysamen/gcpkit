package storage

type FileUpload struct {
	FilePath string
	FileName string
	MimeType string
	Data     []byte
}

func (f *FileUpload) filename() string {
	return f.FilePath + "/" + f.FileName
}

type FileDownload struct {
	FilePath string
	FileName string
	Data     []byte
}

func (f *FileDownload) filename() string {
	return f.FilePath + "/" + f.FileName
}
