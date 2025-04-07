# LockFree - Golang Package

## Overview

The `LockFree` package provides a simple, lock-free linked list implementation using atomic operations in Go. This package allows for adding elements to a linked list and printing the list's contents in a thread-safe manner without the need for locks.

The list is implemented using atomic pointer operations, which ensures that multiple goroutines can safely modify the list concurrently.

## Features

- **Add operation**: Efficiently adds a new element to the head of the list using atomic operations.
- **Print operation**: Safely traverses and prints the values in the list.
- **Lock-free**: Achieves concurrency safety without the use of locks, relying solely on atomic operations.

## Installation

To install this package, simply add it to your Go project:

```sh
go get github.com/yourusername/lockfree
```

## Usage

### Importing the Package

```go
import "github.com/yourusername/lockfree"
```


### Creating and Using a Lock-Free List

```go
package main

import (
	"fmt"
	"github.com/Egorekski/lockfree"
)

func main() {
	// Create a new lock-free list
	list := &lockfree.List{}

	// Add values to the list
	list.Add(1)
	list.Add(2)
	list.Add(3)

	// Print values from the list
	list.Print()
}
```


### Adding Elements to the List

You can add elements to the list using the `Add` method:

```go
list.Add(value int)
```

Where `value` is the integer you want to insert at the head of the list.


### Printing the List

To print the elements of the list, use the `Print` method:

```go
list.Print()
```

This will output the values in the list starting from the head.

### Example Output

```bash
3
2
1
```

## TODO

* Implement a `FindByIndex` function to allow finding elements by their index in the list.

## Concurrency and Safety

This package is designed to be safe for concurrent use. All modifications to the list (like adding a node) are performed atomically, which guarantees that no locks are needed for the operations. This ensures that the list can be safely accessed and modified by multiple goroutines concurrently.

## License

This package is licensed under the MIT License. See the LICENSE file for more details.

## Future Improvements

* Improve the functionality to allow node removal.
* Add support for other operations like finding nodes by index.
