# Hacker News Cards: An Alternative UI

<https://hncards.com/> is an unofficial alternative web UI for [Hacker News](https://news.ycombinator.com/) that displays stories as cards.

[![Hacker News](hncards-ss.png)](https://hncards.com/)

## Overview

This project consists of a Go script to get the [Hacker News](https://news.ycombinator.com/) top stories and [Open Graph](https://ogp.me/) data and a React App hosted in GitHub Pages to display it.

The Go script gets the top stories from the [Hacker News API](https://github.com/HackerNews/API) and adds some Open Graph fields (image, icon, title, and description). It uploads the data file ([hn_topstories.json](https://storage.googleapis.com/hncards/hn_topstories.json)) to [GCS](https://console.cloud.google.com/storage/browser/_details/hncards/hn_topstories.json?project=hncards).

## GitHub Pages

This GitHub repo has GitHub Pages [setup with the custom domain](https://github.com/hughmandeville/hncards/settings/pages) hncard.com which is managed by [Google Domains](https://domains.google.com/registrar/hncards.com/dns).

Run `make gh-deploy` to deploy the React app to GitHub Pages.

```sh
make gh-deploy
```

## GCS

The `gsutil` command was used to [enable CORS on the GCS bucket](https://cloud.google.com/storage/docs/configuring-cors#gsutil).

```sh
cd pkg/gcs/
gsutil cors set gcs-cors.json gs://hncards
```

## Running Locally

Use `make run-node` to start up the React app on [port 3000](http://localhost:3000/).

```sh
make run-node
```

Use `make hn-ts-get` to run the Go program that gets the Hacker News top stories and Open Graph data.

```sh
make hn-ts-get
```

## Misc. Links

- [Facebook Sharing Debugger](https://developers.facebook.com/tools/debug/)
- [Google Analytics](https://analytics.google.com/analytics/web/)
- [Hacker News API](https://github.com/HackerNews/API)
- [Material Design Colors](https://material.io/design/color/the-color-system.html)
- [Open Graph](https://ogp.me/)
- [Twitter Card Validator](https://cards-dev.twitter.com/validator)
