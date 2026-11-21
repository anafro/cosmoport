package server

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/anafro/cosmoport/internal/server/errors"
)

type Backend struct {
	Host string
	Port uint16
}

func (b *Backend) ToURL() (*url.URL, error) {
	return url.Parse(fmt.Sprintf(
		"http://%s:%d",
		b.Host,
		b.Port,
	))
}

func ParseBackend(sbackend string) (*Backend, error) {
	parts := strings.Split(sbackend, ":")

	if len(parts) > 2 {
		return nil, errors.NewErrMalformedBackend(
			sbackend,
			fmt.Sprintf("should be port, or host:port"),
		)
	}

	var host string
	var sport string

	if len(parts) == 1 {
		host = "localhost"
		sport = parts[0]
	} else {
		host = parts[0]
		sport = parts[1]
	}

	port, err := strconv.ParseUint(sport, 10, 16)

	if err != nil {
		return nil, errors.NewErrMalformedBackend(
			sbackend,
			fmt.Sprintf("port must be an integer"),
		)
	}

	return &Backend{
		Host: host,
		Port: uint16(port),
	}, nil
}
