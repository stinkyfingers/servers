package leet

import (
	"github.com/google/gopacket"
)

type CustomLayer struct {
	SomeByte byte
	Data     []byte
}

var CustomLayerType = gopacket.RegisterLayerType(2069, gopacket.LayerTypeMetadata{"CustomLayerType", gopacket.DecodeFunc(decode)})

func (l *CustomLayer) LayerContents() []byte {
	return []byte{l.SomeByte}
}

func (l *CustomLayer) LayerPayload() []byte {
	return l.Data
}

func (l *CustomLayer) LayerType() gopacket.LayerType {
	return CustomLayerType
}

func decode(data []byte, p gopacket.PacketBuilder) error {
	p.AddLayer(&CustomLayer{data[0], data[1:]})
	return p.NextDecoder(gopacket.LayerTypePayload)
}
