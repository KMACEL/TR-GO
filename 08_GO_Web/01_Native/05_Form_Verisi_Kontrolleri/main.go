package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

// CreateUser tip
type CreateUser struct {
	UserName string   `json:"userName"`
	Password string   `json:"password"`
	Age      int      `json:"age"`
	Email    string   `json:"email"`
	City     string   `json:"city"`
	Gender   int      `json:"gender"`
	Hobby    []string `json:"hobby"`
}

func anaSayfa(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("Form Verisi : ", r.Form)
	fmt.Println("URL Yolu : ", r.URL.Path)
	fmt.Println("Şema : ", r.URL.Scheme)
	fmt.Println("Belirli Parametre : ", r.Form["username"])

	fmt.Println("-----------------------------------")
	fmt.Println("Bütün Parametreler")
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Println("-----------------------------------")

	fmt.Fprintf(w, "Merhaba Dünya!")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var createUserVariable CreateUser

	fmt.Println("Sorgu Method : ", r.Method)
	r.ParseForm()
	if r.Method == "GET" {
		fmt.Println("İlk açıldığında")
		t, _ := template.ParseFiles("createUser.gtpl")
		t.Execute(w, nil)
	} else {
		fmt.Println("Butona Tıklanınca")

		//------------------------------------------------------
		// Kullanıcı Adı Kontrol
		userName := r.Form["username"][0]
		if len(userName) == 0 {
			fmt.Fprintf(w, "Kullanıcı Alanı Boş Olamaz!")
			return
		}
		createUserVariable.UserName = userName

		//------------------------------------------------------
		// Parola Adı Kontrol
		password := r.Form["password"][0]
		if len(password) == 0 {
			fmt.Fprintf(w, "Parola Alanı Boş Olamaz!")
			return
		}
		createUserVariable.Password = password

		//------------------------------------------------------
		// Yaş Adı Kontrol
		getAge, err := strconv.Atoi(r.Form.Get("age"))
		if err != nil {
			fmt.Fprintf(w, "Lütfen Yaş Alanına Sayısal İfade Giriniz")
			return
		}

		if getAge > 120 || getAge < 10 {
			fmt.Fprintf(w, "Lütfen Yaş bölümüne 10 ile 120 Arasında Bir Değer Giriniz")
			return
		}
		createUserVariable.Age = getAge

		//------------------------------------------------------
		// Mail Adı Kontrol
		email := r.Form.Get("email")
		if mail, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, email); !mail {
			fmt.Fprintf(w, "Lütfen Geçerli Formatta E-Mail Adresi Giriniz : ornek@test.com")
			return
		}
		createUserVariable.Email = email

		//------------------------------------------------------
		// Şehir Adı Kontrol
		city := r.Form.Get("city")
		var cityControl = false
		cities := []string{"Bursa", "Uşak", "Kocaeli"}
		for _, city := range cities {
			if city == city {
				cityControl = true
			}
		}
		if !cityControl {
			fmt.Fprintf(w, "Lütfen geçerli bir şehir giriniz!")
			return
		}
		createUserVariable.City = city

		//------------------------------------------------------
		// Cinsiyet Kontrol
		var genderControl = false
		gender := []int{1, 2}
		getGender, err := strconv.Atoi(r.Form.Get("gender"))
		if err != nil {
			fmt.Fprintf(w, "Lütfen geçerli bir cinsitet seçiniz!")
		}
		for _, g := range gender {
			if g == getGender {
				genderControl = true
			}
		}

		if !genderControl {
			fmt.Fprintf(w, "Lütfen geçerli bir cinsiyet giriniz!")
			return
		}
		createUserVariable.Gender = getGender

		//------------------------------------------------------
		// Hobi Kontrol
		var hobbyControl []string
		hobbyList := []string{"spor", "muzik", "dans"}
		getHobbies := r.Form["hobby"]

		for _, getHobby := range getHobbies {
			for _, hobby := range hobbyList {
				if hobby == getHobby {
					hobbyControl = append(hobbyControl, hobby)
				}
			}
		}

		if hobbyControl != nil {
			createUserVariable.Hobby = hobbyControl
		}

		//------------------------------------------------------
		// Bilgiler
		userJSON, err := json.Marshal(createUserVariable)
		if err != nil {
			fmt.Println("Bİlgiler JSON Formatına Dönüştürülemedi : ", err.Error())
		}

		fmt.Println("Kullanıcı Bilgileri : ", string(userJSON))

	}
}

func main() {
	http.HandleFunc("/", anaSayfa)
	http.HandleFunc("/create", createUser)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
