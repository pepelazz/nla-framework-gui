package processAst

import (
	"fmt"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/iancoleman/strcase"
	"go/token"
	"os"
)

func (doc *Doc) WriteMainGo() error  {

	//пишем название package
	doc.DstFile.Name.Name = doc.getPackageName()

	//пишем блок с константами
	doc.WriteConstBlock()

	//пишем блок GetFlds
	getDocNode := findFuncDecl(doc.DstFile, "GetDoc")
	// чистим массив, перед началом добавления
	clearArray(getDocNode, "Flds")
	for _, fld := range doc.Flds {
		addToFldsEltsDst(getDocNode, fld.Print())
	}
	f, err := os.Create(fmt.Sprintf("../projectTemplate/docs/%s/main.go", doc.getPackageName()))
	if err != nil {
		return err
	}
	defer f.Close()

	if err := decorator.Fprint(f, doc.DstFile); err != nil {
		return err
	}

	return nil
}

func addToFldsEltsDst(node *dst.FuncDecl, newGetFld *dst.CallExpr) (res []dst.Expr) {
	dst.Inspect(node, func(n dst.Node) bool {
		v, ok := n.(*dst.KeyValueExpr)
		if ok {
			if k, ok := v.Key.(*dst.Ident); ok && k.Name == "Flds" {
				if elts, ok := v.Value.(*dst.CompositeLit); ok {
					elts.Elts = append(elts.Elts, newGetFld)
					if len(elts.Elts) == 1 {
						elts.Elts[len(elts.Elts)-1].Decorations().Before = dst.NewLine
					}
					elts.Elts[len(elts.Elts)-1].Decorations().After = dst.NewLine
				}
				return true
			}
		}
		return true
	})
	return
}

func (doc *Doc) WriteConstBlock() {
	dst.Inspect(doc.DstFile, func(n dst.Node) bool {
		v, ok := n.(*dst.GenDecl)
		if ok && v.Tok == token.CONST {
			for _, spec := range v.Specs {
				if vs, ok := spec.(*dst.ValueSpec); ok {
					key := vs.Names[0].Name
					if key == "name" {
						vs.Values[0].(*dst.BasicLit).Value = fmt.Sprintf("%q", doc.Name)
					}
					if key == "name_ru" {
						vs.Values[0].(*dst.BasicLit).Value = fmt.Sprintf("%q", doc.NameRu)
					}
					if key == "name_ru_plural" {
						vs.Values[0].(*dst.BasicLit).Value = fmt.Sprintf("%q", doc.NameRuPlural)
					}
					if key == "menu_icon" {
						vs.Values[0].(*dst.BasicLit).Value = fmt.Sprintf("%q", doc.MenuIcon)
					}
				}
			}
		}
		return true
	})
}

func (doc *Doc) getPackageName() string  {
	return strcase.ToLowerCamel(doc.Name)
}
