package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("25_files/example.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		panic(err)
	}

	fmt.Println("file name: ", fileInfo.Name())
	fmt.Println("file or folder: ", fileInfo.IsDir())
	fmt.Println("size: ", fileInfo.Size())
	fmt.Println("permission: ", fileInfo.Mode())
	fmt.Println("modified at: ", fileInfo.ModTime())

	// read file
	buf := make([]byte, fileInfo.Size())

	d, err := file.Read(buf)

	if err != nil {
		panic(err)
	}

	for i := 0; i < len(buf); i++ {
		println("data", d, string(buf[i]))
	}

	// read file content
	data, err := os.ReadFile("25_files/example.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println("read file content: ", string(data))

	//read folders

	files, err := os.Open("25_files/.")

	if err != nil {
		panic(err)
	}

	defer files.Close()

	info, err := files.ReadDir(0)

	if err != nil {
		panic(err)
	}

	fmt.Println("read file list")
	for _, inf := range info {
		fmt.Println(inf.Name())
	}

	// create a file and write in it

	createdFile, err := os.Create("25_files/example2.txt")

	if err != nil {
		panic(err)
	}

	defer createdFile.Close()

	createdFile.WriteString("This is goolang project \n")
	createdFile.WriteString("This is goolang project 2 \n")

	bytes := []byte("Hello world")

	createdFile.Write(bytes)

	// read and write to another file (with streaming)
	sourceFile, err := os.Open("25_files/example2.txt")

	if err != nil {
		panic(err)
	}

	defer sourceFile.Close()

	destFile, err := os.Create("25_files/example3.txt")

	if err != nil {
		panic(err)
	}

	defer destFile.Close()

	reader := bufio.NewReader(sourceFile)

	writer := bufio.NewWriter(destFile)

	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err.Error() != "EOF" {
				panic(err)
			}

			break
		}

		wrErr := writer.WriteByte(b)

		if wrErr != nil {
			panic(wrErr)
		}

	}

	writer.Flush()

	fmt.Println("written to new file successfully")

	// delete a file

	dltErr := os.Remove("25_files/example3.txt")

	if dltErr != nil {
		panic(dltErr)
	}

	fmt.Println("file deleted successfully")

}
