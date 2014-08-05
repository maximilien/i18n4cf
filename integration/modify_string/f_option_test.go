package merge_strings_test

import (
	. "github.com/maximilien/i18n4go/integration/test_helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("modify-string ... ... -f fileName", func() {
	var (
		fixturesPath      string
		inputFilesPath    string
		expectedFilesPath string
	)

	BeforeEach(func() {
		_, err := os.Getwd()
		Ω(err).ShouldNot(HaveOccurred())

		fixturesPath = filepath.Join("..", "..", "test_fixtures", "modify_string")
		inputFilesPath = filepath.Join(fixturesPath, "f_option", "input_files")
		expectedFilesPath = filepath.Join(fixturesPath, "f_option", "expected_output")
	})

	Context("valid options", func() {
		Context("rewriting file and updating i18n file", func() {
			BeforeEach(func() {
				session := Runi18n("-c", "modify-string", 
								   "-v", 
								   "-d", filepath.Join(inputFilesPath),
								   "--old-string", oldString,
								   "--new-string", newString,
								   "-f", filepath.Join(inputFilesPath, "f_option_test.go"),
								   "--i18n-strings-filename", filepath.Join(inputFilesPath, "i18n", "f_option_test.en.json"),)
				Ω(session.ExitCode()).Should(Equal(0))
			})

			AfterEach(func() {
				RemoveAllFiles(
					GetFilePath(inputFilesPath, "en.all.json"),
				)
			})

			It("changes all ocurrences of old-string to new-string in the Go source by -f", func() {

				})

			It("changes all ocurrences of old-string to new-string in the i18n JSON file passed by i18n-strings-filename", func() {
				
				})
			})

		Context("rewriting file and updating i18n directory", func() {
			It("changes all ocurrences of old-string to new-string in the Go source", func() {

				})

			It("changes all ocurrences of old-string to new-string in all the i18n JSON file in the i18n-strings-dirname", func() {
				
				})
			})
		})
	})

	Context("invalid options", func() {

	})
})