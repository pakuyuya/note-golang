## instraction

```
go get github.com/urfave/cli
```

### command log

```
PS E:\_work\note-golang\urfave_cli> go run src/example/main.go pre-flag -s qqq
pre-flag
-s
qqq
PS E:\_work\note-golang\urfave_cli> go run src/example/main.go pre-flag -s
pre-flag
-s
PS E:\_work\note-golang\urfave_cli> go run src/example/main.go -s
Incorrect Usage. flag needs an argument: -s

NAME:
   example - this is sample.
  try some option.

USAGE:
   main.exe [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --str value, -s value  input some text.
   --help, -h             show help
   --version, -v          print the version
PS E:\_work\note-golang\urfave_cli>
```

オプション先に指定しないと動かない。なんかいけてない
