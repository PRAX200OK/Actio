# Installation

Install the Actio CLI to create projects, validate sidecars, and run the MCP server.

## Option 1: Download from GitHub Releases (recommended)

Releases are built automatically via **GitHub Actions** and **GoReleaser** when a version tag (e.g. `v0.1.0`) is pushed. Pre-built binaries are published on the [Releases](https://github.com/PRAX200OK/Actio/releases) page.

### Supported platforms

| OS      | Architectures | Archive   |
|---------|---------------|-----------|
| Windows | amd64, arm64  | `.zip`    |
| Linux   | amd64, arm64  | `.tar.gz` |
| macOS   | amd64, arm64  | `.tar.gz` |

### Steps

1. Open [github.com/PRAX200OK/Actio/releases](https://github.com/PRAX200OK/Actio/releases) and download the archive for your OS and architecture (e.g. `actio_Windows_x86_64.zip` or `actio_Linux_arm64.tar.gz`).

2. Extract the archive:
   - **Windows (.zip):** Unzip and you’ll get `actio.exe` (and optionally `setup.bat`, `setup.sh`).
   - **Linux / macOS (.tar.gz):**  
     `tar -xzf actio_<OS>_<arch>.tar.gz`

3. Put the `actio` (or `actio.exe`) binary in your `PATH`, for example:
   - **Windows:** Move `actio.exe` to a folder that’s on your PATH, or add the folder to your user PATH.
   - **Linux / macOS:**  
     `sudo mv actio /usr/local/bin/`  
     or  
     `export PATH="$PATH:$(pwd)"`

4. Verify:

   ```bash
   actio version
   actio --help
   ```

## Option 2: Build from source

You need **Go 1.22+** installed.

1. Clone the repository and build:

   ```bash
   git clone https://github.com/PRAX200OK/Actio.git
   cd Actio/actio
   go build -o actio .
   ```

   On Windows the binary will be `actio.exe`.

2. (Optional) Move the binary into your `PATH`:

   ```bash
   # Linux / macOS
   sudo mv actio /usr/local/bin/

   # Or add current directory to PATH
   export PATH="$PATH:$(pwd)"
   ```

3. Verify:

   ```bash
   actio version
   actio --help
   ```

## How releases are published

- The project uses **GoReleaser** (see [.goreleaser.yml](https://github.com/PRAX200OK/Actio/blob/main/.goreleaser.yml) in the repo).
- When maintainers push a **version tag** (e.g. `v0.1.0`), the **CI** workflow runs tests and then **GoReleaser** builds binaries for all supported OS/arch combinations and publishes them to the [GitHub Releases](https://github.com/PRAX200OK/Actio/releases) page.
- The docs site is deployed separately via **GitHub Actions** (GitHub Pages) on push to `main`/`master`.

## Next

- [Quick start](/docs/getting-started/quick-start) — run your first commands.
- [Create a project](/docs/getting-started/create-project) — scaffold a full Actio-enabled repo.
