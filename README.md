# Go Projects Repository

This repository is a collection of **Go projects** that showcase different programming concepts, algorithms, and applications. All projects include **unit tests** using **Ginkgo** and **Gomega**.

---

## Repository Structure

```
go-projects/
├── project1/
│   ├── main.go
│   ├── operations.go
│   ├── project1_test.go
│   ├── project1_suite_test.go
│   └── README.md
├── project2/
│   ├── main.go
│   ├── utils.go
│   ├── project2_test.go
│   ├── project2_suite_test.go
│   └── README.md
└── README.md
```

* Each project folder contains:

  * **main.go**: The main executable code.
  * **supporting files** (`operations.go`, `utils.go`, etc.): Additional functions.
  * **_test.go files**: Unit tests written with Ginkgo and Gomega.
  * **README.md**: Project-specific documentation.

---

## Features

* ✅ Go projects demonstrating **algorithms, data structures, and utilities**.
* ✅ **Test-driven development** with Ginkgo and Gomega.
* ✅ Well-structured projects for learning and experimentation.

---

## How to Run a Project

### **1. Using Go Run**

Navigate to a project folder and run:

```bash
go run *.go
```

> Note: Exclude `_test.go` files when running the project.

### **2. Build an Executable**

```bash
go build -o project_exe *.go
```

Run the executable:

```bash
./project_exe   # Linux / macOS
project_exe.exe # Windows
```

### **3. Run Tests**

All projects include tests using Ginkgo and Gomega.

* Using Go only:

```bash
go test ./...
```

* Using Ginkgo:

```bash
go run github.com/onsi/ginkgo/v2/ginkgo -v
```

---

## Notes

* Ensure **Go 1.21+** is installed.
* Each project is independent and can be run or tested individually.
* Tests cover both **unit tests** and **behavioral scenarios** where applicable.

---

This repository is ideal for **Go enthusiasts**, students, or developers looking to **practice Go programming with proper testing**.
