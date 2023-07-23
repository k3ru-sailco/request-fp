package lib

import (
	"fmt"
	"sync"
)

var (
	FINGERPRINT_MUTEXT = &sync.Mutex{}
	FINGERPRINTS       = map[string]*Fingerprint{}
)

/*
start listening and appending packet data to the map of fingerprints
for inbound connections to the machine
*/
func StartListening(networkInterface string) error {
	packetSource, err := initSource(networkInterface)
	if err != nil {
		return err
	}
	for packet := range packetSource.Packets() {
		go readPacket(packet)
	}
	return nil
}

/*
used for an incoming request in a handler,
pass the Ip address of the request to get tcp data, and fingerprint (created by hashing parts of data)
*/
func TCP(ipAddress string) (*Fingerprint, error) {
	defer FINGERPRINT_MUTEXT.Unlock()
	FINGERPRINT_MUTEXT.Lock()
	if fingerprint, ok := FINGERPRINTS[ipAddress]; ok {
		delete(FINGERPRINTS, ipAddress)
		return fingerprint, nil
	}
	return nil, fmt.Errorf("no fingerprint for %s", ipAddress)
}
