
# lpm

This is LeoConsole Package Manager of 3rd generation: after `pkg` and `apkg` it
is now time for an *even better* package manager.

lpm is written in Go using the [gilc library](https://github.com/alexcoder04/gilc).
It is still a work in progress, however, when it'll be ready, it will have
following advantages over `apkg`:

 - better `apkg-builder` integration: it is compiled directly into the `lpm`
   binary
 - better code quality and faster development cycle: I'm not a pro at C#, I
   love Go much more, so I produce much better code using this language
 - apkg backwards compatibility: you can use the two package managers alongside
   each other: they use the same config files and packaging system!
 - you can compile lpm without installing dotnet (you have to install Go though :)

## Installation

Terminal:

```sh
git clone https://github.com/alexcoder04/lpm.git
```

Then, in LeoConsole

```text
apkg get-local <folder where you cloned lpm to>
```

