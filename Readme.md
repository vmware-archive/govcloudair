## govcloudair [![Build Status](https://travis-ci.org/vmware/govcloudair.svg?branch=master)](https://travis-ci.org/frapposelli/govcloudair) [![Coverage Status](https://img.shields.io/coveralls/vmware/govcloudair.svg)](https://coveralls.io/r/vmware/govcloudair) [![GoDoc](https://godoc.org/github.com/vmware/govcloudair?status.svg)](http://godoc.org/github.com/vmware/govcloudair)

This package provides the `govcloudair` package which offers an interface to the vCloud Air 5.6 API.

It serves as a foundation for a project currently in development, there are plans to make it a general purpose API in the future.

The API is currently under heavy development, its coverage is extremely limited at the moment.


### vCloud Air Subscription
The package supports the subscription model of authenticating and retrieving compute resources.

Examples to come.

### vCloud Air On Demand
There is a slight difference to how the authentication occurs between the subscription and on-demand models.  However, once you authenticate and retrieve the appropriate ```Vdc``` object you should be able to leverage all of the functionality equivalently between both methods.

The On-Demand services are configured by supplying your vCloud Air credentials by environment variables (VCLOUDAIR_USERNAME and VCLOUDAIR_PASSWORD) or by parameters.  

See the following Go snippets.  A working example of a CLI application can be found at the [Goair](https://github.com/emccode/goair) project.

The first step is to authenticate against the vCloud Air API.  From there the client object can be used to perform other methods and functions.

#### Authenticate

    client, err = govcloudair.NewClient()
    if err != nil {
      return client, fmt.Errorf("error with NewClient: %s", err)
    }

    err = client.AuthenticateOD("user", "pass")
    if err != nil {
      return client, fmt.Errorf("error Authenticating: %s", err)
    }

#### Get Available Instances
The next snippet will retrieve the available instances or regions that can be used.

    instanceList, err := client.GetInstances()
    if err != nil {
      log.Fatalf("error Getting instances: %s", err)
    }

#### Authenticate to the Instance
Once you choose an instance you can then authenticate to the backend vCloud Director instance with appropriate credentials.  You should use the appropriate ```instance.InstanceAttributes``` field which includes the backend vCD instance API field.

    ia := instanceList[0].InstanceAttributes
    instanceAttributes := vcatypes.InstanceAttributes{}
    json.Unmarshal([]byte(ia), &instanceAttributes)

    err = client.GetBackendAuthOD(instanceAttributes)
    if err != nil {
      return fmt.Errorf("error Authenticating: %s", err)
    }

#### Retrieve OrgVdc list
This should allow you to then see the compute resoures that are available in the region.  From there you can perform any number of operations that are relevant to the general ```govcloudair``` functionality.

    links, err := govcloudair.GetOrgVdc(client, &client.VCDORGHREF)
    if err != nil {
      log.Fatalf("Err geting orgvdc: %s", err)
    }
