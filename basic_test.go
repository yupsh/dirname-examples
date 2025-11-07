package dirname_test

import (
	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/dirname"
)

func ExampleDirname_basic() {
	// dirname /path/to/file.txt
	yup.MustRun(
		Dirname("/path/to/file.txt"),
	)
	// Output:
	// /path/to
}

