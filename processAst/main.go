package processAst

import (
	"github.com/dave/dst"
	"io/ioutil"
)

var Project = ProjectType{
	Docs: []Doc{},
}

//func main() {
//	err := readDocFiles()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	err = writeDocMainGoFiles()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//}

func ReadProject() error {
	// сбрасываем старые значения
	Project.Docs = []Doc{}
	dirs, err := ioutil.ReadDir("../projectTemplate/docs")
	if err != nil {
		return err
	}
	for _, d := range dirs {
		doc := Doc{Name: d.Name()}
		err := doc.ReadMainGo()
		if err != nil {
			return err
		}
		Project.Docs = append(Project.Docs, doc)
	}
	return nil
}

func WriteProject() error {
	for _, d := range Project.Docs {
		err := d.WriteMainGo()
		if err != nil {
			return err
		}
	}
	return nil
}

func findFuncDecl(node *dst.File, funcName string) *dst.FuncDecl {
	for _, d := range node.Decls {
		if decl, ok := d.(*dst.FuncDecl); ok {
			if decl.Name.Name == funcName {
				return decl
			}
		}
	}
	return nil
}

func clearArray(node *dst.FuncDecl, keyName string) {
	dst.Inspect(node, func(n dst.Node) bool {
		v, ok := n.(*dst.KeyValueExpr)
		if ok {
			if k, ok := v.Key.(*dst.Ident); ok && k.Name == keyName {
				v.Value.(*dst.CompositeLit).Elts = []dst.Expr{}
				return true
			}
		}
		return true
	})
}

