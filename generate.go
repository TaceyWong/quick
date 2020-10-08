/*
* Functions for generating a project from a project template.
* @Author: TaceyWong
* @Date:   2020-09-23 17:00:57
* @Last Modified by:   TaceyWong
* @Last Modified time: 2020-09-23 17:48:32
*/
package starter
import (
	"github.com/flosch/pongo2/v4"
)

func Example(){
	// Compile the template first (i. e. creating the AST)
tpl, err := pongo2.FromString("Hello {{ name|capfirst }}!")
if err != nil {
    panic(err)
}
// Now you can render the template with the given
// pongo2.Context how often you want to.
out, err := tpl.Execute(pongo2.Context{"name": "florian"})
if err != nil {
    panic(err)
}
fmt.Println(out)
}