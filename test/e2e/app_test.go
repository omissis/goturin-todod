package e2e_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	app string

	MkdirTemp = func(pattern string) string {
		tmpdir, err := os.MkdirTemp("", pattern)
		if err != nil {
			Fail(err.Error())
		}

		return tmpdir
	}

	RunCmd = func(cmd string, args ...string) (string, error) {
		out, err := exec.Command(cmd, args...).CombinedOutput()

		GinkgoWriter.Println(string(out))

		return string(out), err
	}

	_ = BeforeSuite(func() {
		tmpdir := MkdirTemp("app-e2e")

		app = filepath.Join(tmpdir, "app")

		if out, err := RunCmd("go", "build", "-o", app, "../../main.go"); err != nil {
			Fail(fmt.Sprintf("Could not build app: %v\nOutput: %s", err, out))
		}
	})

	_ = Describe("app", func() {
		Context("version display", func() {
			It("should print its version information", func() {
				out, err := RunCmd(app, "version")

				Expect(err).To(Not(HaveOccurred()))
				Expect(out).To(ContainSubstring("buildTime: unknown"))
				Expect(out).To(ContainSubstring("gitCommit: unknown"))
				Expect(out).To(ContainSubstring("goVersion: unknown"))
				Expect(out).To(ContainSubstring("osArch: unknown"))
				Expect(out).To(ContainSubstring("version: unknown"))
			})
		})
	})
)

func TestE2e(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "App E2e Suite")
}
