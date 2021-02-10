# rook&#46;io/v1alpha2

**Note:** This document is generated from code and comments. Do not edit it directly.

## Table of Contents
* [Attachment](#attachment)
* [Volume](#volume)
* [VolumeList](#volumelist)

## Attachment



| Field | Description | Type | Required |
| ----- | ----------- | ---- | -------- |
| node |  | string | true |
| podNamespace |  | string | true |
| podName |  | string | true |
| clusterName |  | string | true |
| mountDir |  | string | true |
| readOnly |  | bool | true |

[Back to TOC](#table-of-contents)

## Volume



| Field | Description | Type | Required |
| ----- | ----------- | ---- | -------- |
| kind | `Volume` | string | false |
| apiVersion | `rook.io/v1alpha2` | string | false |
| metadata |  | [metav1.ObjectMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ObjectMeta) | true |
| attachments |  | [][Attachment](#attachment) | true |

[Back to TOC](#table-of-contents)

## VolumeList



| Field | Description | Type | Required |
| ----- | ----------- | ---- | -------- |
| kind | `VolumeList` | string | false |
| apiVersion | `rook.io/v1alpha2` | string | false |
| metadata |  | [metav1.ListMeta](https://pkg.go.dev/k8s.io/apimachinery/pkg/apis/meta/v1#ListMeta) | true |
| items |  | [][Volume](#volume) | true |

[Back to TOC](#table-of-contents)
