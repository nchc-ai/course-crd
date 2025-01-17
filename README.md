# Course-CRD


## Usage

Use [kube-code-generator](https://github.com/slok/kube-code-generator/tree/main) to build CRD clientset. 

## Update CRD Clientset

1. Start from a clean devcontainer. Avoid package conflict. 

2. Use `make clean` to remove previous generated clientset.

2. Update corresponding [kube-code-generator version](https://github.com/slok/kube-code-generator/tree/main?tab=readme-ov-file#kubernetes-versions) in [Makefile](./Makefile)

3. Update miminal package required in [go.mod](./go.mod). For v0.32.1 example. 
    ```
    module github.com/nchc-ai/course-crd

    go 1.22.0

    toolchain go1.23.5

    require (
        k8s.io/apimachinery v0.32.1
        k8s.io/client-go v0.32.1
    )
    ```
4. Then Run `go get ./...` to update the indirect dependency in `go.mod`

5. Run `make` to generate clientset.

6. After clienset is generated, there will be some missing import error. run `go mod tidy` to get the required dependency for the generated code. 