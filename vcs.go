/*
* @Author: TaceyWong
* @Date:   2020-09-23 16:59:52
* @Last Modified by:   TaceyWong
* @Last Modified time: 2020-09-23 17:14:03
 */
package quick

import (
	"log"
	"os"
	"strings"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/mitchellh/go-homedir"
)

// IdentifyRepo Determine if `repo_url` should be treated as a URL to a git or hg repo.
//
// Repos can be identified by prepending "hg+" or "git+" to the repo URL.
//
// :param repo_url: Repo URL of unknown type.
// :returns: ('git', repo_url), ('hg', repo_url), or None.
func IdentifyRepo(repoURL string) (string, string, error) {
	repoURLValues := strings.Split(repoURL, "+")
	if len(repoURLValues) == 2 {
		repoType := repoURLValues[0]
		if repoType != "git" || repoType != "hg" {
			return "", "", nil // UnknownRepoType
		}
		return repoType, repoURLValues[1], nil
	}
	if strings.Contains(repoURL, "git") {
		return "git", repoURL, nil
	} else if strings.Contains(repoURL, "bitbucket") {
		return "hg", repoURL, nil
	}
	return "", "", nil // UnknownRepoType
}

// Clone Clone a repo to the current directory.
//
// :param repo_url: Repo URL of unknown type.
// :param checkout: The branch, tag or commit ID to checkout after clone.
// :param clone_to_dir: The directory to clone to.
// 					 Defaults to the current directory.
// :param no_input: Suppress all user prompts when calling via API.
func Clone(repoURL string, checkout string, cloneToDir string, noInput bool) string {
	cloneToDir, err := homedir.Expand(cloneToDir)
	if err != nil {
		log.Panic(err)
	}
	MakeSurePathExist(cloneToDir)
	repoType, repoUrl, err := IdentifyRepo(repoURL)
	if err != nil {
		log.Panic(err)
	}
	// "github.com/go-git/go-git/v5/plumbing/transport/http"
	// Auth: &http.BasicAuth{
	// 	Username: username,
	// 	Password: password,
	// },
	// "github.com/go-git/go-git/v5/plumbing/transport/http"
	// Auth: &http.BasicAuth{
	// 	Username: "abc123", // yes, this can be anything except an empty string
	// 	Password: token,
	// },
	publicKeys, _ := ssh.NewPublicKeysFromFile("git", privateKeyFile, "")
	r, _ := git.PlainClone(cloneToDir, false, &git.CloneOptions{
		URL:               repoURL,
		Progress:           os.Stdout,
		Auth:     publicKeys,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
	ref,_ := r.Head()
	commit, _ := r.CommitObject(ref.Hash())

	return ""
}
