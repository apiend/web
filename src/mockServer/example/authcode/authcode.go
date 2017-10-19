package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	//如果你不知道怎么读取图片,请看这里
	image, _ := ioutil.ReadFile("./img/20171017173133.JPG")
	s := base64.StdEncoding.EncodeToString(image)
	fmt.Println(s)
	//以上是读取图片的并且base64的结果
	data := url.Values{}
	data.Add("key", "3b3b571ffbb3f1551cfd7795fd12288a")
	data.Add("codeType", "4006")
	data.Add("dtype", "json")
	// data.Add("base64Str", "R0lGODlhPwASAPcAAAAAAAAAMwAAZgAAmQAAzAAA/wArAAArMwArZgArmQArzAAr/wBVAABVMwBVZgBVmQBVzABV/wCAAACAMwCAZgCAmQCAzACA/wCqAACqMwCqZgCqmQCqzACq/wDVAADVMwDVZgDVmQDVzADV/wD/AAD/MwD/ZgD/mQD/zAD//zMAADMAMzMAZjMAmTMAzDMA/zMrADMrMzMrZjMrmTMrzDMr/zNVADNVMzNVZjNVmTNVzDNV/zOAADOAMzOAZjOAmTOAzDOA/zOqADOqMzOqZjOqmTOqzDOq/zPVADPVMzPVZjPVmTPVzDPV/zP/ADP/MzP/ZjP/mTP/zDP//2YAAGYAM2YAZmYAmWYAzGYA/2YrAGYrM2YrZmYrmWYrzGYr/2ZVAGZVM2ZVZmZVmWZVzGZV/2aAAGaAM2aAZmaAmWaAzGaA/2aqAGaqM2aqZmaqmWaqzGaq/2bVAGbVM2bVZmbVmWbVzGbV/2b/AGb/M2b/Zmb/mWb/zGb//5kAAJkAM5kAZpkAmZkAzJkA/5krAJkrM5krZpkrmZkrzJkr/5lVAJlVM5lVZplVmZlVzJlV/5mAAJmAM5mAZpmAmZmAzJmA/5mqAJmqM5mqZpmqmZmqzJmq/5nVAJnVM5nVZpnVmZnVzJnV/5n/AJn/M5n/Zpn/mZn/zJn//8wAAMwAM8wAZswAmcwAzMwA/8wrAMwrM8wrZswrmcwrzMwr/8xVAMxVM8xVZsxVmcxVzMxV/8yAAMyAM8yAZsyAmcyAzMyA/8yqAMyqM8yqZsyqmcyqzMyq/8zVAMzVM8zVZszVmczVzMzV/8z/AMz/M8z/Zsz/mcz/zMz///8AAP8AM/8AZv8Amf8AzP8A//8rAP8rM/8rZv8rmf8rzP8r//9VAP9VM/9VZv9Vmf9VzP9V//+AAP+AM/+AZv+Amf+AzP+A//+qAP+qM/+qZv+qmf+qzP+q///VAP/VM//VZv/Vmf/VzP/V////AP//M///Zv//mf//zP///wAAAAAAAAAAAAAAACH5BAEAAPwALAAAAAA/ABIAAAj/APcJHEiwoMGDCBMqXMiwocOHECNKnEixosWLGC1uWhNHTRyOcQYW+xhS5Ec4HKOZ5Fiw0sc7FTd5XIMSDsx9xDrGgVNyn8yTa1Ti1MlTYL2aNHtWLFZ037KU9TyqjBpH6MCnVeul9LkGJlWrE+l9FCrzJkFlcUIZLKuw3liLleCo3UfVo5pTItXQXDO3bke8BnMqlYi2p1aUH9cAphTnZNp99DjOVANWX9KLldTMHRp0n6Y4m+g2VvlzH1PFOI+ENim1orvRIpsWU2N2YGSYggUW62rwLcVNcgkuoy0w5+bYXnkPPb5PX1zmEDeCdZuVo1qZaiPHWdacZjS3d8VSLqaLEixEYsoJMlbDE05xoHHMxt17czfi9BJn1x7ImC/BnDQxN99xnwWY0YEIBgQAOw==")
	resp, err := http.PostForm("http://op.juhe.cn/vercode/index", data)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}
