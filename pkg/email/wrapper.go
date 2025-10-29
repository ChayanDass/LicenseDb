// SPDX-FileCopyrightText: 2025 Chayan Das <01chayandas@gmail.com>
// SPDX-License-Identifier: GPL-2.0-only

package email

import (
	"context"
	"time"

	templates "github.com/fossology/LicenseDb/pkg/email/templetes"
)

func NotifyLicenseCreated(toEmail, userName, shortName string) {
	admins, err := FetchAdminEmails()
	if err != nil {
		return
	}
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		if Email == nil || !Email.IsRunning() {
			return
		}

		subject, html := templates.SingleLicenseEmailTemplate(
			userName,
			"Created",
			time.Now(),
			shortName,
		)

		_ = Email.Queue(ctx, EmailData{
			To:      []string{toEmail},
			Cc:      admins,
			Subject: subject,
			HTML:    html,
		})
	}()
}

func NotifyLicenseUpdated(to, userName, licenseName string) {
	admins, err := FetchAdminEmails()
	if err != nil {
		return
	}
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		if Email == nil || !Email.IsRunning() {
			return
		}

		subject, html := templates.SingleLicenseEmailTemplate(
			userName,
			"Updated",
			time.Now(),
			licenseName,
		)

		data := EmailData{
			To:      []string{to},
			Subject: subject,
			Cc:      admins,
			HTML:    html,
		}

		_ = Email.Queue(ctx, data)
	}()
}

func NotifyImportSummary(to, userName, importedType string, total, success, failed int) {
	admins, err := FetchAdminEmails()
	if err != nil {
		return
	}
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		if Email == nil || !Email.IsRunning() {
			return
		}

		subject, html := templates.ImportSummaryEmailTemplate(
			userName,
			importedType,
			total,
			success,
			failed,
			time.Now(),
		)

		data := EmailData{
			To:      []string{to},
			Subject: subject,
			Cc:      admins,
			HTML:    html,
		}

		_ = Email.Queue(ctx, data)
	}()
}
