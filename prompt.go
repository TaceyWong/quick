/*
* Functions for prompting the user for project info.
* @Author: TaceyWong
* @Date:   2020-09-23 16:59:15
* @Last Modified by:   TaceyWong
* @Last Modified time: 2020-09-23 19:24:56
 */
package quick

import (
	"errors"

	"github.com/AlecAivazis/survey/v2"
	"github.com/tidwall/gjson"
)

// ReadUserVariable Prompt user for variable and return the entered value or given default.
//
// :param str varName: Variable of the context to query the user
// :param defaultValue: Value that will be returned if no input happens
func ReadUserVariable(msg string, varName *string, defaultValue string) string {
	prompt := &survey.Input{
		Message: msg,
		Default: defaultValue,
	}
	survey.AskOne(prompt, varName)
	return *varName
}

// ReadUserYesNo Prompt the user to reply with 'yes' or 'no' (or equivalent values).
//
// Note:
//       Possible choices are 'true', '1', 'yes', 'y' or 'false', '0', 'no', 'n'
// :param str question: Question to the user
// :param defaultValue: Value that will be returned if no input happens
func ReadUserYesNo(confirm *bool, question string, defaultValue bool) bool {
	prompt := &survey.Confirm{
		Message: question,
		Default: defaultValue,
	}
	survey.AskOne(prompt, confirm)
	return *confirm
}

// ReadRepoPassword Prompt the user to enter a password.
//
// :param str question: Question to the user
func ReadRepoPassword(password *string, question string) string {
	prompt := &survey.Password{
		Message: question,
	}
	survey.AskOne(prompt, password)
	return *password
}

// ReadUserChoice Prompt the user to choose from several options for the given variable.
//
// The first item will be returned if no input happens.
//
// :param str varName: Variable as specified in the context
// :param list options: Sequence of options that are available to select from
// :return: Exactly one item of ``options`` that has been chosen by the user
func ReadUserChoice(varName *string, question string, options []string) string {
	if len(options) < 1 {
		return ""
	}
	prompt := &survey.Select{
		Message: question,
		Options: options,
		Default: options[0],
	}
	survey.AskOne(prompt, varName)
	return *varName
}

// ProcessJSON Load user-supplied value as a JSON dict.
//
// :param str userValue: User-supplied value to load as a JSON dict
func ProcessJSON(userValue string) (string, error) {
	if !gjson.Valid(userValue) {
		return "", errors.New("Invalid json")
	}
	return userValue, nil
}

// ReadUserDict Prompt the user to provide a dictionary of data.
//
// :param str varName: Variable as specified in the context
// :param defaultValue: Value that will be returned if no input is provided
// :return: A Go map to use in the context.
func ReadUserDict(varName string, defaultValue string) {

}

// RenderVariable Render the next variable to be displayed in the user prompt.
//
// Inside the prompting taken from the cookiecutter.json file, this renders
// the next variable. For example, if a project_name is "Peanut Butter
// Cookie", the repo_name could be be rendered with:
//
//      `{{ quick.project_name.replace(" ", "_") }}`.
//
// This is then presented to the user as the default.
//
// :param Environment env: A Jinja2 Environment object.
// :param raw: The next value to be prompted for by the user.
// :param dict quicDict: The current context as it's gradually
//                       being populated with variables.
// :return: The rendered value for the default variable.
func RenderVariable(env, raw string, quickDict map[string]string) {

}

// PromptChoiceForConfig Prompt user with a set of options to choose from.
//
// Each of the possible choices is rendered beforehand.
func PromptChoiceForConfig(quickDict, env map[string]string, key string, options []string, noInput bool) {

}

// PromptForConfig Prompt user to enter a new config.
func PromptForConfig(context string, noInput bool) {

}
