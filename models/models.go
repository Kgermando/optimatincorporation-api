package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type User struct {
	ID        string         `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	Email     string         `gorm:"unique;not null" json:"email"`
	Password  string         `gorm:"not null" json:"-"`
	Role      string         `gorm:"default:editor" json:"role"` // admin | editor
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type Portfolio struct {
	ID          string         `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	Slug        string         `gorm:"uniqueIndex;not null" json:"slug"`
	Description string         `json:"description"`
	Category    string         `gorm:"not null" json:"category"`
	Year        int            `json:"year"`
	Client      string         `json:"client"`
	CoverURL    string         `json:"cover_url"`
	VideoURL    string         `json:"video_url"`
	GalleryURLs pq.StringArray `gorm:"type:text[]" json:"gallery_urls"`
	Tags        pq.StringArray `gorm:"type:text[]" json:"tags"`
	Published   bool           `gorm:"default:true" json:"published"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type Service struct {
	ID          string         `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	Subtitle    string         `json:"subtitle"`
	Description string         `json:"description"`
	IconURL     string         `json:"icon_url"`
	Features    pq.StringArray `gorm:"type:text[]" json:"features"`
	Order       int            `gorm:"default:0" json:"order"`
	Published   bool           `gorm:"default:true" json:"published"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type Blog struct {
	ID        string         `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Title     string         `gorm:"not null" json:"title"`
	Slug      string         `gorm:"uniqueIndex;not null" json:"slug"`
	Excerpt   string         `json:"excerpt"`
	Content   string         `gorm:"type:text" json:"content"`
	CoverURL  string         `json:"cover_url"`
	Author    string         `json:"author"`
	Category  string         `json:"category"`
	Published bool           `gorm:"default:true" json:"published"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type TeamMember struct {
	ID              string         `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name            string         `gorm:"not null" json:"name"`
	Position        string         `json:"position"`
	Bio             string         `json:"bio"`
	PhotoURL        string         `json:"photo_url"`
	SocialFacebook  string         `json:"social_facebook"`
	SocialInstagram string         `json:"social_instagram"`
	SocialLinkedin  string         `json:"social_linkedin"`
	Order           int            `gorm:"default:0" json:"order"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}

type Contact struct {
	ID        string         `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Name      string         `gorm:"not null" json:"name"`
	Email     string         `gorm:"not null" json:"email"`
	Phone     string         `json:"phone"`
	Subject   string         `json:"subject"`
	Message   string         `gorm:"type:text;not null" json:"message"`
	Read      bool           `gorm:"default:false" json:"read"`
	CreatedAt time.Time      `json:"created_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
