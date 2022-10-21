package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	flagType := flag.String("type", "file", "docker, export, pull")

	flag.Parse()

	imageDataMap, _ := ReadImageJson("image.json")

	if *flagType == string("docker") {

		// var str string = "export %s=%s"

		// for k, v := range imageDataMap {
		// 	fmt.Println(k, v)
		// 	var strline string
		// 	strline = fmt.Sprint(str, strings.ToUpper(k), v) + "\n"
		// }

		var writeStr string = fmt.Sprintf("FROM %s", imageDataMap["image_from"])
		WriteFile("Dockerfile", []byte(writeStr), 0600)

	} else if *flagType == string("export") {
		ExportCommand(imageDataMap)

	} else if *flagType == string("pull") {
		DockerPullCommand(imageDataMap)
	}
}

func ExportCommand(imageDataMap map[string]interface{}) (err error) {

	for k, v := range imageDataMap {
		fmt.Printf(fmt.Sprintf("EXPORT %s=%s", strings.ToUpper(k), v))
		out, _ := exec.Command(fmt.Sprintf("EXPORT %s=%s", strings.ToUpper(k), v)).Output()
		// if err != nil {
		// 	fmt.Printf("Failed to execute command")
		// }
		fmt.Println(string(out))
	}

	return err
}

func DockerPullCommand(imageDataMap map[string]interface{}) (err error) {

	outPull, _ := exec.Command(fmt.Sprintf("/usr/bin/crictl -n k8s.io image pull %s/%s", imageDataMap["registry"], imageDataMap["image_name"])).Output()

	fmt.Println(string(outPull))

	outTag, _ := exec.Command(fmt.Sprintf("/usr/bin/crictl -n k8s.io image tag %s/%s %s", imageDataMap["registry"], imageDataMap["image_name"], imageDataMap["image_from"])).Output()

	fmt.Println(string(outTag))

	return err
}

func ReadImageJson(filename string) (imageDataMap map[string]interface{}, err error) {

	file, err := os.OpenFile(filename, os.O_RDWR, 0666)

	defer file.Close()

	decoder := json.NewDecoder(file)

	jsonDataMap := make(map[string]interface{})

	err = decoder.Decode(&jsonDataMap)

	return jsonDataMap, err
}

func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}
