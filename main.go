package gogadget

import (
	"fmt"
)

/*
#include<sys/types.h>
#include<sys/stat.h>
#include<fcntl.h>
*/
import "C"

func main() {
	backend := NewGumScriptBackend()

	script, err := backend.Create_script("example", `
console.log("load ok");
Interceptor.attach(Module.findExportByName(null, 'open'), {
onEnter: function (args) {
console.log('[*] open("' + Memory.readUtf8String(args[0]) + '")');
}
});
Interceptor.attach(Module.findExportByName(null, "close"), {
  onEnter: function (args) {
    console.log('[*] close(' + args[0].toInt32() + ')');
}
});
`)
	if err != nil {

		panic(err)
	}
	script.On(func(_script *GumScript, _message map[string]interface{}, _data []byte, _userdata uintptr) {
		fmt.Println(_message)
	})
	script.Load()
	ctx:=GMainContext_Get_thread_default()
	for{
		pd:=ctx.Pending()
		if pd==false{
			break
		}
		ctx.Iteration()
	}

}
