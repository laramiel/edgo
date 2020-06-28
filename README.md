# edgo

Elite: Dangerous golang libraries

## Summary

This is very much a work in progress, but in a nutshell this repo does
the following

* Watches for elite dangerous journal events
* Starts an http server to serve static and dynamic content


## Building

Much of the development has been done on WSL.
Build like this:


```
go get github.com/zserge/lorca
go get github.com/fsnotify/fsnotify

go run gen.go
GOOS=windows GOARCH=amd64 go build -ldflags "-H windowsgui"
```

## Example

Currently the service serves a variant of the elite dangerous 3d
galaxy map from <https://github.com/gbiobob/ED3D-Galaxy-Map>

Every time the NavRoute updates the Ed3d.reset() method is called
to render a new path.

## TODO

* Get examples of every event type and validate that the json parses.
* Parse eddb datasets
* Maybe add database capabilities
* Better gui apps

