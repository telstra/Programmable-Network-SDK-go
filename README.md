# Go API client for TelstraTPN

Telstra Programmable Network is a self-provisioning platform that allows its users to create on-demand connectivity services between multiple end-points and add various network functions to those services. Programmable Network enables to connectivity to a global ecosystem of networking services as well as public and private cloud services. Once you are connected to the platform on one or more POPs (points of presence), you can start creating those services based on the use case that you want to accomplish. The Programmable Network API is available to all customers who have registered to use the Programmable Network. To register, please contact your account representative.

## Overview

- API version: 2.1.2
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
*AuthenticationApi* | [**CreateAnAuthenticationToken**](docs/AuthenticationApi.md#createanauthenticationtoken) | **Post** /1.0.0/auth/generatetoken | Create an authentication token
*AuthenticationApi* | [**ValidateAuthenticationToken**](docs/AuthenticationApi.md#validateauthenticationtoken) | **Get** /1.0.0/auth/validatetoken | Validate authentication token
*ContractsApi* | [**CreateNewContractOnSpecifiedLink**](docs/ContractsApi.md#createnewcontractonspecifiedlink) | **Post** /1.0.0/inventory/links/{linkid}/contract | Create new Contract on specified link
*ContractsApi* | [**GetActiveContractByContractID**](docs/ContractsApi.md#getactivecontractbycontractid) | **Get** /1.0.0/inventory/links/{linkid}/contract/{contractid} | Get active Contract by ContractID
*ContractsApi* | [**UpdateActiveContractByContractID**](docs/ContractsApi.md#updateactivecontractbycontractid) | **Put** /1.0.0/inventory/links/{linkid}/contract/{contractid} | Update active Contract by ContractID
*CustomersApi* | [**GetAccountInformationDetails**](docs/CustomersApi.md#getaccountinformationdetails) | **Get** /1.0.0/account/{customeruuid} | Get account information details
*CustomersApi* | [**ListUsers**](docs/CustomersApi.md#listusers) | **Get** /1.0.0/account/{customeruuid}/user | List users
*DatacentresApi* | [**GetListOfAllTheDataCenters**](docs/DatacentresApi.md#getlistofallthedatacenters) | **Get** /1.0.0/inventory/datacenters | Get list of all the data centers
*EndpointsApi* | [**CreatePhysicalPortEndpoint**](docs/EndpointsApi.md#createphysicalportendpoint) | **Post** /1.0.0/inventory/regularendpoint | Create Physical (Port) Endpoint
*EndpointsApi* | [**CreateVNFEndpoint**](docs/EndpointsApi.md#createvnfendpoint) | **Post** /1.0.0/inventory/vnfendpoint | Create VNF Endpoint
*EndpointsApi* | [**GetInformationAboutTheSpecifiedEndpoint**](docs/EndpointsApi.md#getinformationaboutthespecifiedendpoint) | **Get** /1.0.0/inventory/endpoint/{endpointuuid} | Get information about the specified endpoint
*EndpointsApi* | [**GetListOfEndpointsForACustomer**](docs/EndpointsApi.md#getlistofendpointsforacustomer) | **Get** /1.0.0/inventory/endpoints/customeruuid/{customeruuid} | Get list of endpoints for a customer
*LinksApi* | [**CreateLinkAndInitialContract**](docs/LinksApi.md#createlinkandinitialcontract) | **Post** /1.0.0/inventory/link | Create Link and initial Contract
*LinksApi* | [**GetActiveLinks**](docs/LinksApi.md#getactivelinks) | **Get** /1.0.0/inventory/links/customer/{customeruuid} | Get active Links
*LinksApi* | [**GetDetailsOfSpecifiedLink**](docs/LinksApi.md#getdetailsofspecifiedlink) | **Get** /1.0.0/inventory/links/{linkid} | Get details of specified link
*LinksApi* | [**GetLinkHistory**](docs/LinksApi.md#getlinkhistory) | **Get** /1.0.0/inventory/links/history/{linkid} | Get Link history
*TopologiesApi* | [**CreateANamedTopologyTag**](docs/TopologiesApi.md#createanamedtopologytag) | **Post** /ttms/1.0.0/topology_tag | Create a named topology tag
*TopologiesApi* | [**GetInformationAboutTheSpecifiedTopologyTag**](docs/TopologiesApi.md#getinformationaboutthespecifiedtopologytag) | **Get** /ttms/1.0.0/topology_tag/{topotaguuid}/ | Get information about the specified topology tag
*TopologiesApi* | [**ListAllTopologyTags**](docs/TopologiesApi.md#listalltopologytags) | **Get** /ttms/1.0.0/topology_tag | List all topology tags
*TopologiesApi* | [**ListObjectsForTopology**](docs/TopologiesApi.md#listobjectsfortopology) | **Get** /ttms/1.0.0/topology_tag/{topotaguuid}/objects/ | List objects for Topology
*VnfsApi* | [**ListImagesInTheMarketplace**](docs/VnfsApi.md#listimagesinthemarketplace) | **Get** /1.0.0/marketplace/image | List images in the Marketplace
*VportsApi* | [**CreateVNFVPort**](docs/VportsApi.md#createvnfvport) | **Post** /1.0.0/inventory/vnf/vport | Create VNF VPort
*VportsApi* | [**CreateVPortForPhysicalEndpoint**](docs/VportsApi.md#createvportforphysicalendpoint) | **Post** /1.0.0/inventory/regularvport | Create VPort for physical endpoint
*VportsApi* | [**GetInformationAboutTheSpecifiedVPort**](docs/VportsApi.md#getinformationaboutthespecifiedvport) | **Get** /1.0.0/inventory/vport/{vportuuid} | Get information about the specified VPort


## Documentation For Models

 - [Billing](docs/Billing.md)
 - [Classification](docs/Classification.md)
 - [Contract](docs/Contract.md)
 - [Datacenter](docs/Datacenter.md)
 - [Endpoint](docs/Endpoint.md)
 - [EndpointPort](docs/EndpointPort.md)
 - [Endpointlist](docs/Endpointlist.md)
 - [Error70](docs/Error70.md)
 - [Flavor](docs/Flavor.md)
 - [GlanceImage](docs/GlanceImage.md)
 - [Link](docs/Link.md)
 - [Link62](docs/Link62.md)
 - [Meta](docs/Meta.md)
 - [Model100AccountResponse](docs/Model100AccountResponse.md)
 - [Model100AuthGeneratetokenResponse](docs/Model100AuthGeneratetokenResponse.md)
 - [Model100AuthValidatetokenResponse](docs/Model100AuthValidatetokenResponse.md)
 - [Model100InventoryDatacenters401Error](docs/Model100InventoryDatacenters401Error.md)
 - [Model100InventoryDatacentersResponse](docs/Model100InventoryDatacentersResponse.md)
 - [Model100InventoryEndpointResponse](docs/Model100InventoryEndpointResponse.md)
 - [Model100InventoryEndpointsCustomeruuidResponse](docs/Model100InventoryEndpointsCustomeruuidResponse.md)
 - [Model100InventoryLinkRequest](docs/Model100InventoryLinkRequest.md)
 - [Model100InventoryLinkResponse](docs/Model100InventoryLinkResponse.md)
 - [Model100InventoryLinksContractRequest](docs/Model100InventoryLinksContractRequest.md)
 - [Model100InventoryLinksContractRequest35](docs/Model100InventoryLinksContractRequest35.md)
 - [Model100InventoryLinksContractResponse](docs/Model100InventoryLinksContractResponse.md)
 - [Model100InventoryLinksContractResponse31](docs/Model100InventoryLinksContractResponse31.md)
 - [Model100InventoryLinksContractResponse36](docs/Model100InventoryLinksContractResponse36.md)
 - [Model100InventoryLinksHistoryResponse](docs/Model100InventoryLinksHistoryResponse.md)
 - [Model100InventoryLinksResponse](docs/Model100InventoryLinksResponse.md)
 - [Model100InventoryRegularendpointRequest](docs/Model100InventoryRegularendpointRequest.md)
 - [Model100InventoryRegularendpointResponse](docs/Model100InventoryRegularendpointResponse.md)
 - [Model100InventoryRegularvportRequest](docs/Model100InventoryRegularvportRequest.md)
 - [Model100InventoryRegularvportResponse](docs/Model100InventoryRegularvportResponse.md)
 - [Model100InventoryVnfVportRequest](docs/Model100InventoryVnfVportRequest.md)
 - [Model100InventoryVnfVportResponse](docs/Model100InventoryVnfVportResponse.md)
 - [Model100InventoryVnfendpointRequest](docs/Model100InventoryVnfendpointRequest.md)
 - [Model100InventoryVnfendpointResponse](docs/Model100InventoryVnfendpointResponse.md)
 - [Model100MarketplaceImageResponse](docs/Model100MarketplaceImageResponse.md)
 - [ModelError](docs/ModelError.md)
 - [Object50](docs/Object50.md)
 - [Params](docs/Params.md)
 - [Params29](docs/Params29.md)
 - [Params32](docs/Params32.md)
 - [Params37](docs/Params37.md)
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
- **Scopes**: N/A

