package iptvpack

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type IpTv struct {
	IPTVPath string
	AltPath  string
	Origin   string
	Client   *http.Client
	Writer   io.Writer
}

//Request ...
func (i *IpTv) Request(method, wurl string) (*http.Request, error) {
	req, err := http.NewRequest(method, wurl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Origin", i.Origin)
	req.Header.Set("Referer", i.Origin)

	return req, nil
}

//DoRequest ...
func (i *IpTv) DoRequest(Request *http.Request) (*http.Response, error) {
	return i.Client.Do(Request)
}

//Do ...
func (i *IpTv) Do(method, wurl string) (*http.Response, error) {
	req, err := i.Request(method, wurl)
	if err != nil {
		return nil, err
	}
	return i.DoRequest(req)
}

//Res2Byte ...
func (i *IpTv) Res2Byte(res *http.Response) ([]byte, error) {
	defer res.Body.Close()
	return io.ReadAll(res.Body)
}

//New ...
func New(site string, IPTVPath, AltPath string, Client *http.Client, Writer io.Writer) (*IpTv, error) {
	if Client == nil {
		Client = http.DefaultClient
	}

	if Writer == nil {
		return nil, errors.New("writer cannot be empty")
	}
	//spi := strings.Split(AltPath, "/")
	return &IpTv{
		IPTVPath: IPTVPath,
		AltPath:  AltPath,
		Origin:   site,
		Client:   Client,
		Writer:   Writer,
	}, nil
}

func (i *IpTv) WatchIpTvRaw(loop int, wait time.Duration) error {
	return i._watchIpTv(nil, loop, wait)
}

func (i *IpTv) WatchIpTv(logchan chan string, loop int, wait time.Duration) error {
	return i._watchIpTv(logchan, loop, wait)
}

func (i *IpTv) _watchIpTv(logchan chan string, loop int, wait time.Duration) error {
	var list = []string{}
	for ii := 0; ii < loop; ii++ {
		do, err := i.Do("GET", i.IPTVPath)
		if err != nil {
			return err
		}
		b, err := i.Res2Byte(do)
		if err != nil {
			return err
		}
		bs := string(b)
		if logchan != nil {
			logchan <- "BASE:" + bs
		}

		b = nil

		if do.StatusCode != 200 {
			return errors.New(do.Status)
		}
		for _, v := range strings.Split(bs, "\n") {
			if v == "" || v[0] == '#' || isinlistString(list, v) {
				continue
			}
			if logchan != nil {
				logchan <- "REQUEST:" + i.AltPath + v
			}
			res, err := i.Do("GET", i.AltPath+v)
			if err != nil {
				return err
			}
			defer res.Body.Close()

			if res.StatusCode != 200 {
				var s string
				not202, err := io.ReadAll(res.Body)
				if err != nil {
					s = err.Error()
				} else {
					s = string(not202)
				}

				if logchan != nil {
					logchan <- "MEDIA-ERROR:" + s
				}
				not202 = nil
				res.Body.Close()

				return errors.New(res.Status)
			}
			_, err = io.Copy(i.Writer, res.Body)

			res.Body.Close()

			if err != nil {
				if logchan != nil {
					logchan <- "IO-ERROR:" + err.Error()
				}
				return err
			}
			list = append(list, v)
		}

		if ii != loop-1 {
			if logchan != nil {
				logchan <- "STATUS: WAIT " + fmt.Sprint(wait.Seconds())
			}
			time.Sleep(wait)
		}
	}
	return nil
}

func isinlistString(list []string, key string) bool {
	for _, ğ := range list {
		if ğ == key {
			return true
		}
	}
	return false
}
