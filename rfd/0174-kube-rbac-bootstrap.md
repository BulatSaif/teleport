---
authors: Anton Miniailo (anton@goteleport.com)
state: draft
---

# RFD 174 - Kubernetes RBAC Bootstrap

## Required Approvers

- Engineering: `@rosstimothy` && (`@tigrato` || `@hugoShaka`)
- Product: `@klizhentas` || `@xinding33`

## What

Implement the capability to bootstrap and maintain RBAC resources (roles/role bindings) for Kubernetes clusters.

## Why

It will allow users who have complex RBAC setup to not rely on third-party resources or manual actions
when making sure their Kubernetes clusters are working correctly with RBAC setup defined in Teleport.

## Scope

This RFD is focused only on Kubernetes resources related to RBAC functionality, other types of resources are out of scope,
though they might be added in the future.

## Details

We will add a new type of resources, called "KubeProvision", where user will be able to specify Kubernetes RBAC resources they want to
be provisioned onto their Kubernetes clusters.
clusters. A single Teleport resource of that type can define multiple Kubernetes RBAC resources.

```protobuf
import "teleport/header/v1/metadata.proto";

// KubeProvision represents a Kubernetes resources that can be provisioned on the Kubernetes clusters.
// This includes roles/role bindings and cluster roles/cluster role bindings.
// For rationale behind this type, see the RFD 174.
message KubeProvision {
  // The kind of resource represented.
  string kind = 1;
  // Not populated for this resource type.
  string sub_kind = 2;
  // The version of the resource being represented.
  string version = 3;
  // Common metadata that all resources share.
  teleport.header.v1.Metadata metadata = 4;
  // The specific properties of kube provision.
  KubeProvisionSpec spec = 5;
}

// KubeProvisionSpec is the spec for the kube provision message.
message KubeProvisionSpec {
  // resources_data is base64 encoded YAML definitions of the Kubernetes resources.
  string resources_data = 3;
}
```

Every five minutes we will run a reconciliation loop that will compare the current state on Kubernetes clusters under
Teleport management with the desired state and update the cluster's RBAC if there's a difference.

Teleport will mark RBAC resources under its control by "app.kubernetes.io/managed-by: Teleport" label. That way we will
separate resources managed by Teleport and managed by the user manually.

Resource labels will be taken into account when doing the reconciliation - users will be able to match different
Kubernetes clusters for different KubeProvision resources. If a KubeProvision resource doesn't have any labels defined it
will not match any clusters, effectively being disabled.

### Permissions

Not directly related to goals of this RFD, but we will change the default permissions created when enrolling Kubernetes cluster into Teleport.
We will expose default
Kubernetes user-facing roles: cluster-admin, view, edit. When installing kube-agent/creating service account we will
create cluster role bindings for those roles, accordingly linking them to groups "default-cluster-admin", "default-view" and 
"default-edit". It will give users an opportunity to have standard Kubernetes permission groups which they can use together
with Teleport's fine-grained RBAC definitions for Kubernetes. We will add options to the teleport-kube-agent chart and our
kubeconfig generation script to not create those cluster role bindings, but it will be on by default. This change might be 
implemented separately from the implementation of this RFD.

In order to be able to reconcile resources, Teleport will need to perform CRUD operations with Kubernetes RBAC resources.
When performing CRUD operations, Teleport will try to impersonate the "default-cluster-admin" group, if that fails, it will
try to impersonate the "system:masters" group - this will allow us to have backward compatibility with Kubernetes clusters 
which were enrolled before we started exposing default user-facing roles, since "system:masters" group is available on 
majority of installations.

## Security

The introduction of this functionality does not directly impact security of Teleport
itself, however it introduces a new vector of attack for a malicious actor.
Teleport user with sufficient permissions to create/edit KubeProvision resources
will be able to amend RBAC setup on all Kubernetes clusters enrolled in Teleport.
We will emphasize in the user documentation the need to be vigilant when giving out
the necessary permissions.

Even though we will always run the reconciliation loop, by default it will be a noop, since three will
be no resources to provision, so users need to explicitly create KubeProvision resources to start actively using this new feature.
We also will explicitly require labels to be present for the resource to be provisioned to the Kubernetes cluster, making it 
harder to accidentally misuse the feature.

## Alternative

Alternatively, we could take a bit of a different approach regarding permissions and default roles exposure.
We could directly add permissions required to perform CRUD operations on Kubernetes RBAC resources to 
Teleport kube agent/service account credentials on enrollment. Teleport will only try to perform reconciliation using
its kube agent/service account credentials, without trying to impersonate "system:masters", that means that only newly enrolled 
Kubernetes clusters, or cluster where user performed manual upgrade of permissions will be in scope of provisioning the resources.

We will also ship Teleport with default KubeProvision 
resource that defines role binding for the default user facing roles for edit and view. This default resource will not have any
labels defined, so it will not by provisioned anywhere by default. If users want to enable the exposure of default view/edit roles,
they would need to set labels on that resource. They can then use these roles with the fine-grained RBAC definitions for Kubernetes
in Teleport roles. This more conservative scenario will require more explicit decision-making from the user. 
To expose default Kubernetes user-facing roles, users would need to add labels to the default resource we provide,
and KubeProvisioning will only be active for clusters where Teleport has the required permissions added to its credentials.

## Audit

No changes in the audit events will be required.

## Test plan

We will use integration tests to verify reconciliation functionality.