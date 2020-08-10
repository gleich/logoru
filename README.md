<!-- DO NOT REMOVE - contributor_list:data:start:["Matt-Gleich"]:end -->

# go_loguru

🌲 golang port of Delgan's python loguru

![build](https://github.com/Matt-Gleich/go_loguru/workflows/build/badge.svg)
![test](https://github.com/Matt-Gleich/go_loguru/workflows/test/badge.svg)
![lint](https://github.com/Matt-Gleich/go_loguru/workflows/lint/badge.svg)

## 🚀 Install

Run the following command in your terminal:

```txt
go get -u github.com/Matt-Gleich/go_loguru/loguru
```

## 📝 Documentation [![GoDoc](https://godoc.org/github.com/Matt-Gleich/go_loguru/loguru?status.svg)](https://godoc.org/github.com/Matt-Gleich/go_loguru/loguru)

### `func Debug`

```go
func Debug(msg string) error
```

Output a debugging message

#### Example

```go
package main


import "github.com/Matt-Gleich/go_loguru/loguru"

func main() {
    loguru.Debug("Here is a debug message")
}
```

Output:

<img src="./docs/images/debug_example.png" width="500">

### `func Info`

```go
func Info(msg string) error
```

Output an info message

#### Example

```go
package main


import "github.com/Matt-Gleich/go_loguru/loguru"

func main() {
    loguru.Info("Here is an info message")
}
```

Output:

<img src="./docs/images/info_example.png" width="500">

### `func Success`

```go
func Info(msg string) error
```

Output a success message

#### Example

```go
package main


import "github.com/Matt-Gleich/go_loguru/loguru"

func main() {
    loguru.Success("Here is a success message")
}
```

Output:

<img src="./docs/images/success_example.png" width="500">

### `func Warning`

```go
func Info(msg string) error
```

Output a warning message

#### Example

```go
package main


import "github.com/Matt-Gleich/go_loguru/loguru"

func main() {
    loguru.Warning("Here is a warning message")
}
```

Output:

<img src="./docs/images/warning_example.png" width="500">

### `func Error`

```go
func Error(msg string) error
```

Output a error message

#### Example

```go
package main


import "github.com/Matt-Gleich/go_loguru/loguru"

func main() {
    loguru.Error("Here is an error message")
}
```

Output:

<img src="./docs/images/error_example.png" width="500">

### `func Critical`

```go
func Critical(msg string) error
```

Output a critical message

#### Example

```go
package main


import "github.com/Matt-Gleich/go_loguru/loguru"

func main() {
    loguru.Critical("Here is a critical message")
}
```

Output:

<img src="./docs/images/critical_example.png" width="500">

## 🙌 Contributing

Before contributing please read the [CONTRIBUTING.md file](https://github.com/Matt-Gleich/go_loguru/CONTRIBUTING.md)

<!-- DO NOT REMOVE - contributor_list:start -->

## 👥 Contributors

- **[@Matt-Gleich](https://github.com/Matt-Gleich)**

<!-- DO NOT REMOVE - contributor_list:end -->
