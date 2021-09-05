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

	// обяъвляем функцию отдельно, чтобы можно было рекурсивно вызывать ее внутри самой себя
	var readVueMenu func(node dst.Node) []VueMenu

	readVueMenu = func(node dst.Node) []VueMenu {
		var res []VueMenu
		isParent := true // признак, чтобы прочитать только корневой элемент VueMenu и игнорировать вложенные
		dst.Inspect(node, func(n dst.Node) bool {
			v, ok := n.(*dst.CompositeLit)
			if ok {
				if k, ok := v.Type.(*dst.ArrayType); ok {
					if s, ok := k.Elt.(*dst.SelectorExpr); ok {
						if x, ok := s.X.(*dst.Ident); ok {
							if x.Name == "t" && s.Sel.Name == "VueMenu" && isParent {
								isParent = false
								// итерация по массиву пуктов меню
								for _, elt := range v.Elts {
									menuItem := VueMenu{}
									// итерация по аргументам пункта меню
									for _, elt1 := range elt.(*dst.CompositeLit).Elts {
										name := elt1.(*dst.KeyValueExpr).Key.(*dst.Ident).Name

										// чтение вложенного меню
										if name == "LinkList" {
											menuItem.LinkList = readVueMenu(elt1.(*dst.KeyValueExpr).Value)
											continue
										}

										switch value := elt1.(*dst.KeyValueExpr).Value.(type) {
										case *dst.BasicLit:
											//fmt.Printf("name BasicLit: %s %s\n", name, value.Value)
											switch name {
												case "Url": menuItem.Url = replaceQuotes(value.Value)
												case "Text": menuItem.Text = replaceQuotes(value.Value)
												case "Icon": menuItem.Icon = replaceQuotes(value.Value)
												case "DocName": menuItem.DocName = replaceQuotes(value.Value)
											}
										case *dst.Ident:
											//fmt.Printf("name Ident: %s %s\n", name, value.Name)
											switch name {
												case "IsFolder": menuItem.IsFolder = replaceQuotes(value.Name)
											}
										case *dst.CompositeLit:
											//fmt.Printf("name CompositeLit: %s ", name)
											if name == "Roles" {
												menuItem.Roles = []string{}
												for _, elt := range value.Elts {
													switch v1 := elt.(type) {
													case *dst.SelectorExpr:
														menuItem.Roles = append(menuItem.Roles, v1.Sel.Name)
														//fmt.Printf(" %s.%s ", v1.X.(*dst.Ident).Name, v1.Sel.Name)
													}
												}
											}
											//fmt.Printf("\n")
										default:
											fmt.Printf("uknown type %s %s", name, value)
										}
									}
									res = append(res, menuItem)
								}
							}
						}
					}
				}
			}
			return true
		})
		return res
	}

	p.Menu = readVueMenu(p.DstFile)

	return nil
}

func (p *ProjectType) WriteMainGoMenu() error {
	isParent := true // признак, чтобы прочитать только корневой элемент VueMenu и игнорировать вложенные
	dst.Inspect(p.DstFile, func(n dst.Node) bool {
		v, ok := n.(*dst.CompositeLit)
		if ok {
			if k, ok := v.Type.(*dst.ArrayType); ok {
				if s, ok := k.Elt.(*dst.SelectorExpr); ok {
					if x, ok := s.X.(*dst.Ident); ok {
						if x.Name == "t" && s.Sel.Name == "VueMenu" && isParent {
							isParent = false
							// стираем массив
							v.Elts = []dst.Expr{}
							for _, m := range p.Menu {
								v.Elts = append(v.Elts, m.print())
								if len(v.Elts) == 1 {
									v.Elts[len(v.Elts)-1].Decorations().Before = dst.NewLine
								}
								v.Elts[len(v.Elts)-1].Decorations().After = dst.NewLine
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

func (m *VueMenu) print() dst.Expr  {
	res := &dst.CompositeLit{}
	if len(m.DocName) > 0 {
		res.Elts = append(res.Elts, &dst.KeyValueExpr{Key: &dst.Ident{Name: "DocName"}, Value: &dst.BasicLit{Value: fmt.Sprintf("%q", m.DocName)}})
	}
	if len(m.Url) > 0 {
		res.Elts = append(res.Elts, &dst.KeyValueExpr{Key: &dst.Ident{Name: "Url"}, Value: &dst.BasicLit{Value: fmt.Sprintf("%q", m.Url)}})
	}
	if len(m.Text) > 0 {
		res.Elts = append(res.Elts, &dst.KeyValueExpr{Key: &dst.Ident{Name: "Text"}, Value: &dst.BasicLit{Value: fmt.Sprintf("%q", m.Text)}})
	}
	if len(m.Icon) > 0 {
		res.Elts = append(res.Elts, &dst.KeyValueExpr{Key: &dst.Ident{Name: "Icon"}, Value: &dst.BasicLit{Value: fmt.Sprintf("%q", m.Icon)}})
	}
	if len(m.Roles) > 0 {
		var roleElts []dst.Expr
		for _, r := range m.Roles {
			roleElts = append(roleElts, &dst.SelectorExpr{X: &dst.Ident{Name: "utils"}, Sel: &dst.Ident{Name: r}})
		}
		res.Elts = append(res.Elts, &dst.KeyValueExpr{Key: &dst.Ident{Name: "Roles"}, Value: &dst.CompositeLit{
			Type: &dst.ArrayType{Elt: &dst.Ident{Name: "string"}},
			Elts: roleElts,
		}})
	}
	if len(m.IsFolder) > 0 {
		res.Elts = append(res.Elts, &dst.KeyValueExpr{Key: &dst.Ident{Name: "IsFolder"}, Value: &dst.Ident{Name: m.IsFolder}})
	}
	if len(m.LinkList) > 0 {
		var menuElts []dst.Expr
		for _, m := range m.LinkList {
			menuElts = append(menuElts, m.print())
			if len(menuElts) == 1 {
				menuElts[len(menuElts)-1].Decorations().Before = dst.NewLine
			}
			menuElts[len(menuElts)-1].Decorations().After = dst.NewLine
		}
		res.Elts = append(res.Elts, &dst.KeyValueExpr{Key: &dst.Ident{Name: "LinkList"}, Value: &dst.CompositeLit{
			Type: &dst.ArrayType{Elt: &dst.SelectorExpr{X: &dst.Ident{Name: "t"}, Sel:  &dst.Ident{Name: "VueMenu"}}},
			Elts: menuElts,
		}})
	}
	return res
}