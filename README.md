# Cmalu Urli

A small link shortner written in GoLang.

## Run?
To run it using docker, just do `docker compose up -d`. To run it on local, run `go mod download && go run .` (On local the port is 8080)

## How?
Use `curl -d '{"original_url": "URL"}' localhost:42069/cmalu` to recieve a short link which expires in 24 hours.

Example: 
```bash
$ curl -d '{"original_url": "https://google.com"}' localhost:42069/cmalu

http://localhost:42069/c/abcde
```

This link should redirect to the `original_url`.