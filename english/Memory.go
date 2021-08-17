package english

import (
	"bufio"
	"fmt"
	"goutil/list"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var egh = make(map[string]string);
var contentList = make([]string, 0);
var input = bufio.NewScanner(os.Stdin);
func EnglishChose() {
	filePath := getFilePath();
	ipt(filePath);

	worldNum := len(contentList);

	startTime := time.Now();

	rand.Seed(time.Now().UnixNano());

	rdWorld := "";
	line := "";
	for true {
		if len(contentList) == 0 {
			break;
		}

		rd := rand.Intn(len(contentList));
		rdContent := contentList[rd];
		fmt.Println(rdContent);
		rdWorld = egh[rdContent];
		finishScan := input.Scan();
		if finishScan {
			line = input.Text();
		}
		if line == "T" {
			contentList = list.RemoveByIdx(contentList, rd);
			delete(egh, rdContent);
			continue;
		}

		if rdWorld == line {
			contentList = list.RemoveByIdx(contentList, rd);
			delete(egh, rdContent);
			continue;
		}

		fmt.Println("回答错误:" + rdWorld);
		for i := 0; i<5; i++ {
			input.Scan();
		}
	}

	endTime := time.Now();
	fmt.Println("共计:", worldNum, "，耗时:" , endTime.Sub(startTime));
}

func getFilePath() (fp string)  {
	filePath := "";

	if len(os.Args) == 1 {
		fmt.Println("输入文件路径:");
		input.Scan();
		filePath = input.Text();
	} else  {
		filePath = os.Args[1];
	}
	return filePath;
}

func ipt(filePath string) {
	file, err := os.Open(filePath);
	if err != nil {
		fmt.Println("open file error");
		return
	}
	bufReader := bufio.NewReader(file);
	for true {
		lineByte, _, err := bufReader.ReadLine();
		if err == io.EOF {
			break
		}
		fileStr := string(lineByte);
		parseLineAndSave(fileStr);
	}
}

func parseLineAndSave(lineStr string) {

	beginIdx := 0;
	for idx,s := range lineStr {
		if unicode.Is(unicode.Han, s) {
			beginIdx = idx;
			break;
		}
	}

	world := lineStr[0:beginIdx];
	content := lineStr[beginIdx:];

	egh[content] = strings.Trim(world, " ");
	contentList = append(contentList, content);
}
