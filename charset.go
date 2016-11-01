package charset

import (
    "fmt"
    "io/ioutil"
    "bytes"
    "golang.org/x/text/encoding"
    "golang.org/x/text/encoding/charmap"
    "strings"
)

func ChartMap(charset string) encoding.Encoding {
    switch charset {
    default:
        return nil
    case "CodePage437":
        return charmap.CodePage437
    case "CodePage850":
        return charmap.CodePage850
    case "CodePage852":
        return charmap.CodePage852
    case "CodePage855":
        return charmap.CodePage855
    case "CodePage858":
        return charmap.CodePage858
    case "CodePage860":
        return charmap.CodePage860
    case "CodePage862":
        return charmap.CodePage862
    case "CodePage863":
        return charmap.CodePage863
    case "CodePage865":
        return charmap.CodePage865
    case "CodePage866":
        return charmap.CodePage866
    case "ISO8859_1":
        return charmap.ISO8859_1
    case "ISO8859_2":
        return charmap.ISO8859_2
    case "ISO8859_3":
        return charmap.ISO8859_3
    case "ISO8859_4":
        return charmap.ISO8859_4
    case "ISO8859_5":
        return charmap.ISO8859_5
    case "ISO8859_6":
        return charmap.ISO8859_6
    case "ISO8859_6E":
        return charmap.ISO8859_6E
    case "ISO8859_6I":
        return charmap.ISO8859_6I
    case "ISO8859_7":
        return charmap.ISO8859_7
    case "ISO8859_8":
        return charmap.ISO8859_8
    case "ISO8859_8E":
        return charmap.ISO8859_8E
    case "ISO8859_8I":
        return charmap.ISO8859_8I
    case "ISO8859_10":
        return charmap.ISO8859_10
    case "ISO8859_13":
        return charmap.ISO8859_13
    case "ISO8859_14":
        return charmap.ISO8859_14
    case "ISO8859_15":
        return charmap.ISO8859_15
    case "ISO8859_16":
        return charmap.ISO8859_16
    case "KOI8R":
        return charmap.KOI8R
    case "KOI8U":
        return charmap.KOI8U
    case "Macintosh":
        return charmap.Macintosh
    case "MacintoshCyrillic":
        return charmap.MacintoshCyrillic
    case "Windows874":
        return charmap.Windows874
    case "Windows1250":
        return charmap.Windows1250
    case "Windows1251":
        return charmap.Windows1251
    case "Windows1252":
        return charmap.Windows1252
    case "Windows1253":
        return charmap.Windows1253
    case "Windows1254":
        return charmap.Windows1254
    case "Windows1255":
        return charmap.Windows1255
    case "Windows1256":
        return charmap.Windows1256
    case "Windows1257":
        return charmap.Windows1257
    case "Windows1258":
        return charmap.Windows1258
    }
}

func ConvertFromCharset(charset string, non_utf_str string) string {
    if ChartMap(charset) == nil {
        return non_utf_str
    }

    non_utf_str_reader := strings.NewReader(non_utf_str)
    decoder_reader     := ChartMap(charset).NewDecoder().Reader(non_utf_str_reader)

    result, err := ioutil.ReadAll(decoder_reader)
    if err != nil {
        fmt.Println(err)
        return non_utf_str
    }

    return fmt.Sprintf("%s", result)
}

func ConvertToCharset(charset string, utf_str string) string {
    if ChartMap(charset) == nil {
        return utf_str
    }

    utf_str_writer_buffer := new(bytes.Buffer)
    encoder_writer        := ChartMap(charset).NewEncoder().Writer(utf_str_writer_buffer)

    fmt.Fprintf(encoder_writer, utf_str)

    return fmt.Sprintf("%s", utf_str_writer_buffer.Bytes())
}
