# SRE-TakeHome-Test
Site Reliability Engineer Take Home Test

## Explanation

The build process on this repository will have 2 build process, i.e nodejs and golang. CI/CD i use github actions, because it's can free without provide resource and support the step actions process extension (this optional, can use argoCD or jenkins or gitlab). after build process, will uploading to image registry in existing is use AWS ECR (Elastic Container Registry). After push the image, will continuing to deploy process. the deploy process is will retrieve manifest, then will running use kustomization to manage deployment.

makesure in workflow github actions i placed region ap-southeast-3 as asia/jakarta, and should placed the kubernetes cluster name, image registry, kubernetes namespace, put aws access key id and secret key id into github secret.

For monitoring is have multiple choice

- can set notification after build (if failed status only) and after deploy (all status) into discord, slack etc.
- set logging, one of which is loki. Beside loki installed, need an agent. the agent is promtail, promtail will send logs from container into loki
- set metrics monitoring, one of which is prometheus. After installed prometheus, should installed the agent. the agent is node exporter, then node exporter will collect metrics every pod or nodes then send into prometheus
- set dashboard monitoring, one of which is grafana. It's for visualize, create monitoring with prometheus query language to retrieve cpu, memory, disk or 4 golden signal monitoring.

All monitoring service can deployed into kubernetes and officially provide the helm repository.

![monitoring-logging](simple-design-monitoring-logging.png#gh-light-mode-only)

Thanks

