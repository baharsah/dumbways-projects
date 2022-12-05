package webfs

import (
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

func WebFSHandler(f multipart.File, fh *multipart.FileHeader, e error) (string, error) {
	dir, _ := os.Getwd()
	filename := fh.Filename
	fileLocation := filepath.Join(dir, "files", filename)
	targetFile, _ := os.OpenFile(fileLocation, os.O_WRONLY|os.O_CREATE, 0666)
	defer targetFile.Close()
	if _, err := io.Copy(targetFile, f); err != nil {
		log.Println("error ketika upload file")
	}

	return filename, nil

}
