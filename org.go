/*
 * Copyright 2014 VMware, Inc.  All rights reserved.  Licensed under the Apache v2 License.
 */

package govcloudair

import (
	"fmt"
	"net/url"

	types "github.com/vmware/govcloudair/types/v56"
)

type Org struct {
	Org *types.Org
	c   *Client
}

func NewOrg(c *Client) *Org {
	return &Org{
		Org: new(types.Org),
		c:   c,
	}
}

func (o *Org) FindCatalog(catalog string) (Catalog, error) {

	for _, av := range o.Org.Link {
		if av.Rel == "down" && av.Type == "application/vnd.vmware.vcloud.catalog+xml" && av.Name == catalog {
			u, err := url.ParseRequestURI(av.HREF)

			if err != nil {
				return Catalog{}, fmt.Errorf("error decoding org response: %s", err)
			}

			req := o.c.NewRequest(map[string]string{}, "GET", *u, nil)

			resp, err := checkResp(o.c.Http.Do(req))
			if err != nil {
				return Catalog{}, fmt.Errorf("error retreiving catalog: %s", err)
			}

			cat := NewCatalog(o.c)

			if err = decodeBody(resp, cat.Catalog); err != nil {
				return Catalog{}, fmt.Errorf("error decoding catalog response: %s", err)
			}

			// The request was successful
			return *cat, nil

		}
	}

	return Catalog{}, fmt.Errorf("can't find catalog: %s", catalog)
}

//GetOrg returns an Org from a Org URL
func GetOrg(client *Client, u *url.URL) (org *Org, err error) {
	req := client.NewRequest(map[string]string{}, "GET", *u, nil)
	req.Header.Add("Accept", "application/*+xml;version=5.7")
	//req.Header.Add("x-vcloud-authorization", client.VCDToken)

	resp, err := checkResp(client.Http.Do(req))
	if err != nil {
		return &Org{}, fmt.Errorf("error retreiving org: %s", err)
	}

	org = NewOrg(&Client{})

	if err = decodeBody(resp, org.Org); err != nil {
		return &Org{}, fmt.Errorf("error decoding org response: %s", err)
	}

	return org, nil
}

//GetOrgVdc returns a list of OrgVdcs from an Org
func GetOrgVdc(client *Client, u *url.URL) (links []types.Link, err error) {

	org, err := GetOrg(client, u)
	if err != nil {
		return []types.Link{}, fmt.Errorf("Error getting Org: %s", err)
	}

	for _, arg := range org.Org.Link {
		if arg.Type == "application/vnd.vmware.vcloud.vdc+xml" {
			links = append(links, *arg)
		}
	}
	return links, nil
}
