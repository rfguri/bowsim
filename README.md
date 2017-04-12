# Bowsim

[![GoDoc][1]][2] [![Build Status][3]][4] [![Coverage Status][5]][6] [![GitHub release][7]][8] [![GitHub license][9]][10]

Bowsim is a library that provides a way to perform string similarity score calculation using [bag of words model][11] in [Go][12].

[1]: https://godoc.org/github.com/rfguri/bowsim?status.svg
[2]: http://godoc.org/github.com/rfguri/bowsim
[3]: https://travis-ci.org/rfguri/bowsim.svg
[4]: https://travis-ci.org/rfguri/bowsim
[5]: https://coveralls.io/repos/rfguri/bowsim/badge.svg?branch=master&service=github
[6]: https://coveralls.io/github/rfguri/bowsim?branch=master
[7]: https://img.shields.io/github/release/rfguri/bowsim.svg
[8]: https://github.com/rfguri/bowsim/releases
[9]: https://img.shields.io/github/license/mashape/apistatus.svg
[10]: https://github.com/rfguri/bowsim/blob/master/LICENSE
[11]: https://en.wikipedia.org/wiki/Bag-of-words_model
[12]: https://golang.org

## Installing

Install [Go][13], either [from source][14] or [with a prepackaged binary][15].
Then,

```Bash
$ go get github.com/rfguri/bowsim
```

[13]: http://golang.org
[14]: http://golang.org/doc/install/source
[15]: http://golang.org/doc/install

## Testing

```Bash
$ go test -bench=. -v
```

## Using

```Go
package main

import (
  "fmt"
  "github.com/rfguri/bowsim"
)

func main() {
  a := "Jane loves me more than Julie loves me"
  b := "Julie likes me more than Jake loves me"

  s := bowsim.Get(a, b)

  fmt.Println(s)
}
```

## Contributing

Pull requests are very much welcomed. Create your pull request on a non-master branch, make sure a test or example is included that covers your change and your commits represent coherent changes that include a reason for the change.

To run the integration tests, make sure to run `go test`. TravisCI will also run the integration tests.

## Credits

* [Bernat Fages][16]

[16]: https://github.com/bernatfp

## License

Copyright © 2017 Roger Fernandez Guri. It is free software, and may be redistributed under the terms specified in the [LICENSE](https://github.com/rfguri/bowsim/blob/master/LICENSE) file.
