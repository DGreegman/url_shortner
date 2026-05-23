package database

import (
	"context"
	"url_shortner/internal/models"
)

func FindUrlByCode(code string) (*models.URL, error) {
	var url models.URL

	query := `SELECT id, code, target_url, redirect_type, clicks, created_at, expire_at FROM urls WHERE code = $1`

	err := DB.QueryRow(context.Background(), query, code).Scan(&url.ID, &url.Code, &url.TargetUrl, &url.RedirectType, &url.Clicks, &url.CreatedAt, &url.ExpireAt)

	if err != nil {
		return nil, err

	}

	return &url, nil
}
