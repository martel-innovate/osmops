package pkgr

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/martel-innovate/osmops/osmops/util/bytez"
	"github.com/martel-innovate/osmops/osmops/util/file"
)

const ChecksumFileName = "checksums.txt"

func md5string(data []byte) string {
	hash := md5.Sum(data)
	return fmt.Sprintf("%x", hash)
}

func computeChecksum(target file.AbsPath) (string, error) {
	content, err := os.ReadFile(target.Value())
	if err != nil {
		return "", err
	}
	return md5string(content), nil
}

func writeCheckSumFileContent(src PackageSource) io.Reader {
	buf := bytez.NewBuffer()
	for _, filePath := range src.SortedFilePaths() {
		hash := src.FileHash(filePath)
		line := fmt.Sprintf("%s\t%s\n", hash, filePath)
		io.Copy(buf, strings.NewReader(line))
	}
	return buf
}
