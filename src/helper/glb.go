package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func BodyToJsonReq(r *http.Request, data interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return errors.New(fmt.Sprintf("Body unmarshall error %s", string(body)))
	}

	defer r.Body.Close()

	return nil
}

func StrToInt64(aval string) int64 {
	aval = strings.Trim(strings.TrimSpace(aval), "\n")
	i, err := strconv.ParseInt(aval, 10, 64)
	if err != nil {
		return 0
	}
	return i
}
