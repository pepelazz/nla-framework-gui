package webServer

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pepelazz/guiDev/nla_framework_gui/processAst"
	"github.com/pepelazz/guiDev/src/utils"
	"net/http"
	"os"
	"os/exec"
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

