package main

type updateError struct {
	exitCode uint32
	stdOut   string
	stdErr   string
}
