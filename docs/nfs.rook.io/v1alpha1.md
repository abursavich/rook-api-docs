# nfs&#46;rook&#46;io/v1alpha1

**Note:** This document is generated from code and comments. Do not edit it directly.

## Table of Contents
* [AllowedClientsSpec](#allowedclientsspec)
* [ExportsSpec](#exportsspec)
* [NFSServer](#nfsserver)
* [NFSServerList](#nfsserverlist)
* [NFSServerSpec](#nfsserverspec)
* [NFSServerState](#nfsserverstate)
* [NFSServerStatus](#nfsserverstatus)
* [ServerSpec](#serverspec)

## AllowedClientsSpec

AllowedClientsSpec represents the client specs for accessing the NFS export

| Field | Description | Type | Required |
| ----- | ----------- | ---- | -------- |
| name | Name of the clients group | string | false |
| clients | The clients that can access the share Values can be hostname, ip address, netgroup, CIDR network address, or all | []string | false |
| accessMode | Reading and Writing permissions for the client to access the NFS export Valid values are \"ReadOnly\", \"ReadWrite\" and \"none\" Gets overridden when ServerSpec.accessMode is specified | string | false |
| squash | Squash options for clients Valid values are \"none\", \"rootid\", \"root\", and \"all\" Gets overridden when ServerSpec.squash is specified | string | false |

[Back to TOC](#table-of-contents)

## ExportsSpec

ExportsSpec represents the spec of NFS exports

| Field | Description | Type | Required |
| ----- | ----------- | ---- | -------- |
| name | Name of the export | string | false |
| server | The NFS server configuration | [ServerSpec](#serverspec) | false |
| persistentVolumeClaim | PVC from which the NFS daemon gets storage for sharing | [corev1.PersistentVolumeClaimVolumeSource](https://pkg.go.dev/k8s.io/api/core/v1#PersistentVolumeClaimVolumeSource) | false |

[Back to TOC](#table-of-contents)

## NFSServer

NFSServer is the Schema for the nfsservers API

| Field | Description | Type | Required |
| ----- | ----------- | ---- | -------- |
| kind | `NFSServer` | string | false |
| apiVersion | `nfs.rook.io/v1alpha1` | string | false |
| metadata |  | [metav1.ObjectMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ObjectMeta) | false |
| spec |  | [NFSServerSpec](#nfsserverspec) | false |
| status |  | [NFSServerStatus](#nfsserverstatus) | false |

[Back to TOC](#table-of-contents)

## NFSServerList

NFSServerList contains a list of NFSServer

| Field | Description | Type | Required |
| ----- | ----------- | ---- | -------- |
| kind | `NFSServerList` | string | false |
| apiVersion | `nfs.rook.io/v1alpha1` | string | false |
| metadata |  | [metav1.ListMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ListMeta) | true |
| items |  | [][NFSServer](#nfsserver) | true |

[Back to TOC](#table-of-contents)

## NFSServerSpec

NFSServerSpec represents the spec of NFS daemon

| Field | Description | Type | Required |
| ----- | ----------- | ---- | -------- |
| annotations | The annotations-related configuration to add/set on each Pod related object. | [rookiov1.Annotations](https://pkg.go.dev/github.com/rook/rook/pkg/apis/rook.io/v1#Annotations) | false |
| replicas | Replicas of the NFS daemon | int | false |
| exports | The parameters to configure the NFS export | [][ExportsSpec](#exportsspec) | false |

[Back to TOC](#table-of-contents)

## NFSServerState



| Name | Value | Description |
| ---- | ----- | ----------- |
| StateInitializing | Initializing |  |
| StatePending | Pending |  |
| StateRunning | Running |  |
| StateError | Error |  |

[Back to TOC](#table-of-contents)

## NFSServerStatus

NFSServerStatus defines the observed state of NFSServer

| Field | Description | Type | Required |
| ----- | ----------- | ---- | -------- |
| state |  | [NFSServerState](#nfsserverstate) | false |
| message |  | string | false |
| reason |  | string | false |

[Back to TOC](#table-of-contents)

## ServerSpec

ServerSpec represents the spec for configuring the NFS server

| Field | Description | Type | Required |
| ----- | ----------- | ---- | -------- |
| accessMode | Reading and Writing permissions on the export Valid values are \"ReadOnly\", \"ReadWrite\" and \"none\" | string | false |
| squash | This prevents the root users connected remotely from having root privileges Valid values are \"none\", \"rootid\", \"root\", and \"all\" | string | false |
| allowedClients | The clients allowed to access the NFS export | [][AllowedClientsSpec](#allowedclientsspec) | false |

[Back to TOC](#table-of-contents)
