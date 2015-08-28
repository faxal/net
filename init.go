package net

import (
	"database/sql/driver"
	"net"
)

type IP struct {
	net.IP
}

func (ip *IP) Scan(v interface{}) error {
	var (
		vb []byte
		ok bool
	)
	if vb, ok = v.([]byte); !ok || v == nil {
		*ip = IP{net.ParseIP("0.0.0.0")}
		return nil
	}

	*ip = IP{net.IP(vb)}

	return nil
}

func (ip *IP) Value() (driver.Value, error) {
	if ip.IP == nil {
		return []byte(net.ParseIP("0.0.0.0")), nil
	}
	return []byte(ip.IP), nil
}
