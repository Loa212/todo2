# Todo cli

A simple todo list implemented using cobra-cli

## Installation

```bash
go mod tidy
```

( You'll also need an _.env_ file containing the database uri )

## Usage

list the todos

```bash
go run main.go list
```

create a todo

```bash
go run main.go add [title] [?done:boolean]
```

update todo

```bash
go run main.go update [id] [?title:string] [?done:boolean]
```

delete todo

```bash
go run main.go delete [id]
```

## License

[MIT](https://choosealicense.com/licenses/mit/)
