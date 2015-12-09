## GOntact

GOntact is a simple service to send contact form data by email, its propose is to provide a simple service that could be used by any website against multiple transactional email services such as Mandril, MailGun, and SendGrid.

GOntact provides a simple POST endpoint `/contact`, where it expects to receive the following parameters, that will be sent to your `GONTACT_EMAIL`.

- Name
- Email (optional)
- Address
- Subject (optional)
- Message
- Website Address (optional)

GOntact returns `422` if required parameters are not being passed, otherwise it will return `201`, if there is any provider error it returns a `50X` error.

### Security

GOntact is secured by a environment variable `GONTACT_TOKEN` that should be passed on the request's `X-GONTACT-TOKEN` header.

[TODO: JWT]

### Providers

GOntact built in providers:

- Mandril
- SendGrid (comming soon)
- MailGun (comming soon)
- SMTP (comming soon)

#### Copyright
GOntact is Copyright © 2008-2015 Wawandco SAS. It is free software, and may be redistributed under the terms specified in the LICENSE file.
