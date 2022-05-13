[silde](https://docs.google.com/presentation/d/1eSl5V4ir4aezsOqHJ3eW7Ytc8Qc7YCcutH6tcl_e6E4/edit#slide=id.gf3c4a15969_0_55)

## 1. How To Write A Controller

### 1.1. Define Structure

Define in [pkg/apis/foo/v1/type.go](/pkg/apis/foo/v1/type.go), and generate Custom Resource Definition(CRD) by:

```shell
./gen-crd.sh
```

CRD is generated at [manifest/foo_crd.yaml](/manifest/foo_crd.yaml), this can be directly used in any k8s cluster.

### 1.2. Create Custom Resource

1. First need to register the awesome CRD:

    ```shell
    kubectl apply -f ./manifest/foo_crd.yaml
    ```

1. Then create a Custom Resource

    ```shell
    kubectl apply -f ./manifest/foo.yaml
    ```

Nothing will happen to the Custom Resouce, since there is no controller yet.

### 1.3. Implement Controller

1. First generate boilerplates code:

    ```shell
    ./hack/update-codegen.sh
    ```

    This will generate all other source code files under [pkg](/pkg).

1. Implement controller at [cmd/foo-controller](/cmd/foo-controller).

1. Build controller image:

    ```shell
    KO_DOCKER_REPO=localhost:5001/demo ko publish --base-import-paths ./cmd/foo-controller
    ```

### 1.4. Deploy Controller

Note: watch the created custom resource above.

```shell
kubectl apply -f ./manifest/foo-controller-deployment.yaml
kubectl apply -f ./manifest/foo-controller-rbac.yaml
```

After the deployment something magic will happen.


## 2. Knative Serving Demo

### 2.1. Setup Knative

```shell
./hack/knative-setup.sh
```

The script:
1. Create a GKE cluster.
1. Install Knative serving.
1. Install Kourier as network layer.
1. Setup the cluster with a default domain so that we can curl.

### 2.2. Deploy Helloworld

```shell
kubectl apply ./manifest/knative-helloworld.yaml
```

### 2.3. Autoscaling

First get serving address by:
```shell
kubectl get ksvc
```

Then 1M QPS querying the cluster:
```shell
echo "GET http://hello.default.35.239.32.23.sslip.io" | vegeta attack --rate 1000000 --duration 15s>/dev/null
```
