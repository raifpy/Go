package upformfile

// @raifpy | Mon Jan 18 00:32:11 +03 2021

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
)

//FormDataRequest ...
func FormDataRequest(method, url string, uploadFile io.Reader, fieldName, fileName string, data, headers map[string]string) (*http.Request, error) {
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	d, err := writer.CreateFormFile(fieldName, fileName)
	if err != nil {
		return nil, err
	}
	io.Copy(d, uploadFile)
	for key, value := range data {
		err = writer.WriteField(key, value)
		if err != nil {
			return nil, err
		}
	}
	writer.Close()
	request, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", writer.FormDataContentType())
	for key, value := range headers {
		request.Header.Add(key, value)
	}
	buf = nil // free memory
	d = nil

	return request, nil

}
