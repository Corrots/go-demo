package filelisting

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const prefix = "/list/"

type userError string

func (u userError) Error() string {
	return u.Message()
}

func (u userError) Message() string {
	return string(u)
}

func Handler(writer http.ResponseWriter, req *http.Request) error {
	if strings.Index(req.URL.Path, prefix) != 0 {
		return userError("url path must start with " + prefix)
	}
	path := req.URL.Path[len(prefix):]
	file, err := os.OpenFile(path, os.O_EXCL|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	_, err = writer.Write(bytes)
	return err
}
