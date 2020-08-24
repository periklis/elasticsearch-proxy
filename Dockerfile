FROM registry.svc.ci.openshift.org/ocp/builder:rhel-8-golang-openshift-4.6 AS builder
WORKDIR  /go/src/github.com/openshift/elasticsearch-proxy
COPY . .
RUN make

FROM registry.svc.ci.openshift.org/ocp/4.6:base
COPY --from=builder /go/src/github.com/openshift/elasticsearch-proxy/bin/elasticsearch-proxy /usr/bin/
ENTRYPOINT ["/usr/bin/elasticsearch-proxy"]
LABEL io.k8s.display-name="OpenShift ElasticSearch Proxy" \
      io.k8s.description="OpenShift ElasticSearch Proxy component of OpenShift Cluster Logging" 
