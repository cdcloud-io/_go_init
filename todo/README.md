# TODO

ideas located in repo:

<https://github.com/kubernetes/test-infra/blob/master/Makefile>

## Makefile

- [ ] Integrate dynamic variable assignments.

Example Without Overriding
If you run make without setting REGISTRY, the default value will be used:

```sh
REGISTRY ?= gcr.io/k8s-prow
IMAGE_NAME := my-app
TAG := latest

build:
    docker build -t $(REGISTRY)/$(IMAGE_NAME):$(TAG) .

push:
    docker push $(REGISTRY)/$(IMAGE_NAME):$(TAG)
```

Example with Overriding
If you want to override the default registry, you can set the REGISTRY variable in the environment or on the command line:

```shell
make REGISTRY=my-custom-registry.io
```

- [ ] Decision logic of init.  (in alias / or makefile)
  
  - maybe Makefile init checks for internet connection and pulls an initscript from repo templates?
