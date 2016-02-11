docker.io/iameli/stupid-healthcheck
===================================

Sometimes when I'm doing Kubernetes things on AWS, I need a really tiny container that responds to
TCP requests. This is that. It runs an HTTP server on the port specified by the PORT variable and
responds with a string. That's it.
