package models

type (
	GetDictionaryByType struct {
		ID      int    `json:"id"`
		DctName string `json:"dictionary_name"`
	}
	PostDictionary struct {
		DctName string `json:"dictionary_name"`
		DctType string `json:"dictionary_type"`
	}
)
