# Course-CRD


## Usage

```shell=
# make sure k8s.io/code-generator package is downloaded in vendor folder
# go mod vendor

# update auto-generated code. Generated files wiil be placed in $GOPATH
$ hack/update-codegen.sh

# move generated from $GOPATH to project folder
$ cp -r $GOPATH/src/gitlab.com/nchc-ai/course-crd/pkg/client pkg
$ cp $GOPATH/src/gitlab.com/nchc-ai/course-crd/pkg/apis/coursecontroller/v1alpha1/zz_generated.deepcopy.go pkg/apis/coursecontroller/v1alpha1 
```