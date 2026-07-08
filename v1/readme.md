# Fehmi Endpoint Agent (v1)

> Cross-platform enterprise endpoint management agent developed by **Fehmi Corporation**.

The Fehmi Endpoint Agent is a modular, enterprise-grade endpoint management platform designed for Windows, macOS, and Linux. It provides inventory management, monitoring, remote support, security enforcement, automation, software deployment, and centralized management through the Fehmi Cloud Console.

---

# Repository

```
github.com/fehmicorp/agent/v1
```

---

# Project Objectives

- Cross-platform architecture
- Single codebase for UI
- Platform-specific workers
- Modular package architecture
- Lightweight background services
- Offline-first synchronization
- Secure communication
- Enterprise scalability
- Easy installer generation
- Future microservice integration

---

# Project Structure

```
agent/
└── v1/
    │
    ├── app/                     # Universal Wails Desktop Application (UI)
    │
    ├── cmd/                     # Build runners, installers and packaging
    │
    ├── pkg/
    │   ├── win/                 # Windows implementation
    │   ├── mac/                 # macOS implementation
    │   └── lnx/                 # Linux implementation
    │
    ├── internal/               # Shared internal modules
    ├── api/                    # REST/GRPC/Websocket interfaces
    ├── configs/
    ├── scripts/
    ├── assets/
    ├── docs/
    └── build/
```

---

# Folder Overview

## app/

```
D:\Projects\backend\agent\v1\app
```

Universal desktop application.

Technology

- Wails v2
- React
- TypeScript
- TailwindCSS

Responsibilities

- Dashboard
- Device information
- Endpoint status
- Software inventory
- Remote support UI
- Security dashboard
- Settings
- Update manager
- Logs
- User authentication
- Notification center

The UI should never contain platform-specific implementation.

All platform operations must communicate with packages under `pkg`.

---

## pkg/

Contains all operating system implementations.

```
pkg/
```

---

### pkg/win/

Windows-specific implementation.

```
pkg/win/
```

Contains

- Windows Service
- Registry
- Event Log
- Windows Defender
- BitLocker
- WMI
- Active Directory
- Local Users
- Installed Software
- Startup Manager
- Task Scheduler
- Services
- Firewall
- Certificates
- Drivers
- Device Inventory
- Windows Update
- Remote Worker
- Inventory Worker
- Sync Worker
- Script Engine

---

### pkg/mac/

macOS implementation.

```
pkg/mac/
```

Contains

- LaunchDaemon
- LaunchAgent
- Keychain
- System Profiler
- Security Framework
- Software Inventory
- MDM Integration
- Remote Worker
- Sync Worker
- Inventory Worker

---

### pkg/lnx/

Linux implementation.

```
pkg/lnx/
```

Contains

- systemd
- DBus
- Package Managers
- IPTables
- nftables
- SSH
- Device Inventory
- Software Inventory
- Remote Worker
- Sync Worker
- Monitoring

---

# cmd/

```
D:\Projects\backend\agent\v1\cmd
```

Contains executable runners.

Examples

```
cmd/
    windows/
    macos/
    linux/
    installer/
    updater/
    bootstrap/
```

Responsibilities

- Build executable
- Generate installer
- Package resources
- Embed assets
- Code signing
- Release generation
- MSI generation
- DMG generation
- DEB/RPM generation

---

# internal/

Shared business logic.

Example

```
internal/

inventory/
monitor/
security/
sync/
database/
network/
remote/
policy/
assets/
users/
device/
certificate/
vpn/
logging/
events/
scheduler/
storage/
```

No operating system specific code should exist here.

---

# Architecture

```
                 Wails UI
                    │
                    │
            Application Layer
                    │
     ┌──────────────┼──────────────┐
     │              │              │
  Windows        macOS          Linux
     │              │              │
 Platform APIs  Platform APIs  Platform APIs
     │              │              │
     └──────────────┼──────────────┘
                    │
              Common Services
                    │
             Sync / Database
                    │
              REST / WebSocket
                    │
             Fehmi Cloud Console
```

---

# Supported Platforms

- Windows 10
- Windows 11
- Windows Server

- macOS

- Ubuntu
- Debian
- CentOS
- Rocky Linux
- RedHat

---

# Core Modules

- Asset Inventory
- Software Inventory
- Hardware Inventory
- Device Monitoring
- Active Window Tracking
- Browser Activity
- USB Control
- Device Control
- Patch Management
- Script Execution
- Remote Support
- Remote Shell
- Remote File Explorer
- Process Monitoring
- Service Monitoring
- Endpoint Security
- Antivirus Integration
- BitLocker Management
- Firewall Management
- Policy Enforcement
- User Management
- VPN Management
- Event Collection
- Performance Monitoring
- Offline Synchronization
- Notification System
- Configuration Management

---

# Communication

Supported protocols

- HTTPS REST
- WebSocket
- gRPC
- MQTT (Future)
- Message Queue (Future)

---

# Local Storage

- SQLite
- Local Queue
- Cache
- Encrypted Credentials

---

# Cloud Integration

The agent communicates with

- Fehmi Cloud Console
- Notification Service
- Update Service
- Policy Service
- Authentication Service
- Inventory Service

---

# Development Principles

- Platform abstraction
- Modular packages
- Dependency injection
- Shared interfaces
- Offline-first
- Secure by default
- Least privilege
- Event-driven
- Thread-safe workers
- Minimal resource consumption

---

# Build Targets

```
Windows
    MSI
    EXE

macOS
    APP
    DMG

Linux
    DEB
    RPM
    AppImage
```

---

# Future Roadmap

- Zero Trust Integration
- MDM
- RMM
- AI Assistant
- Endpoint Detection & Response (EDR)
- Patch Automation
- Application Deployment
- Compliance Engine
- Cloud Policy Engine
- Multi-Tenant Architecture
- Remote Desktop
- Remote Terminal
- Certificate Management
- VPN Provisioning
- Kubernetes Agent
- Container Monitoring

---

# License

Copyright © Fehmi Corporation.

All Rights Reserved.