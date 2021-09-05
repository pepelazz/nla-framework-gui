package processAst

import (
	"fmt"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"go/token"
	"os"
	"strconv"
	"strings"
)

// пишем основной main.go проекта
func (p *ProjectType) WriteMainGo() error {


	for i := 0; i < len(p.DstFile.Decls); i++ {
		d := p.DstFile.Decls[i]

		switch d.(type) {
		case *dst.GenDecl:
			dd := d.(*dst.GenDecl)
			// IMPORT Declarations
			if dd.Tok == token.IMPORT {
				// удаляем предыдущие импорты связанные с doc
				for j := len(dd.Specs)-1; j > -1 ; j-- {
					if strings.Contains(dd.Specs[j].(*dst.ImportSpec).Path.Value, "projectTemplate/docs") {
						copy(dd.Specs[j:], dd.Specs[j+1:]) // Shift a[i+1:] left one index.
						dd.Specs[len(dd.Specs)-1] = nil     // Erase last element (write zero value).
						dd.Specs = dd.Specs[:len(dd.Specs)-1]     // Truncate slice
					}
				}
				// Add the new import
				for _, doc := range p.Docs {
					iSpec := &dst.ImportSpec{Path: &dst.BasicLit{Value: strconv.Quote(fmt.Sprintf("github.com/pepelazz/guiDev/projectTemplate/docs/%s", doc.getPackageName()))}}
					//p.DstFile.Imports = append(p.DstFile.Imports, &dst.ImportSpec{Path: &dst.BasicLit{Kind: token.STRING, Value: }})
					dd.Specs = append(dd.Specs, iSpec)
				}
			}
		}
	}

	// находим node t.DocType и прописываем legalEntity.GetDoc(p),
	//p.Docs = []t.DocType{...}
	dst.Inspect(p.DstFile, func(n dst.Node) bool {
		v, ok := n.(*dst.CompositeLit)
		if ok {
			if k, ok := v.Type.(*dst.ArrayType); ok {
				if s, ok := k.Elt.(*dst.SelectorExpr); ok {
					if x, ok := s.X.(*dst.Ident); ok {
						if x.Name == "t" && s.Sel.Name == "DocType" {
							// стираем массив
							v.Elts = []dst.Expr{}
							// записываем docs
							for _, doc := range p.Docs {
								v.Elts = append(v.Elts, &dst.CallExpr{Fun: &dst.SelectorExpr{X: &dst.Ident{Name: doc.getPackageName()}, Sel: &dst.Ident{Name: "GetDoc"}}, Args: []dst.Expr{&dst.Ident{Name: "p"}}})
								if len(v.Elts) == 1 {
									v.Elts[len(v.Elts)-1].Decorations().Before = dst.NewLine
								}
								v.Elts[len(v.Elts)-1].Decorations().After = dst.NewLine
							}
							return true
						}
					}
				}
			}
		}
		return true
	})

	// пишем боковое меню для Vue
	//	p.Vue.Menu = []t.VueMenu{
	err := p.WriteMainGoMenu()
	if err != nil {
		return err
	}

	f, err := os.Create("../projectTemplate/main.go")
	if err != nil {
		return err
	}
	defer f.Close()


	if err := decorator.Fprint(f, p.DstFile); err != nil {
		return err
	}

	return nil
}
