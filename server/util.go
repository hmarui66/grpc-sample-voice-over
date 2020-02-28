package main

import (
	"io"
	"log"
)

func closeCloser(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Fatalf("failed to close: %v", err)
	}
}
