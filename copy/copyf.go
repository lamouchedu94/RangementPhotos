package copyf

import (
	"fmt"
	"os"
	"strings"
)

func CopyPictures(Pictures string, CurrentPath string) {
	TabPhoto := strings.Split(Pictures, "/")
	NbRepertoires := len(TabPhoto)
	source, err := os.ReadFile(Pictures)
	//b = bytes.NewBuffer(source)

	if err != nil {
		return
	}

	dst, err := os.Create(CurrentPath + "/" + TabPhoto[NbRepertoires-1])
	dst1 := CurrentPath + "/" + TabPhoto[NbRepertoires-1]
	if err != nil {
		fmt.Println(err)
	}
	defer dst.Close()

	os.WriteFile(dst1, source, 0750)
	dst.Close()
	//fmt.Println("fin")
}
