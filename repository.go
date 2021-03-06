/*
* Starter repository functions.
* @Author: TaceyWong
* @Date:   2020-09-23 16:59:40
* @Last Modified by:   TaceyWong
* @Last Modified time: 2020-09-23 19:36:17
 */
package starter

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Regex  something like git:// ssh:// file:// etc.
//        or
//        something like user@...
var Regex = regexp.MustCompile(`((((git|hg)\+)?(git|ssh|file|https?):(//)?)|(\w+@[\w\.]+))`)

// IsRepoURL Return True if value is a repository URL.
func IsRepoURL(value string) bool {
	return Regex.Match([]byte(value))
}

// IsZipFile Return True if value is a zip file.
func IsZipFile(value string) bool {
	value = strings.ToLower(value)
	if strings.HasSuffix(value, ".zip") {
		return true
	}
	return false
}

// ExpandAbbreviations Expand abbreviations in a template name.
func ExpandAbbreviations(template, abbreviations string) {

}

// RepositoryHasStarterJSON Determine if `repoDirectory` contains a `starter.json` file.
//
// :param repoDirectory: The candidate repository directory.
// :return: True if the `repo_directory` is valid, else False.
func RepositoryHasStarterJSON(repoDirectory string) bool {
	_, err := os.Stat(repoDirectory)
	if err != nil || os.IsNotExist(err) {
		return false
	}
	_, err = os.Stat(repoDirectory + "starter.json")
	if err != nil || os.IsNotExist(err) {
		return false
	}
	return true
}

// DetermineRepoDir Locate the repository directory from a template reference.
//
// Applies repository abbreviations to the template reference.
// If the template refers to a repository URL, clone it.
// If the template is a path to a local repository, use it.
//
// :param template: A directory containing a project template directory,
//                  or a URL to a git repository.
// :param abbreviations: A dictionary of repository abbreviation definitions.
// :param clone_to_dir: The directory to clone the repository into.
// :param checkout: The branch, tag or commit ID to checkout after clone.
// :param no_input: Prompt the user at command line for manual configuration?
// :param password: The password to use when extracting the repository.
// :param directory: Directory within repo where starter.json lives.
// :return: A tuple containing the starter template directory, and
//          a boolean descriving whether that directory should be cleaned up
//          after the template has been instantiated.
// :raises: `RepositoryNotFound` if a repository directory could not be found.
func DetermineRepoDir(template, abbreviations, cloneToDir string, checkout,
	noInput bool, password, directory string) {
	cleanUp := false
	if IsZipFile(template) {
		_ = UnZip("dir", false, "password")
		cleanUp = true
	} else if IsRepoURL(template) {
		cleanUp = false
	} else {
		cleanUp = false
	}
	fmt.Println(cleanUp)

}
