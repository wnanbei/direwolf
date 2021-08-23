package direwolf

import (
	"bytes"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type MultipartForm struct {
	body     *bytes.Buffer
	m        *multipart.Writer
	boundary string
}

func NewMultipartForm() *MultipartForm {
	body := bytes.Buffer{}
	w := multipart.NewWriter(&body)

	mf := MultipartForm{
		body:     &body,
		m:        w,
		boundary: w.Boundary(),
	}
	return &mf
}

func (mf *MultipartForm) WriteField(key, value string) error {
	return mf.m.WriteField(key, value)
}

func (mf *MultipartForm) WriteFile(key, filePath string) error {
	_, fileName := filepath.Split(filePath)

	fw, err := mf.m.CreateFormFile(key, fileName)
	if err != nil {
		return err
	}

	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(fw, f)
	return err
}

func (mf *MultipartForm) Close() error {
	return mf.m.Close()
}

func (mf *MultipartForm) Reader() *bytes.Reader {
	return bytes.NewReader(mf.body.Bytes())
}

func (mf *MultipartForm) Boundary() string {
	return mf.boundary
}

func (mf *MultipartForm) bindRequest(request *Request) error {
	request.MultipartForm = mf
	return nil
}
