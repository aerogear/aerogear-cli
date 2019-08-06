# Aerogear CLI

A command line tool to be used as plugin for both kubectl and oc clis to enable easy integration and management of 
aerogear mobile services.

## Requirements

Openshift 3.11+, Kubernetes 1.11 + or kubectl 1.12+

## Usage

Clone the repo and run:

```
make install
```

This will test, build and copy both `kubectl-ag` and `plugin.yaml` files into their required targets.

The CLI then can be use by either kubectl or oc:

```bash
$ kubectl ag ...
$ oc plugin ag ...
```

It can also be directly invoked by using the `ag` binary.

## Development

Running tests:

```bash
#all tests
$ make test

#unit tests
$ make test/unit

#e2e tests
$ make test/e2e
```

## License

This software is licensed under Apache 2.0, see the "LICENSE" file for a full description of its terms.