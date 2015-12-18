package main

import (
	"fmt"
	"os"
	"path/filepath"
	"encoding/csv"
	"strings"
	"io/ioutil"
	"log"
	_ "github.com/denisenkom/go-mssqldb"
	"database/sql"
)

type file struct {
	name string
	isPosted bool

}
var Fs = []file{}


func connectDb(server string,user string,pass string )(db *sql.DB){
	strCon :=  "server="+server+";user id="+user+";password="+pass+";port=1433"

	fmt.Println(strCon)
	db, err := sql.Open("mssql",strCon)
	checkErr(err)
	return
}


func main() {

	fmt.Println("Scaning file in folder ...")
	sourcePath := "./data/new/"
	dest := "./data/posted/"

	// scan file in source path
	filepath.Walk(sourcePath, Walker)
	fmt.Println("test go upload file")
	//fmt.Println(Fs)
	//Conn := connectDb("192.168.0.7","sa","[ibdkifu")
	//defer Conn.Close()



	for _, val := range Fs {
		fmt.Println(val.name)
		fullpath := sourcePath+val.name
		fmt.Println("Reading  filename : "+fullpath )
		dest = "./data/posted/"

		// Read txt file
		readcsv(fullpath )
		dest = dest+val.name
		fmt.Print("Copy file destination folder : "+dest)


		// after scan csv program to move file original to Posted Folder
		copyfile(fullpath,dest)
	}
	//readcsv()
}


func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}



func copyfile(src string, dst string) {
	// Read all content of src to data
	data, err := ioutil.ReadFile(src)
	checkErr(err)
	// Write data to dst
	err = ioutil.WriteFile(dst, data, 0644)
	if err != nil {
		checkErr(err)
	} else {
		fmt.Println("copy .... completed")
	}

}





func Walker(fn string, fi os.FileInfo, err error ) error {
//
//	if err != nil {
//		fmt.Println("Walker Error: ", err)
//		return nil
//	}
//
	if fi.IsDir() {
		fmt.Println("Directory: ", fn)
	} else {
		fmt.Print("File: ", fn)
		FsInsert(fi.Name())
	}


	//fmt.Println("insert array filename is : "+fi.Name())
	return nil
}


func FsInsert(fname string){
	if strings.Contains(fname,"TXT") 	{
		fmt.Println(fname+" ->Found text file!")
		ff := new(file)
		ff.name = fname
		ff.isPosted = false
		Fs = append(Fs , *ff)
	} else {
		fmt.Println(fname+" ->not a text file!")
	}
}


func readcsv(filename string ){
	csvfile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer csvfile.Close()
	reader := csv.NewReader(csvfile)
	reader.FieldsPerRecord = -1 // see the Reader struct information below

	rawCSVdata, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// sanity check, display to standard output
	for _, each := range rawCSVdata {

		fmt.Printf(" %s  %s %s %s %s %s %s\n", each[0], each[1] ,each[2], each[3],  each[4] ,each[5], each[6])
		//insertRawData(each , conn)


	}


}

func insertRawData(rec []string , conn *sql.DB){
	lccommand := "insert into navalog.dbo.rawdata (f1,f2,f3,f4,f5,f6,f7) values("
	lccommand = lccommand+"'"+rec[0]+"',"
	lccommand = lccommand+"'"+rec[1]+"',"
	lccommand = lccommand+"'"+rec[2]+"',"
	lccommand = lccommand+"'"+rec[3]+"',"
	lccommand = lccommand+"'"+rec[4]+"',"
	lccommand = lccommand+"'"+rec[5]+"',"
	lccommand = lccommand+"'"+rec[6]+"'"
//	lccommand = lccommand+"'"+rec[7]+"',"
//	lccommand = lccommand+"'"+rec[8]+"',"
//	lccommand = lccommand+"'"+rec[9]+"',"
//	lccommand = lccommand+"'"+rec[10]+"',"
//	lccommand = lccommand+"'"+rec[11]+"',"
//	lccommand = lccommand+"'"+rec[12]+"',"
//	lccommand = lccommand+"'"+rec[13]+"',"
//	lccommand = lccommand+"'"+rec[14]+"',"
//	lccommand = lccommand+"'"+rec[15]+"'"
	lccommand = lccommand +")\n"
	_,err :=conn.Exec(lccommand)
	if  err != nil {
		log.Fatalln(err)
	}
	fmt.Print(lccommand)
}
