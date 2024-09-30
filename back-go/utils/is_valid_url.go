package utils

import (
	"fmt"
	"net/url"
	"regexp"
)

// AddHTTPPrefixIfNeeded añade "http://" si la URL no tiene un esquema válido
func AddHTTPPrefixIfNeeded(u string) string {
	parsed, _ := url.Parse(u)

	// Solo añade el prefijo si no hay esquema
	if parsed.Scheme == "" {
		return "http://" + u
	}

	return u
}

// isValidScheme verifica si la URL tiene un esquema permitido
func isValidScheme(u string) bool {
	parsed, _ := url.Parse(u)
	allowedSchemes := []string{"http", "https", "ftp"}

	for _, scheme := range allowedSchemes {
		if parsed.Scheme == scheme {
			return true
		}
	}

	return false
}

// IsValidURL verifica si una URL es válida y tiene un TLD
func IsValidURL(u string) (string, error) {
	// Primero, se añade el prefijo HTTP si es necesario
	parsed := AddHTTPPrefixIfNeeded(u)

	// Verificar si la URL tiene un esquema válido
	if !isValidScheme(parsed) {
		return "", fmt.Errorf("la URL tiene un esquema no permitido")
	}

	// Comprobar si la URL es válida
	urlParsed, err := url.ParseRequestURI(parsed)
	if err != nil {
		return "", err // Devuelve un string vacío y el error
	}

	// Verifica si la URL tiene un TLD válido
	tldRegex := regexp.MustCompile(`\.[a-z]{2,}$`) // TLD debe tener al menos 2 caracteres
	if !tldRegex.MatchString(urlParsed.Host) {
		return "", fmt.Errorf("la URL debe tener un dominio con un TLD válido")
	}

	return parsed, nil // Devuelve la URL válida y nil para el error
}
