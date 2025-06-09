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
}

func NewBuilder() Builder {
	return Builder{
		sourcePort:         60001,
		destinationPort:    200,
		destinationAddress: dstAddr,
	}
}

func (b Builder) Get() *flowexporter.Connection {
	//tuple := flowexporter.Tuple{SourceAddress: srcAddr, DestinationAddress: svcAddr, Protocol: 6, SourcePort: 60001, DestinationPort: 200}
	//antreaServiceFlow := &flowexporter.Connection{
	//	FlowKey: tuple,
	//	Zone:    openflow.CtZone,
	//}
	return &flowexporter.Connection{
		FlowKey: flowexporter.Tuple{
			SourceAddress:      srcAddr,
			DestinationAddress: b.destinationAddress,
			Protocol:           6,
			SourcePort:         b.sourcePort,
			DestinationPort:    b.destinationPort,
		},
		Zone: openflow.CtZone,
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
