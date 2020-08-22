package controller

import (
	"html/template"
	"log"
	"net/http"

	"github.com/bijoyko/globaldemo/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// {{.Firstname}} {{.Lastname}}
type participant struct {
	Btn1  string
	Link1 string
	Btn2  string
	Link2 string
	Btn3  string
	Link3 string
	Btn4  string
	Link4 string
	Btn5  string
	Link5 string
}

func MainPage(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var getLinks1 models.Linking
	var getLinks2 models.Linking
	var getLinks3 models.Linking
	var getLinks4 models.Linking
	var getLinks5 models.Linking

	if err := db.Table("links").Select("btnname, weblink, btnno").Where("btnno = ?", 1).Find(&getLinks1).Error; err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := db.Table("links").Select("btnname, weblink, btnno").Where("btnno = ?", 2).Find(&getLinks2).Error; err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := db.Table("links").Select("btnname, weblink, btnno").Where("btnno = ?", 3).Find(&getLinks3).Error; err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := db.Table("links").Select("btnname, weblink, btnno").Where("btnno = ?", 4).Find(&getLinks4).Error; err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := db.Table("links").Select("btnname, weblink, btnno").Where("btnno = ?", 5).Find(&getLinks5).Error; err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	p := participant{
		Btn1:  getLinks1.Btnname,
		Link1: getLinks1.Weblink,
		Btn2:  getLinks2.Btnname,
		Link2: getLinks2.Weblink,
		Btn3:  getLinks3.Btnname,
		Link3: getLinks3.Weblink,
		Btn4:  getLinks4.Btnname,
		Link4: getLinks4.Weblink,
		Btn5:  getLinks5.Btnname,
		Link5: getLinks5.Weblink,
	}

	t, _ := template.ParseFiles("view/main.html")
	t.Execute(c.Writer, p)
}

func LoadAdmin(c *gin.Context) {
	session := sessions.Default(c)
	var count int
	v := session.Get("count")
	if v == nil {
		count = 0
	} else {
		count = v.(int)
		count++
	}
	t, _ := template.ParseFiles("view/admin.html")
	t.Execute(c.Writer, nil)
}

func VerifyAdmin(c *gin.Context) {

	if c.PostForm("Username") == "shahrzad" && c.PostForm("Password") == "12345" {
		t, err := template.ParseFiles("view/editlinks.html")
		if err != nil {
			log.Println(err)
		}

		db := c.MustGet("db").(*gorm.DB)

		var getLinks1 models.Linking
		var getLinks2 models.Linking
		var getLinks3 models.Linking
		var getLinks4 models.Linking
		var getLinks5 models.Linking

		if err := db.Table("links").Select("btnname, weblink, btnno").Where("btnno = ?", 1).Find(&getLinks1).Error; err != nil {
			log.Fatalln(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}

		if err := db.Table("links").Select("btnname, weblink, btnno").Where("btnno = ?", 2).Find(&getLinks2).Error; err != nil {
			log.Fatalln(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}

		if err := db.Table("links").Select("btnname, weblink, btnno").Where("btnno = ?", 3).Find(&getLinks3).Error; err != nil {
			log.Fatalln(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}

		if err := db.Table("links").Select("btnname, weblink, btnno").Where("btnno = ?", 4).Find(&getLinks4).Error; err != nil {
			log.Fatalln(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}

		if err := db.Table("links").Select("btnname, weblink, btnno").Where("btnno = ?", 5).Find(&getLinks5).Error; err != nil {
			log.Fatalln(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
			return
		}

		p := participant{
			Btn1:  getLinks1.Btnname,
			Link1: getLinks1.Weblink,
			Btn2:  getLinks2.Btnname,
			Link2: getLinks2.Weblink,
			Btn3:  getLinks3.Btnname,
			Link3: getLinks3.Weblink,
			Btn4:  getLinks4.Btnname,
			Link4: getLinks4.Weblink,
			Btn5:  getLinks5.Btnname,
			Link5: getLinks5.Weblink,
		}
		t.Execute(c.Writer, p)
	} else {
		text1 := "You are not authorised to view this page"
		t, err := template.ParseFiles("view/admin.html")
		if err != nil {
			log.Println(err)
		}
		t.Execute(c.Writer, text1)
	}
}

func UpdateLinks(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	var eraseLinks models.Linking
	if err := db.Table("links").First(&eraseLinks).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Table("links").Delete(&eraseLinks)

	insertchanges := participant{
		Btn1:  c.PostForm("Btn1"),
		Link1: c.PostForm("Link1"),
		Btn2:  c.PostForm("Btn2"),
		Link2: c.PostForm("Link2"),
		Btn3:  c.PostForm("Btn3"),
		Link3: c.PostForm("Link3"),
		Btn4:  c.PostForm("Btn4"),
		Link4: c.PostForm("Link4"),
		Btn5:  c.PostForm("Btn5"),
		Link5: c.PostForm("Link5"),
	}

	type rows struct {
		Btnno   int
		Btnname string
		Weblink string
	}

	row1 := rows{
		Btnno:   1,
		Btnname: insertchanges.Btn1,
		Weblink: insertchanges.Link1,
	}

	db.Table("links").Create(&row1)

	row2 := rows{
		Btnno:   2,
		Btnname: insertchanges.Btn2,
		Weblink: insertchanges.Link2,
	}

	db.Table("links").Create(&row2)

	row3 := rows{
		Btnno:   3,
		Btnname: insertchanges.Btn3,
		Weblink: insertchanges.Link3,
	}

	db.Table("links").Create(&row3)

	row4 := rows{
		Btnno:   4,
		Btnname: insertchanges.Btn4,
		Weblink: insertchanges.Link4,
	}

	db.Table("links").Create(&row4)

	row5 := rows{
		Btnno:   5,
		Btnname: insertchanges.Btn5,
		Weblink: insertchanges.Link5,
	}

	db.Table("links").Create(&row5)
}

//calculate the number of buttons from database
// var j int
// if err := db.Table("links").Select("btnname").Count(&j).Error; err != nil {
// 	log.Fatalln(err)
// 	c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 	return
// }

// loop that many times to extract data from the database
// for i := 1; i <= j; i++ {
// 	if err := db.Table("links").Select("btnname, weblink, btnno").Where("btnno = ?", i).Find(&getLinks).Error; err != nil {
// 		log.Fatalln(err)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
// 		return
// 	}

// 	type links struct {
// 		btn  string
// 		link string
// 	}

// p := links{
// 	btn:  getLinks.Btnname,
// 	link: getLinks.Weblink,
// }
// }
