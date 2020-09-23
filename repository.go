/*
* Quick repository functions.
* @Author: TaceyWong
* @Date:   2020-09-23 16:59:40
* @Last Modified by:   TaceyWong
* @Last Modified time: 2020-09-23 19:36:17
*/
package quick


// IsRepoURL Return True if value is a repository URL.
func IsRepoURL(value)bool{
	return false
}


// IsZipFile Return True if value is a zip file.
func IsZipFile(value)bool{
	return false
}

// ExpandAbbreviations Expand abbreviations in a template name.
func ExpandAbbreviations(template,abbreviations string){

}

// RepositoryHasQuickJSON Determine if `repoDirectory` contains a `quick.json` file.
// 
// :param repoDirectory: The candidate repository directory.
// :return: True if the `repo_directory` is valid, else False.
func RepositoryHasQuickJSON(repoDirectory string)bool{
	return false
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
// :param directory: Directory within repo where quick.json lives.
// :return: A tuple containing the quick template directory, and
//          a boolean descriving whether that directory should be cleaned up
//          after the template has been instantiated.
// :raises: `RepositoryNotFound` if a repository directory could not be found.
func DetermineRepoDir(template,abbreviations,cloneToDir string,checkout,
	noInput bool,password , directory=string){

}
