/**
 * Simple analogue of the Linux command: file
 *
 * @author Sergey Iryupin (javageek@email.cz)
 * @version 9.9.2021
 */
package main

import (
	"bytes"
	"io"
	"log"
	"os"
)

const (
	MAX_BYTES_TO_READ = 1024 // 1KB buffer to read file
)

func main() {
	if len(os.Args) == 1 {
		println("Usage: file <FILE_NAME>")
		return
	}

	for _, fileName := range os.Args[1:] {
		file, err := os.Lstat(fileName)
		if err != nil {
			print(fileName + ": " + err.Error())
			continue
		}

		print(fileName + ": ")

		// by link https://pkg.go.dev/io/fs#FileMode
		if file.Mode() & os.ModeDir != 0 {
			print("directory")
		} else if file.Mode() & os.ModeSymlink != 0 {
			link, _ := os.Readlink(fileName)
			print("symbolic link to " + link)
		} else if file.Mode() & os.ModeDevice != 0 {
			print("device file")
		} else if file.Mode() & os.ModeNamedPipe != 0 {
			print("named pipe (FIFO)")

		// TODO add other options

		} else {
			print(regularFile(fileName))
		}
	}
}

func regularFile(filename string) string {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return err.Error()
	}

	// delayed file close
	defer func() {
		err = file.Close()
		// an error is also possible when closing
		if err != nil {
			log.Fatal(err)
		}
	}()

	// creating a buffer to read the file
	contentByte := make([]byte, MAX_BYTES_TO_READ)
	countByte := 0

	countByte, err = file.Read(contentByte)
	if err != nil && err != io.EOF {
		return err.Error()
	}
	contentByte = contentByte[:countByte]

	if hasPrefix(contentByte, "MZ") {

		lfanew := getBytes(contentByte, 0x3C, 4)
		machine := getBytes(contentByte, lfanew+ 4, 2)

		switch machine {
		case 0x014c:
			return "windows/386"
		case 0x0200:
			return "windows/intel64"
		case 0x8664:
			return "windows/amd64"
		default:
			return "windows/unknown"
		}
	}
	if hasPrefix(contentByte, "\x7FELF") {

		// https://ru.wikipedia.org/wiki/Executable_and_Linkable_Format

		osabi := getBytes(contentByte, 7, 1)

		system := "unknown"
		switch osabi {
		case 0:
			system = "linux" // "unix"
		case 2:
			system = "netbsd"
		case 6:
			system = "solaris"
		case 9:
			system = "freebsd"
		case 12:
			system = "openbsd"
		}

		machine := getBytes(contentByte, 16 + 2, 2)

		switch machine {
		case 3:
			return system + "/386"
		case 0x3E:
			return system + "/amd64"
		default:
			return system + "/unknown"
		}
	}

	return "unknown type"
}

func hasPrefix(s []byte, prefix string) bool {
	return len(s) >= len(prefix) && bytes.Compare(s[:len(prefix)], []byte(prefix)) == 0
}

func getBytes(buffer []byte, start uint32, length uint32) uint32 {
	addr := uint32(0)
	for pos := start + length - 1; pos >= start; pos-- {
		addr = addr * uint32(0x100) + uint32(buffer[pos])
	}
	return addr
}