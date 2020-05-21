/*
 * NRF NFManagement Service
 *
 * NRF NFManagement Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type InterfaceUpfInfoItem struct {
	InterfaceType         UpInterfaceType `json:"interfaceType" yaml:"interfaceType" bson:"interfaceType" mapstructure:"InterfaceType"`
	Ipv4EndpointAddresses []string        `json:"ipv4EndpointAddresses,omitempty" yaml:"ipv4EndpointAddresses" bson:"ipv4EndpointAddresses" mapstructure:"Ipv4EndpointAddresses"`
	Ipv6EndpointAddresses []string        `json:"ipv6EndpointAddresses,omitempty" yaml:"ipv6EndpointAddresses" bson:"ipv6EndpointAddresses" mapstructure:"Ipv6EndpointAddresses"`
	EndpointFqdn          string          `json:"endpointFqdn,omitempty" yaml:"endpointFqdn" bson:"endpointFqdn" mapstructure:"EndpointFqdn"`
	NetworkInstance       string          `json:"networkInstance,omitempty" yaml:"networkInstance" bson:"networkInstance" mapstructure:"NetworkInstance"`
}
