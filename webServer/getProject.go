package webServer

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pepelazz/guiDev/nla_framework_gui/processAst"
	"github.com/pepelazz/guiDev/src/utils"
	"github.com/serenize/snaker"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"text/template"
)


func GetProject(c *gin.Context)  {

	err := processAst.ReadProject()
	if err != nil {
		utils.HttpError(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(200, map[string]interface{}{"ok": true, "result": processAst.Project})
}

func SaveProject(c *gin.Context)  {
	params := struct {
		Project processAst.ProjectType `json:"params"`
	}{}
	err := c.BindJSON(&params)
	if err != nil {
		utils.HttpError(c, http.StatusBadRequest, err.Error())
		return
	}

	// копируем в результат doc.DstFile, а потом обновляем весь проект
	params.Project.DstFile = processAst.Project.DstFile
	for i, d := range params.Project.Docs {
		isFound := false
		for _, doc := range processAst.Project.Docs {
			if doc.Name == d.Name {
				params.Project.Docs[i].DstFile = doc.DstFile
				isFound = true
				break
			}
		}
		if !isFound {
			utils.HttpError(c, http.StatusBadRequest, fmt.Sprintf("для doc %q не найден DstFile", d.Name))
			return
		}
	}
	processAst.Project = params.Project
	err = processAst.WriteProject()
	if err != nil {
		utils.HttpError(c, http.StatusBadRequest, err.Error())
		return
	}

	ProjectTemplateGenerate(c)
}

func AddDoc(c *gin.Context)  {
	p := struct {
		Params struct{
			Name string `json:"name"`
			NameRu string `json:"name_ru"`
			NameRuPlural string `json:"name_ru_plural"`
			NameCamelCase string `json:"name_camel_case"`
		} `json:"params"`
	}{}
	err := c.BindJSON(&p)
	if err != nil {
		utils.HttpError(c, http.StatusBadRequest, err.Error())
		return
	}

	p.Params.NameCamelCase = snaker.SnakeToCamelLower(p.Params.Name)

	// читаем темплейт
	docMainTmpl, err := template.New("docMain.go").Delims("[[", "]]").ParseFiles("./tmpl/docMain.go")
	if err != nil {
		utils.HttpError(c, http.StatusBadRequest, "parse template: " + err.Error())
		return
	}
	// генерим темплейт в буфер
	var buf bytes.Buffer
	err = docMainTmpl.Execute(&buf, p.Params)
	if err != nil {
		utils.HttpError(c, http.StatusBadRequest, "docMainTmpl.Execute: " + err.Error())
		return
	}

	pathDir := "../projectTemplate/docs/"+p.Params.NameCamelCase
	_ = os.Mkdir(pathDir, 0644)


	// записываем файл
	err = ioutil.WriteFile(pathDir + "/main.go", []byte(buf.String()), 0644)
	if err != nil {
		utils.HttpError(c, http.StatusBadRequest, err.Error())
		return
	}

	// вносим изменение в main.go

	utils.HttpSuccess(c, "ok")
}

// генерация проекта
func ProjectTemplateGenerate(c *gin.Context) {
	err := os.Chdir("../projectTemplate")
	if err != nil {
		utils.HttpError(c, http.StatusBadRequest, err.Error())
		return
	}
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("go", "run", ".")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		utils.HttpError(c, http.StatusBadRequest, fmt.Sprintf("error: %s\n%s", err, cmd.Stderr))
		return
	}

	c.JSON(200, map[string]interface{}{"ok": true, "result": "ok"})
}

