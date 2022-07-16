## Usage

Make ``conf.json``.

```json
[
  {
    "domain":"test.example.com",
    "host":"host1.com",
    "scheme":"https"
  },
  {
    "domain":"test2.example.com",
    "host":"host2.com",
    "scheme":"https"
  }
]
```

```
.
├── conf.json
├── go.mod
├── go.sum
├── proxy.go
└── readme.md
```

Execute the command below.

```bash
go run .
```