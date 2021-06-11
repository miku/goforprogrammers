# Project layout

## Notions

A couple notions play a role:

* a package is a directory of code (it will be named after the directory)
* one package can have many files
* one projects can have many subpackages
* there a special `internal` subpackage
* prefer short names, without stutter
* entry points live in package main, if you have many of them, shard them under a directory, typically called `cmd`  

A couple of **non-rules**:

* one struct, one file
* one package, one file
* try to avoid too generic, meaningless names, like `utils.go`, use
  `stringutils.go` or something more speaking
 

## Dependency Management

* use go modules, via `go mod init`

## Walkthrough of a simple package

