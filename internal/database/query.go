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


func LogClickEvent(ctx context.Context, linkID int, ip, userAgent, referrer, deviceType, country string) error {
	query := `INSERT INTO click_events (link_id, ip, user_agent, referer, device_type, country) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := DB.Exec(ctx, query, linkID, ip, userAgent, referrer, deviceType, country)

	return err

}


func GetAnalytics(ctx context.Context, code string) (*models.AnalyticsResponse, error) {
	query := `SELECT 

	u.id, u.code, u.target_url, COUNT(ce.event_id) AS total_clicks
	FROM urls u
	LEFT JOIN click_events ce ON u.id = ce.link_id
	WHERE u.code = $1
	GROUP BY u.id, u.code, u.target_url`
	
	var id int 
	var urlCode, targetURL string 
	var totalClicks string
	
	err := DB.QueryRow(ctx, query, code).Scan(&id, &urlCode, &targetURL, &totalClicks)

	if err != nil {
		return nil, err
	}

	// Fetch click events for this URL
	eventsQuery := `
		SELECT event_id, link_id, ip, user_agent, referer, device_type, country, ts
		FROM click_events
		WHERE link_id = $1
		ORDER BY ts DESC
	`

	rows, err := DB.Query(ctx, eventsQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.ClickEvent

	for rows.Next() {
		var event models.ClickEvent
		err := rows.Scan(
			&event.EventID, 
			&event.LinkID, 
			&event.IP, 
			&event.UserAgent, 
			&event.Referer, 
			&event.DeviceType, 
			&event.Country, 
			&event.Timestamp,
		)
		if err != nil {
			return nil, err
		}

		event.Code = urlCode // include code in event for response
		events = append(events, event)
	}

	return &models.AnalyticsResponse{
			Code: urlCode,
			TargetUrl: targetURL,
			TotalClicks: totalClicks,
			Events: events,
	}, nil
}