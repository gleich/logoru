<!-- DO NOT REMOVE - contributor_list:data:start:["Matt-Gleich", "ImgBotApp"]:end -->

# logoru

🌲 golang port of Delgan's python loguru

![build](https://github.com/Matt-Gleich/logoru/workflows/build/badge.svg)
![test](https://github.com/Matt-Gleich/logoru/workflows/test/badge.svg)
![lint](https://github.com/Matt-Gleich/logoru/workflows/lint/badge.svg)

## 🚀 Install

Run the following command in your terminal:

```txt
go get -u github.com/Matt-Gleich/logoru
```

## 📝 Documentation [![GoDoc](https://godoc.org/github.com/Matt-Gleich/logoru?status.svg)](https://godoc.org/github.com/Matt-Gleich/logoru)

### `func Debug`

```go
func Debug(msg string) error
```

Output a debugging message

#### Example

```go
package main


import "github.com/Matt-Gleich/logoru"

func main() {
    logoru.Debug("Here is a debug message")
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


import "github.com/Matt-Gleich/logoru"

func main() {
    logoru.Info("Here is an info message")
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


import "github.com/Matt-Gleich/logoru"

func main() {
    logoru.Success("Here is a success message")
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


import "github.com/Matt-Gleich/logoru"

func main() {
    logoru.Warning("Here is a warning message")
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


import "github.com/Matt-Gleich/logoru"

func main() {
    logoru.Error("Here is an error message")
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


import "github.com/Matt-Gleich/logoru"

func main() {
    logoru.Critical("Here is a critical message")
}
```

Output:

<img src="./docs/images/critical_example.png" width="500">

## 🙌 Contributing

Before contributing please read the [CONTRIBUTING.md file](https://github.com/Matt-Gleich/logoru/blob/master/CONTRIBUTING.md)

<!-- DO NOT REMOVE - contributor_list:start -->
## 👥 Contributors


- **[@Matt-Gleich](https://github.com/Matt-Gleich)**

- **[@ImgBotApp](https://github.com/ImgBotApp)**

<!-- DO NOT REMOVE - contributor_list:end -->
