package dirname_test

import (
	gloo "github.com/gloo-foo/framework"
	. "github.com/yupsh/dirname"
)

func ExampleDirname_basic() {
	// dirname /path/to/file.txt
	gloo.MustRun(
		Dirname("/path/to/file.txt"),
	)
	// Output:
	// /path/to
}
