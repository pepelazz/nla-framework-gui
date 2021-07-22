package processAst

import (
	"fmt"
	"github.com/dave/dst"
	"go/token"
	"strconv"
)

func (fld *Fld) Print() *dst.CallExpr  {
	res := &dst.CallExpr{
		Fun: &dst.SelectorExpr{X: &dst.Ident{Name: "t"}, Sel: &dst.Ident{Name: fld.FuncName}},
	}
	args := []dst.Expr{}
	if len(fld.Name) > 0 {
		args = append(args, &dst.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("%q", fld.Name)})
	}
	if len(fld.NameRu) > 0 {
		args = append(args, &dst.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("%q", fld.NameRu)})
	}
	if fld.FuncName == "GetFldSelectString" ||  fld.FuncName == "GetFldString"{
		args = append(args, &dst.BasicLit{Kind: token.INT, Value: fmt.Sprintf("%v", fld.Size)})
	}
	if fld.FuncName == "GetFldRef" {
		args = append(args, &dst.BasicLit{Kind: token.INT, Value:  fmt.Sprintf("%q", fld.RefTable)})
	}
	// печать RowCol
	if len(fld.RowCol)>0 {
		rows := &dst.CompositeLit{
			Type: &dst.ArrayType{
				Elt: &dst.ArrayType{
					Elt: &dst.Ident{Name: "int"},
				},
			},
			Elts: []dst.Expr{},
		}
		for _, r := range fld.RowCol {
			rows.Elts = append(rows.Elts, &dst.CompositeLit{
				Elts: []dst.Expr {
					&dst.BasicLit{ Kind: token.INT, Value: strconv.Itoa(r[0])},
					&dst.BasicLit{ Kind: token.INT, Value: strconv.Itoa(r[1])},
				},
			},)
		}
		args = append(args, rows)
	}
	if fld.FuncName == "GetFldSelectString" {
		args = append(args, printFldVueOptionsItem(fld))
	}
	if fld.FuncName == "GetFldFiles" {
		args = append(args, printFldVueFilesParams(fld))
	}
	// печать params
	args = append(args, &dst.BasicLit{Kind: token.STRING, Value:  fmt.Sprintf("%q", fld.ColClass)})
	for i, p := range fld.Params {
		// для GET_FLD_TITLE_COMPUTED первый параметр ставим первым аргументом в функции
		if i == 0 && fld.FuncName == GET_FLD_TITLE_COMPUTED {
			args = append([]dst.Expr{&dst.BasicLit{Kind: token.STRING, Value:  fmt.Sprintf("%q", p)}}, args...)
			continue
		}
		args = append(args, &dst.BasicLit{Kind: token.STRING, Value:  fmt.Sprintf("%q", p)})
	}
	res.Args = args

	// добавляем функции-модификаторы
	for i := len(fld.ModifierList)-1; i >= 0; i-- {
		args := []dst.Expr{}
		for _, v := range fld.ModifierList[i].Args {
			args = append(args, &dst.BasicLit{Kind: token.STRING, Value: fmt.Sprintf("%q", v)})
		}
		res = &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X: res,
				Sel: &dst.Ident{
					Name: fld.ModifierList[i].Name,
				},
			},
			Args: args,
		}
	}

	return res
}

func printFldVueOptionsItem(fld *Fld) *dst.CompositeLit {
	el := &dst.CompositeLit{
		Type: &dst.ArrayType{
			Elt: &dst.SelectorExpr{X: &dst.Ident{Name: "t"}, Sel: &dst.Ident{Name: "FldVueOptionsItem"}},
		},
		Elts: []dst.Expr{},
	}
	for _, v := range fld.FldVueOptionsItem {
		if len(v) == 0 {
			continue
		}
		el1 := &dst.CompositeLit{Elts: []dst.Expr{}}
		for label, value := range v {
			el1.Elts = append(el1.Elts, &dst.KeyValueExpr{Key: &dst.Ident{ Name: label}, Value: &dst.BasicLit{Kind: token.STRING, Value: value}})
		}
		el.Elts = append(el.Elts, el1)
	}
	return el
}

func printFldVueFilesParams(fld *Fld) *dst.CompositeLit {
	el := &dst.CompositeLit{
		Type: &dst.SelectorExpr {X: &dst.Ident{Name: "t"}, Sel: &dst.Ident{Name: "FldVueFilesParams"}},
		Elts: []dst.Expr{},
	}
	for label, value := range fld.FldVueFilesParams {
		el.Elts = append(el.Elts, &dst.KeyValueExpr{Key: &dst.Ident{ Name: label}, Value: &dst.BasicLit{Kind: token.STRING, Value: value}})
	}
	return el
}
