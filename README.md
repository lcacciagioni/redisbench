# redisbench

[![Build Status](https://drone.io/github.com/lcacciagioni/redisbench/status.png)](https://drone.io/github.com/lcacciagioni/redisbench/latest)[![codecov.io](https://codecov.io/github/lcacciagioni/redisbench/coverage.svg?branch=master)](https://codecov.io/github/lcacciagioni/redisbench?branch=master)

A POC of a stress tester for redis in golang. Implemented using [redigo](https://github.com/garyburd/redigo) and [redis-go-cluster](https://github.com/chasex/redis-go-cluster).

### Usage
Single node (Default mode)
```bash
$ redisbench -redisNodes=192.168.1.2:6379
```
Cluster
```bash
$ redisbench -redisNodes=127.0.0.1:6379,127.0.0.1:6380,127.0.0.1:6381
```

For more information about options run `$ redisbench -help`

## TODO

The following points are what from my point of view we need to have before having the first beta version.

* [x] test
* [x] cluster aware
* [x] bynary data maybe using `[]byte('str')` in place of a string
* [ ] more information to be displayed - [WIP]
* [x] optimizations to randomization of strings

> God save [stackoverflow](https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang/) for providing the strings randomization just in time.

### Bugs and Improvements
All the contributions are welcome is nice to know that this can be useful for someone else.
