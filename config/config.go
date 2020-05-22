package config

import(
	Viper "github.com/spf13/viper"
	"log"
	"os"
    "path/filepath"
)

const (
  AppName = "doubtnut"
)

var SimilarQuestions = map[string][]questionLinkPair{}

type questionLinkPair struct{
	Question  string
	VideoLink string
}

func ReadConfig() {
    config := ReadConfigs("similar_questions", GenerateFilePath("config"), "json")
    err := config.Unmarshal(&SimilarQuestions)
    if err != nil {
      log.Fatal(err)
    }
}

func ReadConfigs(filename string, filepath string, configType string) *Viper.Viper {
    config := Viper.New()
    config.SetEnvPrefix(AppName)
    config.AddConfigPath(filepath)
    config.SetConfigType(configType)
    config.SetConfigName(filename)
    err := config.ReadInConfig()
    if err != nil {
      log.Fatal(err)
    }
    return config
}

func GenerateFilePath(path ...string) string {
    dir, err := os.Getwd()
    if err != nil {
     log.Fatal(err)
    }
    f := filepath.Join(dir)
    for _, v := range path {
        f = filepath.Join(f, v)
    }
    return f
}