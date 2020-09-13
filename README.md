# embedded-envoy

Uses [s6-overlay](https://github.com/just-containers/s6-overlay) to wrap a go service with [envoy](https://www.envoyproxy.io/docs/envoy/v1.14.3/) proxy. The envoy.yaml is configured to allow retrys when a GET request returns 5xx.

# should I use this?

In a perfect world with infinite time you should stand up envoy as a seperate container and route traffic through it. If you are short on time and do not want to reconfigure your infrastructure right now this will get you by.