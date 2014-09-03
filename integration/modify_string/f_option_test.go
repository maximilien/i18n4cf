package modify_string_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	common "github.com/maximilien/i18n4go/common"
	. "github.com/maximilien/i18n4go/integration/test_helpers"
)

var _ = Describe("modify-string ... ... -f fileName", func() {
	var (
		fixturesPath         string
		inputFilesPath       string
		expectedFilesPath    string
		oldString, newString string
	)

	Context("valid options", func() {
		Context("rewriting file and updating i18n file", func() {
			BeforeEach(func() {
				//DEBUG
				fmt.Println("PATH:" + os.Getenv("PATH"))

				_, err := os.Getwd()
				Ω(err).ShouldNot(HaveOccurred())

				fixturesPath = filepath.Join("..", "..", "test_fixtures", "modify_string")
				inputFilesPath = filepath.Join(fixturesPath, "f_option", "input_files")
				expectedFilesPath = filepath.Join(fixturesPath, "f_option", "expected_output")

				oldString = "Hello modify string"
				newString = "String is modified"

				session := Runi18n("-c",
					"modify-string",
					"-v",
					"-d", filepath.Join(inputFilesPath),
					"--old-string", oldString,
					"--new-string", newString,
					"-f", filepath.Join(inputFilesPath, "test.go"),
					"--i18n-strings-filename", filepath.Join(inputFilesPath, "i18n", "test.go.en.json"),
				)

				Ω(session.ExitCode()).Should(Equal(0))
			})

			AfterEach(func() {
				RemoveAllFiles(
					GetFilePath(inputFilesPath, "en.all.json"),
				)
			})

			It("changes all ocurrences of old-string to new-string in the Go source by -f", func() {
				fileBytes, err := ioutil.ReadFile(filepath.Join(inputFilesPath, "test.go"))
				Ω(err).ToNot(HaveOccurred())
				Ω(string(fileBytes)).ToNot(ContainSubstring(oldString))
				Ω(string(fileBytes)).To(ContainSubstring(newString))
			})

			FIt("changes all ocurrences of old-string to new-string in the i18n JSON file passed by i18n-strings-filename", func() {
				i18nStringInfos, err := common.LoadI18nStringInfos(filepath.Join(inputFilesPath, "i18n", "test.go.en.json"))
				Ω(err).ToNot(HaveOccurred())
				for _, value := range i18nStringInfos {
					Ω(value.ID).ToNot(Equal(oldString))
					Ω(value.ID).To(Equal(newString))
				}
			})

			Context("rewriting file and updating i18n directory", func() {
				It("changes all ocurrences of old-string to new-string in the Go source", func() {

				})

				It("changes all ocurrences of old-string to new-string in all the i18n JSON file in the i18n-strings-dirname", func() {

				})
			})
		})

		Context("invalid options", func() {

		})
	})
})
