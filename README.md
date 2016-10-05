# kube-hostname-wrapper

## Description

This script is necessary because kubernetes doesn't set the pod hostname to a hostname that's resolvable outside of the pod.  

Kubernetes does however assign pod hostnames that are resolvable from others pods, but the algorithm is simple enough.

## Supported flags

* -f: outputs fqdn hostname.  example: 192-168-0-3.test.pod.cluster.local
* -s: outputs short hostname. example: 192-168-0-3
* -i: outputs ip address. example: 192.168.0.3
* default: outputs fqdn hostname unless KUBERNETES_USE_SHORTNAME=true
