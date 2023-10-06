package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

func ParseBody(r *http.Request, x interface{}) error {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(body), x)

	if err != nil {
		return err
	}

	return nil

}

func ParseParamBookId(vars map[string]string) (int64, bool) {

	paramId := vars["bookId"]
	bookId, err := strconv.ParseInt(paramId, 0, 0)

	ok := (err == nil)

	return bookId, ok

}
