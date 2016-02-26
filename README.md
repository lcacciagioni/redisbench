# redisbench
A POC of a stress tester for redis in golang

## TODO

* [ ] test
* [x]cluster aware
  * I have tested only locally with docker-compose so if any of you can test it a little more is welcome.
* [ ] bynary data maybe using `[]byte('str')` in place of a string
* [ ] more information to be displayed
* [x]optimizations to randomization of strings
  * God save [stackoverflow](https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang/)

> All PR are welcome is nice to know that this can be useful for someone else.
