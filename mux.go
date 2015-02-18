package main

import (
	"log"
	"net/http"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/gorelic"
	"github.com/martini-contrib/render"
)

const (
	API_ENDPOINT = "/v1" // https://api.goquadro.com/v1
)

func main() {
	m := martini.Classic()
	if gqConfig.newRelicKey != "" {
		gorelic.InitNewrelicAgent(gqConfig.newRelicKey, gqConfig.newRelicAppName, true)
		m.Use(gorelic.Handler)
	}
	m.Use(render.Renderer(render.Options{
		//Directory:  RESOURCES + "/templates",   // Specify what path to load the templates from.
		//Extensions: []string{".tmpl", ".html"}, // Specify extensions to load for templates.
		//Layout:     "base",                     // Specify a layout template. Layouts can call {{ yield }} to render the current template.
		IndentJSON: true, // Output human readable JSON
		//IndentXML:  true,                       // Output human readable XML
	}))

	//m.Use(JwtCheck)

	m.Post(API_ENDPOINT+"/login", binding.Bind(User{}), ApiUserLogin)
	m.Post(API_ENDPOINT+"/signup", binding.Bind(User{}), ApiUserSignup)

	m.Get(API_ENDPOINT+"/me", ApiUserGetMe)

	m.Get(API_ENDPOINT+"/me/documents", ApiDocumentsGetAll)
	m.Post(API_ENDPOINT+"/me/documents", binding.Bind(Document{}), ApiDocumentsPost)
	m.Get(API_ENDPOINT+"/me/documents/:id", ApiDocumentsGetOne)
	m.Put(API_ENDPOINT+"/me/documents/:id", binding.Bind(Document{}), ApiDocumentsPut)
	m.Delete(API_ENDPOINT+"/me/documents/:id", ApiDocumentsDelete)

	//m.Get(API_ENDPOINT+"/me/topics", ApiTopicsGetAll)

	log.Println("Serving on", gqConfig.serveAddress)
	log.Fatal(http.ListenAndServe(gqConfig.serveAddress, m))
}

func init() {
	gqConfig.serveAddress = getenv("QDOC_WEB_SERVE_ADDRESS", "localhost:8001")
	gqConfig.resources = getenv("QDOC_RESOURCES_DIR", "/var/www/goquadro")
	gqConfig.newRelicAppName = "goquadro"
	gqConfig.newRelicKey = getenv("QDOC_NEWRELIC_KEY", "")
	gqConfig.googleOauth2Id = getenv("QDOC_GOAUTH_ID", "")
	gqConfig.googleOauth2Secret = getenv("QDOC_GOAUTH_SECRET", "")
	gqConfig.jwtSignKey = getenv("QDOC_JWT_PRIVATE_KEY", "")
	gqConfig.jwtVerifyKey = getenv("QDOC_JWT_PUBLIC_KEY", "")
}
