package entity

type AdditionalFile struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	IDDocType int    `json:"id_doctype"`
	File      string `json:"file"`
	UserID    int    `json:"user_id"`
}

// doctype surat izin / sakit / cuti

type AdditionFileRequest struct {
	IDDocType int    `json:"id_doctype"`
	File      string `json:"file"`
	UserID    int    `json:"user_id"`
}

type AdditionFileResponse struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	IDDocType int    `json:"id_doctype"`
	File      string `json:"file"`
	UserID    int    `json:"user_id"`
}
