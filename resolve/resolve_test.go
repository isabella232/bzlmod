package resolve

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestRepoNames(t *testing.T) {
	wsDir := t.TempDir()
	writeLocalModuleBazel(t, wsDir, `
module(name="A")
bazel_dep(name="B", version="1.0", repo_name="BfromA")
bazel_dep(name="C", version="2.0")
bazel_dep(name="E", version="3.0", repo_name="EfromA")
`)
	reg := &myReg{}
	reg.addModuleBazel(t, "B", "1.0", `
module(name="B", version="1.0")
bazel_dep(name="D", version="0.1", repo_name="DfromB")
`)
	reg.addModuleBazel(t, "C", "2.0", `
module(name="C", version="2.0")
bazel_dep(name="D", version="0.1", repo_name="DfromC")
`)
	reg.addModuleBazel(t, "D", "0.1", `
module(name="D", version="0.1")
bazel_dep(name="E", version="3.0")
`)
	reg.addModuleBazel(t, "E", "3.0", `
module(name="E", version="3.0")
`)
	if err := Resolve(wsDir, reg); err != nil {
		t.Fatal(err)
	}
	wsBytes, err := ioutil.ReadFile(filepath.Join(wsDir, "WORKSPACE"))
	if err != nil {
		t.Fatal(err)
	}
	ws := string(wsBytes)
	const expectedWs = `# This file is automatically generated by bzlmod
workspace(
    name = "A",
    watch_files = { "MODULE.bazel": "sha256-fakevalue" },
    refresh_command = ["bzlmod", "resolve"],
)

repo(
    name = "BfromA",
    fetch_command = ["bzlmod", "fetch", "BfromA"],
    fingerprint = "fakefingerprint",
    repo_deps = {
        "DfromB": "D",
    },
)

repo(
    name = "C",
    fetch_command = ["bzlmod", "fetch", "C"],
    fingerprint = "fakefingerprint",
    repo_deps = {
        "DfromC": "D",
    },
)

repo(
    name = "D",
    fetch_command = ["bzlmod", "fetch", "D"],
    fingerprint = "fakefingerprint",
    repo_deps = {
        "E": "EfromA",
    },
)

repo(
    name = "EfromA",
    fetch_command = ["bzlmod", "fetch", "EfromA"],
    fingerprint = "fakefingerprint",
)
`
	if ws != expectedWs {
		t.Fatalf("unexpected WORKSPACE contents: %v", ws)
	}
}
