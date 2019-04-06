# yreq

[![Go Report Card](https://goreportcard.com/badge/github.com/micnncim/yreq)](https://goreportcard.com/report/github.com/micnncim/yreq)
[![Maintainability](https://api.codeclimate.com/v1/badges/f9b2c8472895f4a17a10/maintainability)](https://codeclimate.com/github/micnncim/yreq/maintainability)
[![CodeFactor](https://www.codefactor.io/repository/github/micnncim/yreq/badge)](https://www.codefactor.io/repository/github/micnncim/yreq)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/96cd2c8ffe014e1aab53d33690797a26)](https://www.codacy.com/app/micnncim/yreq?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=micnncim/yreq&amp;utm_campaign=Badge_Grade)
[![MIT](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)

`yreq` is a tool sending multiple HTTP POST requests by `JSON` using `YAML` file.

## Usage

```
$ yreq -req example.yaml -url <ENDPOINT>
```

## Installation

```
$ go get github.com/micnncim/yreq/cmd/yreq
```

## Example

Both of requests and responses are written by `YAML`.

```yaml
- name: Alice
  email: alice@mail.com
  task:
    - title: Walking
      desc: Daily work out
    - title: Study Japanese
- name: Bob
  email: bob@mail.net
  task:
    - title: Swiming
      desc: Daily work out

```

It is equivalent to the following `JSON`s.

```json
{
  "email": "alice@mail.com",
  "name": "Alice",
  "task": [
    {
      "desc": "Daily work out",
      "title": "Walking"
    },
    {
      "title": "Study Japanese"
    }
  ]
}
```

```json
{
  "email": "bob@mail.net",
  "name": "Bob",
  "task": [
    {
      "desc": "Daily work out",
      "title": "Swimming"
    }
  ]
}
```

## LICENSE

[MIT](./LICENSE)
