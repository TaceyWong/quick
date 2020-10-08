/*
* All exceptions used in the Starter code base are defined here.
* @Author: TaceyWong
* @Date:   2020-09-23 17:00:44
* @Last Modified by:   TaceyWong
* @Last Modified time: 2020-09-23 17:30:34
 */
package starter

// StarterException Base exception class.
//
// All Starter-specific exceptions should subclass this class.
type StarterException struct {
}

func (qe *StarterException) Error() string {
	return "Quick Exception"
}

// NonTemplatedInputDirException Exception for when a project's input dir is not templated.
//
// The name of the input directory should always contain a string that is
// rendered to something else, so that input_dir != output_dir.
type NonTemplatedInputDirException struct {
	StarterException
}

type UnknownTemplateDirException struct {
	StarterException
}

// MissingProjectDir Exception for missing generated project directory.
//
// Raised during cleanup when remove_repo() can't find a generated project
// directory inside of a repo.
type MissingProjectDir struct {
	StarterException
}

// ConfigDoesNotExistException Exception for missing config file.
//
// Raised when get_config() is passed a path to a config file, but no file
// is found at that path.
type ConfigDoesNotExistException struct {
	StarterException
}

// InvalidConfiguration Exception for invalid configuration file.
//
// Raised if the global configuration file is not valid YAML or is
// badly constructed.
type InvalidConfiguration struct {
	StarterException
}

// UnknownRepoType Exception for unknown repo types
//
// Raised if a repo's type cannot be determined.
type UnknownRepoType struct {
	StarterException
}

func (ur *UnknownRepoType)Error()string{
	return "Starter Exception:UnknownRepoType"
}

// VCSNotInstalled Exception when version control is unavailable.
//
// Raised if the version control system (git or hg) is not installed.
type VCSNotInstalled struct {
	StarterException
}

// ContextDecodingException Exception for failed JSON decoding.
//
// Raised when a project's JSON context file can not be decoded.
type ContextDecodingException struct {
	StarterException
}

// OutputDirExistsException Exception for existing output directory.
//
// Raised when the output directory of the project exists already.
type OutputDirExistsException struct {
	StarterException
}

// InvalidModeException Exception for incompatible modes.
//
// Raised when cookiecutter is called with both `no_input==True` and
// `replay==True` at the same time.
type InvalidModeException struct {
	StarterException
}

// FailedHookException Exception for hook failures.
//
// Raised when a hook script fails.
type FailedHookException struct {
	StarterException
}

// UndefinedVariableInTemplate Exception for out-of-scope variables.
//
// Raised when a template uses a variable which is not defined in the
// context
type UndefinedVariableInTemplate struct {
	StarterException
}

// UnknownExtension Exception for un-importable extention.
//
// Raised when an environment is unable to import a required extension.
type UnknownExtension struct {
	StarterException
}

// RepositoryNotFound Exception for missing repo.
//
// Raised when the specified cookiecutter repository doesn't exist.
type RepositoryNotFound struct {
	StarterException
}

// RepositoryCloneFailed Exception for un-cloneable repo.
//
// Raised when a cookiecutter template can't be cloned.
type RepositoryCloneFailed struct {
	StarterException
}

// InvalidZipRepository Exception for bad zip repo.
//
// Raised when the specified cookiecutter repository isn't a valid
// Zip archive.
type InvalidZipRepository struct {
	StarterException
}
