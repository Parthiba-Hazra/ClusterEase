
# ClusterEase
ClusterEase is a command-line tool written in Go that provides convenient operations for interacting with a Kubernetes minikube cluster. The tool allows users to obtain information about the cluster, manage resource quotas, scale the cluster, and perform various other cluster-level operations.

## Prerequisites
Make sure you have Go installed on your system and the GOPATH environment variable set.

## Installation
Clone the repository to your local machine:
```bash
git clone https://github.com/Parthiba-Hazra/ClusterEase.git
```
Change directory to the cloned repository:
```bash
cd ClusterEase
````
Build the application:
```bash
go build
````
Install the application:
```bash
go install
```
## Usage
ClusterEase provides the main command clstres for managing your Kubernetes minikube cluster. Here are some of the key functionalities it offers:

### Creating Kubernetes Resources
You can use the create subcommand to create Kubernetes resources from a YAML file.

```bash
clstres create --fp=<yaml_file_path> --ns=<namespace>
--fp: Path to the YAML file containing the resource definition.
--ns: (Optional) Namespace where the resource will be created. If not provided, it will use the default namespace.
```
### Deleting Kubernetes Resources
To delete a Kubernetes resource, use the delete subcommand.

```bash
clstres delete --k=<resource_kind> --rn=<resource_name> --ns=<namespace>
--k: Kind of the resource you want to delete (e.g., Deployment, Service, Pod, etc.).
--rn: Name of the resource to be deleted.
--ns: (Optional) Namespace of the resource. If not provided, it will use the default namespace.
```
### Getting Resource Details
ClusterEase allows you to view specific details of a resource. For example, you can view details of a Deployment, Service, or Pod.

```bash
clstres details deployment -d=<deployment_name> --ns=<namespace>
clstres details service -s=<service_name> --ns=<namespace>
clstres details pod -p=<pod_name> --ns=<namespace>
```
### Namespace Details
You can obtain details about a specific namespace using the namespace subcommand.

```bash
clstres details namespace --ns=<namespace_name>
```
### Running Commands Inside a Pod
ClusterEase provides the podEnter subcommand to run a command inside a pod.

```bash
clstres podEnter --p=<pod_name> --ns=<namespace> --cmd=<command>
--p: Name of the pod where the command will be executed.
--ns: (Optional) Namespace of the pod. If not provided, it will use the default namespace.
--cmd: The command to be executed inside the pod.
```
### Listing All Resources
You can list all resources or specific types of resources using the show subcommand.

```bash
clstres show all --ns=<namespace>
clstres show deploy --ns=<namespace>
clstres show services --ns=<namespace>
clstres show pods --ns=<namespace>
clstres show namespaces
--ns: (Optional) Filter resources by namespace. If not provided, it will show resources from all namespaces.
Remember to use the --ns=<namespace> flag at the root command level to specify the namespace for subsequent commands. This flag will apply to all commands unless explicitly overridden in the subcommands.
```

### Enjoy using ClusterEase to manage your Kubernetes minikube cluster with ease!