# Security Policy

## Supported Versions

| Version | Supported          |
| ------- | ------------------ |
| 1.x     | ✅ Supported       |
| < 1.0   | ❌ Not Supported   |

## Reporting a Vulnerability

**Please do not disclose security vulnerabilities through public GitHub issues.**

If you discover a security vulnerability, please email the maintainers directly or use GitHub's private vulnerability reporting feature.

### Steps to Report:
1. Go to the [Security](https://github.com/dpkrn/gotunnel/security) tab
2. Click "Report a vulnerability"
3. Provide detailed information about the vulnerability
4. Include steps to reproduce if applicable

### What to Include:
- Description of the vulnerability
- Steps to reproduce
- Potential impact
- Your contact information (optional)
- Timeline for public disclosure (if any)

## Response Timeline

- **Initial Response**: Within 7 days
- **Investigation**: 1-2 weeks
- **Fix & Release**: As soon as possible
- **Disclosure**: After fix is released

## Security Best Practices

When using gotunnel:

1. **Keep Dependencies Updated**: Regularly update Go and other dependencies
2. **Use HTTPS**: Always use HTTPS when tunneling sensitive data
3. **Authentication**: Implement proper authentication for your tunneled services
4. **Access Control**: Restrict who can access your tunnel endpoints
5. **Monitor Traffic**: Regular audit and monitoring of tunnel connections
6. **Secure Configuration**: Protect your tunnel configuration and credentials

## Security Considerations

- Gotunnel establishes persistent outbound TCP connections
- Ensure your tunnel server is trusted and properly secured
- Use strong credentials for any authentication mechanisms
- Monitor for unauthorized access attempts
- Regularly rotate API keys and credentials

## Responsible Disclosure

We believe in responsible disclosure and ask that researchers:

1. Give us reasonable time to fix vulnerabilities before public disclosure
2. Avoid public discussion of unpatched vulnerabilities
3. Provide detailed information to help us understand and fix the issue
4. Act in good faith and respect privacy and confidentiality

---

Thank you for helping keep gotunnel secure! 🔒
