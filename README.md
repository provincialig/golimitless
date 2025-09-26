## GoLimitless

A set of tools to solve some GoLang limitations that lead to writing a lot of code.

#### Install
```bash
go get github.com/provincialig/golimitless
```

#### Tools

- **Data structures**
  - **SafeSet**: A thread-safe typed implementation of Set with many helpful methods.
  - **SageMap**: A thread-safe typed implementation of Map with many helpful methods.

- **Slice utils**
  - **Filter**
  - **Map**
  - **Reduce**
  - **ForEach**

- **Sync**
  - **MutexBlock**: A tool used for execute many operations in safe block.
  - **MutexBlockWithValue**: Like **MutexBlock** with return value.