package common

import "errors"

// VcableID is the vcable it of the instance that
// is associated with the ip reservation.
type VcableID string

// Validate checks if the VcableID provided is empty or not
func (v VcableID) Validate() (err error) {
	if v == "" {
		return errors.New("go-oracle-cloud: Empty vcable id")
	}

	return nil
}

// IPPool type describing the
// parent pool of an ip association
type IPPool string

const (
	// PublicIPPool standard ip pool for the oracle provider
	PublicIPPool IPPool = "/oracle/public/ippool"
)

func NewIPPool(name IPPool, prefix IPPrefixType) IPPool {
	return IPPool(prefix) + name
}

type IPPrefixType string

const (
	IPReservationType IPPrefixType = "ipreservation:"
	IPPoolType        IPPrefixType = "ippool:"
)

// IcmpCode is the  ICMP code
// for sec application
type IcmpCode string

const (
	IcmpCodeNetwork  IcmpCode = "network"
	IcmpCodeHost     IcmpCode = "host"
	IcmpCodeProtocol IcmpCode = "protocol"
	IcmpCodePort     IcmpCode = "port"
	IcmpCodeDf       IcmpCode = "df"
	IcmpCodeAdmin    IcmpCode = "admin"
)

// IcmpType is the icmp type for
// sec application
type IcmpType string

const (
	IcmpTypeEcho       IcmpType = "echo"
	IcmpTypeReply      IcmpType = "reply"
	IcmpTypeTTL        IcmpType = "ttl"
	IcmpTypeTraceroute IcmpType = "traceroute"
	IcmpUnreachable    IcmpType = "unreachable"
)

// Protocol is the protocol
// for sec application
type Protocol string

func (p Protocol) Validate() (err error) {
	if p == "" {
		return errors.New(
			"go-oracle-cloud: Empty protocol field",
		)
	}

	return nil
}

const (
	TCP    Protocol = "6"
	ICMP   Protocol = "1"
	IGMP   Protocol = "2"
	IPIP   Protocol = "94"
	RDP    Protocol = "27"
	ESP    Protocol = "50"
	AH     Protocol = "51"
	GRE    Protocol = "47"
	ICMPV6 Protocol = "58"
	OSPF   Protocol = "89"
	PIM    Protocol = "103"
	SCTP   Protocol = "132"
	MPLSIP Protocol = "137"
)

type AdminDistance int

const (
	AdminDistanceZero = 0
	AdminDistanceOne  = 1
	AdminDistanceTwo  = 2
)
