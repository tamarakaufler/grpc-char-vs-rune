# Acceptance tests

## TODO

After the correspoding REST service is implemented

## Implementation

 Acceptance tests run against the scratch cloud environment:

 ```
  a) scratch environment must be set up:
        as a scratch cloud cluster with a particular namespace, where the relevant services that the
        acceptance testing needs, are running

                OR

  b) using docker-compose setting up all the relevant services that the acceptance testing needs running
 ```

They leverage a library with predefined tests related to various services, that other services can use
when implementing their acceptance testing.

Acceptance tests are compiled, so they can run as a binary and a docker image can be created, that other
services, depending on this service, can use as part of their acceptance test suite run.
