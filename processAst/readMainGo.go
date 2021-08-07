package processAst

import (
	"errors"
	"fmt"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"go/token"
	"os"
	"strconv"
	"strings"
)

// чтение Flds из файла main.go
func (doc *Doc) ReadMainGo() error {
	b, err := os.ReadFile(fmt.Sprintf("../projectTemplate/docs/%s/main.go", doc.Name))
	if err != nil {
		return err
	}
	dstFile, err := decorator.Parse(b)
	if err != nil {
		return err
	}
	doc.DstFile = dstFile

	// находим node GetDoc
	getDocNode := findFuncDecl(dstFile, "GetDoc")
	if getDocNode == nil {
		return errors.New("GetDoc not found")
	}
	doc.Flds = []Fld{}

	// итерируем по поляем
	for _, v := range getFldsElts(getDocNode) {
		// дефолтный класс col-4
		fld := Fld{ColClass: "col-4"}
		dst.Inspect(v, func(node dst.Node) bool {
			if n, ok := node.(*dst.CallExpr); ok {
				if fun, ok := n.Fun.(*dst.SelectorExpr); ok {
					if strings.HasPrefix(fun.Sel.Name, "Get") {
						fld.FuncName = fun.Sel.Name
						if len(n.Args) > 0 {
							for i, arg := range n.Args {
								if arg, ok := arg.(*dst.BasicLit); ok {
									// аргументы GetFldTitle обрабатываем отдельно
									if fld.FuncName == GET_FLD_TITLE || fld.FuncName == GET_FLD_TITLE_COMPUTED {
										if arg.Kind == token.STRING {
											argVal := replaceQuotes(arg.Value)
											// отдельно - ширину колонки
											if strings.HasPrefix(argVal, "col-") {
												fld.ColClass = argVal
											} else {
												fld.Params = append(fld.Params, argVal)
											}
										}
									} else {
										if i == 0 && arg.Kind == token.STRING {
											fld.Name = replaceQuotes(arg.Value)
										}
										if i == 1 && arg.Kind == token.STRING {
											fld.NameRu = replaceQuotes(arg.Value)
										}
										if i == 2 && arg.Kind == token.INT {
											fld.Size, _ = strconv.Atoi(arg.Value)
										}
										if i == 2 && arg.Kind == token.STRING && fld.FuncName == GET_FLD_REF {
											fld.RefTable = replaceQuotes(arg.Value)
										}
										// обработка params
										if i > 2 && arg.Kind == token.STRING {
											argVal := replaceQuotes(arg.Value)
											// отдельно - ширину колонки
											if strings.HasPrefix(argVal, "col-") {
												fld.ColClass = argVal
											} else {
												fld.Params = append(fld.Params, argVal)
											}
										}
									}
								}
								// поиск что [][]int{{}}
								if len(fld.RowCol) == 0 {
									rc := findRowCol(arg)
									if len(rc) > 0 {
										fld.RowCol = rc
									}
								}
								if fld.FuncName == GET_FLD_SELECT_STRING || fld.FuncName == GET_FLD_SELECT_MULTIPLE || fld.FuncName == GET_FLD_RADIO_STRING {
									res := findArrType(arg, "FldVueOptionsItem")
									if len(res) > 0 {
										fld.FldVueOptionsItem = res
									}
								}

								if fld.FuncName == GET_FLD_FILES {
									res := findObjectType(arg, "FldVueFilesParams")
									if len(res) > 0 {
										fld.FldVueFilesParams = res
									}
								}

								if fld.FuncName == GET_FLD_IMG || fld.FuncName == GET_FLD_IMG_LIST{
									res := findObjectType(arg, "FldVueImgParams")
									if len(res) > 0 {
										fld.FldVueImgParams = res
									}
								}

							}
						}
					} else {
						fm := FldModifier{Name: fun.Sel.Name}
						for _, arg := range n.Args {
							if arg, ok := arg.(*dst.BasicLit); ok {
								fm.Args = append(fm.Args, replaceQuotes(arg.Value))
							}
						}
						fld.ModifierList = append(fld.ModifierList, fm)
					}
				}
			}
			return true
		})
		if len(fld.FuncName) > 0 {
			doc.Flds = append(doc.Flds, fld)
		}
	}

	// чистим массив Flds в dst.Node
	dst.Inspect(getDocNode, func(n dst.Node) bool {
		v, ok := n.(*dst.KeyValueExpr)
		if ok {
			if k, ok := v.Key.(*dst.Ident); ok && k.Name == "Flds" {
				if elts, ok := v.Value.(*dst.CompositeLit); ok {
					elts.Elts = []dst.Expr{}
				}
				return true
			}
		}
		return true
	})

	// читаем значения констант
	dst.Inspect(doc.DstFile, func(n dst.Node) bool {
		v, ok := n.(*dst.GenDecl)
		if ok && v.Tok == token.CONST {
			for _, spec := range v.Specs {
				if vs, ok := spec.(*dst.ValueSpec); ok {
					key := vs.Names[0].Name
					if key == "name" {
						doc.Name = replaceQuotes(vs.Values[0].(*dst.BasicLit).Value)
					}
					if key == "name_ru" {
						doc.NameRu = replaceQuotes(vs.Values[0].(*dst.BasicLit).Value)
					}
					if key == "name_ru_plural" {
						doc.NameRuPlural = replaceQuotes(vs.Values[0].(*dst.BasicLit).Value)
					}
					if key == "menu_icon" {
						doc.MenuIcon = replaceQuotes(vs.Values[0].(*dst.BasicLit).Value)
					}
				}
			}
		}
		return true
	})

	return nil
}

func getFldsElts(node *dst.FuncDecl) (res []dst.Expr) {
	dst.Inspect(node, func(n dst.Node) bool {
		v, ok := n.(*dst.KeyValueExpr)
		if ok {
			if k, ok := v.Key.(*dst.Ident); ok && k.Name == "Flds" {
				if elts, ok := v.Value.(*dst.CompositeLit); ok {
					res = elts.Elts
				}
				return true
			}
		}
		return true
	})
	return
}

func findRowCol(arg dst.Expr) [][]int {
	var arr [][]int
	dst.Inspect(arg, func(node dst.Node) bool {
		if n, ok := node.(*dst.CompositeLit); ok {
			if len(n.Elts) == 2 {
				if intVal1, ok := n.Elts[0].(*dst.BasicLit); ok {
					if intVal2, ok := n.Elts[1].(*dst.BasicLit); ok {
						if intVal1.Kind == token.INT && intVal2.Kind == token.INT {
							int1, _ := strconv.Atoi(intVal1.Value)
							int2, _ := strconv.Atoi(intVal2.Value)
							arr = append(arr, []int{int1, int2})
							//fmt.Printf("[%v, %v]", intVal1.Value, intVal2.Value)
						}
					}
				}
			}
		}
		return true
	})
	return arr
}

func findArrType(arg dst.Expr, typeName string) (res []map[string]string) {
	dst.Inspect(arg, func(node dst.Node) bool {
		if n, ok := node.(*dst.CompositeLit); ok {
			if elType, ok := n.Type.(*dst.ArrayType); ok {
				if elt, ok := elType.Elt.(*dst.SelectorExpr); ok {
					// проверяем тип массива
					if elt.Sel.Name == typeName {
						dst.Inspect(n, func(node dst.Node) bool {
							if n, ok := node.(*dst.CompositeLit); ok {
								m := map[string]string{}
								for _, elt := range n.Elts {
									if el, ok := elt.(*dst.KeyValueExpr); ok {
										var key, value string
										if k, ok := el.Key.(*dst.Ident); ok {
											key = k.Name
										}
										if v, ok := el.Value.(*dst.BasicLit); ok {
											value = replaceQuotes(v.Value)
										}
										m[key] = value
									}
								}
								if len(m) > 0{
									res = append(res, m)
								}
							}

							return true
						})
					}
				}
			}
		}
		return true
	})
	return res
}

func findObjectType(arg dst.Expr, typeName string) (res map[string]string) {
	res = map[string]string{}
	dst.Inspect(arg, func(node dst.Node) bool {
		if n, ok := node.(*dst.CompositeLit); ok {
			if elType, ok := n.Type.(*dst.SelectorExpr); ok {
				// проверяем тип объекта
				if elType.Sel.Name == typeName {
					dst.Inspect(n, func(node dst.Node) bool {
						if n, ok := node.(*dst.CompositeLit); ok {
							for _, elt := range n.Elts {
								if el, ok := elt.(*dst.KeyValueExpr); ok {
									var key, value string
									if k, ok := el.Key.(*dst.Ident); ok {
										key = k.Name
									}
									if v, ok := el.Value.(*dst.BasicLit); ok {
										value = replaceQuotes(v.Value)
									}
									if v, ok := el.Value.(*dst.Ident); ok {
										value = replaceQuotes(v.Name)
									}
									res[key] = value
								}
							}
						}

						return true
					})
				}
			}
		}
		return true
	})
	return res
}

func replaceQuotes(str string) string {
	return strings.Replace(str, "\"", "", -1)
}
