export const dash = {
    hostname: "DESKTOP-ABC123",
    agentId: "agnt_01HQkasli_w2026",
    status: "Online",
    security: "Protected",
    lastscan: "",
    os: "Windows 11 Pro",
    sec: [
        {title:"Windows Defender", value:true},
        {title:"Firewall", value:true},
        {title:"TPM", value:true},
        {title:"BitLocker", value:false},

    ],
    dev: [
        {title:"Hostname", value:"DESKTOP-ABC123"},
        {title:"Domain", value:"Workgroup"},
        {title:"User", value:"Administrator"},
        {title:"OS", value:"Windows 11 Pro"},
        {title:"Agent Version", value:"v1.0.1"},
    ],
    stat: [
        {title: "CPU Usage", value:64},
        {title: "RAM Usage", value:82},
        {title: "Disk Usage", value:41},
        {title: "Network", value:80}
    ]
}

export const upgradeAgent = {
    version: "v1.0.2",
    tag:"latest"
}

export const software = [
  {
    id: 1,
    tag: "sw-01",
    name: "Wireshark Packet Analyzer",
    developer: "The Wireshark Foundation",
    category: "Networking",
    version: "4.2.5",
    size: "52.4 MB",
    status: "Installed",
    architecture: "x64",
    description: "Deep inspection of hundreds of network protocols with live capture and offline analysis capability."
  },
  {
    id: 2,
    tag: "sw-02",
    name: "Visual Studio Code (Hardened)",
    developer: "Microsoft Corp (SecOps Mirror)",
    category: "Development",
    version: "1.90.1",
    size: "94.1 MB",
    status: "Update Available",
    architecture: "Universal",
    description: "Code editing redefined and optimized with custom enterprise proxy telemetry streaming validation."
  },
  {
    id: 3,
    tag: "sw-03",
    name: "Burp Suite Community Edition",
    developer: "PortSwigger Ltd",
    category: "Security",
    version: "2024.5.1",
    size: "284.0 MB",
    status: "Not Installed",
    architecture: "x64",
    description: "Best-in-class web security testing toolkit for penetration testing and web vulnerability assessments."
  },
  {
    id: 4,
    tag: "sw-04",
    name: "OpenVPN Daemon Client",
    developer: "OpenVPN Inc.",
    category: "Networking",
    version: "2.6.10",
    size: "12.8 MB",
    status: "Installed",
    architecture: "ARM64",
    description: "Full-featured open-source SSL VPN solution configuration routing endpoint binaries wrapper."
  },
  {
    id: 5,
    tag: "sw-05",
    name: "Sysinternals Suite Lite",
    developer: "Sysinternals - Microsoft",
    category: "Utility",
    version: "2024.03",
    size: "45.7 MB",
    status: "Not Installed",
    architecture: "Universal",
    description: "Advanced system utilities to manage, troubleshoot and diagnose internal Windows kernel instances."
  }
]

export const support = {
  limit: {
    minDesc: 10,
    maxDesc: 200,
    file: {
      mandatory: false,
      type: ["png", "jpeg", "jpg"],
      min: 0,
      max: 10,
    },
  },
  categories: [
    {
      id: 1,
      tag: "infra",
      title: "Infrastructure",
      scategory: [
        {
          id: 201,
          tag: "cnod",
          title: "Compute Nodes",
          subject: [
            "High Avail cluster nodes load variance warning"
          ]
        },
        {
          id: 202,
          tag: "ndns",
          title: "Network / DNS",
          subject: [
            "Internal DNS routing resolution timeout"
          ]
        }
      ]
    },
    {
      id: 2,
      tag: "acst",
      title: "Access Control",
      scategory: [
        {
          id: 301,
          tag: "oath",
          title: "OAuth / Tokens",
          subject: [
            "Provisioning Token Expiry Overhaul"
          ]
        }
      ]
    },
    {
      id: 3,
      tag: "pkgm",
      title: "Package Mirror",
      scategory: [
        {
          id: 401,
          tag: "urep",
          title: "Upstream Repos",
          subject: [
            "Upstream Alpine Mirror Synch Failure"
          ]
        }
      ]
    }
  ],
  priority: [
    "Low", 
    "Medium", 
    "High", 
    "Critical"
  ],
  tickets: [
    {
      id: 1,
      tktId: "TKT-8942",
      category: "Package Mirror",
      subcategory: "Upstream Repos",
      subject: "Upstream Alpine Mirror Synch Failure",
      priority: "Critical",
      status: "In Progress",
      created_at: "2026-06-23 09:14 AM",
      updated_at: "2026-06-24 14:02 PM",
      description: "Automated cron execution caught an invalid SHA-256 signature mismatch rule flag when attempting to manifest the global stable tree index updates."
    },
    {
      id: 2,
      tktId: "TKT-8711",
      category: "Access Control",
      subcategory: "OAuth / Tokens",
      subject: "Provisioning Token Expiry Overhaul",
      priority: "Medium",
      status: "Open",
      created_at: "2026-06-24 14:02 PM",
      updated_at: "2026-06-24 14:02 PM",
      description: "Requesting clear architecture verification parameters to extend the base OAuth application expiration lifecycle thresholds past 72 hours."
    },
    {
      id: 3,
      tktId: "TKT-8290",
      category: "Infrastructure",
      subcategory: "Compute Nodes",
      subject: "High Avail cluster nodes load variance warning",
      priority: "High",
      status: "Resolved",
      created_at: "2026-06-20 11:45 AM",
      updated_at: "2026-06-24 14:02 PM",
      description: "Resource scheduler imbalance inside node cluster group 4-B threw automated system alerts."
    }
  ]
};

export const ticketTimeline = [
  {
    id: 564,
    tktId: "TKT-8290",
    author: {
      name: "Sameer",
      id: 501
    },
    created_at: "2026-06-24 14:02 AM",
    description: "Same has been resolved"
  }
]