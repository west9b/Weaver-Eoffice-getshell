# Weaver-Eoffice-getshell exp
泛微 eoffice10 前台 getshell

author:160team.west9B

**仅限用于安全研究人员在授权的情况下使用，遵守网络安全法，产生任何问题，后果自负，与作者无关。**

# 01-基本介绍

泛微 eoffice10 前台getshell，公开日期：2022/7/28

# 02-使用说明

## usage: ./Eoffice10getshell.exe -u url  (加http://)

exp上传冰蝎，默认密码，上传成功返回shell地址

poc：访问/eoffice10/server/public/iWebOffice2015/OfficeServer.php返回200存在漏洞，404及其他不存在。

# Screenshots
![Image text](https://github.com/west9b/Weaver-Eoffice-getshell/blob/main/exp.png)
![Image text](https://github.com/west9b/Weaver-Eoffice-getshell/blob/main/getshell.png)
