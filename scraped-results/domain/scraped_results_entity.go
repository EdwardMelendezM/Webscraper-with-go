package domain

import "time"

type ScrapedResult struct {
	//Description: the id of the topic
	Id string `json:"id" example:"739bbbc9-7e93-11ee-89fd-0242ac110010"`
	//Description: the project id of the topic
	ProjectId string `json:"project_id" example:"739bbbc9-7e93-11ee-89fd-0242ac110010"`
	//Description: the title of the topic
	Title string `json:"title" example:"historias cortas de acoso"`
	//Description: the url of the topic
	Url string `json:"url" example:"https://www.google.com"`
	//Description: the content of the topic
	Content string `json:"content" example:"historias cortas de acoso en la escuela"`
	//Description: the number of the topic
	Number int `json:"number" example:"1"`
	//Description: the title corpus of the topic
	TitleCorpus *string `json:"title_corpus" example:"historias cortas de acoso"`
	//Description: the content corpus of the topic
	ContentCorpus *string `json:"content_corpus" example:"historias cortas de acoso en la escuela"`
	//Description: the word key of the topic
	WordKey *string `json:"word_key" example:"acoso"`
	//Description: the created at of the topic
	CreatedAt *time.Time `json:"created_at" example:"2022-01-01T00:00:00Z"`
}

type SemanticOntologyCountResult struct {
	ProjectID     string  `bson:"project_id"`
	Title         string  `bson:"title"`
	URL           string  `bson:"url"`
	Content       string  `bson:"content"`
	Number        int     `bson:"number"`
	Mensajes      float64 `bson:"mensajes"`
	RedesSociales float64 `bson:"redes_sociales"`
	Chat          float64 `bson:"chat"`
	Video         float64 `bson:"video"`
	Correo        float64 `bson:"correo"`
	Foros         float64 `bson:"foros"`
	Facebook      float64 `bson:"facebook"`
	Instagram     float64 `bson:"instagram"`
	SnapChat      float64 `bson:"snapChat"`
	WhatsApp      float64 `bson:"whatsapp"`
	Twitter       float64 `bson:"twitter"`
	YouTube       float64 `bson:"youtube"`
	TikTok        float64 `bson:"tiktok"`
	LinkedIn      float64 `bson:"linkedin"`
	Blog          float64 `bson:"blog"`
	Email         float64 `bson:"email"`
	Primaria      float64 `bson:"primaria"`
	Secundaria    float64 `bson:"secundaria"`
	Facultad      float64 `bson:"facultad"`
	Bachillerato  float64 `bson:"bachillerato"`
	Universidad   float64 `bson:"universidad"`
	Preparatoria  float64 `bson:"preparatoria"`
	Colegio       float64 `bson:"colegio"`
	Instituto     float64 `bson:"instituto"`
	Media         float64 `bson:"media"`
	Academia      float64 `bson:"academia"`
	Institucion   float64 `bson:"institucion"`
	Acosador      float64 `bson:"acosador"`
	Victima       float64 `bson:"victima"`
	Perpetrador   float64 `bson:"perpetrador"`
	Companeros    float64 `bson:"companeros"`
	Agresor       float64 `bson:"agresor"`
	Testigos      float64 `bson:"testigos"`
	Espia         float64 `bson:"espia"`
	Maton         float64 `bson:"maton"`
	Grupo         float64 `bson:"grupo"`
	Bully         float64 `bson:"bully"`
	Supervisor    float64 `bson:"supervisor"`
	Adolescente   float64 `bson:"adolescente"`
	Joven         float64 `bson:"joven"`
	Niño          float64 `bson:"niño"`
	Infantil      float64 `bson:"infantil"`
	Constante     float64 `bson:"constante"`
	Frecuente     float64 `bson:"frecuente"`
	Persistente   float64 `bson:"persistente"`
	Reiterado     float64 `bson:"reiterado"`
	Ocasional     float64 `bson:"ocasional"`
	Repetitivo    float64 `bson:"repetitivo"`
	Periodico     float64 `bson:"periodico"`
	DeletedAt     bool    `bson:"deleted_at"`
}

type SemanticOntologyTfIdfResult struct {
	ProjectID     string  `bson:"project_id"`
	Title         string  `bson:"title"`
	URL           string  `bson:"url"`
	Content       string  `bson:"content"`
	Number        int     `bson:"number"`
	Mensajes      float64 `bson:"mensajes"`
	RedesSociales float64 `bson:"redes_sociales"`
	Chat          float64 `bson:"chat"`
	Video         float64 `bson:"video"`
	Correo        float64 `bson:"correo"`
	Foros         float64 `bson:"foros"`
	Facebook      float64 `bson:"facebook"`
	Instagram     float64 `bson:"instagram"`
	SnapChat      float64 `bson:"snapChat"`
	WhatsApp      float64 `bson:"whatsapp"`
	Twitter       float64 `bson:"twitter"`
	YouTube       float64 `bson:"youtube"`
	TikTok        float64 `bson:"tiktok"`
	LinkedIn      float64 `bson:"linkedin"`
	Blog          float64 `bson:"blog"`
	Email         float64 `bson:"email"`
	Primaria      float64 `bson:"primaria"`
	Secundaria    float64 `bson:"secundaria"`
	Facultad      float64 `bson:"facultad"`
	Bachillerato  float64 `bson:"bachillerato"`
	Universidad   float64 `bson:"universidad"`
	Preparatoria  float64 `bson:"preparatoria"`
	Colegio       float64 `bson:"colegio"`
	Instituto     float64 `bson:"instituto"`
	Media         float64 `bson:"media"`
	Academia      float64 `bson:"academia"`
	Institucion   float64 `bson:"institucion"`
	Acosador      float64 `bson:"acosador"`
	Victima       float64 `bson:"victima"`
	Perpetrador   float64 `bson:"perpetrador"`
	Companeros    float64 `bson:"companeros"`
	Agresor       float64 `bson:"agresor"`
	Testigos      float64 `bson:"testigos"`
	Espia         float64 `bson:"espia"`
	Maton         float64 `bson:"maton"`
	Grupo         float64 `bson:"grupo"`
	Bully         float64 `bson:"bully"`
	Supervisor    float64 `bson:"supervisor"`
	Adolescente   float64 `bson:"adolescente"`
	Joven         float64 `bson:"joven"`
	Niño          float64 `bson:"niño"`
	Infantil      float64 `bson:"infantil"`
	Constante     float64 `bson:"constante"`
	Frecuente     float64 `bson:"frecuente"`
	Persistente   float64 `bson:"persistente"`
	Reiterado     float64 `bson:"reiterado"`
	Ocasional     float64 `bson:"ocasional"`
	Repetitivo    float64 `bson:"repetitivo"`
	Periodico     float64 `bson:"periodico"`
	DeletedAt     bool    `bson:"deleted_at"`
}
