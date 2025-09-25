package config

import (
    "encoding/json"
    "os"
    "letsgo/internal/domain"
)

type AppConfig struct {
    Coefficients domain.Coefficients `json:"coefficients"`
}

func Load(path string) (AppConfig, error) {
    var cfg AppConfig
    
    data, err := os.ReadFile(path)
    if err != nil {
        return cfg, err
    }
    
    err = json.Unmarshal(data, &cfg)
    return cfg, err
}