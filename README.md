# Telstra Programmable Network 

Telstra Programmable Network is a self-provisioning platform that allows its users to create on-demand connectivity services between multiple end-points and add various network functions to those services. Programmable Network enables to connectivity to a global ecosystem of networking services as well as public and private cloud services. Once you are connected to the platform on one or more POPs (points of presence), you can start creating those services based on the use case that you want to accomplish. The Programmable Network API is available to all customers who have registered to use the Programmable Network. To register, please contact your account representative.

## Overview

- API version: 2.1.3
- Package version: 1.0.0

## Installation
Put the package under your project folder and add the following in import:
```
    "./TelstraTPN"
```

## Documentation for API Endpoints

All URIs are relative to *https://penapi.pacnetconnect.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuthenticationApi* | [**AuthGeneratetokenPost**](docs/AuthenticationApi.md#authgeneratetokenpost) | **Post** /1.0.0/auth/generatetoken | Create an authentication token
*AuthenticationApi* | [**AuthValidatetokenGet**](docs/AuthenticationApi.md#authvalidatetokenget) | **Get** /1.0.0/auth/validatetoken | Validate authentication token
*ContractsApi* | [**InventoryLinksContractByLinkidAndContractidGet**](docs/ContractsApi.md#inventorylinkscontractbylinkidandcontractidget) | **Get** /1.0.0/inventory/links/{linkid}/contract/{contractid} | Get active Contract by ContractID
*ContractsApi* | [**InventoryLinksContractByLinkidAndContractidPut**](docs/ContractsApi.md#inventorylinkscontractbylinkidandcontractidput) | **Put** /1.0.0/inventory/links/{linkid}/contract/{contractid} | Update active Contract by ContractID
*ContractsApi* | [**InventoryLinksContractByLinkidPost**](docs/ContractsApi.md#inventorylinkscontractbylinkidpost) | **Post** /1.0.0/inventory/links/{linkid}/contract | Create new Contract on specified link
*CustomersApi* | [**AccountByCustomeruuidGet**](docs/CustomersApi.md#accountbycustomeruuidget) | **Get** /1.0.0/account/{customeruuid} | Get account information details
*CustomersApi* | [**AccountUserByCustomeruuidGet**](docs/CustomersApi.md#accountuserbycustomeruuidget) | **Get** /1.0.0/account/{customeruuid}/user | List users
*DatacentresApi* | [**InventoryDatacentersGet**](docs/DatacentresApi.md#inventorydatacentersget) | **Get** /1.0.0/inventory/datacenters | Get list of all the data centers
*EndpointsApi* | [**Eis100EndpointsAssignTopologyTagByEndpointuuidPost**](docs/EndpointsApi.md#eis100endpointsassigntopologytagbyendpointuuidpost) | **Post** /eis/1.0.0/endpoints/{endpointuuid}/assign_topology_tag | Assign a Topology Tag to an Endpoint
*EndpointsApi* | [**InventoryEndpointByEndpointuuidGet**](docs/EndpointsApi.md#inventoryendpointbyendpointuuidget) | **Get** /1.0.0/inventory/endpoint/{endpointuuid} | Get information about the specified endpoint
*EndpointsApi* | [**InventoryEndpointsCustomeruuidByCustomeruuidGet**](docs/EndpointsApi.md#inventoryendpointscustomeruuidbycustomeruuidget) | **Get** /1.0.0/inventory/endpoints/customeruuid/{customeruuid} | Get list of endpoints for a customer
*EndpointsApi* | [**InventoryRegularendpointPost**](docs/EndpointsApi.md#inventoryregularendpointpost) | **Post** /1.0.0/inventory/regularendpoint | Create Physical (Port) Endpoint
*EndpointsApi* | [**InventoryVnfendpointPost**](docs/EndpointsApi.md#inventoryvnfendpointpost) | **Post** /1.0.0/inventory/vnfendpoint | Create VNF Endpoint
*LinksApi* | [**InventoryLinkPost**](docs/LinksApi.md#inventorylinkpost) | **Post** /1.0.0/inventory/link | Create Link and initial Contract
*LinksApi* | [**InventoryLinksByLinkidGet**](docs/LinksApi.md#inventorylinksbylinkidget) | **Get** /1.0.0/inventory/links/{linkid} | Get details of specified link
*LinksApi* | [**InventoryLinksCustomerByCustomeruuidGet**](docs/LinksApi.md#inventorylinkscustomerbycustomeruuidget) | **Get** /1.0.0/inventory/links/customer/{customeruuid} | Get active Links
*LinksApi* | [**InventoryLinksHistoryByLinkidGet**](docs/LinksApi.md#inventorylinkshistorybylinkidget) | **Get** /1.0.0/inventory/links/history/{linkid} | Get Link history
*TopologiesApi* | [**Ttms100TopologyTagByTopotaguuidDelete**](docs/TopologiesApi.md#ttms100topologytagbytopotaguuiddelete) | **Delete** /ttms/1.0.0/topology_tag/{topotaguuid}/ | Delete a topology tag
*TopologiesApi* | [**Ttms100TopologyTagByTopotaguuidGet**](docs/TopologiesApi.md#ttms100topologytagbytopotaguuidget) | **Get** /ttms/1.0.0/topology_tag/{topotaguuid}/ | Get information about the specified topology tag
*TopologiesApi* | [**Ttms100TopologyTagByTopotaguuidPut**](docs/TopologiesApi.md#ttms100topologytagbytopotaguuidput) | **Put** /ttms/1.0.0/topology_tag/{topotaguuid}/ | Update a topology tag&#39;s name and/or description
*TopologiesApi* | [**Ttms100TopologyTagGet**](docs/TopologiesApi.md#ttms100topologytagget) | **Get** /ttms/1.0.0/topology_tag | List all topology tags
*TopologiesApi* | [**Ttms100TopologyTagObjectsByTopotaguuidGet**](docs/TopologiesApi.md#ttms100topologytagobjectsbytopotaguuidget) | **Get** /ttms/1.0.0/topology_tag/{topotaguuid}/objects/ | List objects for Topology
*TopologiesApi* | [**Ttms100TopologyTagPost**](docs/TopologiesApi.md#ttms100topologytagpost) | **Post** /ttms/1.0.0/topology_tag | Create a named topology tag
*VnfsApi* | [**MarketplaceImageGet**](docs/VnfsApi.md#marketplaceimageget) | **Get** /1.0.0/marketplace/image | List images in the Marketplace
*VportsApi* | [**InventoryRegularvportPost**](docs/VportsApi.md#inventoryregularvportpost) | **Post** /1.0.0/inventory/regularvport | Create VPort for physical endpoint
*VportsApi* | [**InventoryVnfVportPost**](docs/VportsApi.md#inventoryvnfvportpost) | **Post** /1.0.0/inventory/vnf/vport | Create VNF VPort
*VportsApi* | [**InventoryVportByVportuuidGet**](docs/VportsApi.md#inventoryvportbyvportuuidget) | **Get** /1.0.0/inventory/vport/{vportuuid} | Get information about the specified VPort


## Documentation For Models

 - [AccountResponse](docs/AccountResponse.md)
 - [AuthGeneratetokenResponse](docs/AuthGeneratetokenResponse.md)
 - [AuthValidatetokenResponse](docs/AuthValidatetokenResponse.md)
 - [Billing](docs/Billing.md)
 - [Classification](docs/Classification.md)
 - [Contract](docs/Contract.md)
 - [Datacenter](docs/Datacenter.md)
 - [Eis100EndpointsAssignTopologyTagRequest](docs/Eis100EndpointsAssignTopologyTagRequest.md)
 - [Endpoint](docs/Endpoint.md)
 - [EndpointPort](docs/EndpointPort.md)
 - [Endpointlist](docs/Endpointlist.md)
 - [Error74](docs/Error74.md)
 - [Flavor](docs/Flavor.md)
 - [GlanceImage](docs/GlanceImage.md)
 - [InventoryDatacenters401Error](docs/InventoryDatacenters401Error.md)
 - [InventoryDatacentersResponse](docs/InventoryDatacentersResponse.md)
 - [InventoryEndpointResponse](docs/InventoryEndpointResponse.md)
 - [InventoryEndpointsCustomeruuidResponse](docs/InventoryEndpointsCustomeruuidResponse.md)
 - [InventoryLinkRequest](docs/InventoryLinkRequest.md)
 - [InventoryLinkResponse](docs/InventoryLinkResponse.md)
 - [InventoryLinksContractRequest](docs/InventoryLinksContractRequest.md)
 - [InventoryLinksContractRequest37](docs/InventoryLinksContractRequest37.md)
 - [InventoryLinksContractResponse](docs/InventoryLinksContractResponse.md)
 - [InventoryLinksContractResponse33](docs/InventoryLinksContractResponse33.md)
 - [InventoryLinksContractResponse38](docs/InventoryLinksContractResponse38.md)
 - [InventoryLinksHistoryResponse](docs/InventoryLinksHistoryResponse.md)
 - [InventoryLinksResponse](docs/InventoryLinksResponse.md)
 - [InventoryRegularendpointRequest](docs/InventoryRegularendpointRequest.md)
 - [InventoryRegularendpointResponse](docs/InventoryRegularendpointResponse.md)
 - [InventoryRegularvportRequest](docs/InventoryRegularvportRequest.md)
 - [InventoryRegularvportResponse](docs/InventoryRegularvportResponse.md)
 - [InventoryVnfVportRequest](docs/InventoryVnfVportRequest.md)
 - [InventoryVnfVportResponse](docs/InventoryVnfVportResponse.md)
 - [InventoryVnfendpointRequest](docs/InventoryVnfendpointRequest.md)
 - [InventoryVnfendpointResponse](docs/InventoryVnfendpointResponse.md)
 - [Link](docs/Link.md)
 - [Link66](docs/Link66.md)
 - [MarketplaceImageResponse](docs/MarketplaceImageResponse.md)
 - [Meta](docs/Meta.md)
 - [ModelError](docs/ModelError.md)
 - [Object52](docs/Object52.md)
 - [Params](docs/Params.md)
 - [Params31](docs/Params31.md)
 - [Params34](docs/Params34.md)
 - [Params39](docs/Params39.md)
 - [Product](docs/Product.md)
 - [Role](docs/Role.md)
 - [SuccessFragment](docs/SuccessFragment.md)
 - [Topology](docs/Topology.md)
 - [Ttms100TopologyTagObjectsResponse](docs/Ttms100TopologyTagObjectsResponse.md)
 - [Ttms100TopologyTagRequest](docs/Ttms100TopologyTagRequest.md)
 - [User](docs/User.md)
 - [Vlan](docs/Vlan.md)
 - [VnfTag](docs/VnfTag.md)
 - [Vport](docs/Vport.md)
 - [Vportvalue](docs/Vportvalue.md)


## Documentation For Authorization

## auth
- **Type**: OAuth
- **Flow**: password
- **Authorization URL**: 
- **Scopes**: 
 - **Oauth2**: Oauth2

Example
```
	auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
    r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.
```
	import 	"golang.org/x/oauth2"

    / .. Perform OAuth2 round trip request and obtain a token .. //

    tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
	auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
    r, err := client.Service.Operation(auth, args)
```

## Author



