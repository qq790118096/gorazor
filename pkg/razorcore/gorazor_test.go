package razorcore

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var testdata = "testdata"

func TestCap(t *testing.T) {
	if Capitalize("") != "" {
		t.Error()
	}
	if Capitalize("hello") != "Hello" {
		t.Error()
	}
	if Capitalize("0hello") != "0hello" {
		t.Error()
	}
}

func TestLayManager(t *testing.T) {
	SetLayout("hello", []string{"this", "is", "good"})
	SetLayout("world", []string{"funny"})
	if len(LayoutArgs("hello")) != 3 {
		t.Error()
	}
	if len(LayoutArgs("world")) != 1 {
		t.Error()
	}
	if len(LayoutArgs("NO")) != 0 {
		t.Error()
	}
}

func TestLexer1(t *testing.T) {
	text := "case do func var switch"
	lex := &Lexer{text, Tests}
	res, err := lex.Scan()
	if err != nil {
		t.Error(err)
	}
	if len(res) != 10 {
		t.Error("token number")
	}
	for i, x := range res {
		if i%2 == 0 && x.Type != tkKeyword {
			t.Error("KEYWORD", x)
		}
	}
}
func TestLexer2(t *testing.T) {
	text := "case casex do do3 func func_ var var+ "
	lex := &Lexer{text, Tests}
	res, err := lex.Scan()
	if err != nil {
		t.Error(err)
	}
	if len(res) != 18 {
		t.Error(err)
	}
	for i, x := range res {
		if i == 0 || i == 4 || i == 8 || i == 12 || i == 14 {
			if x.Type != tkKeyword {
				t.Error("KEYWORD")
			}
		} else if x.Type == tkKeyword {
			t.Error("Should NOT KEYWORD", x)
		}
	}
}

func TestDebug(t *testing.T) {
	casedir, _ := filepath.Abs(filepath.Dir("./cases/"))
	outdir, _ := filepath.Abs(filepath.Dir("./" + testdata + "/"))
	option := Option{}
	option.IsDebug = true
	GenFile(casedir+"/var.gohtml", outdir+"/_var.gohtml", option)
}

func TestGenerate(t *testing.T) {
	casedir, _ := filepath.Abs(filepath.Dir("./cases/"))
	sap := string(filepath.Separator)

	visit := func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() { // regular file
			if strings.HasPrefix(filepath.Base(path), ".") {
				return nil
			}
			name := strings.Replace(path, ".gohtml", ".go", 1)
			cmp := strings.Replace(name, sap+"cases"+sap, sap+testdata+sap, -1)
			dirname := filepath.Dir(cmp)
			log := filepath.Join(dirname, "_"+filepath.Base(cmp))
			if !exists(dirname) {
				os.MkdirAll(dirname, 0755)
			}
			option := Option{}

			if strings.HasSuffix(path, "panic.gohtml") {
				defer func() {
					if r := recover(); r != nil {
						panicMsg := fmt.Sprint(r)
						if !strings.HasPrefix(panicMsg, "failed to format template") ||
							!strings.Contains(panicMsg, ">>>> func Panic(totalMessage i nt) string {") {
							t.Error("panic.gohtml test failed")
						}
					}
				}()
				GenFile(path, log, option)
			} else {
				GenFile(path, log, option)
			}

			if !exists(cmp) {
				t.Error("No cmp:", cmp)
			} else if !exists(log) {
				t.Error("No log:", log)
			} else {
				//compare the log file and cmp file
				_cmp, _e1 := ioutil.ReadFile(cmp)
				_log, _e2 := ioutil.ReadFile(log)
				if _e1 != nil || _e2 != nil {
					t.Error("Reading")
				} else if string(_cmp) != string(_log) {
					t.Error("MISMATCH:", log, cmp)
				} else {
					t.Log("PASS")
				}
			}
		}
		return nil
	}
	QuickMode = true
	err := filepath.Walk(casedir, visit)
	if err != nil {
		t.Error("walk")
	}
}
