package direwolf

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
	"os"
	"path/filepath"
)

type MultipartForm struct {
	body *bytes.Buffer
	m    *multipart.Writer
}

func NewMultipartForm() *MultipartForm {
	body := bytes.Buffer{}
	w := multipart.NewWriter(&body)

	mf := MultipartForm{
		body: &body,
		m:    w,
	}
	return &mf
}

func (mf *MultipartForm) WriteField(key, value string) error {
	return mf.m.WriteField(key, value)
}

func (mf *MultipartForm) WriteFile(key, filePath string) error {
	_, fileName := filepath.Split(filePath)

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, key, fileName))

	fw, err := mf.m.CreatePart(h)
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

func (mf *MultipartForm) ContentType() string {
	return mf.m.FormDataContentType()
}

func (mf *MultipartForm) bindRequest(request *Request) error {
	request.MultipartForm = mf
	return nil
}
