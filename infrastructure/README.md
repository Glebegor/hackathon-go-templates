# Infrastructure

This folder contains everything needed to run the project outside of local
`npm start` / `python app.py` style commands. It includes cloud and container
setup, CI/CD, and configuration examples.

## Structure

- `terraform/` – Infrastructure as Code (Terraform configs)
  - `dev/`, `staging/`, `prod/` – example environments

- `docker/` – Docker-related files
  - `backend/`, `frontend/` – Dockerfiles and docker-compose files
  - `reverse-proxy/` – nginx/traefik configs for routing

- `k8s/` – Kubernetes manifests (optional)
  - `base/` – core deployments/services
  - `overlays/` – environment-specific tweaks

- `ci-cd/` – Pipeline definitions for GitHub Actions / GitLab / etc.

- `config/` – Example config and `.env.example` files

- `scripts/` – Helper scripts for setup, migrations, seeding, etc.

- `observability/` – Example configs for logging/monitoring (Grafana, Prometheus, etc.)

> Note: All files here are **templates**. You should adapt them to your own
> cloud provider, deployment platform, and secrets management.

