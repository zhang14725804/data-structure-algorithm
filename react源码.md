### 1、协调(Reconciliation)

react有两套协调算法：stack reconciler(react@15.x) 和 fiber reconciler(react@16.x)

### 1.1 stack reconciler

在 react 16 之前的调度算法被称作 stack reconciler，它的特点是自顶向下的更新方式，比如调用 this.setState() 之后的流程就像这样：

    this.setState() => 生成虚拟 DOM => diff 算法比较 => 找到要更新的元素 => 放到更新队列里


在 stack reconciler 里面更新时同步的，自顶向下的更新方式，只要更新过程开始就会“一条道走到黑”直到所有节点的比对全部完成

### 浏览器帧

    每一帧进行的操作：
    （1）处理用户的交互
    （2）处JS 解析执行
    （3）处帧开始。窗口尺寸变更，页面滚去等的处理
    （4）处rAF(requestAnimationFrame)
    （5）处布局
    （6）处绘制

浏览器这一帧里面要做的事情着实不少，如果 js 引擎的执行占用的时间过长那势必导致其它任务的执行（比如响应用户交互）要延后，这也就是导致卡顿的原因

### 1.2 Fiber reconciler

    能够把可中断的任务切片处理
    能够调整优先级，重置并复用任务
    能够在父元素与子元素之间交错处理，以支持 React 中的布局
    能够在 render() 中返回多个元素
    更好地支持错误边界

要实现这些特性一个很重要的点就是：切片，把一个很大的任务拆分成很多小任务，每个小任务做完之后看看是否需要让位于其它优先级更高的任务，这样就能保证这个唯一的线程不会被一直占用，我们把每一个切片（小任务）就叫做一个 fiber

react 的整个执行流程分为两个大的阶段：

    Phase1: render/reconciliation
    Phase2: commit

第一个阶段负责产生 Virtual DOM -> diff 算法 ->  产生要执行的 update，第二个阶段负责将更新渲染到 DOM Tree 上

#### 如何调度任务执行？

    requestIdleCallback 该方法接收一个 callback，这个回调函数将在浏览器空闲的时候调用
    requestAnimationFrame 该方法接收一个 callback，这个回调函数会在下次浏览器重绘之前调用

它俩的区别是 requestIdleCallback 是需要等待浏览器空闲的时候才会执行而 requestAnimationFrame 是每帧都会执行，所以高优先级的任务交给 requestAnimationFrame 低优先级的任务交给 requestIdleCallback 去处理，但是为了避免因为浏览器一直被占用导致低优先级任务一直无法执行，requestIdleCallback 还提供了一个 timeout 参数指定超过该事件就强制执行回调。

### Fiber中的一些概念

- Fiber

它表示 react 中最小的一个工作单元， 在 react 中 ClassComponent，FunctionComponent，普通 DOM 节点，文本节点都对应一个 Fiber 对象，Fiber 对象的本质其实就是一个 Javascript 对象。

- child，subling，return

child 是 Fiber 对象上的属性，指向的是它的第一个子节点（Fiber）；sibling 是 Fiber 对象上的属性，指向的是它的第一个兄弟节点（Fiber）；eturn 是 Fiber 对象上的属性，指向的是它的父节点（Fiber）

- stateNode

stateNode 是 Fiber 对象上的属性，表示的是 Fiber 对象对应的实例对象，比如 Class 实例、DOM 节点等

- current

current 表示已经完成更新的 Fiber 对象

- workInProgress

workInProgress 表示正在更新的 Fiber 对象

- alternate

alternate 用来指向 current 或 workInProgress，current 的 alternate 指向 workInProgress，而 workInProgress 的 alternate 指向 current（这他么图了个什么）

- FiberRoot，RootFiber（FiberRoot有什么用呢，这么看好像是多余的）

FiberRoot 表示整个应用的起点，它内部保存着 container 信息，对应的 fiber 对象（RootFiber）等；

RootFiber 表示整个 fiber tree 的根节点，它内部的 stateNode 指向 FiberRoot，它的 return 为  null

- fiber tree

链表结构；fiber 的 child 只指向第一个 子节点，但是可以通过 sibling 找到其兄弟节点，通过return指向父节点

- 挂载阶段

第一次渲染阶段也就是挂载阶段，react 会自上而下的创建整个 fiber tree，创建顺序是同 FiberRoot 开始，通过一个叫做 work loop 的循环来不断创建 fiber 节点，循环会先遍历子节点（child），当没有子节点再遍历兄弟节点（sibling），最终达到全部创建的目的，以我们的 demo 为例，fiber 的创建顺序如下图箭头所示。（根据什么创建fiber tree，虚拟DOM么）

- 更新阶段（这里有些晦涩难懂！！！！）

[react-fiber-architecture](https://github.com/acdlite/react-fiber-architecture)

[fiber](https://www.youtube.com/watch?v=ZCuYPiUIONs&t=637s)

更新会放入 List 组件的更新队列里（update queue），在 fiber reconciler 中异步的更新并不会立即处理，它会执行调度（schedule）程序，让调度程序判断什么时候让它执行，调度程序就是通过上面所说的 requestIdleCallback 来判断何时处理更新。

当主进程把控制权交给我们的时候我们就可以开始执行更新了，我们把这个阶段叫做 work loop，work loop 是一个循环，它循环的去执行下一个任务，在执行之前要先判断剩余的时间是否够用，所以对于 work loop 它要追踪两个东西：下一个工作单元和剩余时间。

react 会保留之前生成的 fiber tree 我们管它叫 current fiber tree，然后新生成一个 workInProgress tree，它用来计算出产生的变化，执行 reconciliation 的过程，HostRoot 可以直接通过克隆旧的 HostRoot 产生，此时新的 HostRoot 的 child 还指向老的 List 节点。

当 HostRoot 处理完成之后就会向下寻找它的子节点也就是 List，所以下一个任务就是 List，我们同样可以克隆之前的 List 节点，克隆好之后 react 会判断一下是否还有剩余的时间，发现还有剩余的时间，那么开始执行 List 节点的更新任务。

当我们执行 List 的更新时我们发现 List 的 update queue 里面有 update，所以 react 要处理该更新

当我们处理完 update queue 之后会判断 List 节点上面是否有 effect 要处理，比如 componentDidUpdate ，getSnapshotBeforeUpdate。

因为 List 产生了新的 state，所以 react 会调用它的 render 方法返回新的 VDOM，接下来就用到了 diff 算法了，根据新旧节点的类型来判断是否可以复用，可以复用的话就直接复制旧的节点，否则就删除掉旧的节点创建新的节点

下一个要处理的任务就是 List 的第一个 child：button，我们还是在处理之前先检查一下是否还有剩余的时间

为了显示出 fiber 的作用我们此时假设用户点击了放大字体的按钮，这个逻辑和 react 无关，完全由原生 JS 实现，此时一个 callback 生成需要等待主线程去处理，但是此时主线程并不会立即处理，因为此时距离下一帧还有剩余的时间，主线程还是会先处理 react 相关的任务

对于 button 节点它没有任何更新，而且也没有子节点（文本节点不算），所以我们可以执行完成逻辑（completeUnitOfWork）这里 react 会比对新旧节点属性的变化，记录在 fiber 对象的 upddateQueue 里，接着 react 会找 button 的兄弟节点（sibling）也就是第一个 Item。

接着 work loop 去查看是否有剩余的时间，发现还有剩余时间，那接下来的执行过程其实和 List 是类似的。

接着 react 会找第二个 Item 的兄弟节点（sibling）也就是第三个 Item，此时 work loop 进行 deadline 判断的时候发现已经没有剩余时间了，此时 react 会将执行权交还给主进程，但是 react 还有剩余的任务没有执行完，所以它会在之前结束的地方等待主进程空闲时继续完成剩余工作。

当已经完成对整个 fiber tree 的最后一个节点更新后，react 会开始不断向上寻找父节点执行完成工作（completeUnitOfWork），如果父节点是普通 DOM 节点会比对属性的变化放在 update queue 上，同时将子节点产生的 effect 链挂载在自己的 effect 链上。

当 react 遍历完最后一个 fiber 节点也就是 HostRoot，我们的第一个阶段（Reconciliation phase）就完成了，我们将把这个 HostRoot 交给第二个阶段（Commit phase）进行处理。

在执行 commit 阶段之前我们还会判断一下时间是否够用

接下来开始第二个阶段（Commit phase），react 将会从第一个 effect 开始遍历更新，对于 DOM 元素就执行对应的增删改的操作，对于 Class 组件会将执行对应的声明周期函数：componentDidMount、componentDidUpdate、componentWillUnmount，解除 ref 的绑定等

commit 阶段执行完成后 DOM 已经更新完成，这个时候 workInProgress tree 就是当前 App 的最新状态了，所以此时 react 将会把 current 指向 workInProgress tree

上面的整个流程可以让浏览器的任务不被一直打断，但是还有一个问题没有解决，如果 react 当前处理的任务耗时过长导致后面更紧急的任务无法快速响应，那该怎么办呢？

react 的做法是设置优先级，通过调度算法（schedule）来找到高优先级的任务让它先执行，也就是说高优先级的任务会打断低优先级的任务，等到高优先级的任务执行完成之后再去执行低优先级的任务，注意是从头开始执行。

### Fiber reconciler 流程

fiber reconciler 将一个更新分成两个阶段（Phase）：Reconciliation Phase 和 Commit Phase，第一个阶段也就是我们上面所说的协调的过程，它的作用是找出哪些 dom 是需要更新的，这个阶段是可以被打断的；第二个阶段就是将找出的那些 dom 渲染出来的过程，这个阶段是不能被打断的。对我们有影响的就是这两个阶段会调用的生命周期函数，以 render 函数为界，第一个阶段会调用以下生命周期函数：

    componentWillMount
    componentWillReceiveProps
    shouldComponentUpdate
    componentWillUpdate
    render

第二个阶段会调用的生命周期函数：

    componentDidMount
    componentDidUpdate
    componentWillUnmount

因为 fiber reconciler 会导致第一个阶段被多次执行所以我们需要注意在第一阶段的生命周期函数里不要执行那些只能调用一次的操作。





