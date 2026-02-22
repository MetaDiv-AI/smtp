# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.1] - 2025-02-23

### Fixed

- Use `SetAddressHeader` for From field to properly handle non-ASCII display names (e.g. Chinese). Previously `SetHeader` encoded the entire string which broke `net/mail.ParseAddress` when extracting the address for SMTP.

## [1.0.0] - 2025-02-14

### Added

- Initial release
- Fluent builder API for dialer configuration (`NewDialer`, `Host`, `Port`, `Username`, `Password`)
- Fluent builder API for email composition (`NewEmail`, `Subject`, `Body`, `To`, `Cc`, `Bcc`, `Attachments`, `Values`, `Build`)
- Text and HTML email support (`EmailTypeText`, `EmailTypeHTML`)
- Attachments with automatic MIME type detection
- Template values — replace `{{key}}` placeholders in subject, body, and attachment filenames
- CC and BCC recipient support
- Display name support for sender address
