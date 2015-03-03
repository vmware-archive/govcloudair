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
