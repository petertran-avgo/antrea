// Copyright 2020 Antrea Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	crdv1beta1 "antrea.io/antrea/pkg/apis/crd/v1beta1"
)

type AntreaNetworkPolicySpecBuilder struct {
	Spec      crdv1beta1.NetworkPolicySpec
	Name      string
	Namespace string
}

type ANNPAppliedToSpec struct {
	ExternalEntitySelector         map[string]string
	ExternalEntitySelectorMatchExp []metav1.LabelSelectorRequirement
	PodSelector                    map[string]string
	PodSelectorMatchExp            []metav1.LabelSelectorRequirement
	Group                          string
}

func (b *AntreaNetworkPolicySpecBuilder) Get() *crdv1beta1.NetworkPolicy {
	if b.Spec.Ingress == nil {
		b.Spec.Ingress = []crdv1beta1.Rule{}
	}
	if b.Spec.Egress == nil {
		b.Spec.Egress = []crdv1beta1.Rule{}
	}
	return &crdv1beta1.NetworkPolicy{
		ObjectMeta: metav1.ObjectMeta{
			Name:      b.Name,
			Namespace: b.Namespace,
		},
		Spec: b.Spec,
	}
}

func (b *AntreaNetworkPolicySpecBuilder) SetName(namespace string, name string) *AntreaNetworkPolicySpecBuilder {
	b.Name = name
	b.Namespace = namespace
	return b
}

func (b *AntreaNetworkPolicySpecBuilder) SetPriority(p float64) *AntreaNetworkPolicySpecBuilder {
	b.Spec.Priority = p
	return b
}

func (b *AntreaNetworkPolicySpecBuilder) SetTier(tier string) *AntreaNetworkPolicySpecBuilder {
	b.Spec.Tier = tier
	return b
}

func (b *AntreaNetworkPolicySpecBuilder) SetAppliedToGroup(specs []ANNPAppliedToSpec) *AntreaNetworkPolicySpecBuilder {
	for _, spec := range specs {
		appliedToPeer := b.GetAppliedToPeer(spec.PodSelector, spec.PodSelectorMatchExp, spec.ExternalEntitySelector, spec.ExternalEntitySelectorMatchExp, spec.Group)
		b.Spec.AppliedTo = append(b.Spec.AppliedTo, appliedToPeer)
	}
	return b
}

func (b *AntreaNetworkPolicySpecBuilder) GetAppliedToPeer(podSelector map[string]string,
	podSelectorMatchExp []metav1.LabelSelectorRequirement,
	entitySelector map[string]string,
	entitySelectorMatchExp []metav1.LabelSelectorRequirement,
	appliedToGrp string) crdv1beta1.AppliedTo {
	var ps, ees *metav1.LabelSelector
	if len(entitySelector) > 0 || len(entitySelectorMatchExp) > 0 {
		ees = &metav1.LabelSelector{
			MatchLabels:      entitySelector,
			MatchExpressions: entitySelectorMatchExp,
		}
	}
	if len(podSelector) > 0 || len(podSelectorMatchExp) > 0 {
		ps = &metav1.LabelSelector{
			MatchLabels:      podSelector,
			MatchExpressions: podSelectorMatchExp,
		}
	}
	peer := crdv1beta1.AppliedTo{
		PodSelector:            ps,
		ExternalEntitySelector: ees,
	}
	if appliedToGrp != "" {
		peer.Group = appliedToGrp
	}
	return peer
}

func (b *AntreaNetworkPolicySpecBuilder) AddIngress(rb RuleBuilder) *AntreaNetworkPolicySpecBuilder {
	var ees *metav1.LabelSelector
	var appliedTos []crdv1beta1.AppliedTo
	if b.Spec.Ingress == nil {
		b.Spec.Ingress = []crdv1beta1.Rule{}
	}

	ps := rb.generatePodSelector()
	ns := rb.generateNsSelector()

	if len(rb.EeSelector) > 0 || len(rb.EeSelectorMatchExp) > 0 {
		ees = &metav1.LabelSelector{
			MatchLabels:      rb.EeSelector,
			MatchExpressions: rb.EeSelectorMatchExp,
		}
	}
	var ipBlock *crdv1beta1.IPBlock
	if rb.Cidr != nil {
		ipBlock = &crdv1beta1.IPBlock{
			CIDR: *rb.Cidr,
		}
	}
	for _, at := range rb.ANPRuleAppliedToSpecs {
		appliedTos = append(appliedTos, b.GetAppliedToPeer(at.PodSelector, at.PodSelectorMatchExp, at.ExternalEntitySelector, at.ExternalEntitySelectorMatchExp, at.Group))
	}
	// An empty From/To in ANNP rules evaluates to match all addresses.
	policyPeer := make([]crdv1beta1.NetworkPolicyPeer, 0)
	if ps != nil || ns != nil || ipBlock != nil || rb.RuleGroup != "" || ees != nil {
		policyPeer = []crdv1beta1.NetworkPolicyPeer{{
			PodSelector:            ps,
			NamespaceSelector:      ns,
			ExternalEntitySelector: ees,
			IPBlock:                ipBlock,
			Group:                  rb.RuleGroup,
		}}
	}
	ports, protocols := GenPortsOrProtocols(rb)
	newRule := crdv1beta1.Rule{
		From:        policyPeer,
		Ports:       ports,
		Protocols:   protocols,
		L7Protocols: rb.L7Protocols,
		Action:      &rb.Action,
		Name:        rb.Name,
		AppliedTo:   appliedTos,
	}
	b.Spec.Ingress = append(b.Spec.Ingress, newRule)
	return b
}

func (b *AntreaNetworkPolicySpecBuilder) AddEgress(rb RuleBuilder) *AntreaNetworkPolicySpecBuilder {
	// For simplicity, we just reuse the Ingress code here.  The underlying data model for ingress/egress is identical
	// With the exception of calling the rule `To` vs. `From`.
	c := &AntreaNetworkPolicySpecBuilder{}
	c.AddIngress(RuleBuilder{
		Protoc:                rb.Protoc,
		Port:                  rb.Port,
		PortName:              rb.PortName,
		EndPort:               rb.EndPort,
		IcmpType:              rb.IcmpType,
		IcmpCode:              rb.IcmpCode,
		IgmpType:              rb.IgmpType,
		L7Protocols:           rb.L7Protocols,
		GroupAddress:          rb.GroupAddress,
		Cidr:                  rb.Cidr,
		PodSelector:           rb.PodSelector,
		NsSelector:            rb.NsSelector,
		EeSelector:            rb.EeSelector,
		PodSelectorMatchExp:   rb.PodSelectorMatchExp,
		NsSelectorMatchExp:    rb.NsSelectorMatchExp,
		EeSelectorMatchExp:    rb.EeSelectorMatchExp,
		ANPRuleAppliedToSpecs: rb.ANPRuleAppliedToSpecs,
		Action:                rb.Action,
		RuleGroup:             rb.RuleGroup,
		Name:                  rb.Name,
	})
	theRule := c.Get().Spec.Ingress[0]

	b.Spec.Egress = append(b.Spec.Egress, crdv1beta1.Rule{
		To:          theRule.From,
		Ports:       theRule.Ports,
		Action:      theRule.Action,
		Name:        theRule.Name,
		AppliedTo:   theRule.AppliedTo,
		L7Protocols: theRule.L7Protocols,
	})
	return b
}

func (b *AntreaNetworkPolicySpecBuilder) AddToServicesRule(svcRefs []crdv1beta1.PeerService,
	name string, ruleAppliedToSpecs []ANNPAppliedToSpec, action crdv1beta1.RuleAction) *AntreaNetworkPolicySpecBuilder {
	var appliedTos []crdv1beta1.AppliedTo
	for _, at := range ruleAppliedToSpecs {
		appliedTos = append(appliedTos, b.GetAppliedToPeer(at.PodSelector, at.PodSelectorMatchExp, at.ExternalEntitySelector, at.ExternalEntitySelectorMatchExp, at.Group))
	}
	newRule := crdv1beta1.Rule{
		To:         make([]crdv1beta1.NetworkPolicyPeer, 0),
		ToServices: svcRefs,
		Action:     &action,
		Name:       name,
		AppliedTo:  appliedTos,
	}
	b.Spec.Egress = append(b.Spec.Egress, newRule)
	return b
}

func (b *AntreaNetworkPolicySpecBuilder) AddEgressLogging(logLabel string) *AntreaNetworkPolicySpecBuilder {
	for i, e := range b.Spec.Egress {
		e.EnableLogging = true
		e.LogLabel = logLabel
		b.Spec.Egress[i] = e
	}
	return b
}

func (b *AntreaNetworkPolicySpecBuilder) AddFQDNRule(fqdn string,
	protoc AntreaPolicyProtocol, port *int32, portName *string, endPort *int32, name string,
	specs []ANNPAppliedToSpec, action crdv1beta1.RuleAction) *AntreaNetworkPolicySpecBuilder {
	var appliedTos []crdv1beta1.AppliedTo

	for _, at := range specs {
		appliedTos = append(appliedTos, b.GetAppliedToPeer(at.PodSelector,
			at.PodSelectorMatchExp,
			at.ExternalEntitySelector,
			at.ExternalEntitySelectorMatchExp,
			at.Group))
	}

	policyPeer := []crdv1beta1.NetworkPolicyPeer{{FQDN: fqdn}}
	ports, _ := GenPortsOrProtocols(RuleBuilder{
		Protoc:   protoc,
		Port:     port,
		PortName: portName,
		EndPort:  endPort,
	})
	newRule := crdv1beta1.Rule{
		To:        policyPeer,
		Ports:     ports,
		Action:    &action,
		Name:      name,
		AppliedTo: appliedTos,
	}
	b.Spec.Egress = append(b.Spec.Egress, newRule)
	return b
}
