# Model Mesh

<img src="static/logo.png" alt="Model Mesh Arch" width="500" length="500"/>

Model Mesh is a framework specification for AI model governance in the era of cloud-native AI development. With a focus on abstraction and standardization of multiple AI model governance actions, Model Mesh aims to provide developers with a transparent and streamlined experience in accessing AI models. It supports a consistent governance interface and facilitates various automation tasks, thereby advancing the transition from Infrastructure as Code to Infrastructure Governance as Code within the AI domain.

## Introduction

In the landscape of rapidly evolving AI technologies, the proliferation of AI models across various applications creates intricate dependencies resembling vast interconnected networks. Managing these dependencies presents significant challenges. Different AI-based services demand distinct models, necessitating an efficient approach to manage and govern these models. While Kubernetes simplifies application deployment at scale, and Service Mesh offers solutions for microservice governance in cloud-native architectures, there's currently a lack of a similar solution for AI models.

Model Mesh is proposed as a governance framework tailored to oversee AI models in cloud-native environments. It adheres to the following core principles:
* First-Class Model Resource: All actions derive from the governance of AI models.
* Developer-Centric: Model Mesh abstracts the operational intricacies, allowing developers to focus solely on the model protocol, without worrying about the underlying operational nuances. Simultaneously, it facilitates automation for Site Reliability Engineers (SREs) and Database Administrators (DBAs) in a declarative manner.
* Cloud-Native Compatibility: Embracing an open community approach, Model Mesh facilitates the adoption of AI models from various vendors, promoting cloud-native architecture without vendor lock-in.

The aspirations of Model Mesh include:
* Reducing the developer burden, enabling a sharper focus on business enhancement.
* Establishing a framework that encompasses model traffic, runtime resources, and reliability in a configurable, pluggable, and programmable manner.
* Offering a standardized interface compatible with a variety of AI models, including heterogeneous models, cloud-native models, and distributed models.

### Architecture
![model-mesh-spec-v1](./static/model-mesh-spec-v1.png)

## Specification

Work on the following documents is currently in progress:

|                               |         Latest Release             | 
| :----------------------------: | :--------------------------------: |
| Model Mesh |  [v0.1.0](/SPEC.md) |

## Community


| | |
|:-|:-|
| Mailing List| https://groups.google.com/g/model-mesh |
| Dev Meetings (Starting Feb 16th, 2022), Bi-weekly Wednesday 9:00AM PST|https://meet.google.com/yhv-zrby-pyt |
| Dev Meetings APAC Friendly (Starting April 27th, 2022), Bi-weekly APAC Wednesday 9:00PM GMT+8|https://meeting.tencent.com/dm/6UXDMNsHBVQO |
| Wechat Broker|pisanix|
| Slack |https://join.slack.com/t/modelmesh/shared_invite/zt-19rhvnxkz-USjZ~am~ghd_Q0q~8bAJXA |
| Meetings Notes |https://bit.ly/39Fqt3x |

## Support

Whether you are a user or contributor, feel free to open issues on GitHub:

* [Issues](https://github.com/model-mesh/model-mesh/issues)

## Community Code of Conduct

Model Mesh adheres to the [CNCF Code of Conduct](https://github.com/cncf/foundation/blob/master/code-of-conduct.md).