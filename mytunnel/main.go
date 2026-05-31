package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/dpkrn/gotunnel/pkg/tunnel"
)

func printHelp() {
	fmt.Println("mytunnel — expose your local server to the internet")
	fmt.Println()
	fmt.Println("Usage:")
	fmt.Println("  mytunnel http <port> [flags]")
	fmt.Println("  mytunnel help")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  http <port>   Forward HTTP traffic to localhost:<port>")
	fmt.Println("  help          Show this help message")
	fmt.Println()
	fmt.Println("Flags (http):")
	fmt.Println("  -i, -inspector string   Traffic inspector listen port (default 4040)")
	fmt.Println("  -no-inspector            Disable the traffic inspector")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  mytunnel http 3000")
	fmt.Println("  mytunnel http 8080 -i 4040")
	fmt.Println("  mytunnel http 8080 -inspector 9090")
	fmt.Println("  mytunnel http 8080 -no-inspector")
	fmt.Println()
	fmt.Println("Issues: https://github.com/dpkrn/gotunnel/issues")
}

func normalizeInspectorPort(s string) string {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, ":")
	if s == "" {
		return "4040"
	}
	return s
}

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	command := os.Args[1]

	if command == "help" || command == "--help" || command == "-h" {
		printHelp()
		return
	}

	if command != "http" {
		fmt.Println("Unknown command:", command)
		fmt.Println("Run 'mytunnel help' to see available commands.")
		os.Exit(1)
	}

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: mytunnel http <port> [flags]")
		os.Exit(1)
	}

	port := os.Args[2]

	fs := flag.NewFlagSet("http", flag.ExitOnError)
	fs.SetOutput(os.Stderr)

	var inspectorPort string
	fs.StringVar(&inspectorPort, "i", "4040", "traffic inspector listen `port`")
	fs.StringVar(&inspectorPort, "inspector", "4040", "alias for -i")

	noInspector := fs.Bool("no-inspector", false, "disable traffic inspector")

	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: mytunnel http <port> [flags]\n\n")
		fs.PrintDefaults()
	}

	if err := fs.Parse(os.Args[3:]); err != nil {
		os.Exit(1)
	}

	opts := tunnel.TunnelOptions{
		Inspector:     !*noInspector,
		InspectorAddr: normalizeInspectorPort(inspectorPort),
	}

	_, stop, err := tunnel.StartTunnel(port, opts)
	if err != nil {
		stop()
		fmt.Println("could not start tunnel", err)
		os.Exit(1)
	}
	defer stop()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
	fmt.Fprintln(os.Stderr, "mytunnel: shutting down…")
}
