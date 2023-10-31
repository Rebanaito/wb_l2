package main

import "testing"

func TestValidInput(t *testing.T) {
	host := "localhost"
	port := "8080"
	argv := []string{host, port, "--timeout=15s"}
	parsedHost, parsedPort, err := parseArgs(argv)
	if err != nil {
		t.Fatal(`Error where not expected`, err)
	}
	if parsedHost != host {
		t.Fatalf(`Error parsing the host - want: "%s", got: "%s"`, host, parsedHost)
	}
	if parsedPort != port {
		t.Fatalf(`Error parsing the port - want: "%s", got: "%s"`, port, parsedPort)
	}
	if TIMEOUT != 15 {
		t.Fatalf(`Error parsing timeout - want: "%d", got: "%d"`, 15, TIMEOUT)
	}
}

func TestTooManyArguments(t *testing.T) {
	host := "localhost"
	port := "8080"
	argv := []string{host, port, "--timeout=15s", "extra"}
	want := "telnet: too many arguments"
	_, _, err := parseArgs(argv)
	if err == nil || err.Error() != want {
		t.Fatalf(`Wrong error - want: "%s", got: "%s"`, want, err)
	}
}

func TestInvalidTimeoutArgument(t *testing.T) {
	host := "localhost"
	port := "8080"
	argv := []string{host, port, "--timeout=15d"}
	want := "telnet: --timeout: invalid argument"
	_, _, err := parseArgs(argv)
	if err == nil || err.Error() != want {
		t.Fatalf(`Wrong error - want: "%s", got: "%s"`, want, err)
	}

	argv = []string{host, port, "--timeout=1f5"}
	_, _, err = parseArgs(argv)
	if err == nil {
		t.Fatal(`Want a strconv.Atoi error, got nil`)
	}
}
