/*
* Functions for finding Quick templates and other components.
* @Author: TaceyWong
* @Date:   2020-09-23 17:00:36
* @Last Modified by:   TaceyWong
* @Last Modified time: 2020-09-23 17:47:30
*/

package quick

// FindTemplate Determine which child directory of `repo_dir` is the project template.
// 
// :param repoDir: Local directory of newly cloned repo.
// :returns projectTemplate: Relative path to project template.
func FindTemplate(repoDir string)string{
	return "~/.quick"
}

