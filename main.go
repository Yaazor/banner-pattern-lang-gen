package main

import (
	"banner_pattern_lang_gen/app"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	workingDir, err := os.Getwd()

	if err != nil {
		fmt.Println(fmt.Errorf("%s", err))
		os.Exit(1)
	}

	config := app.PatternConfig{
		LangKey:        "block.minecraft.banner.namespace.key",
		TranslatedName: "Pattern",
		LangFR:         false,
		TranslatedLangs: []string{
			"en_us",
		},
	}

	clearResultDir(workingDir + "/result")

	_ = os.Mkdir(workingDir+"/result/", 0755)

	txtc, err := os.Open(workingDir + "/config.json")

	if err != nil {
		fmt.Println("config.json was not found, creating.")
		txtc, err = os.Create(workingDir + "/config.json")
	}

	if err != nil {
		fmt.Println(fmt.Errorf("%s", err))
		os.Exit(1)
	}

	defer txtc.Close()

	txtConfig, _ := os.ReadFile(workingDir + "/config.json")

	if json.Valid(txtConfig) {
		_ = json.Unmarshal(txtConfig, &config)
	} else {
		jsObj, _ := json.Marshal(config)
		_, err = txtc.Write(jsObj)
		if err != nil {
			fmt.Println(fmt.Errorf("%s", err))
		}
		_ = txtc.Sync()
	}

	for _, lang := range config.TranslatedLangs {
		createLangFile(workingDir+"/result/"+lang, config)
	}

}

func createLangFile(path string, pattern app.PatternConfig) {
	lang, _ := os.Create(path + ".json")
	defer lang.Close()
	translations := make(map[string]string)

	w := app.CreateColorWrapper(pattern)

	for color, t := range w.Colors {
		key := pattern.LangKey + "." + color
		var v string
		if pattern.LangFR {
			v = pattern.TranslatedName + " (" + t + ")"
		} else {
			v = t + " " + pattern.TranslatedName
		}

		translations[key] = v
	}

	i, _ := json.Marshal(translations)

	lang.Write(i)
}

func clearResultDir(dir string) {
	err := os.RemoveAll(dir)
	if err != nil {
		log.Fatal(err)
	}
}
