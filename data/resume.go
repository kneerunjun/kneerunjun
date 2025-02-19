package data

import "fmt"

type ProfilePhoto struct {
	Location string // web location of the photo that appears on the front page splash
	HeightPx int    // height in pixels for the photo
	WidthPx  int    // width in pixels for the photo
}

type ProfileContact struct {
	FBLink     string `json:"fblink" bson:"fblink"`
	GmailLink  string `json:"gmaillink" bson:"gmaillink"`
	LinkedLink string `json:"linkedlink" bson:"linkedlink"`
	GitLink    string `json:"gitlink" bson:"gitlink"`
	Phone      string `json:"phone" bson:"phone"`
	Email      string `json:"email" bson:"email"`
	Address    string `json:"address" bson:"address"`
}

type EducQual struct {
	Start       string `json:"start" bson:"start"`
	End         string `json:"end" bson:"end"`
	Degree      string `json:"degree" bson:"degree"`
	ShortDegree string `json:"sdegree" bson:"sdegree"`
	GovnBody    string `json:"govnbody" bson:"govnbody"`
	Desc        string `json:"desc" bson:"desc"`
	ShortDesc   string `json:"sdesc" bson:"sdesc"`
}
type Skill struct {
	Title string `json:"title" bson:"title"`
	Level int8   `json:"level" bson:"level"`
	Desc  string `json:"desc" bson:"desc"`
	Span  string `json:"span" bson:"span"`
}
type Accolade struct {
	Title string `json:"title" bson:"title"`
	Year  string `json:"year" bson:"year"`
}

type Workexp struct {
	ImgSrc   string `json:"imgsrc" bson:"imgsrc"`
	ImgHt    int    `json:"imght" bson:"imght"`
	ImgWd    int    `json:"imgwd" bson:"imgwd"`
	Desig    string `json:"desig" bson:"desig"`
	Employer string `json:"employer" bson:"employer"`
	Span     string `json:"span" bson:"span"`
}

type Resume struct {
	ID          string         `json:"id" bson:"id"`
	Title       string         `json:"title" bson:"-"`
	FullName    string         `json:"fullname" bson:"fullname"`
	Photo       ProfilePhoto   `json:"photo" bson:"photo"`
	ShortDesc   string         `json:"shortdesc" bson:"shortdesc"`
	ShortDescSm string         `json:"shortdescsm" bson:"shortdescsm"`
	Contact     ProfileContact `json:"contact" bson:"contact,inline"`
	Education   EducQual       `json:"educ" bson:"educ"`
	TopSkills   []Skill        `json:"skills" bson:"skills"`
	Accolades   []Accolade     `json:"accolades" bson:"accolades"`
	Experience  []Workexp      `json:"experience" bson:"experience"`
}

// NiranjanAwati : seeds niranjan's reume to the database
func NiranjanAwati() error {
	res := &Resume{
		ID:       "niranjanawati",
		FullName: "Niranjan Awati",
		Photo: ProfilePhoto{
			Location: "/images/meb_w.jpg",
			HeightPx: 140,
			WidthPx:  140,
		},
		ShortDesc: `Seasoned Go Lang developer with solid 18 years of total experience. An avid IoT junkie, building prototype
		solutions atop single board computers for their sensing capabilities & cloud connectivity. He is adept at
		developing containerized REST API for the web & concurrent applications on IoT devices using Go Lang. He has
		also, in his past contributed extensively to learning functions of his organization.`,
		ShortDescSm: `Seasoned Go Lang developer with solid 18 years of total experience. An avid IoT junkie,He is adept at
		developing containerized REST API for the web & concurrent applications on IoT devices using Go Lang.`,
		Contact: ProfileContact{
			FBLink:     "https://www.facebook.com/kneerunjun/",
			GmailLink:  "mailto:kneerunjun@gmail.com?subject=Reference to your online profile",
			LinkedLink: "https://www.linkedin.com/in/niranjan-awati-a2395856/",
			GitLink:    "https://github.com/kneerunjun",
			Phone:      "+91 8390302623",
			Email:      "kneerunjun@gmail.com",
			Address:    "Sangria, Megapolis Hinjewadi Phase-III, Pune 411057",
		},
		Education: EducQual{
			Start:       "2000",
			End:         "2004",
			Degree:      "Bachelor of Engineering, Mechanical",
			ShortDegree: "B.E. Mechanical",
			GovnBody:    "University of Pune",
			Desc:        "Pursued a 4y bachelor's degree at Maharashtra Institute Technology,Pune. Internal combustion engines as the elective subject in the final year & graduate trainee stint at TATA motors in the year 2004.",
			ShortDesc:   "Pursued a 4y bachelor's degree at Maharashtra Institute Technology,Pune.",
		},
		TopSkills: []Skill{
			{Desc: "Building REST API over HTTP, programming IoT u-controllers using TinyGo", Title: "GoLang", Level: 85, Span: "2017-today"},
			{Desc: "Deep exposure to docker, docker-componse in building portable/scalable apps.", Title: "Docker", Level: 70, Span: "2018-today"},
			{Desc: "Can build single page, responsive apps from ground up.", Title: "AngularJs", Level: 60, Span: "2016-2021"},
			{Desc: "Can build single page, responsive apps from ground up.", Title: "Python", Level: 60, Span: "2016-2021"},
		},
		Accolades: []Accolade{
			{Title: "M.V.P, Infosys", Year: "2007"},
			{Title: "Pride, Boeing", Year: "2008"},
			{Title: "Pride, Boeing", Year: "2009"},
		},
		Experience: []Workexp{
			{ImgSrc: "/images/infy_logo.png", ImgHt: 40, ImgWd: 60, Desig: "Pr. Consultant", Employer: "Infosys Ltd.", Span: "2005-2022"},
			{ImgSrc: "/images/dheeti.jpeg", ImgHt: 45, ImgWd: 40, Desig: "Sr. Developer", Employer: "Dheeti Technologies", Span: "2022-2022"},
			{ImgSrc: "/images/ncs_logo.png", ImgHt: 45, ImgWd: 50, Desig: "Sr. Programmer", Employer: "NCS Technologies", Span: "2022-2023"},
			{ImgSrc: "/images/persistent_logo.png", ImgHt: 45, ImgWd: 50, Desig: "Sr. Architect", Employer: "Persistent", Span: "2023-today"},
		},
	}
	if err := AddResume(res); err != nil {
		return fmt.Errorf("failed to add resume to the database: %s", err)
	}
	return nil
}
