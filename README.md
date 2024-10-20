# Small cli tool to generate pdfs filled with code

> [!NOTE]  
> Barely any time was spent on this project, it is just a tool i had a need for myself that i made quickly.

## Usage

Install the tool, create a `.pdfconfig.yaml` like [this example](https://github.com/Phillezi/pdfgenerator/blob/main/.pdfconfig.yaml) in your cwd where you wish to run it and then run:

```bash
pdfgenerator
```

## Install

## With Go

If go is installed you can run

```bash
go install github.com/Phillezi/pdfgenerator@latest
```

### Download and install binary

#### Mac and Linux

For Mac and Linux, there is an installation script that can be run to install it.

##### Prerequisites

- bash
- curl

```bash
curl -fsSL https://raw.githubusercontent.com/Phillezi/pdfgenerator/main/scripts/install.sh | bash

```

Check out what the script does [here](https://github.com/Phillezi/pdfgenerator/blob/main/scripts/install.sh).

#### Windows

There is a PowerShell installation script that can be run to install the CLI.

```powershell
powershell -c "irm https://raw.githubusercontent.com/Phillezi/pdfgenerator/main/scripts/install.ps1 | iex"

```

Check out what the script does [here](https://github.com/Phillezi/pdfgenerator/blob/main/scripts/install.ps1).
