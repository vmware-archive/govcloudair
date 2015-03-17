package vcatypes

//ServiceGroupIds struct
type ServiceGroupIds struct {
	ServiceGroupID []string `xml:"serviceGroupIds"`
}

//Plan struct
type Plan struct {
	Name                string `xml:"name,attr"`
	ID                  string `xml:"id,attr"`
	Region              string `xml:"region"`
	ServiceName         string `xml:"serviceName"`
	Description         string `xml:"description"`
	PlanVersion         string `xml:"planVersion"`
	PlanAttributes      string `xml:"planAttributes"`
	InstanceSpec        string `xml:"instanceSpec"`
	InstanceDefaultSpec string `xml:"instanceDefaultSpec"`
	BindingSpec         string `xml:"bindingSpec"`
	PlanPolicy          struct {
		CanCreateInstance bool `xml:"canCreateInstance"`
		CanCreateBinding  bool `xml:"canCreateBinding"`
		MaxInstanceCount  bool `xml:"maxInstanceCount"`
	} `xml:"planPolicy"`
}

//PlanList struct
type PlanList struct {
	Plans []Plan `xml:"plans"`
}

//Instance struct
type Instance struct {
	APIURL             string `xml:"apiUrl"`
	DashboardURL       string `xml:"dashboardUrl"`
	Description        string `xml:"description,omitempty"`
	InstanceAttributes string `xml:"instanceAttributes,omitempty"`
	InstanceVersion    string `xml:"instanceVersion,omitempty"`
	PlanID             string `xml:"planId"`
	Region             string `xml:"region"`
	ServiceGroupID     string `xml:"serviceGroupId"`
}

//InstanceAttributes struct
type InstanceAttributes struct {
	OrgName       string `json:"orgName"`
	SessionURI    string `json:"sessionUri"`
	APIVersionURI string `json:"apiVersionUri"`
}

//InstanceList struct
type InstanceList struct {
	Instances []Instance `xml:"instances"`
}

//InstanceSpecParams struct
type InstanceSpecParams struct {
	Name                string `xml:"name"`
	Description         string `xml:"description,omitempty"`
	PlanID              string `xml:"planId"`
	ServiceGroupID      string `xml:"serviceGroupId"`
	InstanceSpec        string `xml:"instanceSpec,omitempty"`
	InstanceDefaultSpec string `xml:"instanceDefaultSpec,omitempty"`
	BindingSpec         string `xml:"bindingSpec,omitempty"`
}

//Users struct
type Users struct {
	User []struct {
		Meta struct {
			Created  string `xml:"created,omitempty"`
			Modified string `xml:"modified,omitempty"`
		} `xml:"meta"`
		Schemas        string `xml:"schemas,omitempty"`
		State          string `xml:"state,omitempty"`
		CompanyID      string `xml:"companyId,omitempty"`
		CustomerNumber string `xml:"customerNumber,omitempty"`
		Email          string `xml:"email,omitempty"`
		FamilyName     string `xml:"familyName,omitempty"`
		GivenName      string `xml:"givenName,omitempty"`
		Roles          struct {
			Role []struct {
				Description string `xml:"description,omitempty"`
				Name        string `xml:"name,omitempty"`
				Rights      struct {
					ID   string `xml:"id,omitempty"`
					Name string `xml:"name,omitempty"`
				} `xml:"rights,omitempty"`
			} `xml:"role"`
		} `xml:"roles"`
		ServiceGroups []struct {
			ServiceGroup struct {
				ServiceGroupID string `xml:"serviceGroupId,omitempty"`
				DisplayName    string `xml:"displayName,omitempty"`
			} `xml:"serviceGroup,omitempty"`
		} `xml:"serviceGroups,omitempty"`
		TosAcceptDate string `xml:"tosAcceptDate,omitempty"`
		TosAccepted   string `xml:"tosAccepted,omitempty"`
		UserName      string `xml:"userName,omitempty"`
	} `xml:"User"`
}

//BillableCosts struct
type BillableCosts struct {
	Cost []struct {
		Type     string `xml:"type,attr"`
		CostItem []struct {
			Properties struct {
				Property []struct {
					Name  string `xml:"name,attr,omitempty"`
					Value string `xml:",chardata"`
				} `xml:"Property"`
			} `xml:"Properties"`
		} `xml:"CostItem"`
	} `xml:"Cost"`
	Currency       string `xml:"Currency"`
	LastUpdateTime string `xml:"LastUpdateTime"`
}

//BilledCosts struct
type BilledCosts struct {
	Month     string `xml:"month,attr,omitempty"`
	Year      string `xml:"year,attr,omitempty"`
	StartTime string `xml:"startTime,attr,omitempty"`
	EndTime   string `xml:"endTime,attr,omitempty"`
	Cost      []struct {
		Type       string `xml:"type,omitempty"`
		Amount     string `xml:"amount,omitempty"`
		Properties []struct {
			Property struct {
				Name string `xml:"name,attr,omitempty"`
			}
		}
	} `xml:"Cost"`
	Currency string `xml:"currency,omitempty"`
}
