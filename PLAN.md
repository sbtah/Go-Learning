# Golang Mastery Learning Plan
## 1 Hour Daily (30min Theory + 30min Practice)

---

## 🎓 RECOMMENDED UDEMY COURSES

Based on extensive research, here are the **TOP 3 Udemy Golang courses** for 2025:

### 🥇 #1 - "Go - The Complete Guide" by Maximilian Schwarzmüller
**⭐ Rating:** 4.7/5 (7,668+ ratings) | **👥 Students:** 42,966+
- **Duration:** 16 hours of video content
- **Best For:** Complete beginners to advanced learners
- **Why Choose This:** Maximilian has taught over 2,500,000 students across all platforms. This course covers Go from fundamentals to advanced topics with multiple demo projects, including a complete REST API with authentication and SQL database.
- **What You'll Build:** REST API, Web Applications, Multiple Demo Projects
- **Course Link:** Search "Go The Complete Guide Maximilian" on Udemy

### 🥈 #2 - "Go: The Complete Developer's Guide (Golang)" by Stephen Grider
**⭐ Rating:** 4.6/5 (41,859+ ratings) | **👥 Students:** 194,226+
- **Duration:** 9 hours
- **Best For:** Those who want to understand Go's concurrency model deeply
- **Why Choose This:** Stephen is known for his ability to simplify complex topics. Strong focus on Go's unique features like goroutines, channels, and interfaces.
- **What You'll Build:** Card Deck Application, Web Status Checker
- **Course Link:** Search "Go Complete Developer Guide Stephen Grider" on Udemy

### 🥉 #3 - "Go Bootcamp: Master Golang with 1000+ Exercises" by Jose Portilla & Inanc Gumus
**⭐ Rating:** 4.5/5 (1,000+ ratings) | **👥 Students:** 7,000+
- **Duration:** 22 hours
- **Best For:** Those who want extensive hands-on practice
- **Why Choose This:** Includes 1000+ exercises and quizzes. Jose Portilla is renowned for his Python courses, bringing that same quality to Go.
- **What You'll Build:** Log Parser, File Scanner, Spam Masker, and more
- **Course Link:** Search "Go Bootcamp Master Golang 1000 Exercises" on Udemy

---

## 💡 HOW TO USE THIS PLAN WITH YOUR UDEMY COURSE

**IMPORTANT:** This plan provides a topic-based learning structure. Once you purchase your Udemy course, you'll need to map the course sections to the topics below. Here's how:

### First Day Setup (Before Day 1):
1. **Purchase your chosen Udemy course** (wait for sale - typically $12-15)
2. **Review the course curriculum** on Udemy
3. **Create your own mapping** - Write down which course sections cover which topics from this plan
4. **Adjust pacing** - Some sections may take more or less than the allocated time

### Recommended Approach:
- **For 16-hour courses (like Maximilian's):** Aim for ~10-12 minutes of video per day (30min session includes pausing, coding along, rewinding)
- **For 9-hour courses (like Stephen's):** Aim for ~6 minutes of video per day
- **Code along as you watch** - Don't just passively watch
- **Take notes** on key concepts

### Sample Daily Session Structure:
- **Minutes 0-5:** Review previous day's notes
- **Minutes 5-25:** Watch course videos (pause frequently to understand)
- **Minutes 25-30:** Note key concepts
- **Minutes 30-60:** Practice exercises (provided below for each day)

---

## Overview
- **Daily Time Commitment:** 1 hour (30 minutes theory video, 30 minutes practice)
- **Phase 1 (Basics):** ~30 days to functional proficiency
- **Phase 2 (Intermediate):** ~30 days to solid competency
- **Phase 3 (Advanced):** ~30 days to mastery
- **Total Timeline:** ~90 days (3 months) to master the language

---

## PHASE 1: FOUNDATIONS (Days 1-30)
**Goal:** Understand Go fundamentals and write basic programs

**Topics to Cover in Your Udemy Course:**
- Getting Started & Installation
- Go Basics (Variables, Types, Functions)
- Control Flow (If/Else, Switch, Loops)
- Data Structures (Arrays, Slices, Maps)
- Structs & Methods
- Pointers
- Interfaces
- Error Handling
- Packages & Modules
- Concurrency Basics (Goroutines & Channels)

---

### Week 1: Getting Started & Go Basics (Days 1-7)

#### Day 1: Introduction & Environment Setup
- **Theory (30min) - Udemy Topics:**
  - Course introduction and Go overview
  - Installing Go and setting up your IDE
  - Your first Go program ("Hello World")
  - Understanding `go run`, `go build`
  
- **Practice (30min):**
  - Install Go from golang.org
  - Set up VS Code with Go extension
  - Create "Hello World" program
  - Run it with `go run main.go`
  - Build it with `go build`
  - Create a GitHub repo called "learning-go"

**✅ Self-Check:** Can you run a Go program on your machine?

---

#### Day 2: Variables, Types & Constants
- **Theory (30min) - Udemy Topics:**
  - Variables and declarations (`var`, `:=`)
  - Basic types (int, float64, string, bool)
  - Type inference
  - Constants
  - Zero values
  
- **Practice (30min):**
  ```go
  // Exercise 1: Create a program that:
  - Declares variables of different types
  - Uses both var and := syntax
  - Declares constants
  - Prints all values with their types
  
  // Exercise 2: Build a simple calculator
  - Takes two numbers
  - Performs add, subtract, multiply, divide
  - Prints results
  ```

**✅ Self-Check:** Can you explain the difference between `var` and `:=`?

---

#### Day 3: Functions
- **Theory (30min) - Udemy Topics:**
  - Function syntax and declarations
  - Parameters and arguments
  - Return values (including multiple returns)
  - Named return values
  - Function scope
  
- **Practice (30min):**
  ```go
  // Exercise 1: Write functions that:
  - Calculate factorial of a number
  - Check if a number is prime
  - Return multiple values (quotient and remainder)
  
  // Exercise 2: Build a mini-math library
  - Create functions: power, square root, absolute value
  - Use multiple return values for error handling
  ```

**✅ Self-Check:** Can you write a function that returns two values?

---

#### Day 4: Control Flow - If/Else & Switch
- **Theory (30min) - Udemy Topics:**
  - If statements and if-else chains
  - Short statement with if
  - Switch statements
  - Switch without condition
  
- **Practice (30min):**
  ```go
  // Exercise 1: Number guessing game
  - Generate random number
  - Take user input
  - Give "higher" or "lower" hints
  - Congratulate when correct
  
  // Exercise 2: Grade calculator
  - Take numeric score as input
  - Convert to letter grade (A, B, C, D, F)
  - Use switch statement
  ```

**✅ Self-Check:** When would you use switch vs if-else?

---

#### Day 5: Loops
- **Theory (30min) - Udemy Topics:**
  - The `for` loop (Go's only loop)
  - While-style loops
  - Infinite loops
  - Break and continue
  - Nested loops
  
- **Practice (30min):**
  ```go
  // Exercise 1: FizzBuzz
  - Print numbers 1 to 100
  - "Fizz" for multiples of 3
  - "Buzz" for multiples of 5
  - "FizzBuzz" for multiples of both
  
  // Exercise 2: Print patterns
  - Right triangle of stars
  - Pyramid of stars
  - Diamond shape
  ```

**✅ Self-Check:** Can you create any loop style using just `for`?

---

#### Day 6: Arrays & Slices - Part 1
- **Theory (30min) - Udemy Topics:**
  - Arrays fundamentals
  - Array declaration and initialization
  - Introduction to slices
  - Differences between arrays and slices
  
- **Practice (30min):**
  ```go
  // Exercise 1: Array basics
  - Create an array of 5 integers
  - Initialize with values
  - Loop through and print
  - Find sum and average
  
  // Exercise 2: Slice basics
  - Create slices using different methods
  - Use make() to create slices
  - Append elements to slices
  - Print length and capacity
  ```

**✅ Self-Check:** Can you explain the difference between an array and a slice?

---

#### Day 7: Slices - Part 2 (Deep Dive)
- **Theory (30min) - Udemy Topics:**
  - Slice internals (length vs capacity)
  - Slicing slices
  - Copy function
  - Slice tricks and common patterns
  
- **Practice (30min):**
  ```go
  // Exercise 1: Slice operations
  - Reverse a slice without creating a new one
  - Remove an element at index i
  - Insert element at index i
  
  // Exercise 2: Mini-project
  - Build a To-Do List CLI
  - Add tasks (append to slice)
  - List all tasks (iterate)
  - Mark complete by index (remove from slice)
  ```

**✅ Self-Check:** Can you manipulate slices confidently?

---

### Week 2: Data Structures & Custom Types (Days 8-14)

#### Day 8: Maps
- **Theory (30min) - Udemy Topics:**
  - Map fundamentals
  - Creating maps
  - Adding, updating, deleting
  - Checking key existence
  - Iterating over maps
  
- **Practice (30min):**
  ```go
  // Exercise 1: Word frequency counter
  - Take a string input
  - Count how many times each word appears
  - Print results
  
  // Exercise 2: Phone book
  - Add contacts (name -> phone number)
  - Search for contacts
  - Delete contacts
  - List all contacts
  ```

**✅ Self-Check:** How do you check if a key exists in a map?

---

#### Day 9: Structs - Part 1
- **Theory (30min) - Udemy Topics:**
  - Struct basics
  - Defining custom types
  - Struct fields
  - Struct initialization
  - Anonymous structs
  
- **Practice (30min):**
  ```go
  // Exercise 1: Create Person struct
  type Person struct {
    Name string
    Age int
    Email string
  }
  - Create multiple persons
  - Initialize using different methods
  
  // Exercise 2: Product inventory
  - Create Product struct (name, price, quantity)
  - Create slice of products
  - Calculate total inventory value
  ```

**✅ Self-Check:** Can you create and initialize structs?

---

#### Day 10: Structs - Part 2 & Methods
- **Theory (30min) - Udemy Topics:**
  - Methods on structs
  - Receiver functions
  - Value receivers vs pointer receivers
  - When to use pointers
  
- **Practice (30min):**
  ```go
  // Exercise 1: Rectangle struct with methods
  type Rectangle struct {
    Width, Height float64
  }
  - Add Area() method
  - Add Perimeter() method
  - Add Scale(factor) method using pointer receiver
  
  // Exercise 2: Bank account
  - Create Account struct
  - Add Deposit() method
  - Add Withdraw() method (with validation)
  - Add GetBalance() method
  ```

**✅ Self-Check:** When should you use a pointer receiver vs value receiver?

---

#### Day 11: Pointers
- **Theory (30min) - Udemy Topics:**
  - What are pointers
  - The & (address) and * (dereference) operators
  - Pass by value vs pass by reference
  - Pointers to structs
  - Nil pointers
  
- **Practice (30min):**
  ```go
  // Exercise 1: Understand pointers
  - Create a variable and print its address
  - Create a pointer to that variable
  - Modify value through pointer
  - See the change in original variable
  
  // Exercise 2: Swap function
  - Write swap(a, b *int) to swap two integers
  - Write a function that modifies struct fields via pointer
  ```

**✅ Self-Check:** Can you explain when and why to use pointers?

---

#### Day 12: Packages & Code Organization
- **Theory (30min) - Udemy Topics:**
  - Understanding packages
  - Creating custom packages
  - Exported vs unexported identifiers
  - The main package
  - Importing packages
  
- **Practice (30min):**
  ```go
  // Exercise 1: Create a math utility package
  - Create folder "mathutils"
  - Add functions: Add, Subtract, Multiply, Divide
  - Export appropriate functions
  - Use from main package
  
  // Exercise 2: Create string utility package
  - Reverse string
  - Is palindrome
  - Count vowels
  - Use from main
  ```

**✅ Self-Check:** What makes an identifier exported vs unexported?

---

#### Day 13: Interfaces - Part 1
- **Theory (30min) - Udemy Topics:**
  - What are interfaces
  - Interface definition
  - Implicit implementation
  - Empty interface `interface{}`
  - Type assertions
  
- **Practice (30min):**
  ```go
  // Exercise 1: Shape interface
  type Shape interface {
    Area() float64
  }
  - Implement for Circle
  - Implement for Rectangle
  - Write function that accepts Shape
  
  // Exercise 2: Writer interface
  - Create interface with Write method
  - Implement for different writers
  - Use polymorphism
  ```

**✅ Self-Check:** How does Go know a type implements an interface?

---

#### Day 14: Interfaces - Part 2 & Review
- **Theory (30min) - Udemy Topics:**
  - Interface composition
  - Standard library interfaces (Stringer, Reader, Writer)
  - Type switches
  - Best practices
  
- **Practice (30min):**
  ```go
  // Mini-Project: Contact Manager CLI
  - Create Contact struct (name, phone, email)
  - Store contacts in slice
  - Functions: Add, List, Search, Delete
  - Save to/load from JSON file
  - Use interfaces where appropriate
  ```

**✅ Phase 1 Week 2 Complete! Review all concepts from Days 8-14.

---

### Week 3: File I/O, Error Handling & Testing (Days 15-21)

#### Day 15: Error Handling - Part 1
- **Theory (30min) - Udemy Topics:**
  - The error type
  - Returning errors
  - Creating errors with errors.New()
  - Error checking patterns
  - The if err != nil pattern
  
- **Practice (30min):**
  ```go
  // Exercise 1: Functions that return errors
  - Division function (error on divide by zero)
  - Square root function (error on negative)
  - Age validator (error on invalid age)
  
  // Exercise 2: Error handling chain
  - Function that calls multiple error-returning functions
  - Proper error propagation
  - Error messages with context
  ```

**✅ Self-Check:** Why does Go use explicit error returns instead of exceptions?

---

#### Day 16: Error Handling - Part 2 (Custom Errors)
- **Theory (30min) - Udemy Topics:**
  - Custom error types
  - Error wrapping with fmt.Errorf and %w
  - errors.Is() and errors.As()
  - Sentinel errors
  
- **Practice (30min):**
  ```go
  // Exercise 1: Custom error types
  type ValidationError struct {
    Field string
    Message string
  }
  - Implement Error() method
  - Use in validation functions
  
  // Exercise 2: Error wrapping
  - Create function chain
  - Wrap errors with context at each level
  - Use errors.Is to check error types
  ```

**✅ Self-Check:** When should you create custom error types?

---

#### Day 17: File I/O - Reading
- **Theory (30min) - Udemy Topics:**
  - Reading files with os.ReadFile
  - Opening files with os.Open
  - Using bufio.Scanner
  - Reading line by line
  - Closing files with defer
  
- **Practice (30min):**
  ```go
  // Exercise 1: File reader
  - Read entire file into string
  - Count lines in a file
  - Count words in a file
  
  // Exercise 2: Log file parser
  - Read log file line by line
  - Parse each line
  - Count error vs info messages
  - Handle file not found errors
  ```

**✅ Self-Check:** Why do we use `defer file.Close()`?

---

#### Day 18: File I/O - Writing & JSON
- **Theory (30min) - Udemy Topics:**
  - Writing files with os.WriteFile
  - Creating files with os.Create
  - JSON marshaling and unmarshaling
  - Struct tags for JSON
  
- **Practice (30min):**
  ```go
  // Exercise 1: Config file manager
  type Config struct {
    AppName string `json:"app_name"`
    Version string `json:"version"`
    Debug bool `json:"debug"`
  }
  - Marshal to JSON
  - Save to file
  - Read from file
  - Unmarshal back to struct
  
  // Exercise 2: CSV writer
  - Create simple CSV writer
  - Write data to CSV file
  ```

**✅ Self-Check:** How do struct tags work with JSON?

---

#### Day 19: Testing - Part 1
- **Theory (30min) - Udemy Topics:**
  - The testing package
  - Writing test functions
  - Running tests with go test
  - Test file naming (_test.go)
  - Basic assertions with t.Error()
  
- **Practice (30min):**
  ```go
  // Exercise 1: Write tests for math functions
  - Test addition function
  - Test division function (including error cases)
  - Test factorial function
  
  // Exercise 2: Table-driven tests
  - Create table of test cases
  - Loop through and test each
  - Test edge cases (0, negative, large numbers)
  ```

**✅ Self-Check:** What makes a good test?

---

#### Day 20: Testing - Part 2
- **Theory (30min) - Udemy Topics:**
  - Test coverage
  - Benchmarks
  - Test helpers
  - Testing best practices
  
- **Practice (30min):**
  ```go
  // Exercise 1: Test coverage
  - Run go test -cover
  - Aim for >80% coverage on previous code
  - Add missing test cases
  
  // Exercise 2: Benchmarks
  - Write benchmark for string operations
  - Write benchmark for slice operations
  - Compare performance of different implementations
  ```

**✅ Self-Check:** How do you check test coverage?

---

#### Day 21: Week 3 Project
- **Theory (30min):**
  - Review Days 15-20
  - Watch advanced tips from your Udemy course
  - Review error handling and testing best practices
  
- **Practice (30min):**
  ```go
  // Project: Note-Taking CLI Application
  - Create Note struct (id, title, content, created date)
  - CRUD operations (Create, Read, Update, Delete)
  - Save notes to JSON file
  - Load notes from JSON file
  - Proper error handling throughout
  - Write tests for all functions
  - Run tests with coverage report
  ```

**✅ Phase 1 Week 3 Complete! You now have solid Go fundamentals.

---

### Week 4: Concurrency (Days 22-30)

#### Day 22: Goroutines
- **Theory (30min) - Udemy Topics:**
  - What is concurrency
  - Goroutines basics
  - The `go` keyword
  - Concurrency vs parallelism
  
- **Practice (30min):**
  ```go
  // Exercise 1: Basic goroutines
  - Create a function that prints numbers
  - Run it as a goroutine
  - Use time.Sleep to see execution
  
  // Exercise 2: Multiple goroutines
  - Launch 10 goroutines
  - Each prints its ID
  - Observe execution order
  ```

**✅ Self-Check:** Why doesn't the program wait for goroutines to finish?

---

#### Day 23: WaitGroups
- **Theory (30min) - Udemy Topics:**
  - The sync.WaitGroup type
  - Add(), Done(), and Wait() methods
  - Coordinating goroutines
  
- **Practice (30min):**
  ```go
  // Exercise 1: Goroutines with WaitGroup
  - Launch multiple goroutines
  - Wait for all to complete
  - Track completion properly
  
  // Exercise 2: Concurrent URL checker
  - Check if 5 websites are up
  - Do it concurrently
  - Wait for all checks to complete
  - Print results
  ```

**✅ Self-Check:** What happens if you forget to call Done()?

---

#### Day 24: Channels - Part 1
- **Theory (30min) - Udemy Topics:**
  - What are channels
  - Creating channels
  - Sending and receiving
  - Channel blocking behavior
  - Buffered vs unbuffered channels
  
- **Practice (30min):**
  ```go
  // Exercise 1: Basic channel usage
  - Create a channel
  - Send value from one goroutine
  - Receive in main
  
  // Exercise 2: Producer-consumer
  - Producer goroutine sends numbers
  - Consumer goroutine receives and processes
  - Use channels for communication
  ```

**✅ Self-Check:** What's the difference between buffered and unbuffered channels?

---

#### Day 25: Channels - Part 2
- **Theory (30min) - Udemy Topics:**
  - Channel direction (send-only, receive-only)
  - Closing channels
  - Ranging over channels
  - The select statement
  
- **Practice (30min):**
  ```go
  // Exercise 1: Channel direction
  - Function with send-only channel parameter
  - Function with receive-only channel parameter
  
  // Exercise 2: Select statement
  - Use select with multiple channels
  - Implement timeout with select
  - Non-blocking receive
  ```

**✅ Self-Check:** When should you close a channel?

---

#### Day 26: Mutexes & Race Conditions
- **Theory (30min) - Udemy Topics:**
  - What are race conditions
  - sync.Mutex
  - Lock() and Unlock()
  - sync.RWMutex
  - When to use mutexes vs channels
  
- **Practice (30min):**
  ```go
  // Exercise 1: Race condition demo
  - Create a race condition
  - Run with go run -race
  - Fix with mutex
  
  // Exercise 2: Thread-safe counter
  - Counter accessed by multiple goroutines
  - Use mutex to protect
  - Compare performance with RWMutex
  ```

**✅ Self-Check:** Channels or mutexes - when to use which?

---

#### Day 27: Concurrency Patterns
- **Theory (30min) - Udemy Topics:**
  - Worker pool pattern
  - Pipeline pattern
  - Fan-out, fan-in
  - Context package basics
  
- **Practice (30min):**
  ```go
  // Exercise 1: Worker pool
  - Create pool of 5 workers
  - Send 20 jobs
  - Workers process jobs concurrently
  - Collect results
  
  // Exercise 2: Pipeline
  - Stage 1: Generate numbers
  - Stage 2: Square them
  - Stage 3: Sum them
  - Connect with channels
  ```

**✅ Self-Check:** Can you explain the worker pool pattern?

---

#### Days 28-29: Concurrency Project
- **Theory (30min each day):**
  - Review all concurrency concepts
  - Study advanced patterns from course
  - Watch supplementary: "Go Concurrency Patterns" (Rob Pike)
  
- **Practice (30min each day):**
  ```go
  // Project: Concurrent Web Scraper
  Day 28:
  - Design the scraper architecture
  - Create URL fetcher function
  - Create parser function
  - Set up worker pool (5 workers)
  
  Day 29:
  - Scrape 10+ web pages concurrently
  - Extract titles and links
  - Store results in thread-safe structure
  - Save results to JSON
  - Add proper error handling
  - Add timeout handling
  ```

**✅ Self-Check:** Does your scraper handle errors and timeouts properly?

---

#### Day 30: Phase 1 Review & Assessment
- **Theory (30min):**
  - Review entire Phase 1 (Days 1-30)
  - Identify weak areas
  - Watch recap videos from your course
  - Review your notes
  
- **Practice (30min):**
  - Refactor one of your early projects
  - Add tests to projects that lack them
  - Push all code to GitHub
  - Write a README for your learning journey
  
  **Self-Assessment Checklist:**
  - [ ] Can you explain variables, types, and functions?
  - [ ] Can you use slices, maps, and structs confidently?
  - [ ] Can you write and use interfaces?
  - [ ] Can you handle errors properly?
  - [ ] Can you read/write files and JSON?
  - [ ] Can you write tests?
  - [ ] Can you use goroutines and channels?
  - [ ] Can you avoid race conditions?

**✅ PHASE 1 COMPLETE! You now have solid Go fundamentals and can write concurrent programs.**

---

**Ready for Phase 2?** Continue to days 31-60 where you'll build REST APIs, work with databases, and create production-ready web applications!

---

## PHASE 2: INTERMEDIATE - WEB DEVELOPMENT & DATABASES (Days 31-60)
**Goal:** Build production-ready REST APIs and web applications

**Topics to Cover in Your Udemy Course:**
- HTTP basics and the net/http package
- Building REST APIs
- Routing and middleware
- Working with SQL databases
- Database ORMs (optional)
- Authentication and authorization (JWT)
- Project: Complete REST API with authentication

---

### Week 5: HTTP & REST API Basics (Days 31-37)

#### Day 31: HTTP Fundamentals
- **Theory (30min) - Udemy Topics:**
  - HTTP protocol basics
  - The `net/http` package
  - HTTP server creation
  - Handling requests and responses
  
- **Practice (30min):**
  ```go
  // Exercise 1: Basic HTTP server
  - Create server listening on port 8080
  - Handle root path "/"
  - Return "Hello, World!"
  - Test with browser or curl
  
  // Exercise 2: Multiple routes
  - Add /about route
  - Add /contact route
  - Return different responses
  ```

**✅ Self-Check:** Can you create and run an HTTP server?

---

#### Day 32: Request Handling
- **Theory (30min) - Udemy Topics:**
  - http.Request object
  - http.ResponseWriter interface
  - Query parameters
  - URL parameters
  - Request methods (GET, POST, etc.)
  
- **Practice (30min):**
  ```go
  // Exercise 1: Query parameters
  - Create /greet route
  - Read "name" query parameter
  - Return personalized greeting
  
  // Exercise 2: Different HTTP methods
  - Handle GET /users (list users)
  - Handle POST /users (create user)
  - Return appropriate responses
  ```

**✅ Self-Check:** Can you parse query parameters and handle different HTTP methods?

---

#### Day 33: JSON APIs
- **Theory (30min) - Udemy Topics:**
  - Working with JSON in HTTP
  - Encoding responses as JSON
  - Decoding JSON request bodies
  - Content-Type headers
  
- **Practice (30min):**
  ```go
  // Exercise 1: JSON response
  type User struct {
    ID int `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
  }
  - Return JSON response
  - Set proper Content-Type
  
  // Exercise 2: Parse JSON request
  - Accept POST with JSON body
  - Decode into struct
  - Validate and respond
  ```

**✅ Self-Check:** Can you send and receive JSON in your API?

---

#### Day 34: Routing Libraries
- **Theory (30min) - Udemy Topics:**
  - Limitations of stdlib router
  - Using gorilla/mux or chi router
  - Route parameters
  - Route groups
  
- **Practice (30min):**
  ```go
  // Exercise 1: Install router
  - go get github.com/gorilla/mux (or chi)
  - Create router
  - Define routes with parameters
  
  // Exercise 2: RESTful routes
  - GET /users -> list all
  - GET /users/{id} -> get one
  - POST /users -> create
  - PUT /users/{id} -> update
  - DELETE /users/{id} -> delete
  ```

**✅ Self-Check:** Can you use route parameters like /users/{id}?

---

#### Day 35: Middleware - Part 1
- **Theory (30min) - Udemy Topics:**
  - What is middleware
  - Middleware pattern in Go
  - Creating custom middleware
  - Chaining middleware
  
- **Practice (30min):**
  ```go
  // Exercise 1: Logging middleware
  - Log every request (method, path, time)
  - Wrap your handlers
  
  // Exercise 2: Recovery middleware
  - Catch panics
  - Log error
  - Return 500 response
  ```

**✅ Self-Check:** Can you explain how middleware works?

---

#### Day 36: Middleware - Part 2
- **Theory (30min) - Udemy Topics:**
  - Common middleware patterns
  - CORS middleware
  - Authentication middleware (preview)
  - Request ID middleware
  
- **Practice (30min):**
  ```go
  // Exercise 1: CORS middleware
  - Add CORS headers
  - Handle preflight requests
  
  // Exercise 2: Chain multiple middleware
  - Logging + Recovery + CORS
  - Apply to all routes
  - Test the chain
  ```

**✅ Self-Check:** Can you chain multiple middleware functions?

---

#### Day 37: Week 5 Mini-Project
- **Theory (30min):**
  - Review Days 31-36
  - API design best practices
  - Status codes
  
- **Practice (30min):**
  ```go
  // Project: Simple Blog API
  type Post struct {
    ID int
    Title string
    Content string
    Author string
  }
  - Store posts in memory (slice)
  - CRUD endpoints
  - Use router and middleware
  - Return proper JSON responses
  - Use appropriate status codes
  ```

**✅ Week 5 Complete! You can build basic REST APIs.

---

### Week 6-7: Databases (Days 38-51)

#### Day 38: SQL Basics in Go
- **Theory (30min) - Udemy Topics:**
  - database/sql package
  - SQL drivers
  - Connecting to PostgreSQL/MySQL
  - Connection strings
  
- **Practice (30min):**
  ```go
  // Exercise 1: Database setup
  - Install PostgreSQL or MySQL locally
  - Install Go driver: go get github.com/lib/pq
  - Create database and table
  
  // Exercise 2: Connect from Go
  - Import database/sql
  - Open connection
  - Ping database
  - Handle errors
  ```

**✅ Self-Check:** Can you connect to a database from Go?

---

#### Day 39: CRUD Operations - Part 1 (Create & Read)
- **Theory (30min) - Udemy Topics:**
  - Executing queries
  - db.Query() vs db.QueryRow()
  - Scanning results
  - Handling NULL values
  
- **Practice (30min):**
  ```go
  // Exercise 1: INSERT operation
  - Create function to insert user
  - Use db.Exec()
  - Return created ID
  
  // Exercise 2: SELECT operations
  - Get single user by ID
  - Get all users
  - Scan rows into structs
  ```

**✅ Self-Check:** Can you insert and query data?

---

#### Day 40: CRUD Operations - Part 2 (Update & Delete)
- **Theory (30min) - Udemy Topics:**
  - UPDATE queries
  - DELETE queries
  - Checking affected rows
  - Prepared statements
  
- **Practice (30min):**
  ```go
  // Exercise 1: UPDATE
  - Update user by ID
  - Check rows affected
  - Return error if not found
  
  // Exercise 2: DELETE
  - Delete user by ID
  - Check rows affected
  - Soft delete vs hard delete
  ```

**✅ Self-Check:** Can you perform all CRUD operations?

---

#### Day 41: Prepared Statements & Transactions
- **Theory (30min) - Udemy Topics:**
  - Why use prepared statements
  - SQL injection prevention
  - Transactions in Go
  - Commit and Rollback
  
- **Practice (30min):**
  ```go
  // Exercise 1: Prepared statements
  - Convert queries to use placeholders
  - Prepare statement
  - Execute multiple times
  
  // Exercise 2: Transactions
  - Begin transaction
  - Execute multiple operations
  - Rollback on error
  - Commit on success
  ```

**✅ Self-Check:** Why are prepared statements important?

---

#### Day 42: Database Best Practices
- **Theory (30min) - Udemy Topics:**
  - Connection pooling
  - Context with database operations
  - Error handling patterns
  - Database configuration
  
- **Practice (30min):**
  ```go
  // Exercise 1: Connection pooling
  - Configure max open connections
  - Configure max idle connections
  - Set connection max lifetime
  
  // Exercise 2: Context usage
  - Add context to database functions
  - Implement timeouts
  - Handle context cancellation
  ```

**✅ Self-Check:** How does connection pooling work?

---

#### Days 43-44: Database Repository Pattern
- **Theory (30min) - Udemy Topics:**
  - Repository pattern
  - Separating database logic
  - Interface-based design
  
- **Practice (30min each day):**
  ```go
  Day 43:
  // Create Repository interface and implementation
  type UserRepository interface {
    Create(user *User) error
    GetByID(id int) (*User, error)
    GetAll() ([]User, error)
    Update(user *User) error
    Delete(id int) error
  }
  
  Day 44:
  // Implement repository
  - Create PostgreSQLUserRepository
  - Implement all interface methods
  - Use in your API handlers
  ```

**✅ Self-Check:** What are the benefits of the repository pattern?

---

#### Days 45-47: ORM Basics (Optional but Recommended)
- **Theory (30min) - Udemy Topics:**
  - What is an ORM
  - Introduction to GORM
  - Models and migrations
  - Basic CRUD with GORM
  
- **Practice (30min each day):**
  ```go
  Day 45:
  - Install GORM: go get -u gorm.io/gorm
  - Define models with GORM tags
  - Auto-migration
  
  Day 46:
  - CRUD operations with GORM
  - Create, First, Find, Save, Delete
  - Compare with raw SQL
  
  Day 47:
  - Relationships (belongs to, has many)
  - Eager loading
  - Query building
  ```

**✅ Self-Check:** When would you use an ORM vs raw SQL?

---

#### Days 48-51: Integrate Database with API
- **Theory (30min) - Udemy Topics:**
  - Connecting API to database
  - Dependency injection
  - Configuration management
  - Environment variables
  
- **Practice (30min each day):**
  ```go
  Day 48:
  // Refactor Blog API to use database
  - Create posts table
  - Implement repository
  - Update handlers to use DB
  
  Day 49:
  // Add more features
  - Add users table
  - User authentication (basic)
  - Associate posts with users
  
  Day 50:
  // Error handling & validation
  - Validate input
  - Return proper errors
  - Handle database errors gracefully
  
  Day 51:
  // Configuration
  - Read database config from env
  - Use godotenv or similar
  - Different configs for dev/prod
  ```

**✅ Self-Check:** Can you build a database-backed API?

---

### Week 8: Authentication & Major Project (Days 52-60)

#### Days 52-53: Authentication - JWT Basics
- **Theory (30min) - Udemy Topics:**
  - What is JWT
  - JWT structure (header, payload, signature)
  - Installing JWT library
  - Token generation and validation
  
- **Practice (30min each day):**
  ```go
  Day 52:
  - Install: go get github.com/golang-jwt/jwt/v5
  - Generate JWT token
  - Sign with secret
  - Set expiration
  
  Day 53:
  - Validate JWT token
  - Extract claims
  - Handle expired tokens
  - Refresh tokens (optional)
  ```

**✅ Self-Check:** What are the three parts of a JWT?

---

#### Days 54-55: Authentication Implementation
- **Theory (30min) - Udemy Topics:**
  - Password hashing with bcrypt
  - User registration
  - User login
  - Protected routes
  
- **Practice (30min each day):**
  ```go
  Day 54:
  // User registration and login
  - Hash passwords with bcrypt
  - Store users in database
  - POST /register endpoint
  - POST /login endpoint (returns JWT)
  
  Day 55:
  // Protected routes
  - Create auth middleware
  - Verify JWT from header
  - Attach user to context
  - Protect specific routes
  ```

**✅ Self-Check:** Can you implement user authentication?

---

#### Days 56-60: MAJOR PROJECT - Complete REST API
- **Theory (30min each day):**
  - Review all Phase 2 concepts
  - API design principles
  - Security best practices
  - Testing strategies
  
- **Practice (30min each day):**
  ```go
  // Build: Task Management API
  
  Day 56: Project Setup & Models
  - Define models: User, Task, Category
  - Database schema
  - Migrations
  - Repository interfaces
  
  Day 57: User Auth
  - Registration endpoint
  - Login endpoint
  - Password hashing
  - JWT generation
  
  Day 58: Task CRUD
  - Create task (auth required)
  - List user's tasks (auth required)
  - Update task (owner only)
  - Delete task (owner only)
  - Filter by category/status
  
  Day 59: Additional Features
  - Categories CRUD
  - Mark task complete
  - Due dates
  - Pagination
  
  Day 60: Testing & Polish
  - Write tests for all endpoints
  - Error handling
  - Input validation
  - API documentation (README)
  - Push to GitHub
  ```

**✅ PHASE 2 COMPLETE!** You can now build production-ready REST APIs with authentication and databases!

---

**Self-Assessment Checklist:**
- [ ] Can you build REST APIs with proper routing?
- [ ] Can you work with SQL databases in Go?
- [ ] Can you implement authentication with JWT?
- [ ] Can you hash passwords securely?
- [ ] Can you structure an API project properly?
- [ ] Can you write middleware?
- [ ] Can you use environment variables for config?
- [ ] Can you test your API endpoints?

**Ready for Phase 3?** Continue to days 61-90 for advanced topics, microservices, and production deployment!

---

## PHASE 3: ADVANCED - PRODUCTION & MASTERY (Days 61-90)
**Goal:** Master advanced topics and build production-ready systems

**Topics for Advanced Learning:**
- Advanced concurrency patterns
- Performance optimization and profiling
- Microservices and gRPC
- Docker and deployment
- CI/CD basics
- Security best practices
- Design patterns in Go
- Production monitoring and logging

**Note:** Phase 3 requires supplementary resources beyond your main Udemy course. Use additional videos, documentation, and tutorials.

---

### Week 9: Advanced Concurrency & Performance (Days 61-67)

#### Days 61-62: Advanced Concurrency Patterns
- **Theory (30min):**
  - Context package deep dive
  - Cancellation and timeouts
  - Worker pools advanced
  - Rate limiting patterns
  
- **Practice (30min):**
  ```go
  Day 61:
  // Context patterns
  - Context with cancel
  - Context with deadline
  - Context with timeout
  - Context with values
  
  Day 62:
  // Rate limiter
  - Implement token bucket algorithm
  - Use golang.org/x/time/rate
  - Rate limit API endpoints
  ```

**✅ Self-Check:** When would you use context.WithCancel vs context.WithTimeout?

---

#### Days 63-64: Performance & Profiling
- **Theory (30min):**
  - Performance profiling with pprof
  - CPU profiling
  - Memory profiling
  - Benchmarking best practices
  
- **Practice (30min):**
  ```go
  Day 63:
  // CPU Profiling
  - Add pprof to your API
  - Generate CPU profile
  - Analyze with go tool pprof
  - Find bottlenecks
  
  Day 64:
  // Benchmarking
  - Write benchmarks for critical functions
  - Use testing.B
  - Compare different implementations
  - Optimize based on results
  ```

**✅ Self-Check:** Can you profile and optimize Go applications?

---

#### Days 65-67: Optimization Techniques
- **Theory (30min):**
  - Memory optimization
  - Sync.Pool for object reuse
  - String optimization
  - Slice and map optimization
  
- **Practice (30min):**
  ```go
  Day 65:
  // Memory optimization
  - Use sync.Pool
  - Avoid unnecessary allocations
  - Reuse buffers
  
  Day 66:
  // String optimization
  - strings.Builder vs concatenation
  - Byte slices vs strings
  - Benchmark comparisons
  
  Day 67:
  // Data structure optimization
  - Pre-allocate slices when size known
  - Choose right data structure
  - Profile and optimize your REST API
  ```

**✅ Self-Check:** What optimization techniques do you know?

---

### Week 10: Microservices & gRPC (Days 68-74)

#### Days 68-69: gRPC Basics
- **Theory (30min):**
  - What is gRPC
  - Protocol Buffers (protobuf)
  - gRPC vs REST
  - When to use gRPC
  
- **Practice (30min):**
  ```go
  Day 68:
  // Setup
  - Install protoc compiler
  - Install Go protobuf plugin
  - Create simple .proto file
  - Generate Go code
  
  Day 69:
  // First gRPC service
  - Define service in proto
  - Implement server
  - Implement client
  - Test communication
  ```

**✅ Self-Check:** What are the advantages of gRPC over REST?

---

#### Days 70-72: Building Microservices
- **Theory (30min):**
  - Microservices architecture
  - Service communication
  - Service discovery
  - API Gateway pattern
  
- **Practice (30min):**
  ```go
  Day 70:
  // Create first microservice
  - User service (gRPC)
  - Define proto
  - Implement CRUD operations
  
  Day 71:
  // Create second microservice
  - Product service (gRPC)
  - Communication with user service
  - Handle errors between services
  
  Day 72:
  // API Gateway
  - REST API gateway
  - Routes to gRPC services
  - Convert REST to gRPC calls
  ```

**✅ Self-Check:** Can you explain microservices architecture?

---

#### Days 73-74: Service Communication Patterns
- **Theory (30min):**
  - Synchronous vs asynchronous communication
  - Message queues (intro to RabbitMQ/Kafka)
  - Circuit breaker pattern
  - Retry strategies
  
- **Practice (30min):**
  ```go
  Day 73:
  // Circuit breaker
  - Implement basic circuit breaker
  - Add to service calls
  - Handle failures gracefully
  
  Day 74:
  // Retry with backoff
  - Exponential backoff
  - Maximum retry attempts
  - Timeout handling
  ```

**✅ Self-Check:** What is a circuit breaker and why use it?

---

### Week 11: Production Ready (Days 75-81)

#### Days 75-76: Logging & Monitoring
- **Theory (30min):**
  - Structured logging
  - Log levels
  - slog package (Go 1.21+)
  - Centralized logging
  
- **Practice (30min):**
  ```go
  Day 75:
  // Structured logging
  - Implement slog or zap
  - Add to all handlers
  - Log with context
  - Different log levels
  
  Day 76:
  // Request tracing
  - Add request IDs
  - Log request/response
  - Track performance
  ```

**✅ Self-Check:** What's the difference between structured and unstructured logging?

---

#### Days 77-78: Docker & Containerization
- **Theory (30min):**
  - Docker basics
  - Creating Dockerfiles for Go
  - Multi-stage builds
  - Docker Compose
  
- **Practice (30min):**
  ```go
  Day 77:
  // Dockerfile
  - Create multi-stage Dockerfile
  - Build small images (alpine)
  - Run container locally
  
  Day 78:
  // Docker Compose
  - Define services (API + DB)
  - Set up networking
  - Environment variables
  - Run with docker-compose up
  ```

**✅ Self-Check:** What are multi-stage Docker builds?

---

#### Days 79-80: CI/CD Basics
- **Theory (30min):**
  - CI/CD concepts
  - GitHub Actions
  - Automated testing
  - Automated deployment
  
- **Practice (30min):**
  ```go
  Day 79:
  // GitHub Actions
  - Create .github/workflows/test.yml
  - Run tests on push
  - Check code coverage
  
  Day 80:
  // Build and push Docker
  - Add build step
  - Push to Docker Hub
  - Basic deployment automation
  ```

**✅ Self-Check:** What is CI/CD?

---

#### Day 81: Security Best Practices
- **Theory (30min):**
  - Input validation
  - SQL injection prevention
  - XSS prevention
  - HTTPS/TLS
  - Security headers
  
- **Practice (30min):**
  ```go
  // Security audit
  - Add input validation everywhere
  - Review use of prepared statements
  - Add security headers
  - Set up HTTPS locally
  - Scan for vulnerabilities
  ```

**✅ Self-Check:** What are the OWASP Top 10?

---

### Week 12: FINAL PROJECT (Days 82-90)

#### Days 82-85: Project Planning & Architecture
- **Theory (30min each day):**
  - System design principles
  - Database schema design
  - API design
  - Deployment strategy
  
- **Practice (30min each day):**
  ```go
  // Plan: E-Commerce Microservices Platform
  
  Day 82: Architecture Design
  - Service breakdown (User, Product, Order, Payment)
  - Database per service
  - Communication patterns
  - Draw architecture diagram
  
  Day 83: Database Design
  - Design schemas for each service
  - Define relationships
  - Plan migrations
  
  Day 84: API Design
  - Design REST endpoints
  - Design gRPC services
  - Plan authentication flow
  
  Day 85: Setup Project
  - Create repository structure
  - Set up microservices
  - Initialize databases
  - Configure Docker
  ```

---

#### Days 86-89: Implementation
- **Theory (30min each day):**
  - Review relevant concepts
  - Best practices for each component
  
- **Practice (30min each day):**
  ```go
  Day 86: User Service
  - Implement authentication
  - User CRUD
  - JWT generation
  - PostgreSQL integration
  
  Day 87: Product Service
  - Product CRUD (gRPC)
  - Category management
  - Search functionality
  
  Day 88: Order Service
  - Create orders
  - Order history
  - Integration with User & Product services
  - Handle transactions
  
  Day 89: API Gateway & Polish
  - REST gateway for all services
  - Error handling
  - Validation
  - Documentation
  ```

---

#### Day 90: Deployment & Final Review
- **Theory (30min):**
  - Review entire learning journey
  - Areas for continued learning
  - Next steps in Go career
  
- **Practice (30min):**
  ```go
  // Final Steps:
  - Write comprehensive README
  - Add API documentation
  - Docker Compose for entire system
  - Push to GitHub
  - Deploy to cloud (Railway/Render/Heroku free tier)
  - Write blog post about your journey
  - Update LinkedIn
  - Celebrate! 🎉
  ```

**✅ PHASE 3 COMPLETE! You've mastered Go!**

---

## 🎯 Post-90 Day Action Plan

### Continue Learning:
- **Read production Go codebases:**
  - Kubernetes
  - Docker
  - Terraform
  - Hugo
  
- **Contribute to open source:**
  - Find Go projects on GitHub
  - Start with "good first issue" labels
  - Submit PRs
  
- **Advanced topics to explore:**
  - WebAssembly with Go
  - Distributed systems
  - Advanced database patterns
  - Cloud-native development
  - GraphQL in Go

### Career Steps:
- [ ] Build 3-5 projects and add to GitHub
- [ ] Create a technical blog
- [ ] Update resume with Go skills
- [ ] Apply for Go developer positions
- [ ] Contribute to Go open source
- [ ] Join Go communities actively
- [ ] Consider Go certification

---

## 🏆 Skills Mastered After 90 Days

**Core Go:**
✅ Variables, types, functions, control flow
✅ Slices, maps, structs, interfaces
✅ Pointers and memory management
✅ Error handling patterns
✅ Package organization
✅ Testing and benchmarking

**Concurrency:**
✅ Goroutines and channels
✅ Sync primitives (Mutex, WaitGroup)
✅ Concurrency patterns
✅ Context package
✅ Race condition prevention

**Web Development:**
✅ HTTP servers and routing
✅ REST API development
✅ Middleware patterns
✅ JSON handling
✅ Authentication (JWT)

**Databases:**
✅ SQL with database/sql
✅ ORMs (GORM)
✅ Transactions
✅ Repository pattern
✅ Database optimization

**Production Skills:**
✅ Docker containerization
✅ CI/CD basics
✅ Logging and monitoring
✅ Performance profiling
✅ Security best practices

**Advanced Topics:**
✅ gRPC and Protocol Buffers
✅ Microservices architecture
✅ Deployment strategies
✅ Design patterns in Go

---

## 📊 Estimated Skill Level After Each Phase

**After Day 30 (Phase 1):**
- Junior Go Developer level
- Can build CLI tools
- Understand Go fundamentals
- Can write concurrent programs

**After Day 60 (Phase 2):**
- Mid-level Go Developer level
- Can build production REST APIs
- Database integration skills
- Ready for junior backend positions

**After Day 90 (Phase 3):**
- Senior Junior / Entry Mid-level
- Can build microservices
- Production deployment skills
- Ready for mid-level positions with mentorship

---

**Congratulations on completing this journey!** Remember: this is just the beginning. Go is a language you'll continue to master throughout your career. Keep building, keep learning, and most importantly, keep coding! 🚀

---

## 📚 Primary Learning Resources

### 🎓 MAIN UDEMY COURSE (Foundation of This Plan)
**"Go - The Complete Guide" by Maximilian Schwarzmüller**
- **Where:** Udemy (search for the course title)
- **Duration:** 16 hours of video
- **Best For:** Your primary learning source - covers everything systematically
- **Price:** Usually $12-15 during Udemy sales (regularly discounted from $119.99)
- **What You Get:** Complete Go fundamentals, concurrency, REST API project, databases, authentication

### 🌟 Supplementary Udemy Courses (Optional but Recommended)
1. **"Go: The Complete Developer's Guide" by Stephen Grider**
   - Excellent for deeper concurrency understanding
   - Great complement to Maximilian's course
   - Focus on goroutines and channels

2. **"Go Bootcamp: Master Golang with 1000+ Exercises"** by Jose Portilla & Inanc Gumus
   - Perfect for extra practice
   - 1000+ coding exercises
   - When you want more hands-on work

### 📺 Free YouTube Channels (Supplementary)
Use these to fill gaps or get different perspectives:

1. **freeCodeCamp** - "Learn Go Programming" (7+ hour course)
   - Great free alternative
   - Good for reinforcement

2. **Tech With Tim** - Go Tutorial Series
   - Short, focused videos
   - Good for specific topics

3. **Anthony GG** - Go Programming Channel
   - Intermediate to advanced
   - Real-world scenarios

4. **Traversy Media** - "Go Crash Course"
   - 2-hour overview
   - Good for quick review

5. **JustForFunc** by Francesc Campoy
   - Deep dives into Go concepts
   - Advanced topics

### 📖 Essential Supplementary Reading
- **Official Go Tour** (go.dev/tour) - Interactive learning
- **Go By Example** (gobyexample.com) - Quick reference with examples
- **Effective Go** (go.dev/doc/effective_go) - Official style guide
- **The Go Blog** (blog.golang.org) - Latest features and best practices

### 📚 Recommended Books (Optional)
- **"Learning Go" by Jon Bodner** - Excellent comprehensive book
- **"Concurrency in Go" by Katherine Cox-Buday** - Deep dive into concurrency
- **"The Go Programming Language"** by Donovan & Kernighan - The definitive reference

### 🎥 Advanced Conference Talks (Phase 3)
- **GopherCon Talks** (YouTube) - Annual Go conference
- **Google I/O Talks** - Rob Pike's concurrency patterns
- **dotGo Conference** - European Go conference talks

---

## 🎯 Daily Practice Resources

1. **Exercism.org** - Go track with mentor feedback
2. **LeetCode** - Practice algorithms in Go
3. **HackerRank** - Go challenges
4. **Codewars** - Go katas
5. **GitHub** - Read popular Go projects source code

---

## ✅ Milestones & Self-Assessment

### After 30 Days (Basics Mastered)
You should be able to:
- ✓ Write command-line programs
- ✓ Understand all basic syntax
- ✓ Use goroutines and channels
- ✓ Handle errors properly
- ✓ Work with files and JSON
- ✓ Write unit tests

### After 60 Days (Intermediate)
You should be able to:
- ✓ Build REST APIs
- ✓ Work with databases
- ✓ Create production-quality code
- ✓ Understand design patterns
- ✓ Build CLI tools

### After 90 Days (Advanced/Mastery)
You should be able to:
- ✓ Design microservices
- ✓ Optimize for performance
- ✓ Use advanced concurrency patterns
- ✓ Deploy production applications
- ✓ Contribute to open source Go projects
- ✓ **START APPLYING FOR GO DEVELOPER POSITIONS**

---

## 💡 Daily Routine Tips

1. **Stick to the schedule:** 30min theory, 30min practice, same time daily
2. **Code along with videos:** Don't just watch passively - pause and type the code yourself
3. **Take notes:** Keep a "Go learning journal" with key concepts and gotchas
4. **Push to GitHub daily:** Track your progress publicly, build your portfolio
5. **Join Go communities:** 
   - r/golang (Reddit)
   - Gophers Slack (invite.slack.golangbridge.org)
   - Go Discord servers
6. **Review on weekends:** Optional extra time for review/catching up or getting ahead
7. **Don't skip practice:** The hands-on work is where real learning happens
8. **Ask questions:** Use the Udemy Q&A section - Maximilian and his team respond actively

---

## 💰 UDEMY COURSE BUYING TIPS

**NEVER pay full price for Udemy courses!** Here's how to get the best deals:

### When to Buy:
- **New Year Sales** (January) - Usually 80-90% off
- **Black Friday/Cyber Monday** (November) - Best prices of the year
- **Monthly Flash Sales** - First week of every month
- **New User Discount** - Create a new account for first-time buyer discount
- **Regular Sales** - Udemy has sales almost every week

### Expected Prices:
- **Regular Price:** $119.99 (DON'T PAY THIS!)
- **Sale Price:** $12.99 - $14.99 (WAIT FOR THIS)
- **Best Price:** $9.99 (rare but happens)

### How to Get Discounts:
1. **Wait for sales:** Add course to cart, wait a few days - usually a sale appears
2. **Use incognito mode:** Check prices in incognito/private browsing
3. **Check different devices:** Sometimes mobile has different prices
4. **Course creator links:** Sometimes instructors share discount links on Twitter/LinkedIn
5. **Browse from different countries:** VPN can sometimes show different prices (use ethically)

### Udemy Personal Plan:
- **Cost:** $30/month (often $20 during promotions)
- **Value:** Access to 11,000+ courses
- **Worth it if:** You want to take multiple courses (Go + React + Python, etc.)
- **Not worth it if:** You only want to learn Go

### Free Alternatives:
If you can't afford Udemy courses right now:
- **freeCodeCamp YouTube** - 7-hour Go course (free)
- **Official Go Tour** - Interactive (free)
- **Go By Example** - Excellent reference (free)
- **Exercism.org** - Practice with mentors (free)

---

## 🚀 Beyond Day 90

After completing this plan:
- Contribute to open-source Go projects
- Build your own substantial project
- Read production Go codebases (Kubernetes, Docker, Terraform)
- Take the "Go Developer" certification
- Start applying for Go positions or freelancing

---

## Summary: Timeline to Proficiency

- **Basic Competency:** 30 days - You can write functional Go programs
- **Job-Ready Skills:** 60 days - You can build real applications (REST APIs, web services)
- **Language Mastery:** 90 days - You understand Go deeply and can build production systems

**Total Time Investment:** 90 hours of focused learning = 3 months of daily 1-hour study

---

## 🎯 Course Integration Notes

This plan is specifically structured around **Maximilian Schwarzmüller's "Go - The Complete Guide"** course, which perfectly covers all the fundamentals you need in Phase 1 and 2. Here's how they align:

- **Days 1-30 (Phase 1):** Corresponds to Course Sections 1-9
- **Days 31-60 (Phase 2):** Corresponds to Course Sections 10-15 + Major REST API Project
- **Days 61-90 (Phase 3):** Advanced topics using supplementary resources + building capstone projects

The beauty of Maximilian's course is that it's:
- ✅ Well-structured and progressive
- ✅ Project-based (learn by building)
- ✅ Updated regularly (as of 2025)
- ✅ Includes a complete REST API project
- ✅ Has active Q&A support

**Alternative approach:** If you prefer Stephen Grider's teaching style, his course is also excellent. The plan structure remains the same - just substitute the course sections accordingly.

---

## 🚀 Getting Started Checklist

Before Day 1, make sure you have:
- [ ] Purchased the Udemy course (wait for a sale!)
- [ ] Installed Go from golang.org
- [ ] Set up VS Code with Go extension
- [ ] Created a GitHub account (for pushing code daily)
- [ ] Joined r/golang community
- [ ] Set a daily alarm/reminder for your 1-hour study session
- [ ] Created a dedicated "learning-go" folder on your computer
- [ ] Prepared a notebook or digital note-taking system

---

**Remember:** Consistency beats intensity. One hour every day beats 7 hours once a week. Show up daily, code along with the videos, and build those projects. You've got this! 🎯

Good luck on your Go journey!