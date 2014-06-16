package main
import (
	"fmt"
	"net/http"
	"flag"
	"os"
	"io"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

var baseDir *string

func hello(c web.C, w http.ResponseWriter, r *http.Request) {
	path := *baseDir+"/"+c.URLParams["path"]
	src,err := os.Open(path)
	if(err!=nil){
		fmt.Println(err)
	}
	defer src.Close()
	io.Copy(w,src)
}

func main() {
	baseDir = flag.String("dir",".","Base directory (default: .)")
	flag.Parse()
    goji.Get("/:path", hello)
    goji.Serve()
}