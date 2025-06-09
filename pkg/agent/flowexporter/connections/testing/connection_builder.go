package testing

import (
	"net/netip"

	flowexporter "antrea.io/antrea/pkg/agent/flowexporter"
	"antrea.io/antrea/pkg/agent/openflow"
)

var (
	srcAddr = netip.MustParseAddr("1.2.3.4")
	dstAddr = netip.MustParseAddr("4.3.2.1")
)

type Builder struct {
	sourcePort         uint16
	destinationPort    uint16
	destinationAddress netip.Addr
	zone               uint16
	protocol           uint8
}

func NewBuilder() Builder {
	return Builder{
		sourcePort:         60001,
		destinationPort:    200,
		destinationAddress: dstAddr,
		zone:               openflow.CtZone,
		protocol:           6,
	}
}

func (b Builder) Get() *flowexporter.Connection {
	return &flowexporter.Connection{
		FlowKey: flowexporter.Tuple{
			SourceAddress:      srcAddr,
			DestinationAddress: b.destinationAddress,
			Protocol:           b.protocol,
			SourcePort:         b.sourcePort,
			DestinationPort:    b.destinationPort,
		},
		Zone: b.zone,
	}
}

func (b Builder) SetSourcePort(sourcePort uint16) Builder {
	b.sourcePort = sourcePort
	return b
}

func (b Builder) SetDestinationPort(destinationPort uint16) Builder {
	b.destinationPort = destinationPort
	return b
}

func (b Builder) SetDestinationAddress(destinationAddress netip.Addr) Builder {
	b.destinationAddress = destinationAddress
	return b
}

func (b Builder) SetZone(zone uint16) Builder {
	b.zone = zone
	return b
}

func (b Builder) SetProtocol(protocol uint8) Builder {
	b.protocol = protocol
	return b
}
