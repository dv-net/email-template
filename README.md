# Email Templates Repository for dv-merchant

This repository contains a collection of email templates used for user notifications, authentication, and transactional emails. The templates are organized in **Mustache** format for easy integration with various backend systems and also include legacy HTML templates for reference.

## Features
- Multi-language support: English, Russian, German, Spanish (i18n folder).
- Modular Mustache templates: Common components like header, footer, and styles are reusable across multiple emails.
- Transactional and notification emails: Covers user registration, password resets, 2FA, email verification, wallet operations, and more.
- Legacy templates: HTML templates from older versions of the system are preserved for backward compatibility.

## Folder Structure

```
./
├── i18n/                   # Localization JSON files
├── mustache/               # Current Mustache templates
│   ├── common/             # Common reusable components
│   └── templates/          # Specific email templates
```

## Usage
- Import the templates into your backend system.
- Render the templates with user-specific data using a Mustache engine.
- Send emails through your preferred email service.

