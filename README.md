# Hacker News Alternative UI

This project contains an [alternative web UI](https://hughmandeville.github.io/hnui/) for [Hacker News](https://news.ycombinator.com/).

<https://hughmandeville.github.io/hnui/>

It consists of a React App and a Go script that gets the top stories from the [Hacker News API](https://github.com/HackerNews/API) and adds some Open Graph fields (image, icon, title, and description).

## GitHub Pages

The React app is deployed to [GitHub pages](https://hughmandeville.github.io/hnui/).

```sh
make get-hn-ts
```

## Running Locally

Use `make run-client` to start up the React app.

```sh
make run-client
```

Use `make get-hn-ts` to run the Go program that gets the top stories data.

```sh
make get-hn-ts
```
