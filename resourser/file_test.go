package resourser

import (
	"github.com/golang/lint"
	"os"
	"testing"
)

func TestFsiter_Init(t *testing.T) {
	os.RemoveAll(`/tmp/test-project-234`)
	iter := &fsiter{id: `test-project-234`, uri: `https://github.com/lintflow/golang-test-project.git`}
	count, err := iter.Init()
	if err != nil {
		t.Error(err)
	}
	if count != 2 {
		t.Errorf(`expected count like 2, but - %d`, count)
	}
	linter := new(lint.Linter)
	iterCount := 0
	for iter.Next() {
		iterCount++
		name, blob := iter.File()
		problems, _ := linter.Lint(name, blob)
		if len(problems) == 0 {
			t.Errorf(`expected problems not equals == 0`)
		}
	}
	if iterCount != count {
		t.Errorf(` %d != %d`, count, iterCount)
	}
}
