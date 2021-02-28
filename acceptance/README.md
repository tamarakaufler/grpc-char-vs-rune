# Acceptance tests

## TODO

After the correspoding REST service is implemented

## Implementation

Running against the scratch cloud environment:
  - scratch environment must be set up:
        as a scratch cloud cluster with a particular namespace, where the relevant services that the acceptance testing needs, are running

                OR

  - using docker-compose setting up all the relevant services that the acceptance testing needs running
        
Using a library with predefined services tests, that other services can use when implementing their acceptance testing.