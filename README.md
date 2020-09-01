# Hamlet Data Generator

Sample data generator.

## Usage

* `GetWords(n)` - gets n words in a string
* `GetWordArray(n, m)` - gets n strings with m words
* `GetObject()` - gets a book object
* `GetObjectArray(n)` - gets n book objects

## Sample

    package main

    import (
        "fmt"

        hamlet "github.com/davidbetz/hamletgo"
    )

    func main() {
        fmt.Printf("%v", hamlet.GetObject())
    }
