/*
* Utility functions for handling and fetching repo archives in zip format
* @Author: TaceyWong
* @Date:   2020-09-23 17:00:26
* @Last Modified by:   TaceyWong
* @Last Modified time: 2020-09-23 18:12:04
*/
package quick


// UnZip Download and unpack a zipfile at a given URI.
// 
// This will download the zipfile to the cookiecutter repository,
// and unpack into a temporary directory.
// 
// :param zipURI: The URI for the zipfile.
// :param isURL: Is the zip URI a URL or a file?
// :param cloneToDir: The cookiecutter repository directory
//                    to put the archive into.
// :param noInput: Suppress any prompts
// :param password: The password to use when unpacking the repository.
func UnZip(zipURI string,isURL bool,cloneToDir string,noInput bool,password string){

}