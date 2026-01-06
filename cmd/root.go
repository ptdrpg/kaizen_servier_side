package cmd

import (
	"log"
	"net/http"
	"os"

	"KageNoEn/app"
	"KageNoEn/controller"
	"KageNoEn/repository"
	"KageNoEn/router"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kaizen",
	Short: "A brief description of your application",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		app.Connexion()
		db := app.DB
		repo := repository.NewRepository(db)
		ctrl := controller.NewController(db, repo)

		r := router.NewRouter(ctrl)
		r.RegisterRouter()

		log.Println("Server is running on port :4400")
		log.Fatal(http.ListenAndServeTLS(":4400", "key/localhost.pem", "key/localhost-key.pem", r.Handler()))
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.sma-back.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
