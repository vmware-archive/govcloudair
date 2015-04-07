package types

//Media has an attribute of HREF
type Media struct {
	// Attributes
	HREF string `xml:"href,attr,omitempty"` // HREF
}

//MediaInsertOrEjectParams has an element of Media
type MediaInsertOrEjectParams struct {
	Xmlns string `xml:"xmlns,attr"`
	Media *Media `xml:"Media,omitempty"` // Media
}

//ExternalIPAddressActionList has an element of Allocation
type ExternalIpAddressActionList struct {
	Xmlns        string        `xml:"xmlns,attr"`
	Allocation   *Allocation   `xml:"Allocation,omitempty"`
	Deallocation *Deallocation `xml:"Deallocation,omitempty"`
}

//Allocation specifies that amount of IP addresses to reserve
type Allocation struct {
	ExternalNetworkName                   string `xml:"ExternalNetworkName,omitempty"`
	ExternalNetworkRef                    string `xml:"ExternalNetworkRef,omitempty"`
	NumberOfExternalIPAddressesToAllocate string `xml:"NumberOfExternalIpAddressesToAllocate,omitempty"`
}

//Deallocation specifies an IP address to deallocate
type Deallocation struct {
	ExternalNetworkName string `xml:"ExternalNetworkName,omitempty"`
	ExternalNetworkRef  string `xml:"ExternalNetworkRef,omitempty"`
	ExternalIPAddress   string `xml:"ExternalIpAddress,omitempty"`
}
