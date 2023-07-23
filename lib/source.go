package lib

import (
	gopacket "github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func initSource(nif string) (*gopacket.PacketSource, error) {
	live, err := pcap.OpenLive(nif, 1600, true, pcap.BlockForever)
	if err != nil {
		return nil, err
	}
	source := gopacket.NewPacketSource(live, live.LinkType())
	return source, nil
}
