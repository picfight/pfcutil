pfcutil
=======

[![Build Status](http://img.shields.io/travis/picfight/pfcutil.svg)]
(https://travis-ci.org/picfight/pfcutil) [![Coverage Status]
(http://img.shields.io/coveralls/picfight/pfcutil.svg)]
(https://coveralls.io/r/picfight/pfcutil?branch=master) [![ISC License]
(http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)]
(http://godoc.org/github.com/picfight/pfcutil)

Package pfcutil provides picfight-specific convenience functions and types.
A comprehensive suite of tests is provided to ensure proper functionality.  See
`test_coverage.txt` for the gocov coverage report.  Alternatively, if you are
running a POSIX OS, you can run the `cov_report.sh` script for a real-time
report.

This package was developed for pfcd, a full-node implementation of PicFight which
is under active development by Company 0.  Although it was primarily written for
pfcd, this package has intentionally been designed so it can be used as a
standalone package for any projects needing the functionality provided.

## Installation and Updating

```bash
$ go get -u github.com/picfight/pfcutil
```

## License

Package pfcutil is licensed under the [copyfree](http://copyfree.org) ISC
License.
