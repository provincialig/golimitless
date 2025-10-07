![License](https://img.shields.io/github/license/provincialig/golimitless)
![Dependency](https://img.shields.io/badge/dependency-0-brightgreen)
![Golangci-Lint version](https://img.shields.io/badge/golangci--lint-2.5.0-brightgreen)
![Test](https://github.com/provincialig/golimitless/actions/workflows/test.yml/badge.svg)

## GoLimitless

A set of tools to solve some GoLang limitations that lead to writing a lot of code.

#### Install
```bash
go get github.com/provincialig/golimitless
```

#### Tools

- **Data structures**
  - **Common**:
    - **SetX**: A thread-safe typed implementation of Set.
    - **MapX**: A thread-safe typed implementation of Map.
    - **Stack**: A thread-safe typed implementation of Stack.
    - **Queue**: A thread-safe typed implementation of Queue.
  - **Extended**:
    - **DoubleMap**: A double layer thread-safe key-value map, with many helpful methods.
    - **ExpireSet**: A thread-safe typed implementation of Set where the elements will removed after retain time.
    - **ISlice**: A thread-safe key-value map where value is a slice, with many helpful methods.

- **Slice utils**
  - **Filter**
  - **Map**
  - **Reduce**
  - **ForEach**
  - **SliceToMap** / **MapToSlice**
  
- **Retrier**: A tool to run a function until an error returns.

- **Sync**
  - **MutexBlock**: A tool used for execute many operations in safe block.
  - **MutexBlockWithValue**: Like **MutexBlock** with return value.
