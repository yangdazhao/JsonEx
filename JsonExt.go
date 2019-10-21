package JsonEx

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
	"bytes"
	"golang.org/x/text/encoding/simplifiedchinese"
)

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func  (jst *JsonStruct)LoadEx(filename string, v interface{}){
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}

	dec := simplifiedchinese.GB18030.NewDecoder()
	buf := new(bytes.Buffer)
	buf.Write(data)

	var read = dec.Reader(buf)
	decBuf, err := ioutil.ReadAll(read)
	if err != nil {
		fmt.Println(err)
		return
	}


	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(decBuf, v)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (jst *JsonStruct) Load(filename string, v interface{}){
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (jst *JsonStruct) Save(filename string, v interface{}) {
	newData, _ := json.Marshal(v)
	
	err := ioutil.WriteFile(filename, newData, 666)
	if err != nil {
		return
	}
}
