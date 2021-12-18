package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"text/template"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var (
	tmplPath   string
	optPath    string
	nameFields []string

	rootCmd = &cobra.Command{
		Use:   "genprofile",
		Short: "Create profiles based on a template and an option file",
		Long: `Create test profiles, e.g. fio profiles, based on a template file(which is in go text/template syntax)
		and an option file(which is based on yaml).
		Option files defined in yaml only support key value pairs and list as below:
		key1: value1
		key2:
		- value1
		- value2 
		`,
		Run: func(cmd *cobra.Command, args []string) {
			options, err := extractOptions(optPath)
			cobra.CheckErr(err)

			for i, option := range options {
				bytes, err := generateConfig(tmplPath, option)
				cobra.CheckErr(err)

				var output string // output profile file name
				if len(nameFields) > 0 {
					var fileds []string
					for _, f := range nameFields {
						fileds = append(fileds, fmt.Sprintf("%v", option[f]))
					}
					output = strings.Join(fileds, "-") + ".profile"
				} else {
					output = fmt.Sprintf("test_%d.profile", i)
				}

				if err := ioutil.WriteFile(output, bytes, 0644); err != nil {
					cobra.CheckErr(err)
				}
			}
		},
	}
)

func init() {
	rootCmd.Flags().StringVarP(&tmplPath, "template", "t", "", "profile template file path")
	rootCmd.Flags().StringVarP(&optPath, "option", "o", "", "option definition file path")
	rootCmd.Flags().StringArrayVarP(&nameFields, "fields", "f", []string{}, "fields based on which the output file name is formed")
	rootCmd.MarkFlagRequired("template")
	rootCmd.MarkFlagRequired("option")
}

func extractOptions(optPath string) ([]map[string]interface{}, error) {
	// return values:
	// - options list: []map[string]interface{}
	// - error
	var options []map[string]interface{}

	contents, err := ioutil.ReadFile(optPath)
	if err != nil {
		return nil, err
	}

	var optRaw map[string]interface{}
	if err := yaml.Unmarshal(contents, &optRaw); err != nil {
		return nil, err
	}

	base := map[string]interface{}{}
	sKeys := []string{}
	sValues := [][]interface{}{}
	for k, v := range optRaw {
		if reflect.TypeOf(v).Kind() == reflect.Slice {
			values := v.([]interface{})
			sKeys = append(sKeys, k)
			sValues = append(sValues, values)
			base[k] = nil
		} else {
			base[k] = v
		}
	}

	var indices []int
	n := len(sKeys)
	for i := 0; i < n; i++ {
		indices = append(indices, 0)
	}
	for {
		option := map[string]interface{}{}
		for k, v := range base {
			option[k] = v
		}
		options = append(options, option)

		for i := 0; i < n; i++ {
			option[sKeys[i]] = sValues[i][indices[i]]
		}

		next := n - 1
		for next >= 0 && (indices[next]+1 >= len(sValues[next])) {
			next -= 1
		}

		if next < 0 {
			break
		}

		indices[next] += 1

		for i := next + 1; i < n; i++ {
			indices[i] = 0
		}

	}
	return options, nil
}

func generateConfig(tmplPath string, data map[string]interface{}) ([]byte, error) {
	contents, err := ioutil.ReadFile(tmplPath)
	if err != nil {
		return nil, err
	}

	tmpl, err := template.New("testconfig").Parse(string(contents))
	if err != nil {
		return nil, err
	}

	buf := bytes.Buffer{}
	if err := tmpl.Execute(&buf, data); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
