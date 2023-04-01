package Jsonconverter

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"unicode"

	"github.com/gin-gonic/gin"
)

// Go-datatypes
var (
	s_struct      = " struct{\n"
	s_array       = " []"
	s_string      = " string\n"
	s_int         = " int\n"
	s_bool        = " bool\n"
	s_interface   = " interface{}\n"
	s_close_curly = "\n}\n"
)

// Checks for different conditions
var arr_alert bool
var arr_struct_alert bool
var garbage bool
var arr_var_check bool
var arr_check bool
var arr_may_end bool
var close_arr bool
var datastore []string

var mystruct string

func ValidateJson(c *gin.Context) {
	databyte, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(500, "Unable to Read json")
		return
	}
	content := string(databyte)
	content = strings.TrimSpace(content)
	_ = ioutil.WriteFile("json_example_write.json", []byte(content), 0644)
	res1 := strings.HasPrefix(content, `{`)
	res2 := strings.HasSuffix(content, `}`)
	if !res1 || !res2 {
		fmt.Println("Error Improper Json")
		c.JSON(400, "Error Improper Json")
		return
	}
	opencurlycount := strings.Count(content, `{`)
	closecurlycount := strings.Count(content, `}`)
	totalcur := opencurlycount - closecurlycount
	openarrbrac := strings.Count(content, `[`)
	closedarrbrac := strings.Count(content, `]`)
	totalarrbrac := openarrbrac - closedarrbrac
	if totalarrbrac != 0 || totalcur != 0 {
		fmt.Println("Error Improper Json")
		c.JSON(400, "Error Improper Json")
		return
	}
	packagename := c.Query("PackageName")
	StructName := c.Query("StructName")

	mystruct = "package " + packagename + "\n\ntype " + StructName + " struct{\n"
	Converter()
	c.JSON(200, "OK")

}
func ReadFile() []string {
	file, err := os.Open("json_example_write.json")
	if err != nil {
		log.Fatalf("failed to open")

	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()
	return text
}
func Converter() {
	text := ReadFile()
	for i := 1; i < len(text); i++ {
		if strings.Contains(text[i], ":") && !arr_alert {
			split := strings.Split(text[i], ":")
			key := split[0]
			key = IsLetter(key)
			Value := split[1]
			Value = checkdatatype(Value)
			mystruct += key + Value
		} else if (arr_check) && strings.Contains(text[i], "]") {
			arr_check = false
			var data string
			for i := 0; i < len(datastore); i++ {
				if datastore[i] == datastore[0] {
					data = datastore[0]
				} else {
					data = s_interface
				}
			}
			mystruct += data
			datastore = nil
			arr_var_check, arr_alert, arr_var_check, garbage = false, false, false, false
		} else if !arr_alert {
			if strings.Contains(text[i], `}`) {
				mystruct += s_close_curly
			} //for array inside struct
		} else if arr_alert && !arr_struct_alert && !arr_var_check && !garbage {
			//array containing struct
			if strings.Contains(text[i], `{`) {
				arr_struct_alert = true
				mystruct += s_struct
			} else {
				arr_var_check = true
				arr_check = true
				data := checkdatatype(text[i])
				datastore = append(datastore, data)
			}
			fmt.Println(arr_struct_alert, arr_var_check)
		} else if !garbage && !arr_var_check && arr_struct_alert {
			if strings.Contains(text[i], `{`) && strings.Contains(text[i-1], `[`) {
				arr_var_check = false
				mystruct += s_struct
			} else if !strings.Contains(text[i], `{`) && strings.Contains(text[i-1], `[`) {
				arr_var_check = true
			}
			if strings.Contains(text[i], ":") {
				split := strings.Split(text[i], ":")
				key := split[0]
				key = IsLetter(key)
				Value := split[1]
				Value = checkdatatype(Value)
				mystruct += key + Value
			} else if strings.Contains(text[i], `}`) {

				mystruct += s_close_curly
				garbage = true
				arr_struct_alert = false
			}
		} else if arr_var_check && arr_alert && arr_struct_alert && !strings.Contains(text[i], `]`) {
			val := strings.TrimSpace(text[i])
			data := checkdatatype(val)
			datastore = append(datastore, data)
		} else if strings.Contains(text[i], `]`) && !garbage {
			var data string
			for i := 0; i < len(datastore); i++ {
				if datastore[i] == datastore[0] {
					data = datastore[0]
				} else {
					data = s_interface
				}
			}
			mystruct += data
			datastore = nil
			arr_var_check = false
		} else if (garbage) && strings.Contains(text[i], `}`) && !strings.Contains(text[i], `,`) {
			arr_may_end = true
		} else if arr_var_check && arr_check {
			fmt.Println("going", text[i])
			val := strings.TrimSpace(text[i])
			data := checkdatatype(val)
			datastore = append(datastore, data)

		} else if arr_may_end {
			if strings.Contains(text[i], `]`) {
				close_arr = true
				mystruct += s_close_curly + "\n"
				arr_var_check, arr_alert, arr_var_check, garbage, arr_check, arr_may_end = false, false, false, false, false, false
			} else {
				fmt.Println("OOPs! it's not an end")
			}
		}
	}
	WriteFile(mystruct)

}

// TODO
// *fix closing
func IsLetter(s string) string {
	str := ""
	s = strings.TrimSpace(s)

	for _, r := range s {
		if unicode.IsLetter(r) {
			str += string(r)
		}

	}
	if str != "" {
		str = strings.Replace(str, string(str[0]), strings.ToUpper(string(str[0])), 1)
	}
	return str
}

// checking datatype
func checkdatatype(s string) string {
	str := ""
	s = strings.TrimSpace(s)
	datacheck := string(s[0])
	if strings.Contains(datacheck, `"`) {
		str = s_string
	} else if regexp.MustCompile(`\d`).MatchString(s) {
		str = s_int
	} else if strings.Contains(datacheck, "t") || strings.Contains(datacheck, "f") {
		str = s_bool
	} else if strings.Contains(datacheck, `[`) {
		str = s_array
		arr_alert = true
	} else if strings.Contains(datacheck, `{`) {
		str += s_struct
	}
	return str
}
func WriteFile(string) {
	validate := string(mystruct)
	opencurlycount := strings.Count(validate, `{`)
	closecurlycount := strings.Count(validate, `}`)
	totalCount := opencurlycount - closecurlycount
	validate = strings.TrimSpace(validate)
	if totalCount != 0 {
		lastchar := validate[len(validate)-1]
		mystruct = strings.Replace(validate, string(lastchar), "", 1)
		_ = os.WriteFile("Go_example_struct.go", []byte(mystruct), 0644)
	} else {
		_ = os.WriteFile("Go_example_struct.go", []byte(mystruct), 0644)
	}
}
