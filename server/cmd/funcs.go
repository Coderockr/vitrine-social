package cmd

import (
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/Coderockr/vitrine-social/server/handlers"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randSeq(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Int63()%int64(len(letters))]
	}
	return string(b)
}

func withEnvironment(run func(*cobra.Command, []string)) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		env := strings.ToLower(env)

		err := godotenv.Load("config/" + env + ".env")
		if err != nil {
			log.Print("Error loading file config/" + env + ".env")
		}

		run(cmd, args)
	}
}

func getJWTOptions() handlers.JWTOptions {
	return handlers.JWTOptions{
		SigningMethod: os.Getenv("VITRINESOCIAL_SIGNING_METHOD"),
		PrivateKey:    []byte(os.Getenv("VITRINESOCIAL_PRIVATE_KEY")), // $ openssl genrsa -out app.rsa keysize
		PublicKey:     []byte(os.Getenv("VITRINESOCIAL_PUBLIC_KEY")),  // $ openssl rsa -in app.rsa -pubout > app.rsa.pub
		Expiration:    60 * time.Minute,
	}
}
