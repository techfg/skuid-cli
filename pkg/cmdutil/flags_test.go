package cmdutil_test

import (
	"fmt"
	"testing"

	"github.com/skuid/skuid-cli/pkg"
	"github.com/skuid/skuid-cli/pkg/cmdutil"
	"github.com/skuid/skuid-cli/pkg/flags"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type addFlag[T flags.FlagType] func(*cobra.Command, *T, *flags.Flag[T])

func TestAddStringFlag(t *testing.T) {
	runAddFlagTests(t, cmdutil.AddStringFlag, "somevalue", "string", false, false)
}

func TestAddBoolFlag(t *testing.T) {
	runAddFlagTests(t, cmdutil.AddBoolFlag, true, "bool", false, false)
}

func TestAddIntFlag(t *testing.T) {
	runAddFlagTests(t, cmdutil.AddIntFlag, 500, "int", false, false)
}

func TestAddStringSliceFlag(t *testing.T) {
	runAddFlagTests(t, cmdutil.AddStringSliceFlag, []string{"foo", "bar"}, "stringSlice", false, false)
}

func TestAddValueFlag(t *testing.T) {
	runAddFlagTests(t, cmdutil.AddValueFlag, 500, "int", true, false)
}

func TestAddValueWithRedactFlag(t *testing.T) {
	runAddFlagTests[int](t, func(cmd *cobra.Command, valPtr *int, flagInfo *flags.Flag[int]) {
		cmdutil.AddValueWithRedactFlag(cmd, valPtr, flagInfo)
	}, 500, "int", true, true)
}

func TestAddValueWithRedactFlagReturnsValue(t *testing.T) {
	p := new(string)
	f := &flags.Flag[string]{Name: "myflag", Usage: "this is the usage", Default: "foo", Redact: func(string) string { return "***" }}
	v := cmdutil.AddValueWithRedactFlag(&cobra.Command{Use: "mycmd"}, p, f)
	assert.NotNil(t, v)
	assert.Equal(t, "string", v.Type())
	assert.Equal(t, "***", v.String())
	assert.Equal(t, "foo", *p)
}

func TestAddAuthFlags(t *testing.T) {
	cmd := &cobra.Command{Use: "mycmd"}
	opts := pkg.AuthorizeOptions{}
	cmdutil.AddAuthFlags(cmd, &opts)
	fs := cmd.Flags()
	err := fs.Parse([]string{"--host", "test.skuidsite.com", "--username", "foo", "--password", "bar"})
	assert.NoError(t, err)
	assert.NotNil(t, fs.Lookup(flags.Host.Name))
	assert.NotNil(t, fs.Lookup(flags.Username.Name))
	assert.NotNil(t, fs.Lookup(flags.Password.Name))
	assert.Equal(t, "https://test.skuidsite.com", opts.Host)
	assert.Equal(t, "foo", opts.Username)
	assert.Equal(t, "***", opts.Password.String())
	assert.Equal(t, "bar", opts.Password.Unredacted().String())
}

func TestEnvVarName(t *testing.T) {
	testCases := []struct {
		testDescription string
		giveName        string
		wantName        string
	}{
		{
			testDescription: "default",
			giveName:        *new(string),
			wantName:        "SKUID_",
		},
		{
			testDescription: "empty string",
			giveName:        "",
			wantName:        "SKUID_",
		},
		{
			testDescription: "simple name",
			giveName:        "dir",
			wantName:        "SKUID_DIR",
		},
		{
			testDescription: "with underscore",
			giveName:        "dir_name",
			wantName:        "SKUID_DIR_NAME",
		},
		{
			testDescription: "with hyphen",
			giveName:        "dir-name",
			wantName:        "SKUID_DIR_NAME",
		},
		{
			testDescription: "with double hyphen",
			giveName:        "dir--name",
			wantName:        "SKUID_DIR__NAME",
		},
		{
			testDescription: "with multiple hyphen",
			giveName:        "log-dir-name",
			wantName:        "SKUID_LOG_DIR_NAME",
		},
		{
			testDescription: "mix case",
			giveName:        "lOG-dir-nAMe",
			wantName:        "SKUID_LOG_DIR_NAME",
		},
		{
			testDescription: "upper case",
			giveName:        "LOG-DIR-NAME",
			wantName:        "SKUID_LOG_DIR_NAME",
		},
		{
			testDescription: "special chars",
			giveName:        "LoG@DIr$NAmE-^#!()",
			wantName:        "SKUID_LOG@DIR$NAME_^#!()",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testDescription, func(t *testing.T) {
			actualName := cmdutil.EnvVarName(tc.giveName)
			assert.Equal(t, tc.wantName, actualName)
		})
	}
}

func expectInvalidFlagNameError(name string) error {
	return fmt.Errorf("flag name %q is invalid", name)
}

func expectInvalidUsageError(name string, usage string) error {
	return fmt.Errorf("flag usage %q is invalid for flag name %q", usage, name)
}

func expectInvalidParseError[T flags.FlagType](name string) error {
	return fmt.Errorf("flag type %T does not support Parse for flag name %q", new(flags.Flag[T]), name)
}

func expectInvalidRedactError[T flags.FlagType](name string) error {
	return fmt.Errorf("flag type %T does not support Redact for flag name %q", new(flags.Flag[T]), name)
}

func emptyParse[T flags.FlagType](val string) (T, error) {
	return *new(T), nil
}

func emptyRedact[T flags.FlagType](val T) string {
	return ""
}

func runAddFlagTests[T flags.FlagType](t *testing.T, addFlag addFlag[T], defaultValue T, expectedType string, supportsParse bool, supportsRedact bool) {
	var parseError error
	if !supportsParse {
		parseError = expectInvalidParseError[T]("test")
	}
	var redactError error
	if !supportsRedact {
		redactError = expectInvalidRedactError[T]("test")
	}

	testCases := []struct {
		testDescription string
		giveFlag        *flags.Flag[T]
		wantPanic       bool
		wantError       error
		wantValue       T
	}{
		{
			testDescription: "name string default",
			giveFlag:        &flags.Flag[T]{Usage: "this is the usage"},
			wantPanic:       true,
			wantError:       expectInvalidFlagNameError(*new(string)),
		},
		{
			testDescription: "name empty",
			giveFlag:        &flags.Flag[T]{Name: "", Usage: "this is the usage"},
			wantPanic:       true,
			wantError:       expectInvalidFlagNameError(""),
		},
		{
			testDescription: "name contains space",
			giveFlag:        &flags.Flag[T]{Name: "foo bar", Usage: "this is the usage"},
			wantPanic:       true,
			wantError:       expectInvalidFlagNameError("foo bar"),
		},
		{
			testDescription: "name contains special char",
			giveFlag:        &flags.Flag[T]{Name: "foo$bar", Usage: "this is the usage"},
			wantPanic:       true,
			wantError:       expectInvalidFlagNameError("foo$bar"),
		},
		{
			testDescription: "name contains upper case",
			giveFlag:        &flags.Flag[T]{Name: "fooBar", Usage: "this is the usage"},
			wantPanic:       true,
			wantError:       expectInvalidFlagNameError("fooBar"),
		},
		{
			testDescription: "name min length",
			giveFlag:        &flags.Flag[T]{Name: "fo", Usage: "this is the usage"},
			wantPanic:       true,
			wantError:       expectInvalidFlagNameError("fo"),
		},
		{
			testDescription: "usage string default",
			giveFlag:        &flags.Flag[T]{Name: "test"},
			wantPanic:       true,
			wantError:       expectInvalidUsageError("test", *new(string)),
		},
		{
			testDescription: "usage empty",
			giveFlag:        &flags.Flag[T]{Name: "test", Usage: ""},
			wantPanic:       true,
			wantError:       expectInvalidUsageError("test", ""),
		},
		{
			testDescription: "usage min length",
			giveFlag:        &flags.Flag[T]{Name: "test", Usage: "012345678"},
			wantPanic:       true,
			wantError:       expectInvalidUsageError("test", "012345678"),
		},
		{
			testDescription: "parse support",
			giveFlag:        &flags.Flag[T]{Name: "test", Usage: "this is the usage", Parse: emptyParse[T]},
			wantPanic:       !supportsParse,
			wantError:       parseError,
		},
		{
			testDescription: "redact support",
			giveFlag:        &flags.Flag[T]{Name: "test", Usage: "this is the usage", Redact: emptyRedact[T]},
			wantPanic:       !supportsRedact,
			wantError:       redactError,
		},
		{
			testDescription: "default value",
			giveFlag:        &flags.Flag[T]{Name: "test", Usage: "this is the usage", Default: defaultValue},
			wantPanic:       false,
			wantError:       nil,
			wantValue:       defaultValue,
		},
		{
			testDescription: "legacy env vars",
			giveFlag:        &flags.Flag[T]{Name: "test", Usage: "this is the usage", LegacyEnvVars: []string{"FOO", "BAR"}},
			wantPanic:       false,
			wantError:       nil,
			wantValue:       *new(T),
		},
		{
			testDescription: "shorthand",
			giveFlag:        &flags.Flag[T]{Name: "test", Usage: "this is the usage", Shorthand: "v"},
			wantPanic:       false,
			wantError:       nil,
			wantValue:       *new(T),
		},
		{
			testDescription: "required",
			giveFlag:        &flags.Flag[T]{Name: "test", Usage: "this is the usage", Required: true},
			wantPanic:       false,
			wantError:       nil,
			wantValue:       *new(T),
		},
		{
			testDescription: "global",
			giveFlag:        &flags.Flag[T]{Name: "test", Usage: "this is the usage", Global: true},
			wantPanic:       false,
			wantError:       nil,
			wantValue:       *new(T),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testDescription, func(t *testing.T) {
			cmd := &cobra.Command{Use: "mycmd"}
			p := new(T)
			if tc.wantPanic {
				assert.PanicsWithError(t, tc.wantError.Error(), func() {
					addFlag(cmd, p, tc.giveFlag)
				})
			} else {
				assert.NotPanics(t, func() {
					addFlag(cmd, p, tc.giveFlag)
				})

				var fs *pflag.FlagSet
				if tc.giveFlag.Global {
					fs = cmd.PersistentFlags()
				} else {
					fs = cmd.LocalNonPersistentFlags()
				}
				actualFlag := fs.Lookup(tc.giveFlag.Name)
				require.NotNil(t, actualFlag)

				// correct type
				assert.Equal(t, expectedType, actualFlag.Value.Type())

				// usage updated to include environment variables
				assert.Contains(t, actualFlag.Usage, cmdutil.EnvVarName(tc.giveFlag.Name))

				// shorthand
				assert.Equal(t, tc.giveFlag.Shorthand, actualFlag.Shorthand)

				// default value
				assert.Equal(t, tc.giveFlag.Default, *p)

				// flag set by skuid cli annotation
				av, ok := actualFlag.Annotations[cmdutil.FlagSetBySkuidCliAnnotation]
				require.True(t, ok)
				assert.True(t, len(av) == 1 && av[0] == "true")

				// legacy env vars annotation
				av, ok = actualFlag.Annotations[cmdutil.LegacyEnvVarsAnnotation]
				require.True(t, ok)
				assert.ElementsMatch(t, tc.giveFlag.LegacyEnvVars, av)

				// if required, since nothing was set it should fail
				err := cmd.ValidateRequiredFlags()
				if tc.giveFlag.Required {
					assert.ErrorContains(t, err, "required")
					assert.ErrorContains(t, err, tc.giveFlag.Name)
				} else {
					require.NoError(t, err)
				}
			}
		})
	}
}
