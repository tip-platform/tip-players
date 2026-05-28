# Security Policy

## Supported versions

| Version             | Security Support |
| ------------------- | ---------------- |
| `main`              | ✅ Active        |
| Any previous branch | ❌ No support    |

Only the `main` branch receives security patches. If you find a vulnerability in a previous version, first check to see if it has already been fixed in `main`.

---

## Report a vulnerability

**Do not open a public issue.** Public issues expose the vulnerability before a fix is available. Use the private GitHub channel:

1. Go to the **Security** tab of the repository
2. Click **"Report a vulnerability"**
3. Describe the problem in as much detail as possible If you prefer by email, write to: `idamendjaume@gmail.com`

---

## What to include in the reporter

So that you can reproduce and fix the problem as soon as possible, include:

- Clear description of the vulnerability
- Steps to reproduce it - Potential impact (what an attacker can do)
- Version or commit affected
- If you have a fix proposal, it is welcome

---

## Process and response times

| Stage                       | Estimated time                                        |
| --------------------------- | ----------------------------------------------------- |
| Confirmation of receipt     | 48 hours                                              |
| Initial severity assessment | 5 business days                                       |
| Fix or mitigation plan      | Depends on severity — critical: 7 days, high: 15 days |
| Publication of the advisory | Coordinated with you before going public              |

Once corrected, a **Security Advisory** will be published in the repository with the CVE if applicable, crediting the reporter (unless you prefer anonymity).

---

## Vulnerabilities in dependencies

This projects uses **Dependabot** to monitor dependencies. If you detect a vulnerability in a dependency before Dependabot reports it, follow the same process described above.
