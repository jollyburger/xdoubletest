##xdoubletest

---

- unit test by docker image which contains test scripts customized 
- performance test 

###Introduction

```
xdoubletest
    |
    |- main.go
    |- app/
        |
        |- config.go //configuration file mapping
    |- examples/     //configure file templates
    |- build/        //binary file
    |- deploy/       //build scripts with cmake and Dockerfile
    |- logic
        |
        |- service/  //http server 
        |- perf/     //performance module
        |- unit/     //unit test module
    |- utils
        |
        |- dockercli //docker client
        |- consul    //consul client
    |- vendor/
```
