// This file is automatically generated by generate-charset-data.
// Do not hand-edit.

package data

import (
	"github.com/joakimthun/go-charset/charset"
	"io"
	"io/ioutil"
	"strings"
)

func init() {
	charset.RegisterDataFile("jisx0201kana.dat", func() (io.ReadCloser, error) {
		r := strings.NewReader("｡｢｣､･ｦｧｨｩｪｫｬｭｮｯｰｱｲｳｴｵｶｷｸｹｺｻｼｽｾｿﾀﾁﾂﾃﾄﾅﾆﾇﾈﾉﾊﾋﾌﾍﾎﾏﾐﾑﾒﾓﾔﾕﾖﾗﾘﾙﾚﾛﾜﾝﾞﾟ")
		return ioutil.NopCloser(r), nil
	})
}
