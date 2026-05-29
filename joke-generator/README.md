# Random Joke Generator

A simple Go application that fetches and displays random jokes using the [Official Joke API](https://official-joke-api.appspot.com/).

## Features

- 🎭 Fetch random jokes
- 📂 Fetch jokes by specific type (general, knock-knock, programming)
- 🎯 Simple and clean API
- 🧪 Comprehensive test coverage
- ⚡ Fast and lightweight

## Supported Joke Types

- `general` - General jokes
- `knock-knock` - Knock-knock jokes
- `programming` - Programming jokes

## Installation

```bash
cd joke-generator
go build
```

## Usage

### Run the application

```bash
./joke-generator
```

### Use as a library

```go
package main

import (
	"fmt"
	"log"
)

func main() {
	// Fetch a random joke
	joke, err := FetchRandomJoke()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Setup: %s\n", joke.Setup)
	fmt.Printf("Delivery: %s\n", joke.Delivery)

	// Fetch a joke by type
	generalJoke, err := FetchRandomJokeByType("general")
	if err != nil {
		log.Fatal(err)
	}

	PrintJoke(generalJoke)
}
```

## API Reference

### FetchRandomJoke()

Fetches a random joke from any category.

```go
joke, err := FetchRandomJoke()
if err != nil {
	log.Fatal(err)
}
PrintJoke(joke)
```

### FetchRandomJokeByType(jokeType string)

Fetches a random joke of a specific type.

```go
joke, err := FetchRandomJokeByType("programming")
if err != nil {
	log.Fatal(err)
}
```

### PrintJoke(joke *Joke)

Prints a joke in a formatted manner.

```go
PrintJoke(joke)
```

## Data Structure

```go
type Joke struct {
	ID       int    `json:"id"`
	Type     string `json:"type"`
	Setup    string `json:"setup"`
	Delivery string `json:"delivery"`
	Error    bool   `json:"error"`
}
```

## Testing

Run the tests:

```bash
go test ./...
```

Run tests with verbose output:

```bash
go test -v ./...
```

Run benchmarks:

```bash
go test -bench=. ./...
```

## Example Output

```
🎭 Random Joke Generator

Fetching a random joke...
------- 😂 Random Joke -------
Type: general
Setup: Why don't scientists trust atoms?
Delivery: Because they make up everything!
------------------------------

Fetching a general joke...
------- 😂 Random Joke -------
Type: general
Setup: What do you call a bear with no teeth?
Delivery: A gummy bear!
------------------------------

Fetching a knock-knock joke...
------- 😂 Random Joke -------
Type: knock-knock
Setup: Knock knock. Who's there? Interrupting cow.
Delivery: Interrupting cow w—MOOOOO!
------------------------------
```

## External API

This application uses the **Official Joke API**:

- **Base URL**: https://official-joke-api.appspot.com/
- **Endpoints**:
  - `/random_joke` - Get a random joke
  - `/jokes/{type}/random` - Get a random joke of a specific type

For more information, visit: https://official-joke-api.appspot.com/

## Error Handling

The application handles various error scenarios:

- Network connection errors
- API response errors
- JSON parsing errors
- Non-200 HTTP status codes

## License

MIT

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](../CONTRIBUTING.md) for guidelines.

## Troubleshooting

### "failed to fetch joke: dial tcp: lookup official-joke-api.appspot.com"

This error indicates a network connectivity issue. Check your internet connection and ensure the API is accessible.

### "API returned status code: 500"

The API server may be temporarily unavailable. Try again later.

### "failed to unmarshal JSON"

This could indicate an API response format change. Check the API documentation for updates.
