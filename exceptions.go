/*
* All exceptions used in the Quick code base are defined here.
* @Author: TaceyWong
* @Date:   2020-09-23 17:00:44
* @Last Modified by:   TaceyWong
* @Last Modified time: 2020-09-23 17:30:34
*/
package quick

// QuickException Base exception class.
// 
// All Quick-specific exceptions should subclass this class.
type QuickException struct{}

// NonTemplatedInputDirException Exception for when a project's input dir is not templated.
// 
// The name of the input directory should always contain a string that is
// rendered to something else, so that input_dir != output_dir.
type NonTemplatedInputDirException struct{
	QuickException
}

type UnknownTemplateDirException struct{
	QuickException
}

// MissingProjectDir Exception for missing generated project directory.
// 
// Raised during cleanup when remove_repo() can't find a generated project
// directory inside of a repo.
type MissingProjectDir struct{
	QuickException
}

// ConfigDoesNotExistException Exception for missing config file.
// 
// Raised when get_config() is passed a path to a config file, but no file
// is found at that path.
type ConfigDoesNotExistException struct {
	QuickException
}

// InvalidConfiguration Exception for invalid configuration file.
// 
// Raised if the global configuration file is not valid YAML or is
// badly constructed.
type InvalidConfiguration struct{
	QuickException
}

// UnknownRepoType Exception for unknown repo types
// 
// Raised if a repo's type cannot be determined.
type UnknownRepoType struct{
	QuickException
}

// VCSNotInstalled Exception when version control is unavailable.
// 
// Raised if the version control system (git or hg) is not installed.
type VCSNotInstalled struct{
	QuickException
}

// ContextDecodingException Exception for failed JSON decoding.
// 
// Raised when a project's JSON context file can not be decoded.
type ContextDecodingException struct{
	QuickException
}

// OutputDirExistsException Exception for existing output directory.
// 
// Raised when the output directory of the project exists already.
type OutputDirExistsException struct {
	QuickException
}

// InvalidModeException Exception for incompatible modes.
// 
// Raised when cookiecutter is called with both `no_input==True` and
// `replay==True` at the same time.
type InvalidModeException struct {
	QuickException
}


// FailedHookException Exception for hook failures.
// 
// Raised when a hook script fails.
type FailedHookException struct {
	QuickException
}

// UndefinedVariableInTemplate Exception for out-of-scope variables.
// 
// Raised when a template uses a variable which is not defined in the
// context
type UndefinedVariableInTemplate struct{
	QuickException
}

// UnknownExtension Exception for un-importable extention.
// 
// Raised when an environment is unable to import a required extension.
type UnknownExtension struct{
	QuickException
}

// RepositoryNotFound Exception for missing repo.
// 
// Raised when the specified cookiecutter repository doesn't exist.
type RepositoryNotFound struct {
	QuickException
}


// RepositoryCloneFailed Exception for un-cloneable repo.
// 
// Raised when a cookiecutter template can't be cloned.
type RepositoryCloneFailed struct {
	QuickException
}

// InvalidZipRepository Exception for bad zip repo.
// 
// Raised when the specified cookiecutter repository isn't a valid 
// Zip archive.
type InvalidZipRepository struct {
	QuickException
}




