package grunt

import (
	"os"
	"os/exec"
	"strings"

	"github.com/revel/revel"
)

type GruntCompiler struct {
	Name       string
	Path       string
	Grunt      string
	notOnStart bool
}

func NewCompiler(name, path, grunt_rule string) *GruntCompiler {
	return &GruntCompiler{
		Name:       name,
		Path:       path,
		Grunt:      grunt_rule,
		notOnStart: true,
	}
}

func NewCompilerOnStart(name, path, grunt_rule string) *GruntCompiler {
	return &GruntCompiler{
		Name:       name,
		Path:       path,
		Grunt:      grunt_rule,
		notOnStart: false,
	}
}

func (c *GruntCompiler) Refresh() *revel.Error {
	// It's start-up or a file changed.  Re-compile...
	if c.notOnStart {
		c.notOnStart = false
		return nil
	}

	revel.INFO.Println("Compiling: ", c.Name)
	os.Chdir(revel.BasePath)

	out, err := exec.Command("grunt", "--no-color", c.Grunt).Output()
	if err != nil {
		revel.ERROR.Println("Failed to compile", c.Path, err)
		revel.ERROR.Println("Output:\n", string(out))
		return nil
	}
	revel.TRACE.Println(string(out))

	os.Chdir(revel.SourcePath)
	return nil
}

func (c *GruntCompiler) WatchDir(info os.FileInfo) bool {
	// Watch all directories, except the ones starting with a dot.
	return !strings.HasPrefix(info.Name(), ".")
}

func (c *GruntCompiler) WatchFile(basename string) bool {
	// Watch all files, except the ones starting with a dot.
	return !strings.HasPrefix(basename, ".")
}
