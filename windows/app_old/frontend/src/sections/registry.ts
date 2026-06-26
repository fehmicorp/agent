import { lazy } from "react";

export const routes = {
  dashboard: lazy(() => import("./Dashboard")),
  scan: lazy(() => import("./Scan")),
  logs: lazy(() => import("./Logs")),
  software: lazy(() => import("./SoftwareCatalogue")),
  support: lazy(() => import("./RequestSupport")),
  agent: lazy(() => import("./Agent")),
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