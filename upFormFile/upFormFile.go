package upformfile

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
)

//Source: https://gist.github.com/andrewmilson/19185aab2347f6ad29f5

//Upload Method URL file(io.Reader) filed filename
// Python: requests.<method>(url,files={filed:open(filename,"rb")})
func Upload(method, url string, file io.Reader, fileid, filename string) (*http.Response, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(fileid, filename)

	if err != nil {
		return nil, err
	}
	io.Copy(part, file)
	writer.Close()

	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", writer.FormDataContentType())
	client := http.Client{}
	return client.Do(request)

}

//MultiUpload abc
func MultiUpload(method, url string, files map[string]io.Reader) (*http.Response, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for key, value := range files {
		part, err := writer.CreateFormFile(key, key)

		if err != nil {
			return nil, err
		}
		io.Copy(part, value)
	}
	writer.Close()

	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", writer.FormDataContentType())
	client := http.Client{}
	return client.Do(request)

}

