
dead lock
```go package main

import (
    "fmt"
    // _ "github.com/go-sql-driver/mysql"
)

func main() {
    c1 := make(chan int, 10)
    fmt.Println(<-c1)
}

```

block：

```go package main

import (
    "fmt"
    // _ "github.com/go-sql-driver/mysql"
)

func main() {
    c1 := make(chan int, 10)
    fmt.Println(<-c1)
}


1、如果不导入其他包，chan是空的，导致死锁的发生，所有G被阻止，没有一个可以运行；
2、如果导入其他包，可能会启动一些未被阻止的其他G，这样就不会有死锁，即使主例程被永久阻止。
      因此两种情况实际上都是阻止， 仅仅是在第一种情况下，编译器可以检测到它是一个“永远被阻塞”的情况，并在程序中发出死锁的警告。