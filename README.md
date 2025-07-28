# Learning-simple-GoDoList

这是一个用来练习 Go Web 开发的简单实践项目。

---

## 一、项目概述与业务场景

**项目名称：GoDoList (Go + To-Do List)**

**业务场景：**

这是一个极简的个人任务管理工具，用户在网页上可以：

* 添加新的待办事项
* 查看所有未完成的待办事项
* 标记为已完成
* 删除任务

这个场景包含最基本的 **CRUD (Create, Read, Update, Delete)** 操作，是学习 Web 开发的精美入门框体。

---

## 二、核心学习要点

通过该项目，你将学习以及系统地应用以下技术：

* Go 的 `net/http`包：启动 HTTP 服务器，处理请求
* URL 路由处理：区分 GET/POST/PUT/DELETE 操作
* JSON 处理：用 `encoding/json` 进行序列化/反序列化
* RESTful API 设计：经典的接口要素和定义
* 静态文件服务：提供 HTML/CSS/JS 资源
* 并发安全：用 `sync.Mutex` 确保全局数据安全
* 前后端交互：前端 fetch API 调用后端接口

---

## 三、后端功能实现 (Go)

### 数据结构

```go
type Todo struct {
    ID        int    `json:"id"`
    Task      string `json:"task"`
    Completed bool   `json:"completed"`
}
```

### 数据存储

```go
var (
    todos  []Todo     // 存储待办列表
    nextID = 1        // 唯一 ID 生成器
    mu     sync.Mutex // 互斥锁
)
```

### API 设计

| 功能   | 方法     | URL             | 请求体                 | 响应体            |
| ---- | ------ | --------------- | ------------------- | -------------- |
| 获取列表 | GET    | /api/todos      | 无                   | 待办列表 JSON      |
| 添加任务 | POST   | /api/todos      | {"task": "xxx"}     | 新任务 JSON       |
| 更新状态 | PUT    | /api/todos/{id} | {"completed": true} | 更新后 JSON       |
| 删除任务 | DELETE | /api/todos/{id} | 无                   | 204 No Content |

---

## 四、前端设计 (HTML/CSS/JS)

### HTML 页面结构

* `<h1>` 标题
* `<form>` 输入框 + 按钮
* `<ul>` 列表区，动态创建 `<li>`

### CSS 核心样式

* 居中布局
* 美化按钮和输入框
* `.completed { text-decoration: line-through; }`

### JavaScript 交互逻辑

* 页面加载：GET `/api/todos` 列表
* 新增任务：POST `/api/todos`
* 更新任务：PUT `/api/todos/{id}`
* 删除任务：DELETE `/api/todos/{id}`

使用事件委托，监听 `<ul>` 上的按钮操作

---

## 五、项目目录经验

```
GoDoList/
├── main.go          // Go 后端
├── static/
│   ├── index.html   // 前端页面
│   ├── style.css    // CSS 样式
│   └── script.js    // JS 逻辑
```

---

## 六、技术栏目总结

| 分层   | 技术              | 说明                    |
| ---- | --------------- | --------------------- |
| 前端   | HTML / CSS / JS | 原生技术栏目                |
| 后端   | Go (`net/http`) | 构建 RESTful API + 静态服务 |
| 并发安全 | sync.Mutex      | 确保全局数据的安全             |
| 数据格式 | JSON            | 前后端数据交换               |

---

## 七、学习目标

* 熟悉 Go 基础 Web 开发
* 理解 RESTful API 和 HTTP 操作
* 熟悉前后端交互流程
* 练习并发编程安全处理
* 为后续实战网站经验打基
