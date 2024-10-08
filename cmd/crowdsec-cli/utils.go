package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/crowdsecurity/crowdsec/pkg/types"
)

func manageCliDecisionAlerts(ip *string, ipRange *string, scope *string, value *string) error {
	/*if a range is provided, change the scope*/
	if *ipRange != "" {
		_, _, err := net.ParseCIDR(*ipRange)
		if err != nil {
			return fmt.Errorf("%s isn't a valid range", *ipRange)
		}
	}

	if *ip != "" {
		ipRepr := net.ParseIP(*ip)
		if ipRepr == nil {
			return fmt.Errorf("%s isn't a valid ip", *ip)
		}
	}

	// avoid confusion on scope (ip vs Ip and range vs Range)
	switch strings.ToLower(*scope) {
	case "ip":
		*scope = types.Ip
	case "range":
		*scope = types.Range
	case "country":
		*scope = types.Country
	case "as":
		*scope = types.AS
	}

	return nil
}
