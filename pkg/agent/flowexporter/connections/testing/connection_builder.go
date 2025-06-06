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
}

func (b Builder) Get() *flowexporter.Connection {
	tuple := flowexporter.Tuple{SourceAddress: srcAddr, DestinationAddress: dstAddr, Protocol: 6, SourcePort: 65280, DestinationPort: 255}
	return &flowexporter.Connection{
		FlowKey: tuple,
		Zone:    openflow.CtZone,
	}
}
