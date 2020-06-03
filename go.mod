module github.com/nuxeo/jxlabs-nos-yaml-patch

go 1.14

require (
	github.com/Azure/go-autorest/tracing v0.5.0 // indirect
	github.com/docker/docker v0.7.3-0.20190327010347-be7ac8be2ae0 // indirect
	github.com/docker/go-units v0.3.3 // indirect
	github.com/evanphx/json-patch/v5 v5.0.0
	github.com/ghodss/yaml v1.0.0
	github.com/jenkins-x/jx/v2 v2.1.56
	github.com/mattbaird/jsonpatch v0.0.0-20171005235357-81af80346b1a
	github.com/spf13/cobra v1.0.0
	gotest.tools v2.2.0+incompatible // indirect
	k8s.io/metrics v0.18.2 // indirect
)

replace github.com/nuxeo/jxlabs-nos-yaml-patch => ./

replace github.com/jenkins-x/jx/v2 => ../jxlabs-nos-jx

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v10.14.0+incompatible
	github.com/banzaicloud/bank-vaults => github.com/banzaicloud/bank-vaults v0.0.0-20190508130850-5673d28c46bd
	k8s.io/api => k8s.io/api v0.0.0-20190528110122-9ad12a4af326

	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190528110544-fa58353d80f3

	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190221084156-01f179d85dbc

	k8s.io/client-go => k8s.io/client-go v0.0.0-20190528110200-4f3abb12cae2

	k8s.io/metrics => k8s.io/metrics v0.0.0-20181128195641-3954d62a524d

)

exclude github.com/Azure/go-autorest/autores v0.9.0
