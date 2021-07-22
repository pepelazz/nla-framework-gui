package processAst

import "github.com/dave/dst"

const (
	GET_FLD_TITLE = "GetFldTitle"
	GET_FLD_TITLE_COMPUTED = "GetFldTitleComputed"
)

type (
	ProjectType struct {
		Docs []Doc `json:"docs"`
	}
	Doc struct {
		Name         string `json:"name"`
		NameRu       string `json:"name_ru"`
		NameRuPlural string `json:"name_ru_plural"`
		MenuIcon     string `json:"menu_icon"`
		Flds         []Fld `json:"flds"`
		DstFile      *dst.File `json:"-"`
	}
	Fld struct {
		FuncName          string `json:"func_name"`
		Name              string `json:"name"`
		NameRu            string `json:"name_ru"`
		RowCol            [][]int `json:"row_col"`
		ColClass          string `json:"col_class"`
		Size              int `json:"size"`
		RefTable          string `json:"ref_table"`
		Params            []string `json:"params"`
		FldVueOptionsItem []map[string]string `json:"fld_vue_options_item"`
		FldVueFilesParams map[string]string `json:"fld_vue_files_params"`
		ModifierList      []FldModifier `json:"modifier_list"`
	}

	FldModifier struct {
		Name string
		Args []string
	}
)
