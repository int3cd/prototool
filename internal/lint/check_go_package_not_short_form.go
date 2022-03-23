package lint

import (
	"strings"

	"github.com/emicklei/proto"
	"github.com/uber/prototool/internal/text"
)

var fileOptionsGoPackageNotShortFormLinter = NewLinter(
	"FILE_OPTIONS_GO_PACKAGE_NOT_SHORT_FORM",
	`Verifies that the file option "go_package" is not of the form "go/import/path".`,
	checkFileOptionsGoPackageNotShortForm,
)

func checkFileOptionsGoPackageNotShortForm(add func(*text.Failure), dirPath string, descriptors []*FileDescriptor) error {
	return runVisitor(&fileOptionsGoPackageNotShortFormVisitor{baseAddVisitor: newBaseAddVisitor(add)}, descriptors)
}

type fileOptionsGoPackageNotShortFormVisitor struct {
	baseAddVisitor

	option *proto.Option
}

func (v *fileOptionsGoPackageNotShortFormVisitor) OnStart(descriptor *FileDescriptor) error {
	v.option = nil
	return nil
}

func (v *fileOptionsGoPackageNotShortFormVisitor) VisitOption(element *proto.Option) {
	if element.Name == "go_package" {
		v.option = element
	}
}

func (v *fileOptionsGoPackageNotShortFormVisitor) Finally() error {
	if v.option == nil {
		return nil
	}
	value := v.option.Constant.Source
	if !strings.Contains(value, ";") {
		v.AddFailuref(v.option.Position, `Option "go_package" cannot be of the short-form "package" and must only be of the long-form "go/import/path;package", but was %q.`, value)
	}
	return nil
}
