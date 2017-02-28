// Copyright 2017 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package api

import (
	"errors"

	"github.com/hoenirvili/go-oracle-cloud/common"
	"github.com/hoenirvili/go-oracle-cloud/response"
)

type SecApplicationParams struct {

	// Description is a description of the security application.
	Description string `json:"description,omitempty"`

	// Dport is the TCP or UDP destination port number.
	// You can also specify a port range, such as 5900-5999 for TCP.
	// If you specify tcp or udp as the protocol, then the dport
	// parameter is required; otherwise, it is optional.
	// This parameter isn't relevant to the icmp protocol.
	// Note: This request fails if the range-end is lower than the range-start.
	// For example, if you specify the port range as 5000-4000.
	Dport string `json:"dport,omitempty"`

	// Icmpcode is the ICMP code.
	// This parameter is relevant only if you specify
	// icmp as the protocol. You can specify one of the following values:
	//
	// common.IcmpCodeNetwork
	// common.IcmpCodeHost
	// common.IcmpCodeProtocol
	// common.IcmpPort
	// common.IcmpCodeDf
	// common.IcmpCodeAdmin
	//
	// If you specify icmp as the protocol and don't
	// specify icmptype or icmpcode, then all ICMP packets are matched.
	Icmpcode common.IcmpCode `json:"icmpcode,omitempty"`

	// Icmptype
	// The ICMP type. This parameter is relevant only if you specify icmp
	// as the protocol. You can specify one of the following values:
	//
	// common.IcmpTypeEcho
	// common.IcmpTypeReply
	// common.IcmpTypeTTL
	// common.IcmpTraceroute
	// common.IcmpUnreachable
	// If you specify icmp as the protocol and
	// don't specify icmptype or icmpcode, then all ICMP packets are matched.
	Icmptype common.IcmpType `json:"icmptype,omitempty"`

	// Name is the name of the secure application
	Name string `json:"name"`

	// Protocol is the protocol to use.
	// The value that you specify can be either a text representation of
	// a protocol or any unsigned 8-bit assigned protocol number
	// in the range 0-254. See Assigned Internet Protocol Numbers.
	// For example, you can specify either tcp or the number 6.
	// The following text representations are allowed:
	// tcp, udp, icmp, igmp, ipip, rdp, esp, ah, gre, icmpv6, ospf, pim, sctp, mplsip, all.
	// To specify all protocols, set this to all.
	Protocol common.Protocol
}

func (s SecApplicationParams) validate() (err error) {
	if s.Name == "" {
		return errors.New(
			"go-oracle-cloud: Empty secure application name",
		)
	}

	if err = s.Protocol.Validate(); err != nil {
		return err
	}

	return nil
}

// CreateSecApplication creates a security application.
// After creating security applications
func (c Client) CreateSecApplication(p SecApplicationParams) (resp response.SecList, err error) {
	if !c.isAuth() {
		return resp, errNotAuth
	}

	if err = p.validate(); err != nil {
		return resp, err
	}

	url := c.endpoints["secapplication"] + "/"

	if err = request(paramsRequest{
		client: &c.http,
		cookie: c.cookie,
		url:    url,
		body:   &p,
		verb:   "POST",
		resp:   &resp,
	}); err != nil {
		return resp, err
	}

	return resp, nil
}
