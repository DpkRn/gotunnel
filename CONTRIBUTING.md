# Contributing to gotunnel

👋 Thank you for your interest in contributing to **gotunnel**!

We appreciate all contributions, from bug reports and feature requests to code improvements and documentation enhancements.

## Getting Started

Before opening an issue or pull request, please take a moment to review this guide.

### Prerequisites
- Read our [README.md](README.md) to understand the project
- Familiarity with Go, HTML, JavaScript, or CSS (depending on your contribution)
- Basic understanding of tunnel/proxy concepts

## How to Contribute

### Reporting Bugs 🐛
If you find a bug, please:
1. Check [existing issues](https://github.com/dpkrn/gotunnel/issues) to avoid duplicates
2. Create a new issue with:
   - Clear title and description
   - Steps to reproduce
   - Expected vs. actual behavior
   - Environment details (OS, Go version, etc.)

### Suggesting Features 💡
1. Check if the feature has already been discussed in [issues](https://github.com/dpkrn/gotunnel/issues)
2. Create an issue with:
   - Clear description of the feature
   - Use case and benefits
   - Possible implementation approach (optional)

### Submitting Code Changes 🔧
1. Fork the repository
2. Create a feature branch: `git checkout -b feature/your-feature-name`
3. Make your changes
4. Follow the code style of the project
5. Commit with clear messages: `git commit -m "Add feature: description"`
6. Push to your fork: `git push origin feature/your-feature-name`
7. Open a Pull Request with:
   - Clear description of changes
   - Reference to related issues
   - Screenshots (if UI changes)

### Code Style
- **Go**: Follow [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- **JavaScript**: Use consistent formatting and meaningful variable names
- **HTML/CSS**: Keep markup semantic and styles organized

### Testing
- Write tests for new Go code
- Test your changes thoroughly before submitting
- Run existing tests to ensure no regressions

## Development Setup

```bash
# Clone the repository
git clone https://github.com/dpkrn/gotunnel.git
cd gotunnel

# Install dependencies
go mod download

# Build the project
make build

# Run tests
make test
```

## Community Guidelines

- Be respectful and constructive in all interactions
- No harassment, discrimination, or hostile behavior
- Help others learn and grow
- Give credit where credit is due

## Pull Request Process

1. Update documentation if needed
2. Add tests for new functionality
3. Ensure all tests pass locally
4. Keep PRs focused on a single feature/fix
5. Provide context and rationale for your changes

## Need Help?

- 💬 Open a discussion or issue
- 📖 Check existing documentation
- 🤝 Reach out to the maintainers

Thank you for making gotunnel better! 🚀
