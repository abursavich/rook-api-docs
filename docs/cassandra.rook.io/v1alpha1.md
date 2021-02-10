# cassandra&#46;rook&#46;io/v1alpha1

**Note:** This document is generated from code and comments. Do not edit it directly.

## Table of Contents
* [Cluster](#cluster)
* [ClusterList](#clusterlist)
* [ClusterMode](#clustermode)
* [ClusterSpec](#clusterspec)
* [ClusterStatus](#clusterstatus)
* [ConditionStatus](#conditionstatus)
* [DatacenterSpec](#datacenterspec)
* [ImageSpec](#imagespec)
* [RackCondition](#rackcondition)
* [RackConditionType](#rackconditiontype)
* [RackSpec](#rackspec)
* [RackStatus](#rackstatus)

## Cluster



| Field | Description | Type | Required |
| ----- | ----------- | ---- | -------- |
| kind | `Cluster` | string | false |
| apiVersion | `cassandra.rook.io/v1alpha1` | string | false |
| metadata |  | [metav1.ObjectMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ObjectMeta) | true |
| spec |  | [ClusterSpec](#clusterspec) | true |
| status |  | [ClusterStatus](#clusterstatus) | true |

[Back to TOC](#table-of-contents)

## ClusterList



| Field | Description | Type | Required |
| ----- | ----------- | ---- | -------- |
| kind | `ClusterList` | string | false |
| apiVersion | `cassandra.rook.io/v1alpha1` | string | false |
| metadata |  | [metav1.ListMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ListMeta) | true |
| items |  | [][Cluster](#cluster) | true |

[Back to TOC](#table-of-contents)

## ClusterMode



| Name | Value | Description |
| ---- | ----- | ----------- |
| ClusterModeCassandra | cassandra |  |
| ClusterModeScylla | scylla |  |

[Back to TOC](#table-of-contents)

## ClusterSpec

ClusterSpec is the desired state for a Cassandra Cluster.

| Field | Description | Type | Required |
| ----- | ----------- | ---- | -------- |
| annotations | The annotations-related configuration to add/set on each Pod related object. | [rookiov1.Annotations](../rook.io/v1#Annotations) | true |
| version | Version of Cassandra to use. | string | true |
| repository | Repository to pull the image from. | *string | false |
| mode | Mode selects an operating mode. | [ClusterMode](#clustermode) | false |
| datacenter | Datacenter that will make up this cluster. | [DatacenterSpec](#datacenterspec) | true |
| sidecarImage | User-provided image for the sidecar that replaces default. | *[ImageSpec](#imagespec) | false |

[Back to TOC](#table-of-contents)

## ClusterStatus

ClusterStatus is the status of a Cassandra Cluster

| Field | Description | Type | Required |
| ----- | ----------- | ---- | -------- |
| racks |  | map[string]*[RackStatus](#rackstatus) | false |

[Back to TOC](#table-of-contents)

## ConditionStatus



| Name | Value | Description |
| ---- | ----- | ----------- |
| ConditionTrue | True | These are valid condition statuses. \"ConditionTrue\" means a resource is in the condition; \"ConditionFalse\" means a resource is not in the condition; \"ConditionUnknown\" means kubernetes can't decide if a resource is in the condition or not. |
| ConditionFalse | False | These are valid condition statuses. \"ConditionTrue\" means a resource is in the condition; \"ConditionFalse\" means a resource is not in the condition; \"ConditionUnknown\" means kubernetes can't decide if a resource is in the condition or not. |
| ConditionUnknown | Unknown | These are valid condition statuses. \"ConditionTrue\" means a resource is in the condition; \"ConditionFalse\" means a resource is not in the condition; \"ConditionUnknown\" means kubernetes can't decide if a resource is in the condition or not. |

[Back to TOC](#table-of-contents)

## DatacenterSpec

DatacenterSpec is the desired state for a Cassandra Datacenter.

| Field | Description | Type | Required |
| ----- | ----------- | ---- | -------- |
| name | Name of the Cassandra Datacenter. Used in the cassandra-rackdc.properties file. | string | true |
| racks | Racks of the specific Datacenter. | [][RackSpec](#rackspec) | true |

[Back to TOC](#table-of-contents)

## ImageSpec

ImageSpec is the desired state for a container image.

| Field | Description | Type | Required |
| ----- | ----------- | ---- | -------- |
| version | Version of the image. | string | true |
| repository | Repository to pull the image from. | string | true |

[Back to TOC](#table-of-contents)

## RackCondition

RackCondition is an observation about the state of a rack.

| Field | Description | Type | Required |
| ----- | ----------- | ---- | -------- |
| type |  | [RackConditionType](#rackconditiontype) | true |
| status |  | [ConditionStatus](#conditionstatus) | true |

[Back to TOC](#table-of-contents)

## RackConditionType



| Name | Value | Description |
| ---- | ----- | ----------- |
| RackConditionTypeMemberLeaving | MemberLeaving |  |

[Back to TOC](.md#table-of-contents)

## RackSpec

RackSpec is the desired state for a Cassandra Rack.

| Field | Description | Type | Required |
| ----- | ----------- | ---- | -------- |
| name | Name of the Cassandra Rack. Used in the cassandra-rackdc.properties file. | string | true |
| members | Members is the number of Cassandra instances in this rack. | int32 | true |
| configMapName | User-provided ConfigMap applied to the specific statefulset. | *string | false |
| jmxExporterConfigMapName | User-provided ConfigMap for jmx prometheus exporter | *string | false |
| storage | Storage describes the underlying storage that Cassandra will consume. | [rookiov1.StorageScopeSpec](../rook.io/v1.md#StorageScopeSpec) | true |
| annotations | The annotations-related configuration to add/set on each Pod related object. | [rookiov1.Annotations](../rook.io/v1.md#Annotations) | true |
| placement | Placement describes restrictions for the nodes Cassandra is scheduled on. | *[rookiov1.Placement](../rook.io/v1.md#Placement) | false |
| resources | Resources the Cassandra Pods will use. | [corev1.ResourceRequirements](https://pkg.go.dev/k8s.io/api/core/v1#ResourceRequirements) | true |

[Back to TOC](#table-of-contents)

## RackStatus

RackStatus is the status of a Cassandra Rack

| Field | Description | Type | Required |
| ----- | ----------- | ---- | -------- |
| members | Members is the current number of members requested in the specific Rack | int32 | true |
| readyMembers | ReadyMembers is the number of ready members in the specific Rack | int32 | true |
| conditions | Conditions are the latest available observations of a rack's state. | [][RackCondition](#rackcondition) | false |

[Back to TOC](#table-of-contents)
