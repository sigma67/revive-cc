package fixtures

import (
	ast "go/ast"
	"bytes"
        "crypto/md5"
	"fmt"
	_ "net/http"
	"strings"
	str "strings"
)

const str = "" // MATCH /The name 'str' shadows an import name/

type myAst struct {
	ast *ast.GenDecl
}

type bytes struct {} // MATCH /The name 'bytes' shadows an import name/

type fmt interface {} // MATCH /The name 'fmt' shadows an import name/

func (ast myAst) foo() {} // MATCH /The name 'ast' shadows an import name/

func md5() {} // MATCH /The name 'md5' shadows an import name/

func bar(_ string) {}

func toto() {
	strings := map[string]string{} // MATCH /The name 'strings' shadows an import name/
}

func titi() {
	v := md5+bytes
	return ast
}
