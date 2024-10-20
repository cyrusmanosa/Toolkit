package toolkit

type Tools struct {
	MaxFileSize        int
	AllowedFileTypes   []string
	MaxJSONSize        int
	AllowUnknownFields bool
}

type UploadedFile struct {
	NewFileName      string
	OriginalFileName string
	FileSize         int64
}
