# Grade Evaluator CLI

A robust command-line interface utility built in Go that processes numeric exam scores, evaluates them against specific performance bounds, and maps them to academic tiers.

This project implements an **expressionless `switch` architecture**, sequential range checks, explicit type assertions with error catching via `strconv`, and controlled conditional cascading using the `fallthrough` keyword.

## Features

*   **Sequential Bound Analysis:** Utilizes a conditional logic stack that safely evaluates numerical variables against relative range thresholds from strictest down to most general.
*   **Cascading Logic via Fallthrough:** Leverages the `fallthrough` keyword to attach descriptive secondary attributes to exact match boundaries (e.g., distinguishing a perfect 100 before merging into general tier assessments).
*   **Defensive Type Parsing:** Intercepts raw terminal string data, isolates single arguments using `strings.Fields`, and validates numerical type compliance via `strconv.Atoi` error checking before running business logic.

## Prerequisites

*   **Go**: Version 1.16 or higher installed on your local machine.

## Project Structure

```text
.
├── main.go      # Numerical parsing, multi-tier validation & switch layout
└── README.md    # Documentation