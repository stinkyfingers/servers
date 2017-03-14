package main

import (
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/stinkyfingers/servers/leet"
)

func send(msg string) error {
	rawbytes := []byte(msg)
	packet := gopacket.NewPacket(rawbytes, leet.CustomLayerType, gopacket.Default)

	customLayer := packet.Layer(leet.CustomLayerType)

	buffer := gopacket.NewSerializeBuffer()
	gopacket.SerializeLayers(buffer, gopacket.SerializeOptions{},
		&layers.Ethernet{},
		&layers.IPv4{},
		&layers.TCP{},
		customLayer, //TODO...
	)

}
