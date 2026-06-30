// sections/navbar.ts

import {
  Home,
  // ScanSearch,
  // Monitor,
  // FileText,
  // Package,
  // LifeBuoy,
  Monitor,
} from "lucide-react";

export const navItems = [
  {
    title: "Home",
    route: "dashboard",
    icon: Home,
  },
  // {
  //   title: "Scan",
  //   route: "scan",
  //   icon: ScanSearch,
  // },
  // {
  //   title: "Logs",
  //   route: "logs",
  //   icon: FileText,
  // },
  // {
  //   title: "Software",
  //   route: "software",
  //   icon: Package,
  // },
  // {
  //   title: "Support",
  //   route: "support",
  //   icon: LifeBuoy,
  // },
  {
    title: "Agent",
    route: "agent",
    icon: Monitor,
  },
];