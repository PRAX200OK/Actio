# Contributing to Actio

Thank you for your interest in contributing to Actio! This document provides guidelines and information for contributors.

## Development Setup

1. Clone the repository:
```bash
git clone https://github.com/PRAX200OK/actio.git
cd actio
```

2. Build the project:
```bash
cd actio
go build .
```

3. Run tests:
```bash
go test ./...
```

## Code Style

- Follow standard Go formatting (`go fmt`)
- Use `gofmt` and `goimports` for consistent formatting
- Write tests for new functionality
- Update documentation for API changes

## Pull Request Process

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Testing

- Ensure all tests pass: `go test ./...`
- Add tests for new features
- Maintain or improve code coverage

## Documentation

- Update README.md for new features
- Add code comments for complex logic
- Update examples and usage instructions

## Issues

- Use GitHub issues for bug reports and feature requests
- Provide detailed information and steps to reproduce bugs
- Suggest improvements with concrete examples

## Code of Conduct

This project follows a code of conduct to ensure a welcoming environment for all contributors.