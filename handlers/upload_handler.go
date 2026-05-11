package handlers

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kgermando/optimatincorporation-api/config"
	"github.com/kgermando/optimatincorporation-api/storage"
)

func UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "No file provided"})
	}

	// Validate file size (50 MB)
	if file.Size > 50*1024*1024 {
		return c.Status(400).JSON(fiber.Map{"error": "File too large (max 50 MB)"})
	}

	// Validate file type
	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowedExts := map[string]string{
		".jpg": "image/jpeg", ".jpeg": "image/jpeg", ".png": "image/png",
		".gif": "image/gif", ".webp": "image/webp",
		".mp4": "video/mp4", ".mov": "video/quicktime", ".avi": "video/x-msvideo",
	}
	contentType, allowed := allowedExts[ext]
	if !allowed {
		return c.Status(400).JSON(fiber.Map{"error": "File type not allowed"})
	}

	// Read file
	fileData, err := file.Open()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not read file"})
	}
	defer fileData.Close()

	buf := make([]byte, file.Size)
	if _, err := fileData.Read(buf); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not read file data"})
	}

	// Folder from query param
	folder := c.Query("folder", "general")
	// Sanitize folder to allowed values
	allowedFolders := map[string]bool{"portfolio": true, "blog": true, "team": true, "general": true}
	if !allowedFolders[folder] {
		folder = "general"
	}

	// Unique filename
	timestamp := time.Now().UnixNano()
	fileName := fmt.Sprintf("%d%s", timestamp, ext)

	cfg := config.Load()
	b2 := &storage.B2Client{
		KeyID:          cfg.B2KeyID,
		ApplicationKey: cfg.B2ApplicationKey,
		BucketID:       cfg.B2BucketID,
		BucketName:     cfg.B2BucketName,
	}

	url, fileID, err := b2.Upload(buf, fileName, contentType, folder)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Upload failed: " + err.Error()})
	}

	return c.JSON(fiber.Map{"data": fiber.Map{"url": url, "file_id": fileID}})
}

func DeleteFile(c *fiber.Ctx) error {
	type deleteReq struct {
		FileID   string `json:"file_id"`
		FileName string `json:"file_name"`
	}
	var req deleteReq
	if err := c.BodyParser(&req); err != nil || req.FileID == "" || req.FileName == "" {
		return c.Status(400).JSON(fiber.Map{"error": "file_id and file_name required"})
	}

	cfg := config.Load()
	b2 := &storage.B2Client{
		KeyID:          cfg.B2KeyID,
		ApplicationKey: cfg.B2ApplicationKey,
		BucketID:       cfg.B2BucketID,
		BucketName:     cfg.B2BucketName,
	}

	if err := b2.Delete(req.FileID, req.FileName); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"success": true})
}
