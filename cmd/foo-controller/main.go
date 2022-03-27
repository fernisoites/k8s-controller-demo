package main

import (
	"context"
	"fmt"
	"log"
	"os"

	v1 "github.com/fernisoites/k8s-controller-demo/pkg/apis/foo/v1"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	"sigs.k8s.io/controller-runtime/pkg/client/config"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

type reconciler struct {
	client.Client
}

func (r *reconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	log.Printf("Reconciling")
	var foo = &v1.Foo{}
	if err := r.Get(ctx, req.NamespacedName, foo); err != nil {
		return reconcile.Result{}, fmt.Errorf("failed to get Foo %s: %w", req.String(), err)
	}
	// if foo.Status == "Pending" {
	// 	log.Println("Skip reconcile as pending state.")
	// 	return reconcile.Result{}, nil
	// }

	foo.Replicas++
	foo.Status = "Pending"
	if err := r.Update(ctx, foo); err != nil {
		return reconcile.Result{}, fmt.Errorf("failed to update Foo %s: %w", req.String(), err)
	}
	log.Printf("Reconcile completed with replicas: %d", foo.Replicas)
	return reconcile.Result{}, nil
}

func main() {
	logf.SetLogger(zap.New())

	var log = logf.Log.WithName("builder-examples")

	mgr, err := manager.New(config.GetConfigOrDie(), manager.Options{})
	if err != nil {
		log.Error(err, "could not create manager")
		os.Exit(1)
	}

	err = builder.
		ControllerManagedBy(mgr). // Create the ControllerManagedBy
		For(&v1.Foo{}).           // ReplicaSet is the Application API
		Owns(&v1.Foo{}).
		Complete(&reconciler{
			Client: mgr.GetClient(),
		})
	if err != nil {
		log.Error(err, "could not create controller")
		os.Exit(1)
	}

	if err := mgr.Start(context.Background()); err != nil {
		log.Error(err, "could not start manager")
		os.Exit(1)
	}
}
