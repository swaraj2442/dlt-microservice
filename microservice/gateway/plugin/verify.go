package plugin

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
)

func NewCustomModifier(
	cfg map[string]interface{},
) func(interface{}) (interface{}, error) {
	return func(input interface{}) (interface{}, error) {
		fmt.Println("Executing new custom modifier...")
		return input, nil
	}
}

func RequestModifier(
	cfg map[string]interface{},
) func(interface{}) (interface{}, error) {
	return func(input interface{}) (interface{}, error) {
		req, ok := input.(RequestWrapper)
		if !ok {
			return nil, errors.New("unknown request type")
		}

		fmt.Println("Processing Request...")

		req.Headers()["X-Custom-Header"] = []string{"CustomValue"}
		fmt.Println("Custom header added.")

		return input, nil
	}
}
func ResponseModifier(
	cfg map[string]interface{},
) func(interface{}) (interface{}, error) {
	return func(input interface{}) (interface{}, error) {
		resp, ok := input.(ResponseWrapper)
		if !ok {
			return nil, errors.New("unknown response type")
		}

		body, err := ioutil.ReadAll(resp.Io())
		if err != nil {
			return nil, errors.New("failed to read response body")
		}

		var data map[string]interface{}
		err = json.Unmarshal(body, &data)
		if err != nil {
			return nil, errors.New("failed to unmarshal response body")
		}

		requestID, ok := resp.Headers()["X-Request-Id"]

		fmt.Printf("The Request ID: %s and Message: %v\n", requestID[0], data)

		return data, nil
	}
}

type RequestWrapper interface {
	Params() map[string]string
	Headers() map[string][]string
	Body() io.ReadCloser
	Method() string
	URL() *url.URL
	Query() url.Values
	Path() string
}

type ResponseWrapper interface {
	Data() map[string]interface{}
	Io() io.Reader
	IsComplete() bool
	StatusCode() int
	Headers() map[string][]string
}
