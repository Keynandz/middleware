package paginate

import "gorm.io/gorm"

type Paginate struct {
	page  int
	limit int
}

type PaginateUpdate struct {
	page  int
	limit int
}

type PaginateCreate struct {
	page  int
	limit int
}

func NewPaginate(page int, limit int) *Paginate {
	return &Paginate{page: page, limit: limit}
}

func UpdatedPaginate(page int, limit int) *PaginateUpdate {
	return &PaginateUpdate{page: page, limit: limit}
}

func CreatedPaginate(page int, limit int) *PaginateCreate {
	return &PaginateCreate{page: page, limit: limit}
}

func (p *Paginate) PaginatedResult(db *gorm.DB) *gorm.DB {
	offset := (p.page - 1) * p.limit

	return db.Order("id DESC").
		Offset(offset).
		Limit(p.limit)
}

func (p *PaginateUpdate) PaginatedUpdate(db *gorm.DB) *gorm.DB {
	offset := (p.page - 1) * p.limit

	return db.Order("updated_at DESC").
		Offset(offset).
		Limit(p.limit)
}

func (p *PaginateCreate) PaginateCreate(db *gorm.DB) *gorm.DB {
	offset := (p.page - 1) * p.limit

	return db.Order("date_added DESC").
		Offset(offset).
		Limit(p.limit)
}
