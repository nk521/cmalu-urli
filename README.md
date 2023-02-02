# Cmalu Urli

A small link shortner written in GoLang.

Use `curl -d '{"original_url": "URL"}' c.raju.dev/cmalu` to recieve a short link which expires in 24 hours.

To run it using docker, just do `docker compose up -d`. To run it on local, run `go mod download && go run .`