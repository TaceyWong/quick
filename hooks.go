/*
* Functions for discovering and executing various Starter hooks.
* @Author: TaceyWong
* @Date:   2020-09-23 16:58:43
* @Last Modified by:   TaceyWong
* @Last Modified time: 2020-09-23 18:07:34
*/
package starter


var Hooks = []string{"pre_gen_project","post_gen_project"} 


// ValidHook Determine if a hook file is valid.
// 
// :param hookFile: The hook file to consider for validity
// :param hookName: The hook to find
// :return: The hook file validity false or true
func ValidHook(hookFile, hookName string)bool{
	return true	
}

// FindHook Return a dict of all hook scripts provided.
// 
// Must be called with the project template as the current working directory.
// Dict's key will be the hook/script's name, without extension, while values
// will be the absolute path to the script. Missing scripts will not be
// included in the returned dict.
// 
// :param hookName: The hook to find
// :param hooksDir: The hook directory in the template
// :return: The absolute path to the hook script or null-str-list
func FindHook(hookName,hookDir string)[]string{
	return []string{}
}


// RunScript Execute a script from a working directory.
// 
// :param scriptPath: Absolute path to the script to run.
// :param cwd: The directory to run the script from.
func RunScript(scriptPath, cwd string){

}


// RunScriptWithContext Execute a script after rendering it with Template Engine.
// 
// :param scriptPath: Absolute path to the script to run.
// :param cwd: The directory to run the script from.
// :param context: Starter project template context.
func RunScriptWithContext(scriptPath,cwd string,context interface{}){

}


// RunHook Try to find and execute a hook from the specified project directory.
// 
// :param hookName: The hook to execute.
// :param projectdir: The directory to execute the script from.
// :param context: Starter project context.
func RunHook(hookName,projectDir string,context interface{}){

}