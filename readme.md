# Go Data Structures & Algorithms Playground

> Implementing classic data structures and algorithms in pure Go 🚀  
> 用 Go 语言实现所有经典数据结构与基础算法，从教材到实战。

## 📘 What is this?

This repository is a summer-long deep-dive into the foundations of computer science, built entirely in Go.  
From arrays and linked lists to trees, graphs, and hash tables, plus classic algorithms like sorting, recursion, and dynamic programming — all implemented from scratch.

## 🔍 这是一个什么项目？

这是一个用 Go 语言手写所有常见数据结构与基础算法的暑期学习项目。

- 覆盖所有基础数据结构：数组、链表、栈、队列、树、图、堆、哈希表等
- 实现常见算法：排序、查找、递归、动态规划、回溯、图搜索等
- 所有代码均为纯 Go 实现，附带注释与单元测试
- 项目结构清晰，适合学习、面试准备、源码阅读等用途

## 📁 Project Structure | 项目结构

```

linkedlist/ // Single and doubly linked lists
stack/ // Stack implementations
queue/ // FIFO queue
tree/ // Binary tree, BST, traversal
hash/ // Hash map, hash set
sort/ // Sorting algorithms
dp/ // Dynamic programming problems
graph/ // DFS, BFS, shortest path
common/ // Reusable utils

```

## ✅ Features

- Fully written in idiomatic Go
- Clean, modular codebase
- Unit tests for every major component
- Self-contained and dependency-free
- Annotated in both English and Chinese (where needed)

## 📦 How to use

```bash
git clone https://github.com/yourusername/go-data-structures.git
cd go-data-structures
go test ./...    # Run all tests
```

## 🌱 About the Author

Created with love and curiosity by [@yourusername](https://github.com/yourusername), a CS student passionate about building things from scratch and understanding how things work under the hood.

## 💬 欢迎交流

如果你也在学习 Go，或者想交流算法实现、系统设计、项目结构等内容，欢迎提 Issue 或发邮件交流！

## ✅ 项目优势总结：

- 📘 教材级别的数据结构全覆盖（Array, LinkedList, Tree, Stack, Queue...）
- 🧠 配套的经典算法（Sorting, Searching, Recursion, DP...）
- 🔨 完整使用 Go 实现，提升语言能力 + 工程能力
- 📚 代码结构清晰，测试齐全，可作为别人学习 Go 的范例项目
- 🚀 自己写 README、注释和测试，训练项目维护和表达能力

---

## 🔍 Highlighted Commit: Fixing Go Interface Nil Trap

We resolved a tricky panic in our LinkedList implementation due to the Go interface `nil` trap.

🧠 Problem: A node interface (`NodeInter`) that _looked_ like `nil` was not truly `nil`, leading to incorrect behavior in `IsNilNode` checks and bugs in `ToSlice()` and list emptiness checks.

✅ Fix Commit: [`55cf479`](https://github.com/ChenYujunjks/go-review/commit/55cf479)

🔧 Solution: Introduced a helper using `reflect` to robustly detect true nil interfaces:

```go
func IsNilNode(n iface.NodeInter) bool {
	if n == nil {
		return true
	}
	return reflect.ValueOf(n).IsNil()
}
```

🧪 Result: All related unit tests now pass, ensuring safe `ToSlice()` and `Length()` operations.
