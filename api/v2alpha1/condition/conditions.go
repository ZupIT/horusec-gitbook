// Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
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

package condition

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

var ComponentMap = map[string]Type{
	"analytic":      AnalyticAvailable,
	"api":           APIAvailable,
	"auth":          AuthAvailable,
	"core":          CoreAvailable,
	"manager":       ManagerAvailable,
	"messages":      MessagesAvailable,
	"vulnerability": VulnerabilityAvailable,
	"webhook":       WebhookAvailable,
}

type Type string

const (
	AnalyticAvailable      Type = "AnalyticAvailable"
	APIAvailable                = "APIAvailable"
	AuthAvailable               = "AuthAvailable"
	CoreAvailable               = "CoreAvailable"
	ManagerAvailable            = "ManagerAvailable"
	MessagesAvailable           = "MessagesAvailable"
	VulnerabilityAvailable      = "VulnerabilityAvailable"
	WebhookAvailable            = "WebhookAvailable"
)

type Reason struct {
	Type    string
	Message string
}

func DatabaseReason(message string) *Reason {
	return &Reason{Type: "DatabaseError", Message: message}
}

func BrokerReason(message string) *Reason {
	return &Reason{Type: "BrokerError", Message: message}
}

func SecretReason(message string) *Reason {
	return &Reason{Type: "SecretError", Message: message}
}

func ConfigReason(message string) *Reason {
	return &Reason{Type: "ConfigError", Message: message}
}

func True(conditionType Type) metav1.Condition {
	return metav1.Condition{
		Type:    string(conditionType),
		Status:  metav1.ConditionTrue,
		Reason:  "AvailableReplicas",
		Message: "Deployment has minimum availability.",
	}
}

func False(conditionType Type, reason *Reason) metav1.Condition {
	return metav1.Condition{
		Type:    string(conditionType),
		Status:  metav1.ConditionFalse,
		Reason:  reason.Type,
		Message: reason.Message,
	}
}

func Unknown(conditionType Type, reason *Reason) metav1.Condition {
	return metav1.Condition{
		Type:    string(conditionType),
		Status:  metav1.ConditionUnknown,
		Reason:  reason.Type,
		Message: reason.Message,
	}
}
