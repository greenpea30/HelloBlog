package dao

import (
	"helloblog/internal/dto"

	"gorm.io/gorm"
)

type SearchDAO struct {
	db *gorm.DB
}

func NewSearchDAO(db *gorm.DB) *SearchDAO {
	return &SearchDAO{db: db}
}

func (d *SearchDAO) FullTextSearch(query string, limit int) ([]dto.SearchResultItem, error) {
	var items []dto.SearchResultItem
	like := "%" + query + "%"
	err := d.db.Raw(
		`SELECT
		   p."id"   AS "post_id",
		   p."title",
		   p."summary",
		   p."created_at"::text,
		   CASE
		     WHEN p."title"   ILIKE ? THEN 3
		     WHEN p."summary" ILIKE ? THEN 2
		     WHEN p."content" ILIKE ? THEN 1
		     ELSE 0
		   END AS "score"
		 FROM "posts" p
		 WHERE p."status" = 'normal'
		   AND (p."title"   ILIKE ? OR
		        p."summary" ILIKE ? OR
		        p."content" ILIKE ?)
		 ORDER BY "score" DESC, p."created_at" DESC
		 LIMIT ?`,
		like, like, like,
		like, like, like,
		limit,
	).Scan(&items).Error
	if items == nil {
		items = []dto.SearchResultItem{}
	}
	return items, err
}
