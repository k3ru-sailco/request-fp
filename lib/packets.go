package lib

import (
	gopacket "github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func readPacket(p gopacket.Packet) {
	tcpLayer := p.Layer(layers.LayerTypeTCP)
	if tcpLayer == nil {
		return
	}
	tcp, _ := tcpLayer.(*layers.TCP)
	ipLayer := p.Layer(layers.LayerTypeIPv4)
	if ipLayer == nil {
		return
	}
	tlsLayer := p.Layer(layers.LayerTypeTLS)
	var tls *layers.TLS
	if tlsLayer != nil {
		tls, _ = tlsLayer.(*layers.TLS)
	}
	ip, _ := ipLayer.(*layers.IPv4)
	savePacket(ip.SrcIP.String(), &Fingerprint{
		TCP: tcp,
		TLS: tls,
	})
}

func savePacket(ip string, f *Fingerprint) {
	defer FINGERPRINT_MUTEXT.Unlock()
	FINGERPRINT_MUTEXT.Lock()
	FINGERPRINTS[ip] = f
}
