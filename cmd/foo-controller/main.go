package main

import (
	"context"
	"fmt"
	"log"

	v1 "github.com/fernisoites/k8s-controller-demo/pkg/apis/foo/v1"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type reconciler struct {
	client client.Client
}

func (r *reconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	var foo v1.Foo
	if err := r.client.Get(ctx, req.NamespacedName, &foo); err != nil {
		return nil, fmt.Errorf("failed to get Foo %s: %w", req.String(), err)
	}
	log.Printf("Get foo: %v", foo)
	return nil, nil
}

func main() {
	cfg, err := rest.InClusterConfig()
	mgr, err := manager.New(cfg, manager.Options{Namespace: "default"})
	err = builder.
		ControllerManagedBy(mgr).
		For(&v1.Foo{}).
		Complete(&reconciler{
			client: mgr.GetClient(),
		})

	if err := mgr.Start(); err != nil {
		log.Fatal("Failed start manager: %v", err)
	}
}
