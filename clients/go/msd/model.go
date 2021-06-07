//
// Code generated by rdl 1.5.2 DO NOT EDIT.
//

package msd

import (
	"encoding/json"
	"fmt"
	rdl "github.com/ardielle/ardielle-go/rdl"
)

var _ = rdl.Version
var _ = json.Marshal
var _ = fmt.Printf

//
// SimpleName - Copyright The Athenz Authors Licensed under the terms of the
// Apache version 2.0 license. See LICENSE file for terms. Common name types
// used by several API definitions A simple identifier, an element of compound
// name.
//
type SimpleName string

//
// CompoundName - A compound name. Most names in this API are compound names.
//
type CompoundName string

//
// DomainName - A domain name is the general qualifier prefix, as its
// uniqueness is managed.
//
type DomainName string

//
// EntityName - An entity name is a short form of a resource name, including
// only the domain and entity.
//
type EntityName string

//
// EntityList - An Entity list is comma separated compound Names
//
type EntityList string

//
// ServiceName - A service name will generally be a unique subdomain.
//
type ServiceName string

//
// ActionName - An action (operation) name.
//
type ActionName string

//
// ResourceName - A resource name Note that the EntityName part is optional,
// that is, a domain name followed by a colon is valid resource name.
//
type ResourceName string

//
// YBase64 - The Y-specific URL-safe Base64 variant.
//
type YBase64 string

//
// YEncoded - YEncoded includes ybase64 chars, as well as = and %. This can
// represent a user cookie and URL-encoded values.
//
type YEncoded string

//
// AuthorityName - Used as the prefix in a signed assertion. This uniquely
// identifies a signing authority.
//
type AuthorityName string

//
// PathElement - A uri-safe path element
//
type PathElement string

//
// TransportPolicyEnforcementState - Types of transport policy enforcement
// states
//
type TransportPolicyEnforcementState int

//
// TransportPolicyEnforcementState constants
//
const (
	_ TransportPolicyEnforcementState = iota
	ENFORCE
	REPORT
)

var namesTransportPolicyEnforcementState = []string{
	ENFORCE: "ENFORCE",
	REPORT:  "REPORT",
}

//
// NewTransportPolicyEnforcementState - return a string representation of the enum
//
func NewTransportPolicyEnforcementState(init ...interface{}) TransportPolicyEnforcementState {
	if len(init) == 1 {
		switch v := init[0].(type) {
		case TransportPolicyEnforcementState:
			return v
		case int:
			return TransportPolicyEnforcementState(v)
		case int32:
			return TransportPolicyEnforcementState(v)
		case string:
			for i, s := range namesTransportPolicyEnforcementState {
				if s == v {
					return TransportPolicyEnforcementState(i)
				}
			}
		default:
			panic("Bad init value for TransportPolicyEnforcementState enum")
		}
	}
	return TransportPolicyEnforcementState(0) //default to the first enum value
}

//
// String - return a string representation of the enum
//
func (e TransportPolicyEnforcementState) String() string {
	return namesTransportPolicyEnforcementState[e]
}

//
// SymbolSet - return an array of all valid string representations (symbols) of the enum
//
func (e TransportPolicyEnforcementState) SymbolSet() []string {
	return namesTransportPolicyEnforcementState
}

//
// MarshalJSON is defined for proper JSON encoding of a TransportPolicyEnforcementState
//
func (e TransportPolicyEnforcementState) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

//
// UnmarshalJSON is defined for proper JSON decoding of a TransportPolicyEnforcementState
//
func (e *TransportPolicyEnforcementState) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err == nil {
		s := string(j)
		for v, s2 := range namesTransportPolicyEnforcementState {
			if s == s2 {
				*e = TransportPolicyEnforcementState(v)
				return nil
			}
		}
		err = fmt.Errorf("Bad enum symbol for type TransportPolicyEnforcementState: %s", s)
	}
	return err
}

//
// TransportPolicyProtocol - Types of transport policy protocols
//
type TransportPolicyProtocol int

//
// TransportPolicyProtocol constants
//
const (
	_ TransportPolicyProtocol = iota
	TCP
	UDP
)

var namesTransportPolicyProtocol = []string{
	TCP: "TCP",
	UDP: "UDP",
}

//
// NewTransportPolicyProtocol - return a string representation of the enum
//
func NewTransportPolicyProtocol(init ...interface{}) TransportPolicyProtocol {
	if len(init) == 1 {
		switch v := init[0].(type) {
		case TransportPolicyProtocol:
			return v
		case int:
			return TransportPolicyProtocol(v)
		case int32:
			return TransportPolicyProtocol(v)
		case string:
			for i, s := range namesTransportPolicyProtocol {
				if s == v {
					return TransportPolicyProtocol(i)
				}
			}
		default:
			panic("Bad init value for TransportPolicyProtocol enum")
		}
	}
	return TransportPolicyProtocol(0) //default to the first enum value
}

//
// String - return a string representation of the enum
//
func (e TransportPolicyProtocol) String() string {
	return namesTransportPolicyProtocol[e]
}

//
// SymbolSet - return an array of all valid string representations (symbols) of the enum
//
func (e TransportPolicyProtocol) SymbolSet() []string {
	return namesTransportPolicyProtocol
}

//
// MarshalJSON is defined for proper JSON encoding of a TransportPolicyProtocol
//
func (e TransportPolicyProtocol) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

//
// UnmarshalJSON is defined for proper JSON decoding of a TransportPolicyProtocol
//
func (e *TransportPolicyProtocol) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err == nil {
		s := string(j)
		for v, s2 := range namesTransportPolicyProtocol {
			if s == s2 {
				*e = TransportPolicyProtocol(v)
				return nil
			}
		}
		err = fmt.Errorf("Bad enum symbol for type TransportPolicyProtocol: %s", s)
	}
	return err
}

//
// TransportPolicySubject - Subject for a transport policy
//
type TransportPolicySubject struct {

	//
	// Name of the domain
	//
	DomainName DomainName `json:"domainName"`

	//
	// Name of the service
	//
	ServiceName EntityName `json:"serviceName"`
}

//
// NewTransportPolicySubject - creates an initialized TransportPolicySubject instance, returns a pointer to it
//
func NewTransportPolicySubject(init ...*TransportPolicySubject) *TransportPolicySubject {
	var o *TransportPolicySubject
	if len(init) == 1 {
		o = init[0]
	} else {
		o = new(TransportPolicySubject)
	}
	return o
}

type rawTransportPolicySubject TransportPolicySubject

//
// UnmarshalJSON is defined for proper JSON decoding of a TransportPolicySubject
//
func (self *TransportPolicySubject) UnmarshalJSON(b []byte) error {
	var m rawTransportPolicySubject
	err := json.Unmarshal(b, &m)
	if err == nil {
		o := TransportPolicySubject(m)
		*self = o
		err = self.Validate()
	}
	return err
}

//
// Validate - checks for missing required fields, etc
//
func (self *TransportPolicySubject) Validate() error {
	if self.DomainName == "" {
		return fmt.Errorf("TransportPolicySubject.domainName is missing but is a required field")
	} else {
		val := rdl.Validate(MSDSchema(), "DomainName", self.DomainName)
		if !val.Valid {
			return fmt.Errorf("TransportPolicySubject.domainName does not contain a valid DomainName (%v)", val.Error)
		}
	}
	if self.ServiceName == "" {
		return fmt.Errorf("TransportPolicySubject.serviceName is missing but is a required field")
	} else {
		val := rdl.Validate(MSDSchema(), "EntityName", self.ServiceName)
		if !val.Valid {
			return fmt.Errorf("TransportPolicySubject.serviceName does not contain a valid EntityName (%v)", val.Error)
		}
	}
	return nil
}

//
// TransportPolicyCondition - Transport policy condition. Used to specify
// additional restrictions for the subject of a transport policy
//
type TransportPolicyCondition struct {

	//
	// State of transport policy enforcement ( ENFORCE / REPORT )
	//
	EnforcementState TransportPolicyEnforcementState `json:"enforcementState"`

	//
	// Acts as restrictions. If present, this transport policy should be
	// restricted to only mentioned instances.
	//
	Instances []string `json:"instances,omitempty" rdl:"optional"`
}

//
// NewTransportPolicyCondition - creates an initialized TransportPolicyCondition instance, returns a pointer to it
//
func NewTransportPolicyCondition(init ...*TransportPolicyCondition) *TransportPolicyCondition {
	var o *TransportPolicyCondition
	if len(init) == 1 {
		o = init[0]
	} else {
		o = new(TransportPolicyCondition)
	}
	return o
}

type rawTransportPolicyCondition TransportPolicyCondition

//
// UnmarshalJSON is defined for proper JSON decoding of a TransportPolicyCondition
//
func (self *TransportPolicyCondition) UnmarshalJSON(b []byte) error {
	var m rawTransportPolicyCondition
	err := json.Unmarshal(b, &m)
	if err == nil {
		o := TransportPolicyCondition(m)
		*self = o
		err = self.Validate()
	}
	return err
}

//
// Validate - checks for missing required fields, etc
//
func (self *TransportPolicyCondition) Validate() error {
	return nil
}

//
// TransportPolicyPort - Transport policy port
//
type TransportPolicyPort struct {

	//
	// Start port of the port range. port and endPort will have same values for a
	// single port definition.
	//
	Port int32 `json:"port"`

	//
	// End port of the port range. port and endPort will have same values for a
	// single port definition.
	//
	EndPort int32 `json:"endPort"`

	//
	// Protocol for this transport policy
	//
	Protocol TransportPolicyProtocol `json:"protocol"`
}

//
// NewTransportPolicyPort - creates an initialized TransportPolicyPort instance, returns a pointer to it
//
func NewTransportPolicyPort(init ...*TransportPolicyPort) *TransportPolicyPort {
	var o *TransportPolicyPort
	if len(init) == 1 {
		o = init[0]
	} else {
		o = new(TransportPolicyPort)
	}
	return o
}

type rawTransportPolicyPort TransportPolicyPort

//
// UnmarshalJSON is defined for proper JSON decoding of a TransportPolicyPort
//
func (self *TransportPolicyPort) UnmarshalJSON(b []byte) error {
	var m rawTransportPolicyPort
	err := json.Unmarshal(b, &m)
	if err == nil {
		o := TransportPolicyPort(m)
		*self = o
		err = self.Validate()
	}
	return err
}

//
// Validate - checks for missing required fields, etc
//
func (self *TransportPolicyPort) Validate() error {
	return nil
}

//
// TransportPolicyMatch - Selector for the subject of a transport policy
//
type TransportPolicyMatch struct {

	//
	// Subject where this transport policy applies
	//
	AthenzService *TransportPolicySubject `json:"athenzService"`

	//
	// List of additional requirements for restrictions. Requirements are ANDed.
	//
	Conditions []*TransportPolicyCondition `json:"conditions"`
}

//
// NewTransportPolicyMatch - creates an initialized TransportPolicyMatch instance, returns a pointer to it
//
func NewTransportPolicyMatch(init ...*TransportPolicyMatch) *TransportPolicyMatch {
	var o *TransportPolicyMatch
	if len(init) == 1 {
		o = init[0]
	} else {
		o = new(TransportPolicyMatch)
	}
	return o.Init()
}

//
// Init - sets up the instance according to its default field values, if any
//
func (self *TransportPolicyMatch) Init() *TransportPolicyMatch {
	if self.AthenzService == nil {
		self.AthenzService = NewTransportPolicySubject()
	}
	if self.Conditions == nil {
		self.Conditions = make([]*TransportPolicyCondition, 0)
	}
	return self
}

type rawTransportPolicyMatch TransportPolicyMatch

//
// UnmarshalJSON is defined for proper JSON decoding of a TransportPolicyMatch
//
func (self *TransportPolicyMatch) UnmarshalJSON(b []byte) error {
	var m rawTransportPolicyMatch
	err := json.Unmarshal(b, &m)
	if err == nil {
		o := TransportPolicyMatch(m)
		*self = *((&o).Init())
		err = self.Validate()
	}
	return err
}

//
// Validate - checks for missing required fields, etc
//
func (self *TransportPolicyMatch) Validate() error {
	if self.AthenzService == nil {
		return fmt.Errorf("TransportPolicyMatch: Missing required field: athenzService")
	}
	if self.Conditions == nil {
		return fmt.Errorf("TransportPolicyMatch: Missing required field: conditions")
	}
	return nil
}

//
// TransportPolicyPeer - Source or destination for a transport policy
//
type TransportPolicyPeer struct {

	//
	// List of transport policy subjects
	//
	AthenzServices []*TransportPolicySubject `json:"athenzServices"`

	//
	// List of network traffic port part of this transport policy
	//
	Ports []*TransportPolicyPort `json:"ports"`
}

//
// NewTransportPolicyPeer - creates an initialized TransportPolicyPeer instance, returns a pointer to it
//
func NewTransportPolicyPeer(init ...*TransportPolicyPeer) *TransportPolicyPeer {
	var o *TransportPolicyPeer
	if len(init) == 1 {
		o = init[0]
	} else {
		o = new(TransportPolicyPeer)
	}
	return o.Init()
}

//
// Init - sets up the instance according to its default field values, if any
//
func (self *TransportPolicyPeer) Init() *TransportPolicyPeer {
	if self.AthenzServices == nil {
		self.AthenzServices = make([]*TransportPolicySubject, 0)
	}
	if self.Ports == nil {
		self.Ports = make([]*TransportPolicyPort, 0)
	}
	return self
}

type rawTransportPolicyPeer TransportPolicyPeer

//
// UnmarshalJSON is defined for proper JSON decoding of a TransportPolicyPeer
//
func (self *TransportPolicyPeer) UnmarshalJSON(b []byte) error {
	var m rawTransportPolicyPeer
	err := json.Unmarshal(b, &m)
	if err == nil {
		o := TransportPolicyPeer(m)
		*self = *((&o).Init())
		err = self.Validate()
	}
	return err
}

//
// Validate - checks for missing required fields, etc
//
func (self *TransportPolicyPeer) Validate() error {
	if self.AthenzServices == nil {
		return fmt.Errorf("TransportPolicyPeer: Missing required field: athenzServices")
	}
	if self.Ports == nil {
		return fmt.Errorf("TransportPolicyPeer: Missing required field: ports")
	}
	return nil
}

//
// TransportPolicyEntitySelector - Entity to which a transport policy applies.
// Describes the subject and port(s) for a transport policy.
//
type TransportPolicyEntitySelector struct {

	//
	// Requirements for selecting the subject for this transport policy.
	//
	Match *TransportPolicyMatch `json:"match"`

	//
	// List of network traffic port of the subject eligible for the transport
	// policy
	//
	Ports []*TransportPolicyPort `json:"ports"`
}

//
// NewTransportPolicyEntitySelector - creates an initialized TransportPolicyEntitySelector instance, returns a pointer to it
//
func NewTransportPolicyEntitySelector(init ...*TransportPolicyEntitySelector) *TransportPolicyEntitySelector {
	var o *TransportPolicyEntitySelector
	if len(init) == 1 {
		o = init[0]
	} else {
		o = new(TransportPolicyEntitySelector)
	}
	return o.Init()
}

//
// Init - sets up the instance according to its default field values, if any
//
func (self *TransportPolicyEntitySelector) Init() *TransportPolicyEntitySelector {
	if self.Match == nil {
		self.Match = NewTransportPolicyMatch()
	}
	if self.Ports == nil {
		self.Ports = make([]*TransportPolicyPort, 0)
	}
	return self
}

type rawTransportPolicyEntitySelector TransportPolicyEntitySelector

//
// UnmarshalJSON is defined for proper JSON decoding of a TransportPolicyEntitySelector
//
func (self *TransportPolicyEntitySelector) UnmarshalJSON(b []byte) error {
	var m rawTransportPolicyEntitySelector
	err := json.Unmarshal(b, &m)
	if err == nil {
		o := TransportPolicyEntitySelector(m)
		*self = *((&o).Init())
		err = self.Validate()
	}
	return err
}

//
// Validate - checks for missing required fields, etc
//
func (self *TransportPolicyEntitySelector) Validate() error {
	if self.Match == nil {
		return fmt.Errorf("TransportPolicyEntitySelector: Missing required field: match")
	}
	if self.Ports == nil {
		return fmt.Errorf("TransportPolicyEntitySelector: Missing required field: ports")
	}
	return nil
}

//
// TransportPolicyIngressRule - Transport policy ingress rule
//
type TransportPolicyIngressRule struct {

	//
	// Assertion id associated with this transport policy
	//
	Id int64 `json:"id"`

	//
	// Last modification timestamp of this transport policy
	//
	LastModified rdl.Timestamp `json:"lastModified"`

	//
	// Describes the entity to which this transport policy applies
	//
	EntitySelector *TransportPolicyEntitySelector `json:"entitySelector"`

	//
	// Source of network traffic
	//
	From *TransportPolicyPeer `json:"from"`
}

//
// NewTransportPolicyIngressRule - creates an initialized TransportPolicyIngressRule instance, returns a pointer to it
//
func NewTransportPolicyIngressRule(init ...*TransportPolicyIngressRule) *TransportPolicyIngressRule {
	var o *TransportPolicyIngressRule
	if len(init) == 1 {
		o = init[0]
	} else {
		o = new(TransportPolicyIngressRule)
	}
	return o.Init()
}

//
// Init - sets up the instance according to its default field values, if any
//
func (self *TransportPolicyIngressRule) Init() *TransportPolicyIngressRule {
	if self.EntitySelector == nil {
		self.EntitySelector = NewTransportPolicyEntitySelector()
	}
	if self.From == nil {
		self.From = NewTransportPolicyPeer()
	}
	return self
}

type rawTransportPolicyIngressRule TransportPolicyIngressRule

//
// UnmarshalJSON is defined for proper JSON decoding of a TransportPolicyIngressRule
//
func (self *TransportPolicyIngressRule) UnmarshalJSON(b []byte) error {
	var m rawTransportPolicyIngressRule
	err := json.Unmarshal(b, &m)
	if err == nil {
		o := TransportPolicyIngressRule(m)
		*self = *((&o).Init())
		err = self.Validate()
	}
	return err
}

//
// Validate - checks for missing required fields, etc
//
func (self *TransportPolicyIngressRule) Validate() error {
	if self.LastModified.IsZero() {
		return fmt.Errorf("TransportPolicyIngressRule: Missing required field: lastModified")
	}
	if self.EntitySelector == nil {
		return fmt.Errorf("TransportPolicyIngressRule: Missing required field: entitySelector")
	}
	if self.From == nil {
		return fmt.Errorf("TransportPolicyIngressRule: Missing required field: from")
	}
	return nil
}

//
// TransportPolicyEgressRule - Transport policy egress rule
//
type TransportPolicyEgressRule struct {

	//
	// Assertion id associated with this transport policy
	//
	Id int64 `json:"id"`

	//
	// Last modification timestamp of this transport policy
	//
	LastModified rdl.Timestamp `json:"lastModified"`

	//
	// Entity to which this transport policy applies
	//
	EntitySelector *TransportPolicyEntitySelector `json:"entitySelector"`

	//
	// Destination of network traffic
	//
	To *TransportPolicyPeer `json:"to"`
}

//
// NewTransportPolicyEgressRule - creates an initialized TransportPolicyEgressRule instance, returns a pointer to it
//
func NewTransportPolicyEgressRule(init ...*TransportPolicyEgressRule) *TransportPolicyEgressRule {
	var o *TransportPolicyEgressRule
	if len(init) == 1 {
		o = init[0]
	} else {
		o = new(TransportPolicyEgressRule)
	}
	return o.Init()
}

//
// Init - sets up the instance according to its default field values, if any
//
func (self *TransportPolicyEgressRule) Init() *TransportPolicyEgressRule {
	if self.EntitySelector == nil {
		self.EntitySelector = NewTransportPolicyEntitySelector()
	}
	if self.To == nil {
		self.To = NewTransportPolicyPeer()
	}
	return self
}

type rawTransportPolicyEgressRule TransportPolicyEgressRule

//
// UnmarshalJSON is defined for proper JSON decoding of a TransportPolicyEgressRule
//
func (self *TransportPolicyEgressRule) UnmarshalJSON(b []byte) error {
	var m rawTransportPolicyEgressRule
	err := json.Unmarshal(b, &m)
	if err == nil {
		o := TransportPolicyEgressRule(m)
		*self = *((&o).Init())
		err = self.Validate()
	}
	return err
}

//
// Validate - checks for missing required fields, etc
//
func (self *TransportPolicyEgressRule) Validate() error {
	if self.LastModified.IsZero() {
		return fmt.Errorf("TransportPolicyEgressRule: Missing required field: lastModified")
	}
	if self.EntitySelector == nil {
		return fmt.Errorf("TransportPolicyEgressRule: Missing required field: entitySelector")
	}
	if self.To == nil {
		return fmt.Errorf("TransportPolicyEgressRule: Missing required field: to")
	}
	return nil
}

//
// TransportPolicyRules - Transport policy containing ingress and egress rules
//
type TransportPolicyRules struct {

	//
	// List of ingress rules
	//
	Ingress []*TransportPolicyIngressRule `json:"ingress"`

	//
	// List of egress rules
	//
	Egress []*TransportPolicyEgressRule `json:"egress"`
}

//
// NewTransportPolicyRules - creates an initialized TransportPolicyRules instance, returns a pointer to it
//
func NewTransportPolicyRules(init ...*TransportPolicyRules) *TransportPolicyRules {
	var o *TransportPolicyRules
	if len(init) == 1 {
		o = init[0]
	} else {
		o = new(TransportPolicyRules)
	}
	return o.Init()
}

//
// Init - sets up the instance according to its default field values, if any
//
func (self *TransportPolicyRules) Init() *TransportPolicyRules {
	if self.Ingress == nil {
		self.Ingress = make([]*TransportPolicyIngressRule, 0)
	}
	if self.Egress == nil {
		self.Egress = make([]*TransportPolicyEgressRule, 0)
	}
	return self
}

type rawTransportPolicyRules TransportPolicyRules

//
// UnmarshalJSON is defined for proper JSON decoding of a TransportPolicyRules
//
func (self *TransportPolicyRules) UnmarshalJSON(b []byte) error {
	var m rawTransportPolicyRules
	err := json.Unmarshal(b, &m)
	if err == nil {
		o := TransportPolicyRules(m)
		*self = *((&o).Init())
		err = self.Validate()
	}
	return err
}

//
// Validate - checks for missing required fields, etc
//
func (self *TransportPolicyRules) Validate() error {
	if self.Ingress == nil {
		return fmt.Errorf("TransportPolicyRules: Missing required field: ingress")
	}
	if self.Egress == nil {
		return fmt.Errorf("TransportPolicyRules: Missing required field: egress")
	}
	return nil
}

//
// Workload - workload type describing workload associated with an identity
//
type Workload struct {

	//
	// name of the domain, optional for getWorkloadsByService API call
	//
	DomainName DomainName `json:"domainName"`

	//
	// name of the service, , optional for getWorkloadsByService API call
	//
	ServiceName EntityName `json:"serviceName"`

	//
	// unique identifier for the workload, usually defined by provider
	//
	Uuid string `json:"uuid"`

	//
	// list of IP addresses associated with the workload, optional for
	// getWorkloadsByIP API call
	//
	IpAddresses []string `json:"ipAddresses"`

	//
	// hostname associated with the workload
	//
	Hostname string `json:"hostname"`

	//
	// infrastructure provider e.g. k8s, AWS, Azure, openstack etc.
	//
	Provider string `json:"provider"`

	//
	// most recent update timestamp in the backend
	//
	UpdateTime rdl.Timestamp `json:"updateTime"`

	//
	// certificate expiry time (ex: getNotAfter)
	//
	CertExpiryTime rdl.Timestamp `json:"certExpiryTime"`
}

//
// NewWorkload - creates an initialized Workload instance, returns a pointer to it
//
func NewWorkload(init ...*Workload) *Workload {
	var o *Workload
	if len(init) == 1 {
		o = init[0]
	} else {
		o = new(Workload)
	}
	return o.Init()
}

//
// Init - sets up the instance according to its default field values, if any
//
func (self *Workload) Init() *Workload {
	if self.IpAddresses == nil {
		self.IpAddresses = make([]string, 0)
	}
	return self
}

type rawWorkload Workload

//
// UnmarshalJSON is defined for proper JSON decoding of a Workload
//
func (self *Workload) UnmarshalJSON(b []byte) error {
	var m rawWorkload
	err := json.Unmarshal(b, &m)
	if err == nil {
		o := Workload(m)
		*self = *((&o).Init())
		err = self.Validate()
	}
	return err
}

//
// Validate - checks for missing required fields, etc
//
func (self *Workload) Validate() error {
	if self.DomainName == "" {
		return fmt.Errorf("Workload.domainName is missing but is a required field")
	} else {
		val := rdl.Validate(MSDSchema(), "DomainName", self.DomainName)
		if !val.Valid {
			return fmt.Errorf("Workload.domainName does not contain a valid DomainName (%v)", val.Error)
		}
	}
	if self.ServiceName == "" {
		return fmt.Errorf("Workload.serviceName is missing but is a required field")
	} else {
		val := rdl.Validate(MSDSchema(), "EntityName", self.ServiceName)
		if !val.Valid {
			return fmt.Errorf("Workload.serviceName does not contain a valid EntityName (%v)", val.Error)
		}
	}
	if self.Uuid == "" {
		return fmt.Errorf("Workload.uuid is missing but is a required field")
	} else {
		val := rdl.Validate(MSDSchema(), "String", self.Uuid)
		if !val.Valid {
			return fmt.Errorf("Workload.uuid does not contain a valid String (%v)", val.Error)
		}
	}
	if self.IpAddresses == nil {
		return fmt.Errorf("Workload: Missing required field: ipAddresses")
	}
	if self.Hostname == "" {
		return fmt.Errorf("Workload.hostname is missing but is a required field")
	} else {
		val := rdl.Validate(MSDSchema(), "String", self.Hostname)
		if !val.Valid {
			return fmt.Errorf("Workload.hostname does not contain a valid String (%v)", val.Error)
		}
	}
	if self.Provider == "" {
		return fmt.Errorf("Workload.provider is missing but is a required field")
	} else {
		val := rdl.Validate(MSDSchema(), "String", self.Provider)
		if !val.Valid {
			return fmt.Errorf("Workload.provider does not contain a valid String (%v)", val.Error)
		}
	}
	if self.UpdateTime.IsZero() {
		return fmt.Errorf("Workload: Missing required field: updateTime")
	}
	if self.CertExpiryTime.IsZero() {
		return fmt.Errorf("Workload: Missing required field: certExpiryTime")
	}
	return nil
}

//
// Workloads - list of workloads
//
type Workloads struct {

	//
	// list of workloads
	//
	WorkloadList []*Workload `json:"workloadList"`
}

//
// NewWorkloads - creates an initialized Workloads instance, returns a pointer to it
//
func NewWorkloads(init ...*Workloads) *Workloads {
	var o *Workloads
	if len(init) == 1 {
		o = init[0]
	} else {
		o = new(Workloads)
	}
	return o.Init()
}

//
// Init - sets up the instance according to its default field values, if any
//
func (self *Workloads) Init() *Workloads {
	if self.WorkloadList == nil {
		self.WorkloadList = make([]*Workload, 0)
	}
	return self
}

type rawWorkloads Workloads

//
// UnmarshalJSON is defined for proper JSON decoding of a Workloads
//
func (self *Workloads) UnmarshalJSON(b []byte) error {
	var m rawWorkloads
	err := json.Unmarshal(b, &m)
	if err == nil {
		o := Workloads(m)
		*self = *((&o).Init())
		err = self.Validate()
	}
	return err
}

//
// Validate - checks for missing required fields, etc
//
func (self *Workloads) Validate() error {
	if self.WorkloadList == nil {
		return fmt.Errorf("Workloads: Missing required field: workloadList")
	}
	return nil
}
