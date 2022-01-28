# Hacker News Alternative UI

This project contains an unofficial [alternative web UI](https://hughmandeville.github.io/hnui/) for [Hacker News](https://news.ycombinator.com/).

<https://hughmandeville.github.io/hnui/>

[![Hacker News](hnui-ss.png)](https://hughmandeville.github.io/hnui)

## Overview

This project consists of a Go script to get the [Hacker News](https://news.ycombinator.com/) top stories and [Open Graph](https://ogp.me/) data and a React App hosted in GitHub Pages to display it.

The Go script gets the top stories from the [Hacker News API](https://github.com/HackerNews/API) and adds some Open Graph fields (image, icon, title, and description). It calls the [GitHub API](https://docs.github.com/en/rest) to upload the data file ([hn_topstories.json](client/public/hn_topstories.json)) to GitHub.

## GitHub Pages

Run `make gh-deploy` to deploy the React app to [GitHub Pages](https://hughmandeville.github.io/hnui/).

```sh
make gh-deploy
```

## Running Locally

Use `make run-node` to start up the React app on [port 3000](http://localhost:3000/hnui).

```sh
make run-node
```

Use `make get-hn-ts` to run the Go program that gets the Hacker News top stories and Open Graph data.

```sh
make get-hn-ts
```

## Misc. Links

- [Facebook Sharing Debugger](https://developers.facebook.com/tools/debug/)
- [Hacker News API](https://github.com/HackerNews/API)
- [Material Design Colors](https://material.io/design/color/the-color-system.html)
- [Open Graph](https://ogp.me/)
- [Twitter Card Validator](https://cards-dev.twitter.com/validator)
