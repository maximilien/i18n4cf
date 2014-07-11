package cmds

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"go/ast"
	"go/parser"
	"go/token"

	"github.com/maximilien/i18n4go/common"
)

type ShowMissingStrings struct {
	options common.Options

	I18nStringInfos     []common.I18nStringInfo
	TranslatedStrings   []string
	I18nStringsFilename string
	Directory           string
}

func NewShowMissingStrings(options common.Options) ShowMissingStrings {
	return ShowMissingStrings{
		options:             options,
		Directory:           options.DirnameFlag,
		I18nStringsFilename: options.I18nStringsFilenameFlag,
		TranslatedStrings:   []string{},
	}
}

func (sms *ShowMissingStrings) Options() common.Options {
	return sms.options
}

func (sms *ShowMissingStrings) Println(a ...interface{}) (int, error) {
	if sms.options.VerboseFlag {
		return fmt.Println(a...)
	}

	return 0, nil
}

func (sms *ShowMissingStrings) Printf(msg string, a ...interface{}) (int, error) {
	if sms.options.VerboseFlag {
		return fmt.Printf(msg, a...)
	}

	return 0, nil
}

func (sms *ShowMissingStrings) Run() error {
	return sms.showMissingStrings()
}

func (sms *ShowMissingStrings) showMissingStrings() error {

	//Load en_US.all.json

	stringInfos, err := common.LoadI18nStringInfos(sms.I18nStringsFilename)
	if err != nil {
		return err
	}
	sms.I18nStringInfos = stringInfos

	//Run AST to get list of strings
	sourceFiles, _ := getFilesAndDir(sms.Directory)
	for _, sourceFile := range sourceFiles {
		err = sms.inspectFile(sourceFile)
		if err != nil {
			return err
		}
	}

	//Compare list of strings with en_US.all.json
	missingStrings := false
	for _, codeString := range sms.TranslatedStrings {
		if !stringInStringInfos(codeString, sms.I18nStringInfos) {
			fmt.Println("Missing:", codeString)
			missingStrings = true
		}
	}

	if missingStrings {
		return errors.New("Missing Strings!")
	}
	return sms.showExtraStrings()
}

func (sms *ShowMissingStrings) showExtraStrings() error {
	//Compare list of strings with en_US.all.json
	additionalStrings := false
	for _, stringInfo := range sms.I18nStringInfos {
		if !stringInTranslatedStrings(stringInfo.ID, sms.TranslatedStrings) {
			fmt.Println("Additional:", stringInfo.ID)
			additionalStrings = true
		}
	}
	if additionalStrings {
		return errors.New("Additional Strings!")
	}
	return nil
}

func stringInTranslatedStrings(str string, list []string) bool {
	for _, translatedStr := range list {
		if translatedStr == str {
			return true
		}
	}

	return false
}

func stringInStringInfos(str string, list []common.I18nStringInfo) bool {
	for _, stringInfo := range list {
		if stringInfo.ID == str {
			return true
		}
	}

	return false
}

func (sms *ShowMissingStrings) inspectFile(filename string) error {
	fset := token.NewFileSet()

	var absFilePath = filename
	if !filepath.IsAbs(absFilePath) {
		absFilePath = filepath.Join(os.Getenv("PWD"), absFilePath)
	}

	fileInfo, err := common.GetAbsFileInfo(absFilePath)
	if err != nil {
		sms.Println(err)
	}

	if strings.HasPrefix(fileInfo.Name(), ".") {
		sms.Println("WARNING ignoring file:", absFilePath)
		return nil
	}

	astFile, err := parser.ParseFile(fset, absFilePath, nil, parser.ParseComments|parser.AllErrors)
	if err != nil {
		sms.Println(err)
		return err
	}

	return sms.extractString(astFile, fset)
}

func (sms *ShowMissingStrings) extractString(f *ast.File, fset *token.FileSet) error {
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.CallExpr:
			switch x.Fun.(type) {
			case *ast.Ident:
				funName := x.Fun.(*ast.Ident).Name

				if funName == "T" {
					translatedString, err := strconv.Unquote(x.Args[0].(*ast.BasicLit).Value)
					if err != nil {
						panic(err.Error())
					}

					sms.TranslatedStrings = append(sms.TranslatedStrings, translatedString)
				}
			default:
				//Skip!
			}
		}
		return true
	})

	return nil
}
