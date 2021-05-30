# WIP: Go for Programmers @ Spartakiade 2021

* [https://spartakiade.org/](https://spartakiade.org/)

Date and location: 2021-06-11, 09-17 (09-12, 13-15, 15:15-17), remote workshop;
language: de/en (depending on participants; slides and material will be in
English)

## Abstract

Go is eine moderne Programmiersprache, die - ein wenig wie Englisch - leicht zu
lernen, aber nicht ganz so leicht zu meistern ist. Dieser Workshop stellt Go
für Entwicklerinnen und Entwickler vor: wir führen die Sprache ein und
konzentrieren uns dann auf verschiedene Aspekte, die Go ausmachen: Design and
Projektstruktur, Concurrency und HTTP Applikationen. Wir schauen uns viele
Beispiele an und schreiben selbst einige kleine Programme.

Voraussetzung: Laptop, am besten mit einer [Go
Installation](https://golang.org/doc/install); Editor und
[git](https://git-scm.com/) für das Clonen des Repos.

Für wen ist der Workshop?

> Alle, die schon programmieren, aber noch wenig oder gar keine Erfahrung mit
> Go haben und die Sprache in wenigen Stunden von vielen Seiten beleuchtet
> sehen möchten: mit echten Code-Beispielen, typischen Mustern und
> Lösungsansätzen.

----

Go is a modern programming language, which - a bit like English - is easy to
learn, but hard to master. This workshop is aimed at developers wanting to
take a look at Go: We introduce the language, and focus on topics specific to
Go, like program design, project layout, concurrecy, testing and HTTP
applications. We will study example code and write small programs ourselves.

Requirements: Laptop (or any computer) with [Go
installed](https://golang.org/doc/install), some editor and
[git](https://git-scm.com/).

Who is this workshop for?

> Anyone, who already programs and would like to learn more about what makes Go
> work, from language features to project patterns and real-world examples.

----

## Outline

* [Language Overview](1-language)

> The basic data types, control structures and program structure.

* [Tools](2-tools)

> Test, build, lint and run Go applications with the Go tool.

* [Program design](3-program-design)

> An entry point are the data structures, custom interfaces emerge along design
and development; keeping interfaces small helps composability.

* [Project layout](4-projects)

> Where does code live, how do you keep track of dependencies, how to structure
services or command line tools.

* [Deployment options](5-deployments)

> What's in the binary and how do you run this thing?

* [Concurrency](6-concurrency)

> Go supports classic concurrency primitives, like locks, but also channels to
communicate between threads of execution. Concurrency is often wrapped inside a
library and keeps most code sequential.

* [Testing](7-testing)

> Simple test are table driven. The Go tool supports testing, coverage and benchmarking.

* [HTTP Applications](8-web-apps)

> Building lean web applications with `net/http` and friends.

* [Misc](9-misc)

> Other topics, like: What are people building with Go.
