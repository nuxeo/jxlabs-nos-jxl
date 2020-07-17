module github.com/jenkins-x-labs/jxl

go 1.13

require (
	github.com/blang/semver v3.5.1+incompatible
	github.com/go-delve/delve v1.4.0 // indirect
	github.com/jenkins-x-labs/helmboot v0.0.90
	github.com/jenkins-x-labs/jwizard v0.0.15
	github.com/jenkins-x-labs/step-go-releaser v0.0.18
	github.com/jenkins-x-labs/trigger-pipeline v0.0.4
	github.com/jenkins-x/helm-unit-tester v0.0.6
	github.com/jenkins-x/jx/v2 v2.1.97
	github.com/nuxeo/jxlabs-nos-yaml-patch v0.0.3
	github.com/pkg/errors v0.9.1
	github.com/spf13/cobra v1.0.0
	github.com/stretchr/testify v1.4.0
)


replace github.com/jenkins-x-labs/helmboot => github.com/nuxeo/jxlabs-nos-helmboot v0.0.3

replace github.com/jenkins-x-labs/jwizard => github.com/nuxeo/jxlabs-nos-jwizard v0.0.8

replace github.com/jenkins-x-labs/step-go-releaser => github.com/nuxeo/jxlabs-nos-step-go-releaser v0.0.1

replace github.com/jenkins-x-labs/trigger-pipeline => github.com/nuxeo/jxlabs-nos-trigger-pipeline v0.0.9

replace code.gitea.io/sdk => github.com/go-gitea/go-sdk v0.0.0-20180702024448-79a281c4e34a

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0

replace k8s.io/api => k8s.io/api v0.0.0-20190528110122-9ad12a4af326

replace k8s.io/metrics => k8s.io/metrics v0.0.0-20181128195641-3954d62a524d

replace k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190221084156-01f179d85dbc

replace k8s.io/client-go => k8s.io/client-go v0.0.0-20190528110200-4f3abb12cae2

replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190528110544-fa58353d80f3

replace git.apache.org/thrift.git => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999

replace github.com/sirupsen/logrus => github.com/jtnord/logrus v1.4.2-0.20190423161236-606ffcaf8f5d

replace github.com/Azure/azure-sdk-for-go => github.com/Azure/azure-sdk-for-go v21.1.0+incompatible

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v10.14.0+incompatible

replace github.com/banzaicloud/bank-vaults => github.com/banzaicloud/bank-vaults v0.0.0-20190508130850-5673d28c46bd

replace github.com/russross/blackfriday => github.com/russross/blackfriday v1.5.2
