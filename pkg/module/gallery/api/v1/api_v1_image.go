package v1

import (
	"io"
	"net/http"
	"os"
)

func (api *apiV1) UploadImageHandlerV1(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)

	file, header, err := r.FormFile("file")
	if err != nil {

	}
	defer file.Close()

	filename := r.FormValue("filename")
	downloadFile, err := os.OpenFile(header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {

	}
	defer downloadFile.Close()

	_, _ = io.WriteString(w, "File '"+filename+"' uploaded successfully")
	_, _ = io.Copy(downloadFile, file)
}
