# yreq

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
