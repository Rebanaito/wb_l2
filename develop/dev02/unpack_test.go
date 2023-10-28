package unpack

import "testing"

func TestValidSimple(t *testing.T) {
	input := "a4bc2d5e"
	want := "aaaabccddddde"
	str, err := Unpack(input)
	if err != nil {
		t.Fatal("Error where one is not expected")
	}
	if str != want {
		t.Fatalf(`Input: "%s", want: "%s", got: "%s"`,
			input, want, str)
	}

	input = "abcd"
	want = "abcd"
	str, err = Unpack(input)
	if err != nil {
		t.Fatal("Error where one is not expected")
	}
	if str != want {
		t.Fatalf(`Input: "%s", want: "%s", got: "%s"`,
			input, want, str)
	}

	input = ""
	want = ""
	str, err = Unpack(input)
	if err != nil {
		t.Fatal("Error where one is not expected")
	}
	if str != want {
		t.Fatalf(`Input: "%s", want: "%s", got: "%s"`,
			input, want, str)
	}
}

func TestValidEscape(t *testing.T) {
	input := "qwe\\4\\5"
	want := "qwe45"
	str, err := Unpack(input)
	if err != nil {
		t.Fatal("Error where one is not expected")
	}
	if str != want {
		t.Fatalf(`Input: "%s", want: "%s", got: "%s"`,
			input, want, str)
	}

	input = "qwe\\45"
	want = "qwe44444"
	str, err = Unpack(input)
	if err != nil {
		t.Fatal("Error where one is not expected")
	}
	if str != want {
		t.Fatalf(`Input: "%s", want: "%s", got: "%s"`,
			input, want, str)
	}

	input = "qwe\\\\5"
	want = "qwe\\\\\\\\\\"
	str, err = Unpack(input)
	if err != nil {
		t.Fatal("Error where one is not expected")
	}
	if str != want {
		t.Fatalf(`Input: "%s", want: "%s", got: "%s"`,
			input, want, str)
	}
}

func TestInvalid(t *testing.T) {
	input := "45"
	_, err := Unpack(input)
	if err != InvalidString {
		t.Fatalf(`"%s" is an invalid string, Unpack() must return InvalidString error`, input)
	}

	input = "4\\5"
	_, err = Unpack(input)
	if err != InvalidString {
		t.Fatalf(`"%s" is an invalid string, Unpack() must return InvalidString error`, input)
	}
}
