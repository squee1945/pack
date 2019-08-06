package acceptance

import (
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"runtime"
	"testing"
)

func BuildPack(t *testing.T, packCmdPath string) string {
	packTmpDir, err := ioutil.TempDir("", "pack.acceptance.binary.")
	if err != nil {
		t.Fatal(err)
	}

	packPath := filepath.Join(packTmpDir, "pack")
	if runtime.GOOS == "windows" {
		packPath = packPath + ".exe"
	}

	if txt, err := exec.Command("go", "build", "-o", packPath, packCmdPath).CombinedOutput(); err != nil {
		t.Fatal("building pack cli:\n", string(txt), err)
	}

	return packPath
}