package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)


func main() {
	file, err := os.Open(`C:\Users\龍二\Desktop\study_directory\password_hash\test.txt`)
	bu := bufio.NewReaderSize(file, 1024)
	if err !=nil{
		log.Fatal(err)
	}
	defer file.Close()
	var column=0;
	var flg int=0;
	var str string;
	for{
		line,_,err:=bu.ReadLine()
		slice:=strings.Split(string(line),"")
		len:=len(slice)
		for i:=0;i<len;i++{
			if slice[i]=="\""{
				if flg==1{
					flg=2
				}
				continue
			}
			if slice[i]=="{"||slice[i]=="["{
				continue
			}
			if slice[i]==":"{
				flg=1;
				continue
			}
			if  flg==1{
				match,_:=regexp.MatchString("[0-9]{1}",slice[i])
				 if match{
				str=str+slice[i]
				 }else if flg==1&&!match{
					fmt.Println("ファイルを確認してください")
					return
				 }
				 continue
			}
			if  flg==2{
				str=str+slice[i]
				if i==len-1&&slice[i]!="\""{
                  fmt.Println(column+1,"行目の文字の最後は\"にして下さい。")
				  return
				}		
			}
	}
	flg=0;
	column++;
	if err==io.EOF{
		break
	}
}
	fmt.Println("str:",str)
}