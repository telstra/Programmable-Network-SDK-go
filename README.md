# Go API client for TelstraTPN

Telstra Programmable Network is a self-provisioning platform that allows its users to create on-demand connectivity services between multiple end-points and add various network functions to those services. Programmable Network enables to connectivity to a global ecosystem of networking services as well as public and private cloud services. Once you are connected to the platform on one or more POPs (points of presence), you can start creating those services based on the use case that you want to accomplish. The Programmable Network API is available to all customers who have registered to use the Programmable Network. To register, please contact your account representative.

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
*AuthenticationApi* | [**100AuthGeneratetokenPost**](docs/AuthenticationApi.md#100authgeneratetokenpost) | **Post** /1.0.0/auth/generatetoken | Create an authentication token
*AuthenticationApi* | [**100AuthValidatetokenGet**](docs/AuthenticationApi.md#100authvalidatetokenget) | **Get** /1.0.0/auth/validatetoken | Validate authentication token
*ContractsApi* | [**100InventoryLinksContractByLinkidAndContractidGet**](docs/ContractsApi.md#100inventorylinkscontractbylinkidandcontractidget) | **Get** /1.0.0/inventory/links/{linkid}/contract/{contractid} | Get active Contract by ContractID
*ContractsApi* | [**100InventoryLinksContractByLinkidAndContractidPut**](docs/ContractsApi.md#100inventorylinkscontractbylinkidandcontractidput) | **Put** /1.0.0/inventory/links/{linkid}/contract/{contractid} | Update active Contract by ContractID
*ContractsApi* | [**100InventoryLinksContractByLinkidPost**](docs/ContractsApi.md#100inventorylinkscontractbylinkidpost) | **Post** /1.0.0/inventory/links/{linkid}/contract | Create new Contract on specified link
*CustomersApi* | [**100AccountByCustomeruuidGet**](docs/CustomersApi.md#100accountbycustomeruuidget) | **Get** /1.0.0/account/{customeruuid} | Get account information details
*CustomersApi* | [**100AccountUserByCustomeruuidGet**](docs/CustomersApi.md#100accountuserbycustomeruuidget) | **Get** /1.0.0/account/{customeruuid}/user | List users
*DatacentresApi* | [**100InventoryDatacentersGet**](docs/DatacentresApi.md#100inventorydatacentersget) | **Get** /1.0.0/inventory/datacenters | Get list of all the data centers
*EndpointsApi* | [**100InventoryEndpointByEndpointuuidGet**](docs/EndpointsApi.md#100inventoryendpointbyendpointuuidget) | **Get** /1.0.0/inventory/endpoint/{endpointuuid} | Get information about the specified endpoint
*EndpointsApi* | [**100InventoryEndpointsCustomeruuidByCustomeruuidGet**](docs/EndpointsApi.md#100inventoryendpointscustomeruuidbycustomeruuidget) | **Get** /1.0.0/inventory/endpoints/customeruuid/{customeruuid} | Get list of endpoints for a customer
*EndpointsApi* | [**100InventoryRegularendpointPost**](docs/EndpointsApi.md#100inventoryregularendpointpost) | **Post** /1.0.0/inventory/regularendpoint | Create Physical (Port) Endpoint
*EndpointsApi* | [**100InventoryVnfendpointPost**](docs/EndpointsApi.md#100inventoryvnfendpointpost) | **Post** /1.0.0/inventory/vnfendpoint | Create VNF Endpoint
*EndpointsApi* | [**Eis100EndpointsAssignTopologyTagByEndpointuuidPost**](docs/EndpointsApi.md#eis100endpointsassigntopologytagbyendpointuuidpost) | **Post** /eis/1.0.0/endpoints/{endpointuuid}/assign_topology_tag | Assign a Topology Tag to an Endpoint
*LinksApi* | [**100InventoryLinkPost**](docs/LinksApi.md#100inventorylinkpost) | **Post** /1.0.0/inventory/link | Create Link and initial Contract
*LinksApi* | [**100InventoryLinksByLinkidGet**](docs/LinksApi.md#100inventorylinksbylinkidget) | **Get** /1.0.0/inventory/links/{linkid} | Get details of specified link
*LinksApi* | [**100InventoryLinksCustomerByCustomeruuidGet**](docs/LinksApi.md#100inventorylinkscustomerbycustomeruuidget) | **Get** /1.0.0/inventory/links/customer/{customeruuid} | Get active Links
*LinksApi* | [**100InventoryLinksHistoryByLinkidGet**](docs/LinksApi.md#100inventorylinkshistorybylinkidget) | **Get** /1.0.0/inventory/links/history/{linkid} | Get Link history
*TopologiesApi* | [**Ttms100TopologyTagByTopotaguuidDelete**](docs/TopologiesApi.md#ttms100topologytagbytopotaguuiddelete) | **Delete** /ttms/1.0.0/topology_tag/{topotaguuid}/ | Delete a topology tag
*TopologiesApi* | [**Ttms100TopologyTagByTopotaguuidGet**](docs/TopologiesApi.md#ttms100topologytagbytopotaguuidget) | **Get** /ttms/1.0.0/topology_tag/{topotaguuid}/ | Get information about the specified topology tag
*TopologiesApi* | [**Ttms100TopologyTagByTopotaguuidPut**](docs/TopologiesApi.md#ttms100topologytagbytopotaguuidput) | **Put** /ttms/1.0.0/topology_tag/{topotaguuid}/ | Update a topology tag&#39;s name and/or description
*TopologiesApi* | [**Ttms100TopologyTagGet**](docs/TopologiesApi.md#ttms100topologytagget) | **Get** /ttms/1.0.0/topology_tag | List all topology tags
*TopologiesApi* | [**Ttms100TopologyTagObjectsByTopotaguuidGet**](docs/TopologiesApi.md#ttms100topologytagobjectsbytopotaguuidget) | **Get** /ttms/1.0.0/topology_tag/{topotaguuid}/objects/ | List objects for Topology
*TopologiesApi* | [**Ttms100TopologyTagPost**](docs/TopologiesApi.md#ttms100topologytagpost) | **Post** /ttms/1.0.0/topology_tag | Create a named topology tag
*VnfsApi* | [**100MarketplaceImageGet**](docs/VnfsApi.md#100marketplaceimageget) | **Get** /1.0.0/marketplace/image | List images in the Marketplace
*VportsApi* | [**100InventoryRegularvportPost**](docs/VportsApi.md#100inventoryregularvportpost) | **Post** /1.0.0/inventory/regularvport | Create VPort for physical endpoint
*VportsApi* | [**100InventoryVnfVportPost**](docs/VportsApi.md#100inventoryvnfvportpost) | **Post** /1.0.0/inventory/vnf/vport | Create VNF VPort
*VportsApi* | [**100InventoryVportByVportuuidGet**](docs/VportsApi.md#100inventoryvportbyvportuuidget) | **Get** /1.0.0/inventory/vport/{vportuuid} | Get information about the specified VPort


## Documentation For Models

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
 - [Link](docs/Link.md)
 - [Link66](docs/Link66.md)
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
 - [Model100InventoryLinksContractRequest37](docs/Model100InventoryLinksContractRequest37.md)
 - [Model100InventoryLinksContractResponse](docs/Model100InventoryLinksContractResponse.md)
 - [Model100InventoryLinksContractResponse33](docs/Model100InventoryLinksContractResponse33.md)
 - [Model100InventoryLinksContractResponse38](docs/Model100InventoryLinksContractResponse38.md)
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



