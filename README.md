# 🥭 MangoTemplate v1

**A minimal HTML template engine for Go** – variable interpolation only, zero dependencies, designed for micro‑frameworks like Flusk.

---

## ✨ Features

- 🍃 **Simple variable replacement** – `{{ key }}` → value from a `map[string]any`
- 📁 **File‑based rendering** – load HTML templates from disk
- 🚀 **Zero dependencies** – pure Go standard library
- 🔌 **Embeddable** – perfect for lightweight web frameworks
- 🧹 **Minimal & readable** – under 150 lines of code
- 📦 **MIT Licensed** – use anywhere

---

## 📦 Installation

```bash
go get github.com/AlexanderXinarxZenDev/mango
```

Or manually add to your project.

---

## 🚀 Quick start

### 1. Create a template file (`index.html`)

```html
<h1>{{ title }}</h1>
<p>{{ message }}</p>
<span>{{ missing_key }}</span>
```

### 2. Render in Go

```go
package main

import (
    "fmt"
    "log"
    "github.com/AlexanderXinarxZenDev/mango_template"
)

func main() {
    data := map[string]any{
        "title":   "Flusk App",
        "message": "Hello MangoTemplate",
    }

    result, err := mango.Render("index.html", data)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(result)
}
```

### 3. Output

```html
<h1>Flusk App</h1>
<p>Hello MangoTemplate</p>
<span></span>
```

> **Note:** Missing keys become empty strings – no errors, just silence.

---

## 📚 API

### `Render(filePath string, data map[string]any) (string, error)`

Reads an HTML file, replaces all `{{ key }}` placeholders, and returns the rendered string.

- If a key exists in `data`, its value is converted to a string and inserted.
- If a key does **not** exist, it is replaced with an empty string.
- Malformed placeholders (e.g., `{{ key }` or `{{key` ) are left untouched.

---

## 🔗 Integration with Flusk

MangoTemplate is designed to work seamlessly with the [Flusk](https://github.com/AlexanderXinarxZenDev/flusk) web framework:

```go
// Flusk context extension
func (c *Context) HTML(status int, file string, data map[string]any) {
    rendered, err := mango.Render(file, data)
    if err != nil {
        c.Text(500, err.Error())
        return
    }
    c.Writer.Header().Set("Content-Type", "text/html")
    c.Writer.WriteHeader(status)
    c.Writer.Write([]byte(rendered))
}

// Usage
app.GET("/", func(c *flusk.Context) {
    c.HTML(200, "index.html", map[string]any{
        "title": "Home",
    })
})
```

---

## 🧠 Design philosophy

- **Minimalism first** – no loops, conditionals, or complex parsing in v1.
- **Explicit over magical** – you see exactly what the engine does.
- **Framework‑friendly** – small enough to embed, predictable enough to trust.
- **Ready to evolve** – the simple parser can be extended later (include files, functions) without breaking existing templates.

---

## 📁 Project structure

```
mango/
├── render.go       # Public Render function
├── parser.go       # Placeholder replacement logic
└── go.mod
```

---

## 🚫 What MangoTemplate is NOT

- A full‑featured template engine (use `html/template` for that)
- A regex‑based monster – it’s a simple state machine
- A tool for complex logic – keep logic in Go, keep templates dumb

---

## 🧪 Running the example

```bash
cd example
go run main.go
```

---

## 📄 License

MIT – see [LICENSE](LICENSE) file.

---

## 🤝 Contributing

Keep it minimal. Keep it simple. Open an issue or PR if you have an idea that fits the v1 spirit.

---

**Built with 🥭 for Go developers who want just enough templating.**
