### ✅ 简单定义：

- ADT 是 **Abstract Data Type（抽象数据类型）** 的缩写。

> **ADT 是一种定义了数据操作方式的模型，而不关心底层是如何实现的。**

你可以把它想象成一个“接口”或“操作规范”。

---

### 📦 举个例子：Stack（栈）这个 ADT

#### Stack 的定义：

- 你只能从**一端**添加或删除元素。
- 有两个操作：

  - `push(x)`：压入元素 `x`
  - `pop()`：弹出最近压入的元素

> 不管你用 **数组** 还是 **链表** 来实现，只要满足上面这两个操作的行为逻辑，它就是一个 Stack。

---

### 📦 再看 Queue（队列）这个 ADT：

- 操作规则是 **先进先出（FIFO）**

  - `enqueue(x)`：加入队尾
  - `dequeue()`：移除队首

- 同样可以用 array 或 linked list 实现

---

太好了你已经理解了 array vs. list 的区别！现在我们来看你要的内容：**所有常见 ADT 的规范** —— 即：**“怎么判断一个结构是某个 ADT”**。

---

## ✅ 什么是 ADT 规范（interface specification）？

每个 ADT 都通过其：

1. **数据特性（结构）**
2. **支持的操作（操作接口）**
3. **操作的语义（行为逻辑）**

来定义。只要一个结构符合这三点，我们就可以说“它是某个 ADT 的实现”。

---

## 🧩 常见 ADT 列表及其规范

### 1. **List（线性表）**

- **结构特征**：有序、可重复元素
- **核心操作**：

  - `insert(index, element)`
  - `remove(index)`
  - `get(index)`
  - `set(index, element)`
  - `length()`

- **语义**：支持随机访问、插入删除，按顺序排列

### 2. **Stack（栈）**

- **结构特征**：**先进后出（LIFO）**
- **核心操作**：

  - `push(element)`：压栈
  - `pop()`：弹栈
  - `peek()` / `top()`：查看栈顶元素
  - `isEmpty()`

- **语义**：

  - `pop()` 必定返回最后一个被 `push()` 的元素
  - 后进先出，访问顺序倒转

---

### 3. **Queue（队列）**

- **结构特征**：**先进先出（FIFO）**
- **核心操作**：

  - `enqueue(element)`：入队
  - `dequeue()`：出队
  - `peek()` / `front()`：查看队首元素
  - `isEmpty()`

- **语义**：

  - `dequeue()` 总是取出最早进入的元素

---

### 4. **Deque（双端队列）**

- **结构特征**：两端都可以插入删除
- **核心操作**：

  - `addFront(element)` / `addBack(element)`
  - `removeFront()` / `removeBack()`
  - `peekFront()` / `peekBack()`

- **语义**：

  - 可模拟栈或队列，也可左右两边都操作

---

### 5. **Set（集合）**

- **结构特征**：**无序，不重复**
- **核心操作**：

  - `add(element)`
  - `remove(element)`
  - `contains(element)`
  - `size()`

- **语义**：

  - 元素不能重复
  - 不关心顺序

---

### 6. **Map / Dictionary（映射）**

- **结构特征**：**key → value 对**
- **核心操作**：

  - `put(key, value)`
  - `get(key)`
  - `remove(key)`
  - `containsKey(key)`

- **语义**：

  - 每个 `key` 对应一个唯一的 `value`
  - 类似现实生活中的词典

---

### 7. **Priority Queue（优先队列）**

- **结构特征**：**元素有优先级**
- **核心操作**：

  - `insert(element, priority)`
  - `removeMin()` / `removeMax()`
  - `peek()`

- **语义**：

  - 总是取出优先级最高（或最低）的元素，而不是最早或最后加入的

---

### 8. **Graph（图）**

- **结构特征**：由节点（Vertex）和边（Edge）组成
- **核心操作**：

  - `addVertex(v)`
  - `addEdge(v1, v2)`
  - `getNeighbors(v)`
  - `removeEdge(v1, v2)`

- **语义**：

  - 可表示路径、连通性等复杂结构

---

### 9. **Tree（树）**

- **结构特征**：分层结构，节点有父子关系，没有环
- **核心操作**：

  - `insert(parent, child)`
  - `traverse()`（如 DFS、BFS）
  - `delete(node)`
  - `getChildren(node)`

- **语义**：

  - 一般从一个根节点开始
  - 每个节点最多一个父节点，可有多个子节点
