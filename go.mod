module github.com/abursavich/rook-api-docs

go 1.14

require (
	github.com/machinezone/configmapsecrets v0.4.1-0.20210210223234-9b74e2147d58
	github.com/rook/rook v1.5.6
	golang.org/x/sync v0.0.0-20201207232520-09787c993a3a
	k8s.io/apimachinery v0.20.2
)

replace (
	k8s.io/client-go => k8s.io/client-go v0.20.2
	k8s.io/kubernetes => k8s.io/kubernetes v0.20.2
)

exclude (
	k8s.io/client-go v1.4.0
	k8s.io/client-go v1.5.0
	k8s.io/client-go v1.5.1
	k8s.io/client-go v10.0.0+incompatible
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/client-go v2.0.0+incompatible
	k8s.io/client-go v3.0.0+incompatible
	k8s.io/client-go v4.0.0+incompatible
	k8s.io/client-go v5.0.0+incompatible
	k8s.io/client-go v5.0.1+incompatible
	k8s.io/client-go v6.0.0+incompatible
	k8s.io/client-go v7.0.0+incompatible
	k8s.io/client-go v8.0.0+incompatible
	k8s.io/client-go v9.0.0+incompatible
)
