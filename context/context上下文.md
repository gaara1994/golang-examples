

# context(上下文)



## 1.是什么

`Context` 是一个接口类型。

`Context` 定义在标准库 context/context.go 中，是在Go1.7版本之后引入的。



## 2.有什么用

`Context` 可以在 **goroutine** 之间专递、共享信息。

也就是说`Context`专门用来处理  **goroutine** 相关的。



## 3.长什么样

context.Context 接口定义了四个需要实现的方法

```go
type Context interface {
    //获取截止时间
	Deadline() (deadline time.Time, ok bool)

    //监听完成信号的通道
	Done() <-chan struct{}

    //获取取消原因
	Err() error

    //检索键值对
	Value(key any) any
}
```

### 1. Deadline()

需要返回当前`Context`被取消的时间，也就是完成工作的截止时间（deadline）

### 2. Done()

需要返回一个`Channel`，这个Channel会在当前工作完成或者上下文被取消之后关闭，多次调用`Done`方法会返回同一个Channel

### 3. Err()

返回当前`Context`结束的原因，它只会在`Done`返回的Channel被关闭时才会返回非空的值

 如果当前`Context`被取消就会返回`Canceled`错误； * 如果当前`Context`超时就会返回`DeadlineExceeded`错误

### 4. Value(key any)

从`Context`中返回键对应的值，对于同一个上下文来说，多次调用`Value` 并传入相同的`Key`会返回相同的结果，该方法仅用于传递跨API和进程间跟请求域的数据；



## 4. 根上下文 Background contex

上下文(context) 总是要有个头的，就像linux目录总有一个根目录 / 作为一切的起始目录。golang上下文的根节点就叫 Background contex 。

根 `context` 很简单，不具备任何功能。所有其他的上下文都应该从它衍生（Derived）出来。

```shell
            +-----------------+
            | context.Background| --- 根节点
            +-----------------+
                          |
                 +---------v---------+
                 |  WithCancel/Ctx  |
                 +---------+---------+
                          |
               +----------v----------+
               |  WithValue/Ctx     |
               |  Key: "User", Val: "Alice" |
               +----------+----------+
                          |
            +---------v---------+
            |  WithDeadline/Ctx  |
            |   Deadline: 5s     |
            +---------+---------+
                          |
                 +--------v--------+
                 |    Child Task    |
                 +-----------------+
```



### 1. 根节点的创建方法

context.Background() 函数用来创建根节点上下文。

```go
package main

import (
	"context"
	"fmt"
)

func main() {
	baseCtx := context.Background()
	fmt.Println(baseCtx)
}
```



我们看一下 context.Background() 的源码。

```go
// Background returns a non-nil, empty [Context]. It is never canceled, has no
// values, and has no deadline. It is typically used by the main function,
// initialization, and tests, and as the top-level Context for incoming
// requests.
func Background() Context {
	return backgroundCtx{}
}
```

下面是这段注释的详细解释：

- **非-nil，空的[Context]**：`Background`函数返回一个非零值（即非-nil）的上下文实例。这里的“空”指的是这个上下文没有附带任何特定的值（通过`WithValue`设置的键值对）、截止时间或取消功能，是一个非常基础的上下文对象。

- **它永远不会被取消**：由于没有关联的取消函数，由`Background`创建的上下文将永远不会被显式取消。这使得它适合用在程序的主流程、初始化阶段或测试环境中，这些情境下通常不需要根据外部条件提前终止操作。

- **没有值，没有截止时间**：强调了这个上下文的简洁性，它既不能存储额外的数据（尽管可以通过后续的`WithValue`方法添加），也没有预设的执行期限，提供了最大的灵活性。

- **典型用途**：
  - **主函数**：在程序的入口点，当没有特定的上下文需求时，可以使用`Background`来初始化上下文。
  - **初始化**：系统或模块的初始化过程中，如果需要一个基础的上下文作为起点，也会用到`Background`。
  - **测试**：编写单元测试或集成测试时，为了简化测试逻辑，测试代码经常使用`Background`作为测试操作的上下文。
  - **顶层Context**：对于处理外部请求（如HTTP服务器接收的请求）的应用，如果没有从请求中继承特定的上下文，可以首先使用`Background`作为处理该请求的顶级上下文。

- **实现细节**：注释最后提到函数返回的是一个`backgroundCtx{}`实例，这是`context`包内部定义的一个类型，用来表示一个不可取消、无截止时间且无值的上下文。这样的设计保证了`Background`的高效和轻量级。



用途举例，前面已经说了，根上下文没有任何功能，只是传递下去而已，ctx没有被使用，它也没有可以被使用的功能。

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 直接使用 Background 作为上下文执行操作
	baseCtx := context.Background()

	// 假设我们有一个函数需要一个上下文参数，但不需要取消或超时功能，也不需要传递值
	go simpleTask(baseCtx)

	// 等待一段时间后程序结束，这里仅为演示，实际中可能有更复杂的流程控制
	time.Sleep(2 * time.Second)
}

func simpleTask(ctx context.Context) {
	// 这里我们的任务很简单，只是打印一条消息
	fmt.Println("正在执行一个简单的任务...")
	// 假设有一些耗时操作
	time.Sleep(1 * time.Second)
	fmt.Println("任务完成")
}
```



### 2. context.TODO() 

看源码

```go
// TODO returns a non-nil, empty [Context]. Code should use context.TODO when
// it's unclear which Context to use or it is not yet available (because the
// surrounding function has not yet been extended to accept a Context
// parameter).
func TODO() Context {
	return todoCtx{}
}
```

这段注释来自Go语言标准库`context`包中关于`TODO`函数的说明。下面是这段注释的详细解释：

- **非-nil，空的[Context]**：类似于`Background`，`TODO`函数也返回一个非零值（即非-nil）的基本上下文实例。它同样没有关联的截止时间、取消功能或存储的值。

- **用途场景**：与`Background`不同，`TODO`主要用于开发过程中的过渡阶段，当你在编写代码时还没有确定应该使用哪个具体的上下文，或者当前的函数、方法尚未被修改以接受`Context`参数时。使用`context.TODO()`可以帮助快速推进代码编写，同时标记出未来需要回过头来进一步明确或优化上下文使用的地方。

- **代码标记**：在代码审查或后续的开发迭代中，`TODO`上下文的存在提醒开发者这里需要仔细考虑上下文的正确使用，可能涉及到添加取消逻辑、设置截止时间或是传递特定的上下文值。

- **实现细节**：注释最后提到函数返回的是一个`todoCtx{}`实例，这是`context`包内部定义的一个类型，用来代表一个临时的、待确定用途的上下文。与`Background`相比，虽然它们都返回一个基础的上下文，但`TODO`明确指示了这是一种开发过程中的权宜之计，应当在未来被更加具体和合适的上下文所替换。

总之，`TODO()`函数提供了一个便捷的方式来填充那些上下文使用尚不明确的代码位置，确保了代码的编译和运行不会因为空指针错误而中断，同时也是一个明确的标记，提示开发者后续需要优化上下文的使用。



context.TODO()  和 context.Background() 返回的是同一个类型的 context，都是`struct{ emptyCtx }`

```go
type todoCtx struct{ emptyCtx }
```

```go
type backgroundCtx struct{ emptyCtx }
```



## 5. 派生的上下文



### 1. WithCancel

**取消控制**

`WithCancel`返回带有新Done通道的父节点的副本。当调用返回的cancel函数或当关闭父上下文的Done通道时，将关闭返回上下文的Done通道，无论先发生什么情况。

取消此上下文将释放与其关联的资源，因此代码应该在此上下文中运行的操作完成后立即调用cancel。

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//创建一个根上下文，通常在main函数或者初始化函数中创建
	ctx := context.Background()

	//使用WithCancel从父上下文创建一个可以取消的上下文
	ctx,cancle := context.WithCancel(ctx)

	//启动一个goroutine，把新的上下文传递进去，让goroutine监听这个ctx，并相应的退出
	go doSomething(ctx)

	//等待3秒之后关闭上下文
	time.Sleep(3 * time.Second)
	cancle()

	// 等待一段时间让goroutine有机会退出
	time.Sleep(2 * time.Second)
}

func doSomething(ctx context.Context)  {
	//通过select 监听上下文的Done通道
	for{
		select {
		case <-ctx.Done():
			//当上下文取消时，这里会接收到信号
			fmt.Println("接收到信号停止运行")
			return
		default:
			fmt.Println("工作中...")
			time.Sleep(1 * time.Second)
		}
	}
}
```



### 2. WithDeadline

**超时控制**

`context.WithDeadline` 是 Go 语言标准库中的一个函数，用于创建一个新的带有截止时间（deadline）的上下文（Context）。这个新上下文是基于给定的父上下文（parent context）派生的，并且会在指定的截止时间到达时自动取消。如果截止时间在此之前就已经到达，那么创建的上下文会立即进入已取消状态。

当上下文（Context）过期或者被取消时，`<-ctx.Done()` 会发生变化。具体来说，`ctx.Done()` 会返回一个只读通道，这个通道在上下文被取消或者达到其设定的截止时间时会被关闭。因此，当你通过 `case <-ctx.Done()` 在 `select` 语句中等待时，一旦上下文被取消，这个通道就会有数据可读（尽管读取到的数据通常会被忽略，因为我们只关心通道是否可读以判断上下文是否已结束），这将导致对应的 case 分支被执行，从而你可以根据这个信号做出相应的处理，比如停止执行某个操作、清理资源或返回错误等。

函数签名如下：

```go
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc)
```

参数解释：
- `parent Context`：这是原始的上下文，它可以携带取消信号或者截止时间信息。当`parent`被取消时，即便`deadline`未到，新创建的上下文也会被取消。
- `deadline time.Time`：指定的截止时间点，绝对时间。这个时间必须大于当前时间。

返回值：
- `Context`：一个新创建的具有截止时间属性的上下文。
- `CancelFunc`：一个函数，调用它会立即取消这个上下文，无论`deadline`是否已经到达。一旦函数被调用，或者`deadline`到达，所有监听上下文   `Done()`通道的goroutine将会接收到信号并可以据此退出。

使用`WithDeadline`可以帮助你控制那些可能运行时间过长的操作，比如网络请求或者数据库操作，以防止资源泄露或长时间阻塞。在启动需要限制执行时间的goroutine时，应该将这个带有截止时间的上下文传递进去，并在操作中定期检查上下文的`Done()`通道，以便在截止时间到达或上下文被取消时及时退出。



```go
package main

import (
	"context"
	"fmt"
	"time"
)

// simulateLongRunningOperation 模拟一个可能运行时间较长的操作
func simulateLongRunningOperation(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("操作正在进行...")
		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				fmt.Println("由于超过最后期限，操作被取消.",ctx.Err())
			} else if ctx.Err() == context.Canceled {
				fmt.Println("操作已取消:", ctx.Err())
			}
			return
		}
		// 注意：这里不再直接比较i和duration
	}
	// fmt.Println("Operation completed successfully.") 这行现在是多余的，因为我们通过context控制退出了
}

func main() {
	// 设置一个具体的截止时间，比如当前时间后3秒
	deadline := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	go simulateLongRunningOperation(ctx)

	// 让main函数等待一段时间，确保上面的goroutine有足够时间执行
	time.Sleep(10 * time.Second)
}
```



### 3. WithTimeout

**超时控制**

WithTimeout 与 context.WithDeadline 类似，不过它可以更方便的设置超时时间，WithDeadline 需要手动计算当前时间+持续时间。

WithTimeout 的第二个参数直接填写 相对时间。

```go
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
	return WithDeadline(parent, time.Now().Add(timeout))
}
```

参数解释：

- `parent Context`：这是基础的上下文，和`WithDeadline`一样，它也可以携带取消信号。
- `timeout time.Duration`：希望的操作超时时间长度，相对时间。这个参数定义了新上下文在多久之后会自动取消。

返回值与`WithDeadline`相同：

- `Context`：基于提供的超时时间创建的新上下文。
- `CancelFunc`：一个函数，调用它可以立即取消这个上下文，无论超时是否已到。调用此函数后，所有监听这个上下文`Done()`通道的goroutine会收到信号。

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func simulateTask(ctx context.Context) {
	for {
		select {
		case <-time.After(1 * time.Second):
			fmt.Println("Task running...")
		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				fmt.Println("Task canceled due to timeout.")
			} else {
				fmt.Println("Task canceled:", ctx.Err())
			}
			return
		}
	}
}

func main() {
	// 设置超时时间为2秒
	parentCtx := context.Background()
	ctx, cancel := context.WithTimeout(parentCtx, 2*time.Second)
	defer cancel() // 确保上下文在函数结束时被取消

	go simulateTask(ctx)

	// 让main函数等待一段时间，确保上面的goroutine有足够时间执行
	time.Sleep(5 * time.Second)
}
```



### 4. WithValue

**携带数据**



```go
func WithValue(parent Context, key, val interface{}) Context
```

参数解释：

- `parent Context`：这是基础的上下文，新创建的上下文将继承它的取消和截止时间属性。
- `key interface{}`：一个用于检索值的唯一键。推荐使用不可变类型（如字符串、整型常量等）作为键，以避免误用或冲突。
- `val interface{}`：与键关联的值，可以是任何类型。

返回值：

- `Context`：一个包含了指定键值对的新上下文。这个上下文是基于`parent`创建的，所以如果`parent`被取消或过期，新上下文也会被取消或过期。



```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 设置超时时间为2秒
	parentCtx := context.Background()

	ctx := context.WithValue(parentCtx,"user_id","10010")

	getUserName(ctx)

	// 让main函数等待一段时间，确保上面的goroutine有足够时间执行
	time.Sleep(5 * time.Second)
}

func getUserName(ctx context.Context) {
	id := ctx.Value("user_id")
	fmt.Println("用户的id是：",id)
}
```

