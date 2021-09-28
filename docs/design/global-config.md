# Global configuration

The feature of global configuration is to expose the parameters of distributed applications in the entire cluster mirror. 
It is highly recommended to expose only a few parameters that users need to care about.

If too many parameters need to be exposed, for example, the entire helm's values ​​want to be exposed, 
then it is recommended to build a new image and put the configuration in to overwrite it.

Using dashboard as an example, we made a cluster mirror of dashboard, 
but different users want to use different port numbers while installing. 
In this scenario, sealer provides a way to expose this port number parameter to the environment variable of Clusterfile. 

Use global configuration capabilities
For the image builder, this parameter needs to be extracted when making the image. 
Take the yaml of the dashboard as an example:

dashboard.yaml:

```yaml
...
kind: Service
apiVersion: v1
metadata:
  labels:
    k8s-app: kubernetes-dashboard
  name: kubernetes-dashboard
  namespace: kubernetes-dashboard
spec:
  ports:
    - port: 443
      targetPort: {{ DashBoardPort }}
  selector:
    k8s-app: kubernetes-dashboard
...
```

To write kubefile, you need to copy yaml to the manifests directory at this time, 
sealer only renders the files in this directory:

```yaml
FROM kubernetes:1.16.9
COPY dashobard.yaml manifests/
CMD kubectl apply -f manifests/dashobard.yaml
```

For users, they only need to specify the cluster environment variables:

```shell script
sealer run -e DashBoardPort=8443 mydashboard:latest
```
Or specify in Clusterfile:

```yaml
apiVersion: sealer.aliyun.com/v1alpha1
kind: Cluster
metadata:
  name: my-cluster
spec:
  image: mydashobard:latest
  provider: BAREMETAL
  env:
    DashBoardPort: 6443 # Specify a custom port here, which will be rendered into the mirrored yaml
  ssh:
    passwd:
    pk: xxx
...
```

## Use with helm

The sealer will also generate a very complete Clusterfile file to the etc directory when it is running, 
which means that these parameters can be obtained in a certain way in the helm chart.

The chart values ​​of the dashboard can be written like this:

```yaml
spec:
  env:
    DashboardPort: 6443
```
Kubefile:

```shell script
FROM kubernetes:v1.16.9
COPY dashboard-chart.
CMD helm install dashboard dashboard-chart -f etc/global.yaml
```
In this way, the value in global.yaml will override the default port parameter in the dashboard.

## Development Document

Before applying guest, perform template rendering on the files in the manifest directory, 
and render environment variables and annotations to the configuration file 
https://github.com/alibaba/sealer/blob/main/guest/guest.go#L28, 
guest module It is to deal with instructions such as RUN CMD in Kubefile.
Generate the global.yaml file to the etc directory