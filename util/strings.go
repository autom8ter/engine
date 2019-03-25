package util

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/Masterminds/sprig"
	"google.golang.org/grpc/grpclog"
	html "html/template"
	"io"
	"strings"
	"text/template"
)

// ToPrettyJsonString encodes an item into a pretty (indented) JSON string
func ToPrettyJsonString(obj interface{}) string {
	output, _ := json.MarshalIndent(obj, "", "  ")
	return fmt.Sprintf("%s", output)
}

// ToPrettyJson encodes an item into a pretty (indented) JSON
func ToPrettyJson(obj interface{}) []byte {
	output, _ := json.MarshalIndent(obj, "", "  ")
	return output
}

func ReadAsCSV(val string) ([]string, error) {
	if val == "" {
		return []string{}, nil
	}
	stringReader := strings.NewReader(val)
	csvReader := csv.NewReader(stringReader)
	return csvReader.Read()
}

func ScanAndReplace(r io.Reader, replacements ...string) string {
	scanner := bufio.NewScanner(r)
	rep := strings.NewReplacer(replacements...)
	var text string
	for scanner.Scan() {
		text = rep.Replace(scanner.Text())
	}
	return text
}

func Render(s string, data interface{}) string {
	if strings.Contains(s, "{{") {
		t, err := template.New("").Funcs(sprig.GenericFuncMap()).Parse(s)
		if err != nil {
			grpclog.Warningln(err.Error())
		}
		buf := bytes.NewBuffer(nil)
		if err := t.Execute(buf, data); err != nil {
			grpclog.Warningln(err.Error())
		}
		return buf.String()
	}
	return s
}

func RenderHTML(s string, data interface{}) string {
	if strings.Contains(s, "{{") {
		t, err := html.New("").Funcs(sprig.GenericFuncMap()).Parse(s)
		if err != nil {
			grpclog.Warningln(err.Error())
		}
		buf := bytes.NewBuffer(nil)
		if err := t.Execute(buf, data); err != nil {
			grpclog.Warningln(err.Error())
		}
		return buf.String()
	}
	return s
}
