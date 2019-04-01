# yreq

`yreq` is a tool sending multiple HTTP POST requests by JSON using YAML file.

## Usage

```
$ yreq -req example.yaml -url <ENDPOINT>
```

## Example

```yaml
- name:  Hanako
  email: flower@mail.com
- name:  Sumire
  email: garnet@mail.net
```

This means it sends the following requests.

```json
{
  "name": "Hanako",
  "email": "flower@mail.com"
} 
```

```json
{
  "name": "Sumire",
  "email": "garnet@mail.net"
} 
```

## LICENSE

[MIT](./LICENSE)
