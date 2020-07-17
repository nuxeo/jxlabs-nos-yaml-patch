module github.com/nuxeo/jxlabs-nos-yaml-patch

go 1.13

require (
	github.com/evanphx/json-patch/v5 v5.0.0
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32
	github.com/jenkins-x/jx-logging v0.0.10
	github.com/jenkins-x/jx/v2 v2.1.97
	github.com/mattbaird/jsonpatch v0.0.0-20171005235357-81af80346b1a
	github.com/spf13/cobra v1.0.0
	k8s.io/metrics v0.18.2 // indirect
)

replace github.com/jenkins-x/jx/v2 => github.com/nuxeo/jxlabs-nos-jx/v2 v2.1.1

replace k8s.io/api => k8s.io/api v0.16.5

replace k8s.io/metrics => k8s.io/metrics v0.0.0-20190819143841-305e1cef1ab1

replace k8s.io/apimachinery => k8s.io/apimachinery v0.16.5

replace k8s.io/client-go => k8s.io/client-go v0.16.5

replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190819143637-0dbe462fe92d

replace git.apache.org/thrift.git => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999

replace github.com/sirupsen/logrus => github.com/jtnord/logrus v1.4.2-0.20190423161236-606ffcaf8f5d

replace github.com/Azure/azure-sdk-for-go => github.com/Azure/azure-sdk-for-go v23.2.0+incompatible

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.3.1+incompatible

replace github.com/banzaicloud/bank-vaults => github.com/banzaicloud/bank-vaults v0.0.0-20191212164220-b327d7f2b681

replace github.com/banzaicloud/bank-vaults/pkg/sdk => github.com/banzaicloud/bank-vaults/pkg/sdk v0.0.0-20191212164220-b327d7f2b681

replace k8s.io/test-infra => github.com/jenkins-x/test-infra v0.0.0-20200611142252-211a92405c22

replace gomodules.xyz/jsonpatch/v2 => gomodules.xyz/jsonpatch/v2 v2.0.1
