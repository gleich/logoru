<!-- DO NOT REMOVE - contributor_list:data:start:["Matt-Gleich", "ImgBotApp", "dependabot-preview[bot]"]:end -->

# logoru

🌲 golang port of [Delgan's python loguru](https://github.com/Delgan/loguru)

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
func Debug(msg interface{})
```

Output a debugging message

#### Example

```go
package main


import "github.com/Matt-Gleich/logoru"

func main() {
    logoru.Debug("Here is a debug message")
    logoru.Debug("Hello", "World")
}
```

Output:

<img src="./docs/images/debug_example.png" width="500">

### `func Info`

```go
func Info(msg interface{})
```

Output an info message

#### Example

```go
package main


import "github.com/Matt-Gleich/logoru"

func main() {
    logoru.Info("Here is an info message")
    logoru.Info("Hello", "World")
}
```

Output:

<img src="./docs/images/info_example.png" width="500">

### `func Success`

```go
func Success(msg interface{})
```

Output a success message

#### Example

```go
package main


import "github.com/Matt-Gleich/logoru"

func main() {
    logoru.Success("Here is a success message")
    logoru.Success("Hello", "World")
}
```

Output:

<img src="./docs/images/success_example.png" width="500">

### `func Warning`

```go
func Warning(msg interface{})
```

Output a warning message

#### Example

```go
package main


import "github.com/Matt-Gleich/logoru"

func main() {
    logoru.Warning("Here is a warning message")
    logoru.Warning("Hello", "World")
}
```

Output:

<img src="./docs/images/warning_example.png" width="500">

### `func Error`

```go
func Error(msg interface{})
```

Output a error message

#### Example

```go
package main


import "github.com/Matt-Gleich/logoru"

func main() {
    logoru.Error("Here is an error message")
    logoru.Error("Hello", "World")
}
```

Output:

<img src="./docs/images/error_example.png" width="500">

### `func Critical`

```go
func Critical(msg interface{})
```

Output a critical message

#### Example

```go
package main


import "github.com/Matt-Gleich/logoru"

func main() {
    logoru.Critical("Here is a critical message")
    logoru.Critical("Hello", "World")
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

- **[@dependabot-preview[bot]](https://github.com/apps/dependabot-preview)**

<!-- DO NOT REMOVE - contributor_list:end -->
