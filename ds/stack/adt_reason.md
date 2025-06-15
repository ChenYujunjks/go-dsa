这是一个非常经典、重要又很有深度的问题。你问的是：

> 既然我可以直接用 array 或 linked list，为什么还要用 Stack、Queue 这种抽象数据类型（ADT）？看不到底层细节反而多此一举？

---

## ✳️ 简答：

> **ADT（抽象数据类型）的目的，不是隐藏细节，而是为了“**清晰地表达意图**”和“**解耦实现与使用**”。**

---

## 🔍 更深入的解释：

### 🧱 1. **抽象数据类型是“契约”而非“实现”**

- `Stack` 定义了操作规则（只能从顶端添加和删除）
- `Queue` 定义了先进先出规则（FIFO）

这些规则让使用者**不用关心实现细节**，就可以安全使用这些数据结构。

👉 举个例子：

```go
s := NewStack()
s.Push("a")
s.Push("b")
s.Pop() // 你不需要知道底层是 array 还是 linked list
```

你写这段代码只要知道：

- `Push` 会加到栈顶
- `Pop` 会移除最新的元素

你**不关心它是 slice append 的，还是 linked list 的头节点。**

---

### 🧼 2. **模块化 + 可维护性：底层随时可以换**

假设你用 `LinkedList` 实现了 `Queue`，后来你发现：

- 对于你的项目，`array` 性能更好（比如插入操作不多，遍历很多）

如果你一开始就**把所有逻辑直接写死在 `LinkedList` 上**，你要改动所有使用它的地方。

但如果你封装成 `Queue` 抽象：

```go
type Queue struct {
    impl *LinkedList
}
```

后来你只需要换成：

```go
type Queue struct {
    impl []interface{}
}
```

使用 Queue 的代码 **一行都不需要改** ✅

> **抽象使得代码更灵活、可维护、可扩展。**

---

### 📖 3. **教育目的：强迫你关注行为而不是存储方式**

在算法学习和设计中，我们更关心：

- 这个问题要用「什么策略」来解决
- 它需要的是**先进先出**还是**后进先出**行为
- 它对插入删除的时间复杂度有何要求

你不应该一上来就纠结：

> 是用 `slice` 还是 `linked list`？

> 是不是得先定义个结构体 Head Tail Node？

这些在选型阶段是次要的。

---

### 🧪 4. **便于测试与验证**

如果你给别人提供一个 `Queue` 的抽象：

```go
type Queue interface {
  Enqueue(x any)
  Dequeue() any
  Peek() any
  Len() int
}
```

你就可以给任何实现这个接口的东西写通用测试：

```go
func TestQueueBehavior(t *testing.T, q Queue) {
  q.Enqueue(1)
  q.Enqueue(2)
  if q.Dequeue() != 1 {
    t.Error("Queue did not behave correctly")
  }
}
```

**测试逻辑和实现彻底解耦。**
