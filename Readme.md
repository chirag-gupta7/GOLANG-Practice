# Go Practice — Hitesh

A progressive series of Go exercises (01–23) covering fundamentals from Hello World to building a REST API.

| #  | Topic                   | Details |
|----|-------------------------|---------|
| 01 | Hello World             | Basic program printing "Hello, World!" |
| 02 | Variables & Types       | Declare string, bool, uint8, float64, rune variables; use `%T` to inspect types |
| 03 | User Input              | Read name and age from stdin using `bufio.NewReader` and `strings.TrimSpace` |
| 04 | Type Conversion         | Convert string input to float64 using `strconv.ParseFloat`; handle errors |
| 05 | Date & Time             | Get current time with `time.Now()`, format with `Format()`, create dates with `time.Date()` |
| 06 | Pointers                | Store address with `&`, dereference with `*` |
| 07 | Arrays                  | Fixed-size `[4]string` array, assign and print elements |
| 08 | Slices                  | Dynamic slices with `append()`, create with `make()`, remove element by index |
| 09 | Maps                    | Create map with `make()`, iterate with `for range` |
| 10 | Structs                 | Define `Info` struct (Name, Age, Email), access fields with dot notation |
| 11 | If/Else                 | Check if a number is even or odd using `if/else` |
| 12 | Switch                  | Dice roll switch with `rand.Intn()`, multiple cases, default case |
| 13 | Loops                   | For loop variants, `range` loop, `break`, `continue`, `goto` with labels |
| 14 | Functions               | Regular function `greet()`, variadic function `adder(x ...int)` |
| 15 | Methods                 | Method `getMail()` on `Info` struct with value receiver |
| 16 | Defer                   | `defer` keyword for delayed execution; deferred calls run in LIFO order |
| 17 | File I/O                | Create file with `os.Create`, write with `io.WriteString`, read with `os.ReadFile` |
| 18 | Web Requests            | HTTP GET request with `http.Get()`, read body with `io.ReadAll`, close with `defer` |
| 19 | URL Parsing             | Parse URL with `url.Parse()`, extract Scheme, Host, Path, Port, Query params |
| 20 | HTTP Client Requests    | Perform GET, POST JSON, and POST form requests with `http.Get`, `http.Post`, `http.PostForm` |
| 21 | JSON Encoding/Decoding  | Marshal/Unmarshal with `encoding/json`, struct tags, `json.Valid()` checking |
| 22 | Custom Modules          | Gorilla Mux HTTP server with custom router, vendored dependencies |
| 23 | REST API                | CRUD API using Gorilla Mux, JSON handlers for courses resource |

Each directory is a standalone program. Run with `go run main.go` (or the appropriate `.go` file).
