# go-organya

This module is a Go implementation of the [Organya music format][organya-specs].
It allows you to parse an Organya file and view its data.

## Getting Started

**Get the module:**

```
go get pkg.nimblebun.works/go-organya
```

**Import it:**

```go
import (
  "pkg.nimblebun.works/go-organya"
)
```

**Create an Organya object:**

```go
org, err := organya.Open("/path/to/file.org")
if err != nil {
  panic(err)
}

// Do something with `org`
```

**Create a playback session:**

```go
session := org.NewSession()

// ...
session.Click()
```

**Convert your Organya data into JSON:**

```go
data, err := org.JSON()
if err != nil {
  panic(err)
}

err = ioutil.WriteFile("organya.json", data, 0644)
// ...
```

## License and Attributions

The project is licensed under the [MIT][license] license. Huge thanks to the
C library this project is based on - [COrg][corg-link].

[organya-specs]: https://gist.github.com/fdeitylink/7fc9ddcc54b33971e5f505c8da2cfd28
[license]: https://github.com/nimblebun/tsc-language-server/blob/master/LICENSE
[corg-link]: https://github.com/DpEpsilon/COrg
