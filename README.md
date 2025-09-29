## GoLimitless

A set of tools to solve some GoLang limitations that lead to writing a lot of code.

#### Install
```bash
go get github.com/provincialig/golimitless
```

#### Tools

- **Data structures**
  - **Stack**: A thread-safe typed implementation of Stack.
  - **Set**: A thread-safe typed implementation of Set with many helpful methods.
  - **Map**: A thread-safe typed implementation of key-value map with many helpful methods.
  - **DoubleMap**: A double layer thread-safe key-value map, with many helpful methods.
  - **IndexedSlice**: A thread-safe key-value map where value is a slice, with many helpful methods.

- **Slice utils**
  - **Filter**
  - **Map**
  - **Reduce**
  - **ForEach**
  - **SliceToMap** / **MapToSlice**

- **Sync**
  - **MutexBlock**: A tool used for execute many operations in safe block.
  - **MutexBlockWithValue**: Like **MutexBlock** with return value.