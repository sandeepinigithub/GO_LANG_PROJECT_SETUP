# DevsMailGo Integration Guide

## What Has Been Implemented

### ✅ 1. OpenLDAP Integration
- **LDAP Service** (`service/ldap_service.go`): Complete LDAP client implementation
- **Authentication Support**: Dual authentication (local + LDAP)
- **User Search**: Search users by UID, email, name
- **Group Management**: Retrieve user groups and roles
- **Configuration**: LDAP connection settings in config
- **User Sync**: Sync LDAP users to local database

### ✅ 2. Mail Stack Integration (Postfix & Dovecot)
- **Mail Service** (`service/mail_service.go`): Complete mail stack management
- **Mailbox Operations**: Create, delete, manage mailboxes
- **Quota Management**: Set and update user quotas
- **Service Management**: Reload Postfix/Dovecot configurations
- **Queue Management**: Monitor mail queues
- **Health Checks**: Verify mail delivery status

### ✅ 3. Roundcube Integration
- **Roundcube Service** (`service/roundcube_service.go`): Roundcube user management
- **User Synchronization**: Auto-sync users to Roundcube
- **Configuration Management**: Manage Roundcube settings
- **Statistics**: Roundcube usage statistics

### ✅ 4. Enhanced Authentication
- **Dual Auth Support**: Local database + OpenLDAP authentication
- **Role-based Access**: Admin and user roles based on LDAP groups
- **JWT Integration**: Enhanced JWT tokens with user roles and groups

### ✅ 5. API Endpoints
- **Mail Management**: Complete mailbox CRUD operations
- **Service Management**: Mail service reload and health checks
- **Enhanced Auth**: LDAP authentication endpoints

### ✅ 6. Docker Compose Stack
- **Complete Stack**: All services in one docker-compose file
- **OpenLDAP**: LDAP server with phpLDAPadmin
- **Mail Stack**: Postfix + Dovecot containers
- **Roundcube**: Webmail interface
- **MySQL**: Database server
- **Redis**: Caching and sessions

## Current Architecture

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

## Next Steps to Complete the Integration

### 1. Environment Setup

#### Create Environment File
```bash
cp env.example .env
```

#### Configure LDAP Settings
```env
LDAP_HOST=localhost
LDAP_PORT=389
LDAP_BASE_DN=dc=example,dc=com
LDAP_BIND_DN=cn=admin,dc=example,dc=com
LDAP_BIND_PASSWORD=admin_password
LDAP_USER_FILTER=(uid=%s)
LDAP_GROUP_FILTER=(memberUid=%s)
LDAP_USE_SSL=false
LDAP_USE_TLS=false
```

### 2. LDAP Initial Setup

#### Create LDAP Directory Structure
```bash
mkdir -p ldap/init
```

#### Create Initial Users LDIF
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

### 3. Start the Complete Stack

```bash
# Start all services
docker-compose up -d

# Check service status
docker-compose ps

# View logs
docker-compose logs -f api
```

### 4. Test the Integration

#### Test LDAP Authentication
```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "admin",
    "password": "admin_password",
    "auth_type": "ldap"
  }'
```

#### Test Mailbox Creation
```bash
# First get a JWT token
TOKEN=$(curl -s -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "admin_password", "auth_type": "ldap"}' \
  | jq -r '.token')

# Create a mailbox
curl -X POST http://localhost:8080/api/mailbox \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{
    "email": "user@example.com",
    "quota": 104857600
  }'
```

#### Test Mail Service Health
```bash
curl -X GET http://localhost:8080/api/mail/check \
  -H "Authorization: Bearer $TOKEN"
```

### 5. Access Web Interfaces

- **Roundcube Webmail**: http://localhost:8082
- **phpLDAPadmin**: http://localhost:8081
- **API Health Check**: http://localhost:8080/api/health

## Configuration Files Needed

### 1. Postfix Configuration
The API will automatically configure Postfix, but you may need to customize:
- Virtual domains
- Virtual aliases
- SMTP authentication
- Mail routing

### 2. Dovecot Configuration
The API will configure Dovecot for:
- IMAP/POP3 services
- User authentication
- Mailbox storage
- Quota management

### 3. Roundcube Configuration
Roundcube will be automatically configured with:
- Database connection
- SMTP/IMAP server settings
- User authentication

## Security Considerations

### 1. Change Default Passwords
- Update all default passwords in production
- Use strong, unique passwords for each service

### 2. Enable SSL/TLS
- Configure SSL certificates for LDAP
- Enable TLS for mail services
- Use HTTPS for web interfaces

### 3. Firewall Configuration
- Open only necessary ports
- Restrict access to admin interfaces
- Configure proper network segmentation

### 4. Regular Maintenance
- Keep all components updated
- Monitor logs regularly
- Implement backup strategies

## Troubleshooting

### Common Issues

1. **LDAP Connection Failed**
   ```bash
   # Check LDAP server
   docker-compose logs openldap
   
   # Test LDAP connection
   docker-compose exec openldap ldapsearch -x -H ldap://localhost -b dc=example,dc=com -D "cn=admin,dc=example,dc=com" -w admin_password
   ```

2. **Mail Delivery Issues**
   ```bash
   # Check Postfix logs
   docker-compose logs postfix
   
   # Check mail queue
   docker-compose exec postfix mailq
   ```

3. **Database Connection Issues**
   ```bash
   # Check MySQL logs
   docker-compose logs mysql
   
   # Test database connection
   docker-compose exec mysql mysql -u devsmailgo -p devsmailgo
   ```

### Log Locations
- **API Logs**: `docker-compose logs api`
- **MySQL Logs**: `docker-compose logs mysql`
- **LDAP Logs**: `docker-compose logs openldap`
- **Postfix Logs**: `docker-compose logs postfix`
- **Dovecot Logs**: `docker-compose logs dovecot`

## API Usage Examples

### Authentication
```bash
# Local authentication
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"username": "user@example.com", "password": "password", "auth_type": "local"}'

# LDAP authentication
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "admin_password", "auth_type": "ldap"}'
```

### User Management
```bash
# Create user
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"email": "newuser@example.com", "name": "New User", "password": "password"}'

# List users
curl -X GET http://localhost:8080/api/users \
  -H "Authorization: Bearer $TOKEN"
```

### Mail Management
```bash
# Create mailbox
curl -X POST http://localhost:8080/api/mailbox \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"email": "user@example.com", "quota": 104857600}'

# Get mailbox info
curl -X GET http://localhost:8080/api/mailbox/user@example.com \
  -H "Authorization: Bearer $TOKEN"

# Update quota
curl -X PUT http://localhost:8080/api/mailbox/user@example.com/quota \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  -d '{"quota": 209715200}'
```

## Production Deployment

### 1. Environment Configuration
- Use production-grade passwords
- Enable SSL/TLS for all services
- Configure proper logging levels
- Set up monitoring and alerting

### 2. Backup Strategy
- Regular database backups
- Configuration file backups
- Mail data backups
- LDAP data backups

### 3. Monitoring
- Set up health checks
- Monitor service logs
- Track API usage
- Monitor mail delivery

### 4. Scaling
- Consider load balancing for API
- Database replication
- Mail server clustering
- Caching strategies

## Support and Maintenance

### Regular Tasks
1. **Daily**: Check service health and logs
2. **Weekly**: Review security logs and update packages
3. **Monthly**: Backup verification and performance review
4. **Quarterly**: Security audit and configuration review

### Emergency Procedures
1. **Service Outage**: Check logs and restart services
2. **Data Loss**: Restore from backups
3. **Security Breach**: Isolate affected services and investigate
4. **Performance Issues**: Scale resources or optimize configuration

This integration guide provides a complete roadmap for implementing and maintaining your DevsMailGo mail stack with OpenLDAP, Postfix, Dovecot, and Roundcube integration. 