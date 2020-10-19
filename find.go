/*
* Functions for finding Starter templates and other components.
* @Author: TaceyWong
* @Date:   2020-09-23 17:00:36
* @Last Modified by:   TaceyWong
* @Last Modified time: 2020-09-23 17:47:30
 */

package starter

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// FindTemplate Determine which child directory of `repo_dir` is the project template.
//
// :param repoDir: Local directory of newly cloned repo.
// :returns projectTemplate: Relative path to project template.
func FindTemplate(repoDir string) string {
	log.Println(fmt.Sprintf("Searching %s for the project template.", repoDir))
	repo_dir_contents, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}
	projectTemplate := ""
	for _, item := range repo_dir_contents {
		name := item.Name()
		if strings.Contains(name, "starter") && strings.Contains(name, "{{") &&
			strings.Contains(name, "}}") {
			projectTemplate = name
			break
		}
	}
	if projectTemplate != "" {
		projectTemplate = repoDir + projectTemplate
		log.Println(fmt.Sprintf("The project template appears to be %s", projectTemplate))
		return projectTemplate
	}
	return "~/.starter" // NonTemplatedInputDirException
}
