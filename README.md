# Go-CLI 

## Overview

This project provides a command-line interface (CLI) application with various functionalities, including:

* **Interactive Quiz:** Test your knowledge on a variety of topics (e.g., capitals of the world, historical dates) in a user-friendly quiz format.
* **Coffee Simulation:** Enjoy a simulated coffee brewing experience, specifying sugar level, milk preference, and strength.
* **Date and Time Display:** Get the current date, time, and day of the week conveniently at your fingertips.

## Features

* **User-Friendly Interface:** The application offers a clear and easy-to-use text-based interface for interacting with its features.
* **Customization:** Adjust quiz settings, coffee preferences, and choose how you want to receive information (date/time display).
* **Validation:** The application validates user input to ensure a smooth and error-free experience.

## Build

**Prerequisites:**

* **Go Programming Language:** To build the application from source, you'll need the Go programming language installed on your system. Download and install Go from the official website ([https://golang.org/dl/](https://golang.org/dl/)).

**Building from Source:**

1. **Clone the Repository:** If you have the source code hosted on a version control platform like Git, clone the repository using a command like:

   ```bash
   git clone https://github.com/OleksiyM/Go-CLI.git
   ```

2. **Navigate to Project Directory:** Use the `cd` command to enter the project directory:

   ```bash
   cd Go-CLI
   ```

3. **Build the Application:** Run the following command in the terminal:

   ```bash
   go build -ldflags="-s -w" go-cli.go
   ```

    - `go build`: Instructs the Go compiler to build the application.
    - `-ldflags="-s -w"`: Optimization flags for smaller binary size (`-s`) and removal of debugging information (`-w`).

**Using Pre-Built Binaries:**

If you don't want to build from source, you can download pre-built binaries for your operating system from the project's [Releases](https://github.com/OleksiyM/Go-CLI/releases) section. Once downloaded, make the binary executable (if necessary) and place it in a directory accessible from your system's PATH environment variable.

## Usage

To start the application, simply run the built executable (e.g., `./go-cli` on Linux/macOS or `go-cli.exe` on Windows) from the command line. The application will display a list of available commands. Select the desired command by typing its name and pressing Enter. Follow any further prompts or instructions provided for specific features (e.g., entering quiz parameters, configuring coffee preferences).

## Project Structure

The project likely adheres to a common Go project structure. Here's a general outline:

- **go-cli.go:** The main entry point of the application.
- **(Optional) Other Go source files:** Additional functionalities might be implemented in separate Go source files for better organization.
- **(Optional) configuration files:** The application may utilize configuration files (e.g., JSON, YAML) to store settings or data.

## Technologies Used

* **Go Programming Language:** The application is built using the Go programming language, known for its efficiency and ease of use.
* **Command-Line Interface (CLI):** The application interacts with the user through a text-based interface, providing control over functionalities.
* **(Optional) Additional Libraries:** Depending on specific features, the application might use additional Go libraries for tasks like input validation, random number generation, etc.
