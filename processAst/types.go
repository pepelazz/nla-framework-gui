package processAst

import "github.com/dave/dst"

const (
	GET_FLD_TITLE = "GetFldTitle"
	GET_FLD_TITLE_COMPUTED = "GetFldTitleComputed"
	GET_FLD_STRING = "GetFldString"
	GET_FLD_SELECT_STRING = "GetFldSelectString"
	GET_FLD_SELECT_MULTIPLE = "GetFldSelectMultiple"
	GET_FLD_RADIO_STRING = "GetFldRadioString"
	GET_FLD_REF = "GetFldRef"
	GET_FLD_FILES = "GetFldFiles"
	GET_FLD_IMG = "GetFldImg"
	GET_FLD_IMG_LIST = "GetFldImgList"
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
		FldVueImgParams map[string]string `json:"fld_vue_img_params"`
		ModifierList      []FldModifier `json:"modifier_list"`
	}

	FldModifier struct {
		Name string
		Args []string
	}
)
