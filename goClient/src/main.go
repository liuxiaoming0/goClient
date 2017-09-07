package main

import (
	"comm"
	"fmt"
)


func main(){

	confdata := comm.NewConf();
	confdata.ReadConfFile("./client.conf");
	
	fmt.Println(confdata);
	fmt.Println("log.filepath:", confdata.Log.Filepath);

	log, _ := comm.NewLog("./test.log");
	log.Info("test info");

}