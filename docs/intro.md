---
sidebar_position: 1
---

# 简介 

`Edgewize ModelMesh` 是一款面向 Kubernetes 的推理服务治理框架。`Edgewize ModelMesh` 通过 SQL 感知的流量治理、审计、安全和扩展性等能力实现 [Database Mesh](https://www.database-mesh.io) 风格的推理服务治理体验。

## 概述

`Edgewize ModelMesh` 关注如下几个问题:

* SQL 感知的流量治理：支持 SQL 流量负载均衡、访问控制和可观测性。 
* 运行时资源管理: 支持多种可扩展的资源控制能力 
* 推理服务可靠性工程：简化 Kubernetes 环境下推理服务的治理 

`Edgewize ModelMesh` 的架构图如下：

![`Edgewize ModelMesh` Arch](/img/pisanix-arch.png)

三个组件的功能分别为：

* ***ModelMesh-Controller***: 用 Go 实现的控制面，提供对数据面组件的管控，如 Sidecar 注入、配置转换和下发，是 `Edgewize ModelMesh` 所有配置的入口。

* ***ModelMesh-Proxy***: 用 Go 实现的高性能量代理，通过 MySQL 协议获取应用的推理服务访问流量，并基于此实现 SQL 流量治理、访问控制、防火墙、可观测性等各种治理能力。

* ***ModelMesh-Broker***: 推理服务数据面，部署在集群中每个节点上并通过宿主机内核的各种能力提供可编程资源管理，如 TrafficQoS 等。


## 特性

### 推理服务流量治理 

应用通过 SQL 访问推理服务，`Edgewize ModelMesh` 可以劫持所有的流量。

### 可观测性 

推理服务的监控指标通常从相关实例处获取，借助 `Edgewize ModelMesh` 可以透视多种推理服务访问指标。

### 多模式

`Edgewize ModelMesh` 支持多种方式访问，未来还会支持 `MQTT` 等方式。


## 快速开始 

- [简介](https://www.pisanix.io/docs)
- [快速开始](https://www.pisanix.io/docs/quickstart)

## 文档 

所有文档可以在 [`Edgewize ModelMesh` 站点查看](https://www.pisanix.io/).

## 社区和支持 

|||
|:-|:-|
| 邮件列表| https://groups.google.com/g/database-mesh |
| 英文社区会议(开始于 2022-02-16), 周三 9:00AM PST|https://meet.google.com/yhv-zrby-pyt |
| 中文社区会议 (开始于 2022-04-27), 周三 9:00PM GMT+8|https://meeting.tencent.com/dm/6UXDMNsHBVQO |
| Slack | https://join.slack.com/t/databasemesh/shared_invite/zt-19rhvnxkz-USjZ~am~ghd_Q0q~8bAJXA |
| 会议记录 |https://bit.ly/39Fqt3x |


- 微信交流群: 添加小助手微信邀请进群 

![Wechat user group broker](/img/wechat-user-group-broker.jpeg)
