character
=========

[![Continuous Integration](https://secure.travis-ci.org/philpennock/character.svg?branch=master)](http://travis-ci.org/philpennock/character)

This is a tool for various manipulations on characters, as characters rather
than full strings, to show names, encodings and more.

The tool is structured as a top-level command, options, and sub-commands which
can do different things.  Many sub-commands will take a `-v` verbose option,
which gives more detail in a pretty-printed table.

```console
$ character help
[... lists all available sub-commands ...]
$ character version
[...]
$ character name ✓
CHECK MARK
$ character named -h
[...]
$ character named 'CHECK MARK'
✓
$ character named -j CHECK MARK
✓
$ character named -v/ check
[... table of results of substring search ...]
$ character browse -b 'Alchemical Symbols'
[... table of results; browse is always a table ...]
$ character transform fraktur Hello world
ℌ𝔢𝔩𝔩𝔬 𝔴𝔬𝔯𝔩𝔡
$ character named -1c 'INFORMATION DESK PERSON' \
    'EMOJI MODIFIER FITZPATRICK TYPE-5'
💁🏾
```

In the last example, note that `-c` copies to clipboard; using `-vc` shows the
results in a table but copies only the characters to the clipboard.  Without
`--oneline` (`-1`) each non-verbose character is shown on its own line.  In
this example we're using an emoji modifier which needs to immediately follow
the modified character, so `-1c` _should_ show you the same thing that is
copied to the clipboard where `-c` on its own would show you the individual
parts while copying the modified/combined whole to the clipboard.

[Licensed](./LICENSE.txt) under a MIT-style license.  
[Accumulated licenses of all dependencies](./LICENSES_all.txt) are available
too.  
Patches welcome.


Building
--------

This code should be `go get` compatible, if you don't care about versions of
library dependencies which have been confirmed to work.  If you encounter
problems, please try one of the vendored library mechanisms listed below
instead.

So `go get github.com/philpennock/character` should download and install the
tool into the `bin` dir inside the first directory listed in `$GOPATH`, which
defaults to being `$HOME/go`.

With no Go environment variables set, that go command should thus give you an
output executable at `$HOME/go/bin/character`.


### GOPATH avoidance

At present, we require that the code be checked out "correctly", per Go, which
probably means (if you've not set Go-specific environment variables, etc) that
it should be checked out into `~/go/src/github.com/philpennock/character/`.

We had some shuffle logic, I've removed it.  If this affects you, please open
an issue.

My intention is to move to Go's new workspace system, when that's published,
which gets rid of the need for `GOPATH`, as this is an end tool.


### Dependencies and vendoring.

The dependent code is not checked into this repository, for various
philosophical reasons which I hold to be good.  If one of the dependencies
disappears, please file an Issue and I'll take care of managing the
dependency.

This is Golang, the binaries shouldn't be built live from source in
Production, it's acceptable for the disappearance of a project to make the
code briefly harder to compile.  Reproducible builds are only guaranteed for
official non-git releases, which _will_ include all dependencies.

Specifics available below, after the Building sub-section


### Building

Two simple options, in preference order:

1. Run `make` (or `gmake`) -- using GNU make, build with any extra steps such
   as embedding version identifiers in the resulting binary.
2. Run `go build` -- idiomatic Go, should always work.

You're done.


### go get

This is the "roll the dice and gamble" approach, which _usually_ works, but
not always.  Just fetch the current dependencies, whatever version they might
be, and build.

```console
$ go get -v
```

And then build.


### godep

`dep` is a vendor-locking tool which uses human directives in `Gopkg.toml` to
let the maintainer update known versions in the tooling file `Gopkg.lock`,
which is used to check out content and make it available in the `vendor`
directory.

Note that `dep ensure` checks out the repositories in a common area,
`~/go/pkg/dep/sources` (or elsewhere with `$GOPATH` set) and copies in the
specific files for this revision.  Thus there will be no git/mercurial
repository data within the vendor tree.

```console
$ dep ensure
$ make || go build
```


### Another dependency version manager

I'm happy to supply dependency versioning lock-files for more tools, as
long as they don't conflict with any other tool, and they don't pull code
maintained by others into this repository.

I don't want code of multiple copyrights and licenses in one git repository.
I don't want someone else's misbehaviour coming to light to force me to
rewrite my git history to remove copyright violating code: that's worse for
provenance than having to hunt around for another clone of a dependency which
does at least match the available checksums.


Table packages
--------------

Rendering content to tables requires a table package.  We default to my own
package, `go.pennock.tech/tabular`.  We originally used
`github.com/apcera/termtables` and briefly tried
`github.com/olekukonko/tablewriter` before switching to writing my own.

You can use Go build tags to switch the table package used.  I might remove
support for this in the future.


Optional Tools
--------------

If you're just installing, you don't need these.  If you're developing, you
might.

* `go get -v github.com/kisielk/godepgraph`  
   Dependency graphing tool
* `go get -v github.com/golang/lint/golint`  
  Linter for Go

