# DiscordGo selfbot fork

[![Go Reference](https://pkg.go.dev/badge/github.com/lltr/selfdiscordgo.svg)](https://pkg.go.dev/github.com/lltr/selfdiscordgo) [![Go Report Card](https://goreportcard.com/badge/github.com/lltr/selfdiscordgo)](https://goreportcard.com/report/github.com/lltr/selfdiscordgo)

This fork aims to provide Discord client capabilities for use by selfbots. Currently receving messages works properly.

## Getting Started

### Installing

This assumes you already have a working Go environment, if not please see
[this page](https://golang.org/doc/install) first.

`go get` *will always pull the latest tagged release from the master branch.*

```sh
go get github.com/lltr/selfdiscordgo
```

### Usage

Import the package into your project.

```go
import "github.com/lltr/selfdiscordgo"
```

Construct a new Discord client which can be used to access the variety of 
Discord API functions and to set callback functions for Discord events.

```go
discord, err := discordgo.New("User authentication token")
```

See Documentation and Examples below for more detailed information.


## Documentation

**NOTICE**: This library and the Discord API are unfinished.
Because of that there may be major changes to library in the future.

The DiscordGo code is fairly well documented at this point and is currently
the only documentation available. Go reference (below) presents that information in a nice format.

- [![Go Reference](https://pkg.go.dev/badge/github.com/bwmarrin/discordgo.svg)](https://pkg.go.dev/github.com/lltr/selfdiscordgo) 
- Hand crafted documentation coming eventually.


## Examples

Below is a list of examples and other projects using DiscordGo. No examples for this selfbot fork.

- [DiscordGo Examples](https://github.com/bwmarrin/discordgo/tree/master/examples) - A collection of example programs written with DiscordGo
- [Awesome DiscordGo](https://github.com/bwmarrin/discordgo/wiki/Awesome-DiscordGo) - A curated list of high quality projects using DiscordGo

## Troubleshooting
For help with common problems please reference the 
[Troubleshooting](https://github.com/bwmarrin/discordgo/wiki/Troubleshooting) 
section of the project wiki.


## Contributing
Contributions are very welcomed, however please follow the below guidelines.

- First open an issue describing the bug or enhancement so it can be
discussed.  
- Try to match current naming conventions as closely as possible.  
- This package is intended to be a low level direct mapping of the Discord API, 
so please avoid adding enhancements outside of that scope without first 
discussing it.
- Create a Pull Request with your changes against the master branch.