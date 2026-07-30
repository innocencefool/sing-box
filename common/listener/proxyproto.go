package listener

import (
	"net"

	"github.com/pires/go-proxyproto"
)

func newProxyProtoListener(listener net.Listener) net.Listener {
	return &proxyproto.Listener{Listener: listener}
}
