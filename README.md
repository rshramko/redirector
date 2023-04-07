# Redirector

A simple tool to help moving web-sites to a new URL.

It starts an http service on a given port, which returns an html page with automatic redirect to a given new address.

### How to build

You are going to need to have the [Go](https://go.dev/) compiler installed. Execute from the root of the repo:

```
$ go build -ldflags "-s -w"
```

### Example of usage:

```
$ ./redirector -port 8085 \
               -title "Nice Service" \
               -wait 15 \
               -url "https://new-address.example.local"
```
