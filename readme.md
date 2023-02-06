# goenv

## Install

```bash
go get -u github.com/tsivinsky/goenv
```

## Example

```go
package main

import (
    "github.com/tsivinsky/goenv"
)

type Env struct {
    APP_NAME string `env:"APP_NAME"`
}

func main() {
    env := new(Env)

    goenv.Load(env)
}
```

Under the hood, `goenv` use [godotenv](https://github.com/joho/godotenv) for loading variables from .env file

## Rules

### `required`

```go
type Env struct {
    VAR string `env:"VAR,required"`
}
```

This will result in runtime error if VAR env doesn't exist

### `default=`

```go
type Env struct {
    VAR string `env:"VAR,default=var`
}
```

This will set value for `Env.VAR` as `var` if no other value is provided (including empty string)
