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

package resources

import (
	"fmt"

	coreV1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	"github.com/ZupIT/horusec-operator/api/v2alpha1"
	"github.com/ZupIT/horusec-operator/internal/resources/analytic"
	"github.com/ZupIT/horusec-operator/internal/resources/api"
	"github.com/ZupIT/horusec-operator/internal/resources/auth"
	"github.com/ZupIT/horusec-operator/internal/resources/core"
	"github.com/ZupIT/horusec-operator/internal/resources/manager"
	"github.com/ZupIT/horusec-operator/internal/resources/messages"
	"github.com/ZupIT/horusec-operator/internal/resources/vulnerability"
	"github.com/ZupIT/horusec-operator/internal/resources/webhook"
)

func (b *Builder) ServicesFor(resource *v2alpha1.HorusecPlatform) ([]coreV1.Service, error) {
	desired := b.listServices(resource)
	for index := range desired {
		if err := b.ensureServices(resource, &desired[index]); err != nil {
			return nil, err
		}
	}
	return desired, nil
}

func (b *Builder) ensureServices(resource *v2alpha1.HorusecPlatform, desired *coreV1.Service) error {
	if err := controllerutil.SetControllerReference(resource, desired, b.scheme); err != nil {
		return fmt.Errorf("failed to set service %q owner reference: %v", desired.GetName(), err)
	}

	return nil
}

func (b *Builder) listServices(resource *v2alpha1.HorusecPlatform) []coreV1.Service {
	services := []coreV1.Service{
		analytic.NewService(resource),
		api.NewService(resource),
		auth.NewService(resource),
		core.NewService(resource),
		manager.NewService(resource),
		vulnerability.NewService(resource),
		webhook.NewService(resource),
	}
	msg := resource.GetMessagesComponent()
	if msg.Enabled {
		services = append(services, messages.NewService(resource))
	}
	return services
}
