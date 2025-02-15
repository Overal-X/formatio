# Formatio

Formatio is an open-source cloud platform designed to simplify application deployment and management. It provides a streamlined experience for developers by automating infrastructure setup, scaling, and monitoring. With Formatio, you can easily deploy applications, manage resources, and ensure high availability without deep infrastructure expertise.

## Roadmap
https://github.com/Overal-X/formatio/milestones

## Development

### Prerequisites
Ensure you have the following installed:

- [Go](https://go.dev/dl/)
- [Node.js](https://nodejs.org/en/download/) (for frontend development)
- [Svelte](https://svelte.dev/)
- [Yarn](https://yarnpkg.com/getting-started/install) (for package management)
- [Docker](https://www.docker.com/) & [Docker Compose](https://docs.docker.com/compose/install/)

### Clone the Repository
```sh
 git clone https://github.com/overal-x/formatio.git
 cd formatio
```

### Backend Setup
```sh
cd server
cp .env.example .env  # Configure environment variables
docker-compose up -d  # Start Postgres, Redis, and RabbitMQ
make run  # Start the backend
```

### Frontend Setup
```sh
cd client
yarn install  # Install dependencies
yarn dev  # Start the frontend
```

## Contributing

We welcome contributions from the community! Here’s how you can get started:

1. **Understand the Monorepo Structure:**
   - `server/` - Contains the backend code (Golang, Postgres, Redis, RabbitMQ)
   - `client/` - Contains the frontend code (Svelte 5)

2. **Setup Your Development Environment:**
   - Follow the installation steps above.

3. **Follow the Contribution Guidelines:**
   - Fork the repository and create a feature branch.
   - Write clear and concise commit messages.
   - Ensure your code passes tests and linting.

4. **Submit a Pull Request:**
   - Describe your changes in detail.
   - Reference related issues (if any).

5. **Join Discussions:**
   - Open an issue for feature requests or bugs.
   - Engage in discussions via GitHub or our community channels.

## License
Formatio is released under the [MIT License](LICENSE).

## Community & Support
- GitHub Issues: [https://github.com/overal-x/formatio/issues](https://github.com/overal-x/formatio/issues)
- Discussions: [https://github.com/overal-x/formatio/discussions](https://github.com/overal-x/formatio/discussions)

