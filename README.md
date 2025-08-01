# DevsMailGo - Complete Mail Stack API

A comprehensive Go-based API service that integrates with Roundcube webmail, MySQL/MariaDB, OpenLDAP, and the mail stack (Postfix & Dovecot).

## Architecture Overview

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Roundcube     │    │   Go API        │    │   OpenLDAP      │
│   Webmail       │◄──►│   Service       │◄──►│   Server        │
│   (UI)          │    │                 │    │                 │
└─────────────────┘    └─────────────────┘    └─────────────────┘
                                │
                                ▼
                       ┌─────────────────┐
                       │   MySQL/MariaDB │
                       │   Database      │
                       └─────────────────┘
                                │
                                ▼
                       ┌─────────────────┐
                       │   Mail Stack    │
                       │   Postfix       │
                       │   Dovecot       │
                       └─────────────────┘
```

## Features

### ✅ Authentication & Authorization
- **Dual Authentication**: Local database + OpenLDAP integration
- **JWT Token Management**: Secure token-based authentication
- **Role-based Access Control**: Admin and user roles
- **Group Management**: LDAP group synchronization

### ✅ User Management
- **User CRUD Operations**: Create, read, update, delete users
- **Domain Management**: Multi-domain support
- **Alias Management**: Email alias configuration
- **Quota Management**: User mailbox quotas

### ✅ Mail Stack Integration
- **Postfix Integration**: MTA configuration and management
- **Dovecot Integration**: IMAP server management
- **Mailbox Operations**: Create, delete, manage mailboxes
- **Mail Queue Management**: Monitor and manage mail queues

### ✅ Roundcube Integration
- **User Synchronization**: Auto-sync users to Roundcube
- **Configuration Management**: Roundcube settings management
- **Statistics**: Roundcube usage statistics

### ✅ Security Features
- **Rate Limiting**: API rate limiting
- **CORS Support**: Cross-origin resource sharing
- **Input Validation**: Request validation and sanitization
- **Logging**: Comprehensive logging system

## Quick Start

### Prerequisites

- Docker and Docker Compose
- Go 1.23+ (for development)
- Make (optional, for using Makefile commands)

### 1. Clone and Setup

```bash
git clone <repository-url>
cd GO_LANG_PROJECT_SETUP
```

### 2. Environment Configuration

```bash
cp env.example .env
# Edit .env with your configuration
```

### 3. Start the Complete Stack

```bash
# Start all services
docker-compose up -d

# Or use Makefile
make up
```

### 4. Access Services

- **API Documentation**: http://localhost:8080/api/health
- **Roundcube Webmail**: http://localhost:8082
- **phpLDAPadmin**: http://localhost:8081
- **MySQL**: localhost:3306
- **LDAP**: localhost:389

## API Endpoints

### Authentication
```
POST /api/login
{
  "username": "user@example.com",
  "password": "password",
  "auth_type": "ldap"  // or "local"
}
```

### User Management
```
GET    /api/users                    # List users
POST   /api/users                    # Create user
GET    /api/users/{id}               # Get user
PUT    /api/users/{id}               # Update user
DELETE /api/users/{id}               # Delete user
```

### Domain Management
```
GET    /api/domains                  # List domains
POST   /api/domain/{domain}          # Create domain
GET    /api/domain/{domain}          # Get domain
PUT    /api/domain/{domain}          # Update domain
DELETE /api/domain/{domain}          # Delete domain
```

### Mail Management
```
POST   /api/mailbox                  # Create mailbox
DELETE /api/mailbox/{email}          # Delete mailbox
GET    /api/mailbox/{email}          # Get mailbox info
PUT    /api/mailbox/{email}/quota    # Update quota
POST   /api/mail/reload              # Reload mail services
GET    /api/mail/check               # Check mail delivery
GET    /api/mail/queue               # Get mail queue
```

### Alias Management
```
GET    /api/aliases                  # List aliases
POST   /api/alias/{address}          # Create alias
GET    /api/alias/{address}          # Get alias
PUT    /api/alias/{address}          # Update alias
DELETE /api/alias/{address}          # Delete alias
```

## Configuration

### Environment Variables

#### Database
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=devsmailgo
DB_PASSWORD=your_secure_password
DB_NAME=devsmailgo
```

#### LDAP
```env
LDAP_HOST=localhost
LDAP_PORT=389
LDAP_BASE_DN=dc=example,dc=com
LDAP_BIND_DN=cn=admin,dc=example,dc=com
LDAP_BIND_PASSWORD=your_ldap_admin_password
LDAP_USER_FILTER=(uid=%s)
LDAP_GROUP_FILTER=(memberUid=%s)
LDAP_USE_SSL=false
LDAP_USE_TLS=false
```

#### Server
```env
SERVER_PORT=8080
SERVER_HOST=0.0.0.0
ENVIRONMENT=development
JWT_SECRET=your-super-secret-jwt-key-change-in-production
JWT_EXPIRY_HOURS=24
```

## Development

### Local Development Setup

1. **Install Dependencies**
```bash
go mod download
```

2. **Run Database Migrations**
```bash
go run cmd/main.go
```

3. **Start Development Server**
```bash
go run cmd/main.go
```

### Testing

```bash
# Run all tests
go test ./...

# Run specific test
go test ./tests/integration_test.go
```

### Building

```bash
# Build for current platform
go build -o bin/devsmailgo cmd/main.go

# Build for Docker
docker build -t devsmailgo .
```

## Docker Commands

```bash
# Start all services
docker-compose up -d

# Stop all services
docker-compose down

# View logs
docker-compose logs -f api

# Rebuild and restart
docker-compose up -d --build

# Access MySQL
docker-compose exec mysql mysql -u devsmailgo -p devsmailgo

# Access LDAP
docker-compose exec openldap ldapsearch -x -H ldap://localhost -b dc=example,dc=com -D "cn=admin,dc=example,dc=com" -w admin_password
```

## LDAP Setup

### Initial LDAP Structure

Create `ldap/init/01-users.ldif`:

```ldif
dn: ou=users,dc=example,dc=com
objectClass: organizationalUnit
ou: users

dn: uid=admin,ou=users,dc=example,dc=com
objectClass: inetOrgPerson
objectClass: posixAccount
objectClass: shadowAccount
uid: admin
sn: Administrator
givenName: System
cn: System Administrator
displayName: System Administrator
uidNumber: 1000
gidNumber: 1000
userPassword: {SSHA}admin_password_hash
homeDirectory: /home/admin
loginShell: /bin/bash
mail: admin@example.com

dn: ou=groups,dc=example,dc=com
objectClass: organizationalUnit
ou: groups

dn: cn=admins,ou=groups,dc=example,dc=com
objectClass: posixGroup
cn: admins
gidNumber: 1001
memberUid: admin
```

## Mail Stack Configuration

### Postfix Configuration

The API will automatically configure Postfix for:
- Virtual domains
- Virtual aliases
- SMTP authentication
- Mail routing

### Dovecot Configuration

The API will configure Dovecot for:
- IMAP/POP3 services
- User authentication
- Mailbox storage
- Quota management

## Security Considerations

1. **Change Default Passwords**: Update all default passwords in production
2. **Use SSL/TLS**: Enable SSL/TLS for LDAP and mail services
3. **Firewall Configuration**: Configure firewall rules appropriately
4. **Regular Updates**: Keep all components updated
5. **Backup Strategy**: Implement regular backups for data and configuration

## Troubleshooting

### Common Issues

1. **LDAP Connection Failed**
   - Check LDAP server is running
   - Verify credentials and base DN
   - Check network connectivity

2. **Mail Delivery Issues**
   - Verify Postfix configuration
   - Check DNS records (MX, SPF, DKIM)
   - Review mail logs

3. **Database Connection Issues**
   - Verify MySQL is running
   - Check database credentials
   - Ensure database exists

### Logs

```bash
# API logs
docker-compose logs api

# MySQL logs
docker-compose logs mysql

# LDAP logs
docker-compose logs openldap

# Postfix logs
docker-compose logs postfix

# Dovecot logs
docker-compose logs dovecot
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Support

For support and questions:
- Create an issue on GitHub
- Check the documentation
- Review the troubleshooting section 