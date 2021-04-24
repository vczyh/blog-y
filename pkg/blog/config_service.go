package blog

func GetConfigValueService(name string) (string, error) {
	var config Config
	if err := db.Where("name = ?", name).Find(&config).Error; err != nil {
		return "", err
	}

	if config.CurrentValue != "" {
		return config.CurrentValue, nil
	}

	return config.InitialValue, nil
}