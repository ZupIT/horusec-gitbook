package horusec

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	"github.com/ZupIT/horusec-operator/api/v2alpha1"
	"github.com/ZupIT/horusec-operator/internal/horusec/auth"
	"github.com/ZupIT/horusec-operator/internal/inventory"
	"github.com/ZupIT/horusec-operator/internal/operation"
)

type Adapter struct {
	scheme *runtime.Scheme
	svc    *Service

	resource *v2alpha1.HorusecPlatform
}

func (a *Adapter) EnsureInitialization(ctx context.Context) (*operation.Result, error) {
	if a.resource.Status.Conditions != nil {
		return operation.ContinueProcessing()
	}
	a.resource.Status.Conditions = []v2alpha1.Condition{}
	a.resource.Status.State = v2alpha1.StatusPending
	err := a.svc.UpdateHorusecPlatformStatus(ctx, a.resource)
	if err != nil {
		return operation.RequeueWithError(err)
	}
	return operation.StopProcessing()
}

//nolint:funlen // to improve in the future
func (a *Adapter) EnsureAuthDeployments(ctx context.Context) (*operation.Result, error) {
	r := a.resource
	desired := auth.NewDeployment(r)
	if err := controllerutil.SetControllerReference(r, desired, a.scheme); err != nil {
		return nil, fmt.Errorf("failed to set Deployment %q owner reference: %v", desired.GetName(), err)
	}

	deps, err := a.svc.ListAuthDeployments(ctx, r.Namespace)
	if err != nil {
		return nil, err
	}

	inv := inventory.ForDeployments(deps.Items, []appsv1.Deployment{*desired})
	err = a.svc.Apply(ctx, inv)
	if err != nil {
		return nil, err
	}

	return operation.ContinueProcessing()
}

func (a *Adapter) EnsureDatabaseConnectivity(ctx context.Context) (*operation.Result, error) {
	panic("implement me") // TODO
}

func (a *Adapter) EnsureBrokerConnectivity(ctx context.Context) (*operation.Result, error) {
	panic("implement me") // TODO
}

func (a *Adapter) EnsureSMTPConnectivity(ctx context.Context) (*operation.Result, error) {
	panic("implement me") // TODO
}

func (a *Adapter) EnsureDatabaseMigrations(ctx context.Context) (*operation.Result, error) {
	panic("implement me") // TODO
}

func (a *Adapter) EnsureDeployments(ctx context.Context) (*operation.Result, error) {
	panic("implement me") // TODO
}

func (a *Adapter) EnsureServices(ctx context.Context) (*operation.Result, error) {
	panic("implement me") // TODO
}

func (a *Adapter) EnsureServicesAccounts(ctx context.Context) (*operation.Result, error) {
	servicesAccounts, err := a.svc.ListAuthServiceAccounts(ctx, a.resource.GetNamespace())
	if err != nil {
		return nil, err
	}

	desired := auth.NewServiceAccount(a.resource)
	if err = controllerutil.SetControllerReference(a.resource, desired, a.scheme); err != nil {
		return nil, fmt.Errorf("failed to set Service Account %q owner reference: %v", desired.GetName(), err)
	}

	inv := inventory.ForServiceAccount(servicesAccounts.Items, []core.ServiceAccount{*desired})
	if err := a.svc.Apply(ctx, inv); err != nil {
		return nil, err
	}
	return operation.ContinueProcessing()
}

func (a *Adapter) EnsureAutoscalers(ctx context.Context) (*operation.Result, error) {
	panic("implement me") // TODO
}

func (a *Adapter) EnsureHPA(ctx context.Context) (*operation.Result, error) {
	panic("implement me") // TODO
}

func (a *Adapter) EnsureIngressRules(ctx context.Context) (*operation.Result, error) {
	panic("implement me") // TODO
}

func (a *Adapter) EnsureEverythingIsRunning(ctx context.Context) (*operation.Result, error) {
	panic("implement me") // TODO
}