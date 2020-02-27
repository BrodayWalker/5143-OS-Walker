## Shell Project
#### Operating Systems Spring 2020

#### Group Members

| Name                                             | Email                     | Github Username |
| ------------------------------------------------ | ------------------------- | --------------- |
| [Jeremy Glebe](https://github.com/jeremyglebe)   | jeremyglebe@gmail.com     | jeremyglebe     |
| [Corbin Matamoros](https://github.com/Pirhomega) | corbinmatamoros@yahoo.com | Pirhomega       |
| [Broday Walker](https://github.com/BrodayWalker) | brodaywalker@gmail.com    | BrodayWalker    |

#### Project Files

| Command     | FileName                              | Author | Works |
| ----------- | ------------------------------------- | ------ | ----- |
|             | [README.md](README.md)                | Broday | :100: |
| cat         | [cat.go](gosh/cat.go)                 | Corbin | :100: |
| cd          | [cd.go](gosh/cd.go)                   | Broday | :100: |
| chmod       | [chmod.go](gosh/chmod.go)             | Corbin | :x:   |
| cp          | [cp.go](gosh/cp.go)                   | Corbin | :100: |
| echo        | [echo.go](gosh/echo.go)               | Broday | :100: |
| exclamation | [exclamation.go](gosh/exclamation.go) | Corbin | :100: |
|             | [gosh.go](gosh/gosh.go)               | All    | :100: |
| head        | [head.go](gosh/head.go)               | Broday | :100: |
| history     | [history.go](gosh/history.go)         | Corbin | :100: |
| ls          | [ls.go](gosh/ls.go)                   | Broday | :100: |
| mkdir       | [mkdir.go](gosh/mkdir.go)             | Broday | :100: |
| nix_mkdir   | [nix_mkdir.go](gosh/nix_mkdir.go)     | Jeremy | :100: |
| nix_mv      | [nix_mv.go](gosh/nix_mv.go)           | Jeremy | :100: |
| pipe        | [pipe.go](gosh/pipe.go)               | Jeremy | :100: |
| pwd         | [pwd.go](gosh/pwd.go)                 | Broday | :100: |
| rm          | [rm.go](gosh/rm.go)                   | Broday | :100: |
| split       | [split.go](gosh/split.go)             | Corbin | 99%   |
| tail        | [tail.go](gosh/tail.go)               | Broday | :100: |
| touch       | [touch.go](gosh/touch.go)             | Broday | :100: |
|             | [utils.go](gosh/utils.go)             | Broday | :100: |
| wc          | [wc.go](gosh/wc.go)                   | Jeremy | :100: |

#### Directory Structure

```
/5143-gosh
├── README.md
├── gosh
│   ├── cat.go
│   ├── cd.go
│   ├── chmod.go
│   ├── echo.go
│   ├── exclamation.go
│   ├── gosh.go
│   ├── head.go
│   ├── history.go
│   ├── ls.go
│   ├── mkdir.go
│   ├── mv.go
│   ├── nix_mkdir.go
│   ├── nix_mv.go
│   ├── pipe.go
│   ├── pwd.go
│   ├── rm.go
│   ├── split.go
│   ├── tail.go
│   ├── touch.go
│   ├── utils.go
│   ├── wc.go
```

## How gosh Works
Currently, the ``gosh`` folder contains all ``.go`` files. ``gosh.go`` contains the ``main`` function, which is used to accept instructions and call the appropriate ``<command>.go`` file. <br>

## Naming Conventions
``gosh`` uses the standard Go naming conventions. Each supported command is contained in an appropriately named ``.go`` file. For example, the command ``mkdir`` is implemented in the ``mkdir.go`` file. <br>
 All exported commands should be capitalized and commented appropriately. <br>

Each ``.go`` file that contains executable code will have the ``package main`` declaration at the top of the file. <br>
#### Example ``.go`` File
```go
package main

import (
	"fmt"
)

// NamePrinter This function prints a name.
func  NamePrinter() {
	name := "Frank Sinatra"
	fmt.Printf("Hello, %s.\n", name)
}
```

The function ``NamePrinter`` will be visible and usable in other ``.go`` files in the folder without having to import anything.
## Building gosh
If you are building `gosh` for the first time, download the following prerequisite package using the command line: <br> <br>
`go get github.com/buger/goterm` <br> <br>
Build ``gosh`` from the ``gosh`` folder  with ``go build .`` The result will be an executable called ``gosh.exe`` .

## Running gosh
Windows: Run ``gosh`` with ``./gosh.exe`` <br> <br>
Linux: Run ``gosh`` with ``./gosh``
