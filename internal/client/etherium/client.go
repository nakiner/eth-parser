package etherium

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type Client struct {
	cl http.Client
}

func NewClient() *Client {
	return &Client{
		cl: http.Client{
			Timeout: time.Second * 30,
		},
	}
}

func toHex(val int64) string {
	return fmt.Sprintf("0x%x", val)
}

func fromHex(val string) (int64, error) {
	if len(val) < 3 {
		return 0, nil
	}

	if val[0] != '0' && val[1] != 'x' {
		return 0, errors.New("wrong hex value")
	}

	num := val[2:]

	intVal, err := strconv.ParseInt(num, 16, 64)
	if err != nil {
		return 0, errors.Wrap(err, "could not ParseInt")
	}

	return intVal, nil
}
