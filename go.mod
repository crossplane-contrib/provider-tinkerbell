module github.com/tinkerbell/crossplane-provider-tinkerbell

go 1.13

require (
	github.com/crossplane/crossplane-runtime v0.16.1
	github.com/crossplane/crossplane-tools v0.0.0-20201007233256-88b291e145bb
	github.com/google/go-cmp v0.5.6
	github.com/pkg/errors v0.9.1
	github.com/tinkerbell/tink v0.0.0-20201106154020-384fe30ce2a4
	google.golang.org/grpc v1.41.0
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	k8s.io/api v0.23.0
	k8s.io/apimachinery v0.23.0
	sigs.k8s.io/controller-runtime v0.11.0
	sigs.k8s.io/controller-tools v0.8.0
	sigs.k8s.io/yaml v1.3.0
)
