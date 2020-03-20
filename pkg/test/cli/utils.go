package cli

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"sync"

	stdlog "log"

	"github.com/rs/zerolog/log"

	shellwords "github.com/mattn/go-shellwords"
)

func CaptureOutput(f func()) string {
	// Setup custom writer to use as Stdout/Stderr
	reader, writer, err := os.Pipe()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	// Revert and cleanup
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		stdlog.SetOutput(os.Stderr)
	}()

	// Use custom writer
	os.Stdout = writer
	os.Stderr = writer
	stdlog.SetOutput(writer)
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		_, err := io.Copy(&buf, reader)
		if err != nil {
			log.Fatal().Err(err).Msg("")
		}
		out <- buf.String()
	}()
	wg.Wait()

	// Run actual function
	f()
	// if err := f(); err != nil {
	// 	t.Errorf("Function call failed: %v", err)
	// }

	writer.Close()
	return <-out
}

func ProvideStdin(f func(), input string) {
	// Setup custom file to use as Stdin
	content := []byte(input)
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal().Err(err).Msg("")
	}
	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal().Err(err).Msg("")
	}

	// Revert and cleanup
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin
	defer os.Remove(tmpfile.Name())        // clean up

	// Use custom file
	os.Stdin = tmpfile

	// Run actual function
	f()
	// if err := f(); err != nil {
	// 	t.Errorf("Function call failed: %v", err)
	// }

	if err := tmpfile.Close(); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}

func ParseShellArgs(input string) []string {
	args, errArgs := shellwords.Parse(input)

	if errArgs != nil {
		log.Fatal().Err(errArgs).Msg("")
	}

	return args
}
