package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func(ctx context.Context) {
		conn, err := net.ListenPacket("udp", ":8080")
		if err != nil {
			log.Fatal(err)
		}

		defer func() {
			_ = conn.Close()
		}()

		buf := make([]byte, 1518)

		for {
			select {
			default:
				n, _, err := conn.ReadFrom(buf)
				if err == io.EOF {
					return
				}

				if err != nil {
					log.Fatal(err)
				}

				fmt.Println(string(buf[:n]))
			case <-ctx.Done():
				return
			}
		}
	}(ctx)

	conn, err := net.Dial("udp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = conn.Close()
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Text()
		if msg == "end" {
			return
		}

		_, _ = fmt.Fprint(conn, msg)
	}
	if err := scanner.Err(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
