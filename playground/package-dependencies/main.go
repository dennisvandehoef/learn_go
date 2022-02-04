package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

var (
	root_dir      *string
	main_location *string
	findings      = make(map[string][]string)
	package_name  string
)

func init() {
	root_dir = flag.String("p", ".", "Base path to the directory")
	main_location = flag.String("m", "", "Path to files containing the main package (if not in root, it is likely you have multiple)")
	flag.Parse()
}

func main() {
	setPackageOwnName()
	parseFiles(*root_dir)

	print_package("main", 0)
}

// Print the package and calls itself for its imported packages
func print_package(name string, depth int8) {
	prefix := ""

	for i := 0; int8(i) < depth; i++ {
		if i == 0 {
			prefix += " |"
		} else {
			prefix += "   |"
		}
	}

	if depth > 0 {
		prefix += "- "
	}
	fmt.Println(prefix + name)
	depth++
	for _, imported_name := range findings[name] {
		print_package(imported_name, depth)
	}
}

// Parse all files from the project and fill up findings
func parseFiles(current_dir string) {
	files, err := ioutil.ReadDir(current_dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if f.IsDir() && f.Name() != "vendor" {
			parseFiles(current_dir + "/" + f.Name())
		} else if strings.HasSuffix(f.Name(), ".go") && !strings.HasSuffix(f.Name(), "_test.go") {
			file_path := current_dir + "/" + f.Name()
			package_in_file := getPackageFromFile(file_path)

			// This main is not the one we are looking for
			if package_in_file == "main" && current_dir != *root_dir+"/"+*main_location {
				continue
			}

			findings[package_in_file] = unique(append(findings[package_in_file], getLocalPackagesUsedFromFile(file_path)...))
		}
	}
}

// Read the package name the file is contributing to
func getPackageFromFile(path string) string {
	r, err := regexp.Compile("package ([a-z_-]+)")
	if err != nil {
		log.Fatal(err)
	}

	return r.FindStringSubmatch(getFileContent(path))[1]
}

// Read file content as a string
func getFileContent(path string) string {
	file_content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return string(file_content)
}

// Get a slice with all the local packages used in a file
func getLocalPackagesUsedFromFile(path string) []string {
	unfiltered_results := getPackagesUsedFromFile(path)
	var filtered_results []string

	for _, name := range unfiltered_results {
		if strings.HasPrefix(name, package_name) {
			filtered_results = append(filtered_results, strings.Replace(name, package_name+"/", "", 1))
		}
	}

	return filtered_results
}

// get all packages imported in a file
func getPackagesUsedFromFile(path string) []string {
	var result []string
	file_content := getFileContent(path)

	// single line imports
	r, err := regexp.Compile("import \"([^\"]+)\"")
	if err != nil {
		log.Fatal(err)
	}

	for _, m := range r.FindAllStringSubmatch(file_content, -1) {
		result = append(result, m[1])
	}

	// block imports
	r, err = regexp.Compile(`import \((\s+"[^"]+")+\s\)`)
	if err != nil {
		log.Fatal(err)
	}

	for _, block := range r.FindAllString(file_content, -1) {
		sub_r, err := regexp.Compile("\"([^\"]+)\"")
		if err != nil {
			log.Fatal(err)
		}
		for _, m := range sub_r.FindAllStringSubmatch(block, -1) {
			result = append(result, m[1])
		}
	}

	return result
}

// make a string-slice unique
// source https://www.golangprograms.com/remove-duplicate-values-from-slice.html
func unique(input []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range input {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// set the full name of the package we are processing based on teh go.mod file
func setPackageOwnName() {
	r, err := regexp.Compile(`module (\S+)`)
	if err != nil {
		log.Fatal(err)
	}

	package_name = r.FindStringSubmatch(getFileContent(*root_dir + "/go.mod"))[1]
}
