package comm


import(
	"encoding/json"
	"io/ioutil"
	"fmt"
)

type ConfData struct {
	Client struct {
		Ip string `json:"ip"` //struct tag???
		Port int
	}
	HTTP struct {
		Url string
	}
	TCP struct {
		Ip string
		Port int
		Num int
	}
	UDP struct {
		Ip string
		Num int
	}
	Log struct {
		Filepath string
		Loglevel string
	}
}

func NewConf() *ConfData {
	return &ConfData{};
}



func (this *ConfData)ReadConfFile(filename string) error{
	
	bytes, err := ioutil.ReadFile(filename);
	if err != nil {
		fmt.Println("ReadConfFile:", err.Error())
		return err;
	}

	return json.Unmarshal(bytes, this);
}