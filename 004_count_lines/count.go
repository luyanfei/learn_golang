package count

import (
	"bufio"
	"io"
)

/*
CountLines统计了输入文本中的行数，然而下面的代码在实现时，判断error时比较啰嗦。使用bufio中的Scanner来改写下面的函数，使之更加简洁。
func CountLines(r io.Reader) (int, error) {
	var (
		br    = bufio.NewReader(r)
		lines int
		err   error
	)

	for {
		_, err = br.ReadString('\n')
		lines++
		if err != nil {
			break
		}
	}

	if err != io.EOF {
		return 0, err
	}
	return lines, nil
}
*/

func CountLines(r io.Reader) (int, error) {

}
