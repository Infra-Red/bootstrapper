package {{ .Buildpack | RemoveHyphens}}_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/paketo-buildpacks/packit"
	{{ .Buildpack | RemoveHyphens}} "github.com/{{ .Organization }}/{{ .Buildpack }}"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testDetect(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		workingDir string
		detect     packit.DetectFunc
	)

	it.Before(func() {
		var err error
		workingDir, err = ioutil.TempDir("", "working-dir")
		Expect(err).NotTo(HaveOccurred())

		detect = {{ .Buildpack | RemoveHyphens}}.Detect()
	})

	it.After(func() {
		Expect(os.RemoveAll(workingDir)).To(Succeed())
	})

	context("when conditions for detect true are met", func() {
		it("detects", func() {
			result, err := detect(packit.DetectContext{
				WorkingDir: workingDir,
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(result.Plan).To(Equal(packit.BuildPlan{}))
		})
	})
}
