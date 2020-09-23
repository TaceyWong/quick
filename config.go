/*
* Global configuration handling.
* @Desc: 
* @Author: TaceyWong
* @Date:   2020-09-23 16:58:12
* @Last Modified by:   TaceyWong
* @Last Modified time: 2020-09-23 17:43:25
*/

package quick

const UserConfigPath string = ""

var BuiltinAbbreviations = map[string]string{
	"gh": "https://github.com/{0}.git",
    "gl": "https://gitlab.com/{0}.git",
    "bb": "https://bitbucket.org/{0}",
}

var DefaultConfig = map[string]interface{}{
	"quick_dir": "~/.quick/",
    "replay_dir": "~/.quick_replay/",
    "default_context"': map[string][]string{},
    "abbreviations"': BuiltinAbbreviations,
}


// map[string]interface Recursively update a dict with the key/value pair of another.
// 
// Dict values that are dictionaries themselves will be updated, whilst preserving existing keys.
func MergeConfigs(defaultC,overwriteC map[string]interface{})map[string]interface{}{
	return nil
}

// GetConfig Retrieve the config from the specified path, returning a config dict.
func GetConfig(configPath string)map[string]interface{}{
	return nil
}

// GetUserConfig Return the user config as a dict.
// 
//  If ``default_config`` is True, ignore ``config_file`` and return default
//  values for the config parameters.
//  
//  If a path to a ``config_file`` is given, that is different from the default
//  location, load the user config from that.
//  
//  Otherwise look up the config file path in the ``COOKIECUTTER_CONFIG`
//  environment variable. If set, load the config from this path. This will
//  raise an error if the specified path is not valid.
//  
//  If the environment variable is not set, try the default config file path
//  before falling back to the default config values.
func GetUserConfig(configFile string,deaultConfig bool)map[string]interface{}{
	return nil
}