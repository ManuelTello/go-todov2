package lib

import (
	"io"
	"strings"

	"todo.proj/src/helpers"
)

func MapBody(b io.ReadCloser) helpers.Body {
	var bodymap helpers.Body = helpers.Body{Context: make(map[string]string)}
	bodyraw, err := io.ReadAll(b)

	if err != nil {
		panic(err)
	}

	// Could be this two
	var bodycleaned string = strings.ReplaceAll(string(bodyraw), "+", " ")
	bodycleaned = strings.ReplaceAll(string(bodyraw), "%20", " ")

	var bodyparsed []string = strings.Split(bodycleaned, "&")

	for i := 0; i < len(bodyparsed); i++ {
		var kv []string = strings.Split(bodyparsed[i], "=")
		bodymap.Context[kv[0]] = kv[1]
	}

	return bodymap
}
