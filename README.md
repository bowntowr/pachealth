<a id="readme-top"></a>

<!-- PROJECT SHIELDS -->

[![Go][go-shield]][go-url]
[![License][license-shield]][license-url]
[![Issues][issues-shield]][issues-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">

<h1 align="center">pachealth</h1>

<p align="center">
  A lightweight Arch Linux system health checker.
  <br />
  <br />
  <a href="#usage">Usage</a>
  &middot;
  <a href="https://github.com/YOUR_USERNAME/pachealth/issues">Report Bug</a>
  &middot;
  <a href="https://github.com/YOUR_USERNAME/pachealth/issues">Request Feature</a>
</p>

</div>



<!-- TABLE OF CONTENTS -->
<details>
<summary>Table of Contents</summary>

- [About The Project](#about-the-project)
- [Features](#features)
- [Built With](#built-with)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)

</details>



<!-- ABOUT THE PROJECT -->
## About The Project

`pachealth` is an Arch Linux maintenance and health checker designed to give users a quick overview of their system's package and configuration state.

Instead of manually running multiple commands to check package updates, orphaned packages, `.pacnew` files, failed services, and other maintenance tasks, `pachealth` provides a single, easy-to-understand report.

The goal is not to replace `pacman`, but to provide a friendly diagnostic layer on top of the Arch ecosystem.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



## Features

Current/planned checks:

- Package update availability
- `.pacnew` and `.pacsave` file detection
- Orphaned package detection
- Package integrity checks
- Failed systemd service detection
- Package cache size monitoring
- Detailed health reports
- JSON output for scripting and integrations

<p align="right">(<a href="#readme-top">back to top</a>)</p>



## Built With

- [![Go][go-shield]][go-url]
- Arch Linux tools:
  - pacman
  - checkupdates
  - pacman-contrib utilities
  - systemd utilities

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->
## Getting Started

### Prerequisites

- Arch Linux
- Go 1.XX or newer (for building from source)

Install required dependencies:

```bash
sudo pacman -S go pacman-contrib
````

### Installation

#### Build from source

```bash
git clone https://github.com/YOUR_USERNAME/pachealth.git
cd pachealth

go build -o pachealth ./cmd/pachealth
```

Run:

```bash
./pachealth
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE EXAMPLES -->

## Usage

Run a system health check:

```bash
pachealth
```

Show additional details:

```bash
pachealth --verbose
```

Output JSON:

```bash
pachealth --json
```

Check `.pacnew` files:

```bash
pachealth pacnew
```

Example output:

```
pachealth

System Status: ATTENTION REQUIRED

✔ Package integrity OK
⚠ 12 package updates available
⚠ 2 .pacnew files found
✔ No failed systemd services
⚠ 3 orphaned packages found
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ROADMAP -->

## Roadmap

* [ ] Basic health check framework
* [ ] Package update detection
* [ ] `.pacnew` detection
* [ ] Orphan package detection
* [ ] Failed service checks
* [ ] JSON output
* [ ] Improved `.pacnew` workflow
* [ ] AUR package release

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- CONTRIBUTING -->

## Contributing

Contributions, bug reports, and feature suggestions are welcome.

1. Fork the project
2. Create your feature branch:

```bash
git checkout -b feature/my-feature
```

3. Commit your changes:

```bash
git commit -m "Add my feature"
```

4. Push your branch:

```bash
git push origin feature/my-feature
```

5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- LICENSE -->

## License

Distributed under the MIT License. See `LICENSE` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS -->

## Acknowledgments

* Arch Linux community and documentation
* `pacman-contrib` utilities
* The open-source Linux ecosystem

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS -->

[go-shield]: https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white
[go-url]: https://go.dev/
[license-shield]: https://img.shields.io/github/license/YOUR_USERNAME/pachealth?style=for-the-badge
[license-url]: https://github.com/YOUR_USERNAME/pachealth/blob/main/LICENSE
[issues-shield]: https://img.shields.io/github/issues/YOUR_USERNAME/pachealth?style=for-the-badge
[issues-url]: https://github.com/YOUR_USERNAME/pachealth/issues
