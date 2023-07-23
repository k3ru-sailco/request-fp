package lib

import "github.com/google/gopacket/layers"

type Fingerprint struct {
	TCP *layers.TCP `json:"tcp_data"`
	TLS *layers.TLS `json:"tls_data"`
}
