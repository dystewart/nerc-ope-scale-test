package hostlistcreator

import (
	"fmt"
	"os"
	"strconv"
)

func GenerateHostList(numNotebooks int) {
	fileName := strconv.Itoa(numNotebooks) + "-hosts.txt" 

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	for i := 1; i <= numNotebooks; i++ {
		entry := "ope-" + strconv.Itoa(i)
		_, err := file.WriteString(entry + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Println("File created successfully:", fileName)
}