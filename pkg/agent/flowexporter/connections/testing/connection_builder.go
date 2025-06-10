package testing

import (
	"net/netip"
	"time"

	flowexporter "antrea.io/antrea/pkg/agent/flowexporter"
	"antrea.io/antrea/pkg/agent/openflow"
)

type Builder struct {
	sourcePort                 uint16
	destinationPort            uint16
	destinationAddress         netip.Addr
	sourceAddress              netip.Addr
	zone                       uint16
	protocol                   uint8
	timeout                    uint32
	startTime                  time.Time
	stopTime                   time.Time
	isPresent                  bool
	statusFlag                 uint32
	mark                       uint32
	originalDestinationAddress netip.Addr
	originalDestinationPort    uint16
	originalPackets            uint64
	originalBytes              uint64
	reversePackets             uint64
	reverseBytes               uint64
}

func NewBuilder() Builder {
	return Builder{
		sourcePort:         60001,
		destinationPort:    200,
		destinationAddress: netip.MustParseAddr("4.3.2.1"),
		zone:               openflow.CtZone,
		protocol:           6,
		sourceAddress:      netip.MustParseAddr("1.2.3.4"),
	}
}

func (b Builder) Get() *flowexporter.Connection {
	return &flowexporter.Connection{
		FlowKey: flowexporter.Tuple{
			SourceAddress:      b.sourceAddress,
			DestinationAddress: b.destinationAddress,
			Protocol:           b.protocol,
			SourcePort:         b.sourcePort,
			DestinationPort:    b.destinationPort,
		},
		Timeout:                    b.timeout,
		StartTime:                  b.startTime,
		StopTime:                   b.stopTime,
		Zone:                       b.zone,
		IsPresent:                  b.isPresent,
		StatusFlag:                 b.statusFlag,
		Mark:                       b.mark,
		OriginalDestinationAddress: b.originalDestinationAddress,
		OriginalDestinationPort:    b.originalDestinationPort,
		OriginalPackets:            b.originalPackets,
		OriginalBytes:              b.originalBytes,
		ReversePackets:             b.reversePackets,
		ReverseBytes:               b.reverseBytes,
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

func (b Builder) SetOriginalDestinationPort(originalDestinationPort uint16) Builder {
	b.originalDestinationPort = originalDestinationPort
	return b
}

func (b Builder) SetDestinationAddress(destinationAddress netip.Addr) Builder {
	b.destinationAddress = destinationAddress
	return b
}

func (b Builder) SetOriginalDestinationAddress(originalDestinationAddress netip.Addr) Builder {
	b.originalDestinationAddress = originalDestinationAddress
	return b
}

func (b Builder) SetSourceAddress(sourceAddress netip.Addr) Builder {
	b.sourceAddress = sourceAddress
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

func (b Builder) SetTimeout(timeout uint32) Builder {
	b.timeout = timeout
	return b
}

func (b Builder) SetStartTime(startTime time.Time) Builder {
	b.startTime = startTime
	return b
}

func (b Builder) SetStopTime(stopTime time.Time) Builder {
	b.stopTime = stopTime
	return b
}

func (b Builder) SetPresent() Builder {
	b.isPresent = true
	return b
}

func (b Builder) SetStatusFlag(statusFlag uint32) Builder {
	b.statusFlag = statusFlag
	return b
}

func (b Builder) SetMark(mark uint32) Builder {
	b.mark = mark
	return b
}

func (b Builder) SetOriginalPackets(packets uint64) Builder {
	b.originalPackets = packets

	return b
}

func (b Builder) SetOriginalBytes(bytes uint64) Builder {
	b.originalBytes = bytes

	return b
}

func (b Builder) SetReversePackets(packets uint64) Builder {
	b.reversePackets = packets
	return b
}

func (b Builder) SetReverseBytes(bytes uint64) Builder {
	b.reverseBytes = bytes
	return b
}
