package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template
var tpl2 *template.Template

type StructTest struct {
	Isim    string
	Soyisim string
	Yas     int
}

var funcmap = template.FuncMap{
	"buyult": strings.ToUpper,
	"kes":    kes,
}

func kes(s string) string {
	s = strings.TrimSpace(s)
	s = s[:2]
	return s
}

func init() {
	//	tpl = template.Must(template.ParseGlob("*.txt"))
	tpl2 = template.Must(template.New("").Funcs(funcmap).ParseFiles("test_funcmap.txt"))

}

func main() {
	/*tmplt, err := template.ParseFiles("test.txt")

	if err != nil {
		log.Fatalln(err.Error())
	}

	err = tmplt.Execute(os.Stdout, "Mert")
	if err != nil {
		log.Fatalln(err.Error())
	}

	//###################################################################

	tmplt, err = template.ParseFiles("test.txt", "test2.txt")

	if err != nil {
		log.Fatalln(err.Error())
	}

	err = tmplt.ExecuteTemplate(os.Stdout, "test2.txt", "Mert")
	if err != nil {
		log.Fatalln(err.Error())
	}

	//###################################################################

	tmplt, err = template.ParseGlob("*.txt")

	if err != nil {
		log.Fatalln(err.Error())
	}
	err = tmplt.ExecuteTemplate(os.Stdout, "test.txt", "Erdem")
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = tmplt.ExecuteTemplate(os.Stdout, "test2.txt", "Mert")
	if err != nil {
		log.Fatalln(err.Error())
	}

	//###################################################################

	err = tpl.ExecuteTemplate(os.Stdout, "test.txt", "Bilge")
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = tpl.ExecuteTemplate(os.Stdout, "test2.txt", "Kübra")
	if err != nil {
		log.Fatalln(err.Error())
	}

	//###################################################################
	err = tpl.ExecuteTemplate(os.Stdout, "test_degisken.txt", "Uraz")
	if err != nil {
		log.Fatalln(err.Error())
	}
	//###################################################################
	dizi := []string{"Erdem", "Uraz", "Bilge"}
	err = tpl.ExecuteTemplate(os.Stdout, "test_dizi.txt", dizi)
	if err != nil {
		log.Fatalln(err.Error())
	}

	//###################################################################
	m := map[string]string{
		"isim":    "Mert",
		"soyisim": "Acel",
	}

	err = tpl.ExecuteTemplate(os.Stdout, "test_map.txt", m)
	if err != nil {
		log.Fatalln(err.Error())
	}

	//###################################################################
	s := StructTest{Isim: "Mert", Soyisim: "Acel", Yas: 25}
	err = tpl.ExecuteTemplate(os.Stdout, "test_struct.txt", s)
	if err != nil {
		log.Fatalln(err.Error())
	}
	//###################################################################
	sa := []StructTest{
		{Isim: "Mert", Soyisim: "Acel", Yas: 25},
		{Isim: "Erdem", Soyisim: "Acel", Yas: 0},
		{Isim: "Kübra", Soyisim: "Acel", Yas: 25},
		{Isim: "Bilge", Soyisim: "Acel", Yas: -7},
	}

	err = tpl.ExecuteTemplate(os.Stdout, "test_struct_dizi.txt", sa)
	if err != nil {
		log.Fatalln(err.Error())
	}
	//###################################################################
	st := []StructTest{
		{Isim: "Mert", Soyisim: "Acel", Yas: 25},
		{Isim: "Kübra", Soyisim: "Acel", Yas: 25},
	}

	type StructTest2 struct {
		AnneIsim string
		BabaIsim string
	}

	st2 := []StructTest2{{AnneIsim: "Fatma", BabaIsim: "Ozcan"}, {AnneIsim: "Nuray", BabaIsim: "Mustafa"}}

	structOrtak := struct {
		S1 []StructTest
		S2 []StructTest2
	}{
		st,
		st2,
	}

	err = tpl.ExecuteTemplate(os.Stdout, "test_struct_karma.txt", structOrtak)
	if err != nil {
		log.Fatalln(err.Error())
	}
	//###################################################################*/

	stFuncMap := []StructTest{
		{Isim: "Mert", Soyisim: "Acel", Yas: 25},
		{Isim: "Kübra", Soyisim: "Acel", Yas: 25},
	}

	err := tpl2.Execute(os.Stdout, stFuncMap)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
