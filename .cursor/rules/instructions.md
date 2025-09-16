---
description: Project overview
globs:
alwaysApply: true
---

# simplemath

Purpose: provide the concrete, current facts needed to work effectively in this repo.

Philosophy: keep code and UI simple, printable, and framework-light.

Quick facts
- Tech: Go 1.25 backend + static HTML/CSS/JS frontend.
- HTTP framework: Echo v4 (already in use).
- Product goal: generate printable elementary-school math problems (add/sub/mul/div). No user accounts or PII.
- Monetization intent: Google AdSense (informational only).
- UI, UX: simple, A4-printable, minimal interactivity. No heavy JS frameworks. Use a friendly, approachable style suitable for parents and children.

Repository structure
- `cmd/` — application entrypoint (`main.go`) and route wiring (POST `/`, GET `/`).
- `handler/` — HTTP layer:
  - `handler.go` — request orchestration and HTML output
  - `formdata.go` — form binding and validation
- `operator/` — `Operator` type with `String()` and `Symbol()` helpers.
- `gen/` — generation utilities: `RandomWithDigits`, `PowerOfTen`, `JoinOperands`.
- `statics/` — static HTML (form UI) and assets.
- `go.mod` — module config and deps.
- `README.md` — minimal project description.

What is present (discoverable)
- `cmd/main.go` — Echo server entrypoint.
  - Routes: `GET /` serves `statics/index.html`; `POST /` accepts the form and returns generated problems.
  - RNG: seeded and injected into the handler.
- `handler/` — request binding/validation (`formdata.go`) and response rendering (`handler.go`).
- `operator/` — `Operator` type with `.String()` and `.Symbol()`.
- `gen/` — helpers: `RandomWithDigits`, `PowerOfTen`, `JoinOperands`.
- `statics/index.html` — form UI with dynamic per-operand digit inputs and a print button.
- `go.mod` — module `simplemath`, Go 1.25, Echo v4, `golang.org/x/text`.
- `README.md` — minimal.

Actionable guidance for agents
- Keep server wiring in `cmd/`, HTTP logic in `handler/`, helpers in `gen/` and `operator/`. Frontend assets live in `statics/`.
- Prefer deterministic generation for tests by injecting a fixed `rand.Source`.
- Preserve printable output; use client-side print with CSS (`@media print`, optional `statics/print.css`).
- Default to stateless, in-memory behaviour (no DB).

Developer workflows (discoverable / recommended)
- Typical Go commands:
  - `go run ./cmd` or `go run ./cmd/main.go` — run the server locally
  - `go build ./...` — build packages
  - `go test ./...` — run unit tests (once tests are added)
  - `gofmt -w .` and `go vet ./...` — format and vet
- If adding CI, prefer a GitHub Actions workflow that runs `go test ./...` and `gofmt`.

Project-specific conventions
- Keep the question generator deterministic for tests by using a seedable RNG (e.g. inject `rand.Source`), then test with a fixed seed.
- Small surface area: avoid heavy third-party frameworks. Prefer stdlib `net/http`, `html/template`, and small third-party helpers only with owner approval.
- Printable UI: templates should produce simple pages (A4-friendly) and include explicit sample outputs that match the examples in `README.md`.

Integration points & external dependencies
- Echo v4 is the HTTP framework used. No databases or external services.
- Future: potential Google AdSense snippet on the frontend (informational only).

Known gaps / next steps
- Add unit tests for `gen` and `handler` validation.
- Add operator-specific rules as needed (e.g., subtraction non-negative, division constraints).
- Add print-specific CSS in `statics/print.css` to refine client-side print.
- Optional: CLI/env to fix RNG seed during development/tests.
- 

## Style Guidelines:

- Primary color: Light blue (#ADD8E6), reminiscent of sky, to imply hosting, approachable and reliable.
- Background color: Very light gray (#F5F5F5) to provide a clean, unobtrusive backdrop.
- Accent color: Soft orange (#FFB347), an analogous color for the button
- Body and headline font: 'PT Sans' (sans-serif) for a clean, modern, accessible aesthetic in both headers and body text.
- Use simple, line-based icons to represent different Firebase Hosting functionalities.
- Maintain a clean, single-column layout with ample whitespace for easy readability and navigation.
- Use subtle animations, like fade-ins and transitions, to enhance the user experience without being distracting.
