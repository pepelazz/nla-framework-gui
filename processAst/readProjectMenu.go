package processAst

import (
	"errors"
	"fmt"
	"github.com/dave/dst"
)

func (p *ProjectType) ReadMainGoMenu() error {
	if p.DstFile == nil {
		return errors.New("readMainGoMenu error: p.DstFile = nil")
	}
	dst.Inspect(p.DstFile, func(n dst.Node) bool {
		v, ok := n.(*dst.CompositeLit)
		if ok {
			if k, ok := v.Type.(*dst.ArrayType); ok {
				if s, ok := k.Elt.(*dst.SelectorExpr); ok {
					if x, ok := s.X.(*dst.Ident); ok {
						if x.Name == "t" && s.Sel.Name == "VueMenu" {
							for _, elt := range v.Elts {
								fmt.Printf("-----\n")
								for _, elt1 := range elt.(*dst.CompositeLit).Elts {
									name := elt1.(*dst.KeyValueExpr).Key.(*dst.Ident).Name
									value := elt1.(*dst.KeyValueExpr).Value
									fmt.Printf("name: %s %s\n", name, value)
								}
							}
						}
					}
				}
			}
		}
		return true
	})
	return nil
}
