package repo

import (
	"crypto/tls"
	"log"
	"net"
	"os"
	"strconv"
	"sync"

	mgo "gopkg.in/mgo.v2"
)

var (
	mongoSession *mgo.Session
	mongoURL     string

	//MongoDB nome do banco de conexão do mongo
	MongoDB string

	// UseTLS ..
	UseTLS bool
)

//GetMongoSession retorna uma cópia da sessão de conexão do mongo.
//Caso seja a primeira execução, então abre a conexão e retorna uma cópia.
func GetMongoSession() *mgo.Session {
	var once sync.Once

	onceBody := func() {
		err := initSession()

		if err != nil {
			panic(err)
		}
	}

	once.Do(onceBody)

	return mongoSession.Clone()
}

func initSession() (err error) {
	loadOSEnvs()
	log.Println("Fazendo dial do mongo ", mongoURL)

	if UseTLS {
		tlsConfig := &tls.Config{}

		tlsConfig.InsecureSkipVerify = true

		dialInfo, err := mgo.ParseURL(mongoURL)

		if err != nil {
			panic(err)
		}

		dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
			return conn, err
		}

		mongoSession, err = mgo.DialWithInfo(dialInfo)
	} else {
		mongoSession, err = mgo.Dial(mongoURL)
	}

	if err != nil {
		panic(err)
	}

	mongoSession.SetMode(mgo.Monotonic, true)

	return
}

//Carrega os valores a partir dos parâmetros
func loadOSEnvs() {
	mongoURL = os.Getenv("MONGO_URL")
	MongoDB = os.Getenv("MONGO_DB")
	UseTLS, _ = strconv.ParseBool(os.Getenv("MONGO_USE_TLS"))
}
