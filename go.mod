module github.com/solo-io/gloo-mesh

go 1.16

replace (
	// pinned to solo-io's fork of cue version 95a50cebaffb4bdba8c544601d8fb867990ad1ad
	cuelang.org/go => github.com/solo-io/cue v0.4.1-0.20210622213027-95a50cebaffb

	// github.com/Azure/go-autorest/autorest has different versions for the Go
	// modules than it does for releases on the repository. Note the correct
	// version when updating.
	github.com/docker/distribution => github.com/docker/distribution v0.0.0-20191216044856-a8371794149d
	github.com/docker/docker => github.com/moby/moby v0.7.3-0.20190826074503-38ab9da00309
	// github.com/envoyproxy/go-control-plane => github.com/envoyproxy/go-control-plane v0.9.9-0.20210115003313-31f9241a16e6
	github.com/spf13/viper => github.com/istio/viper v1.3.3-0.20190515210538-2789fed3109c

	helm.sh/helm/v3 => helm.sh/helm/v3 v3.5.3

	k8s.io/api => k8s.io/api v0.20.4
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.20.4
	k8s.io/apimachinery => k8s.io/apimachinery v0.20.4
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.20.4
	k8s.io/client-go => k8s.io/client-go v0.20.4

	k8s.io/kube-openapi => github.com/howardjohn/kube-openapi v0.0.0-20210104181841-c0b40d2cb1c8
	k8s.io/kubectl => k8s.io/kubectl v0.20.4
	k8s.io/utils => k8s.io/utils v0.0.0-20201110183641-67b214c5f920
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.7.0
)

require (
	cuelang.org/go v0.4.0
	github.com/Masterminds/semver v1.5.0
	github.com/Masterminds/sprig v2.22.0+incompatible // indirect
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d
	github.com/aws/aws-app-mesh-controller-for-k8s v1.1.1
	github.com/aws/aws-sdk-go v1.38.51
	github.com/docker/distribution v2.7.1+incompatible
	github.com/envoyproxy/go-control-plane v0.9.10-0.20210708144103-3a95f2df6351
	github.com/evanphx/json-patch v4.9.0+incompatible
	github.com/fatih/color v1.12.0
	github.com/gertd/go-pluralize v0.1.1
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32
	github.com/go-openapi/swag v0.19.6
	github.com/go-test/deep v1.0.2
	github.com/gobuffalo/packr v1.30.1
	github.com/gobuffalo/packr/v2 v2.8.1 // indirect
	github.com/gogo/protobuf v1.3.2
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.2
	github.com/google/go-cmp v0.5.6
	github.com/google/go-github/v32 v32.0.0
	github.com/hashicorp/go-multierror v1.1.1
	github.com/hashicorp/go-version v1.3.0
	github.com/iancoleman/strcase v0.1.3
	github.com/kr/pretty v0.2.1
	github.com/linkerd/linkerd2 v0.5.1-0.20200402173539-fee70c064bc0
	github.com/logrusorgru/aurora/v3 v3.0.0
	github.com/manifoldco/promptui v0.8.0
	github.com/mitchellh/hashstructure v1.0.0
	github.com/olekukonko/tablewriter v0.0.4
	github.com/onsi/ginkgo v1.16.2
	github.com/onsi/gomega v1.13.0
	github.com/openservicemesh/osm v0.3.0
	github.com/pkg/browser v0.0.0-20180916011732-0a3d74bf9ce4
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_model v0.2.0
	github.com/prometheus/common v0.25.0
	github.com/pseudomuto/protoc-gen-doc v1.4.1
	github.com/pseudomuto/protokit v0.2.0
	github.com/rotisserie/eris v0.4.0
	github.com/servicemeshinterface/smi-sdk-go v0.4.1
	github.com/sirupsen/logrus v1.8.1
	github.com/solo-io/anyvendor v0.0.3
	github.com/solo-io/external-apis v0.1.7-0.20210720151641-09b1d2bb1f72
	github.com/solo-io/go-list-licenses v0.1.3
	github.com/solo-io/go-utils v0.21.9
	github.com/solo-io/k8s-utils v0.0.3
	github.com/solo-io/protoc-gen-ext v0.0.15
	github.com/solo-io/skv2 v0.18.3
	github.com/solo-io/solo-apis v1.6.30
	github.com/solo-io/solo-kit v0.16.0
	github.com/spf13/afero v1.5.1
	github.com/spf13/cobra v1.2.1
	github.com/spf13/pflag v1.0.5
	github.com/stoewer/go-strcase v1.2.0
	go.uber.org/atomic v1.7.0
	go.uber.org/zap v1.16.0
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.26.0
	gopkg.in/yaml.v2 v2.4.0
	helm.sh/helm/v3 v3.6.0
	istio.io/api v0.0.0-20210713195055-3a340e4f154e
	istio.io/client-go v1.10.3
	istio.io/istio v0.0.0-20210720120633-f8c8dba44cb7
	istio.io/pkg v0.0.0-20210716112334-7ebec3684725
	istio.io/tools v0.0.0-20210420211536-9c0f48df3262
	k8s.io/api v0.21.2
	k8s.io/apiextensions-apiserver v0.21.2
	k8s.io/apimachinery v0.21.2
	k8s.io/client-go v8.0.0+incompatible
	k8s.io/kubernetes v1.13.0
	k8s.io/utils v0.0.0-20210527160623-6fdb442a123b
	sigs.k8s.io/controller-runtime v0.9.0-beta.5
)
