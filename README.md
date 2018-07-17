# simpleservice - an implementation of a service in the MESG Network

## Check that builds

```
make
```

## Test

```
make test
```

## Service configuration

[mesg.yml](mesg.yml)

## Test results

[RESULTS.md](RESULTS.md)

## Features

- Tasks as subpackages implementing the task interface
- Tasks fully tested
- `ExecuteMany` task implements conditional parallel POST requests
- Dockerfile relies on vendored dependencies: builds faster and more reliably
- UUID generator generates monotonically increasing UUIDs: ExecutionID string order is meaningful
