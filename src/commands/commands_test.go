package commands

import (
	"flag"
	"fmt"
	"testing"
)

func TestCommands(t *testing.T) {

	testCommand := commandOptions{
		name:    "test",
		summary: "a test command",
		usage:   "[-x] y <n>",
	}

	testFlags := flag.NewFlagSet("test", flag.ExitOnError)
	testFlags.Bool("-foo", false, "foo flag")
	testFlags.Bool("f", false, "")
	testFlags.Bool("b", false, "baz flag")
	testFlags.Bool("-bar", false, "bar flag")
	testCommandWithTestFlags := commandOptions{
		flags:   testFlags,
		name:    "test",
		summary: "a test command",
		usage:   "[-x] y <n>",
	}

	t.Run("Get", func(c *testing.T) {
		c.Parallel()
		actual, err := Get("add")
		_, ok := actual.(Command)

		if err != nil || !ok {
			c.Fail()
		}

		actual, err = Get("foo")
		_, ok = actual.(Command)

		if err.Error() != "'foo' is not a git-env command" || ok {
			c.Fail()
		}

		return
	})

	t.Run("Name", func(c *testing.T) {
		c.Parallel()
		expected := "test"
		actual := testCommand.Name()

		if actual != expected {
			fmt.Println(expected)
			fmt.Println("!=")
			fmt.Println(actual)
			c.Fail()
		}
		return
	})

	t.Run("Summary", func(c *testing.T) {
		c.Parallel()
		expected := "a test command"
		actual := testCommand.Summary()

		if actual != expected {
			fmt.Println(expected)
			fmt.Println("!=")
			fmt.Println(actual)
			c.Fail()
		}
		return
	})

	t.Run("Usage", func(c *testing.T) {
		c.Parallel()
		expected := "usage: git env [-x] y <n>\n"
		actual := testCommand.Usage()

		if actual.String() != expected {
			fmt.Println(expected)
			fmt.Println("!=")
			fmt.Println(actual.String())
			c.Fail()
		}

		expected = "usage: git env [-x] y <n>\n" +
			"\n" +
			"    --bar                 bar flag      \n" +
			"    -f --foo              foo flag      \n" +
			"    -b                    baz flag      \n"
		actual = testCommandWithTestFlags.Usage()
		if actual.String() != expected {
			fmt.Println(expected)
			fmt.Println("!=")
			fmt.Println(actual.String())
			c.Fail()
		}

		return
	})
}
