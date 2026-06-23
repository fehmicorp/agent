import { lazy } from "react";

export const routes = {
  dashboard: lazy(() => import("./Dashboard")),
  scan: lazy(() => import("./Scan")),
  agent: lazy(() => import("./Agent")),
  logs: lazy(() => import("./Logs")),
  software: lazy(() => import("./SoftwareCatalogue")),
  support: lazy(() => import("./RequestSupport")),
  admin: lazy(() => import("./Admin")),
} as const;

export type AppRoute = keyof typeof routes;

export function getCurrentRoute(): AppRoute {
  const hash = window.location.hash
    .replace("#", "")
    .toLowerCase();

  if (hash in routes) {
    return hash as AppRoute;
  }

  return "dashboard";
}