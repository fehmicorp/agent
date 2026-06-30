import React, {
  Suspense,
  useEffect,
  useState,
} from "react";

import {
  routes,
  AppRoute,
  getCurrentRoute,
} from "./sections/registry";
import data from "./sections/data";

import { useSystemTheme } from "./utils/theme";
import Navbar from "./component/Navbar";
import Header from "./component/Header";
import Notification from "./component/notification";
import Queue, { QueueJob } from "./component/queue";

const INITIAL_NOTIFICATIONS = [
  { id: 1, text: "Security baseline audit validation trace successfully completed.", time: "2 mins ago" },
  { id: 2, text: "High memory paging pool usage flagged on Node-04.", time: "1 hr ago" },
  { id: 3, text: "System diagnostic patch package automated routine deployed.", time: "5 hrs ago" },
];

const INITIAL_JOBS: QueueJob[] = [
  { id: "JOB-992", name: "Security Patch Engine Upgrade v2.4.1", type: "update", status: "running", timestamp: "Just now" },
  { id: "JOB-771", name: "Host Agent Node Core Vulnerability Scan", type: "scan", status: "queued", timestamp: "1 min ago" },
  { id: "JOB-402", name: "Deploy @fehmicorp/middleware Bundle to npm Registry", type: "installation", status: "success", timestamp: "12 mins ago" },
  { id: "JOB-109", name: "Synology NAS Target LUN Multipath Sync", type: "backup", status: "failed", timestamp: "1 hr ago", details: "iSCSI connection reset by host fcupepgdb01 pipeline error." },
];

export default function App() {
  useSystemTheme();
  const [route, setRoute] = useState<AppRoute>(getCurrentRoute());
  const [isNotifOpen, setIsNotifOpen] = useState(false);
  const [isQueueOpen, setIsQueueOpen] = useState(false);

  const [notifications, setNotifications] = useState(INITIAL_NOTIFICATIONS);
  const [jobs, setJobs] = useState<QueueJob[]>(INITIAL_JOBS);

  useEffect(() => {
    const handler = () => {
      setRoute(getCurrentRoute());
    };

    window.addEventListener("hashchange", handler);
    handler();

    return () => {
      window.removeEventListener("hashchange", handler);
    };
  }, []);

  const Page = routes[route];
  const runningJobsCount = jobs.filter((j) => j.status === "running").length;

  return (
    <div className="h-screen w-screen bg-white text-slate-900 dark:bg-slate-950 dark:text-white relative overflow-hidden font-mono flex flex-col">
      <Suspense fallback={
        <div className="p-4">
          Getting things ready...
        </div>
      }>
        <div className="bg-stone-100 dark:bg-stone-950 flex flex-col flex-1 h-full min-h-0 overflow-hidden">
          {
            route === "agent" 
              ? <></> 
              : <Header 
                  title={data.name} 
                  onOpenNotifications={() => setIsNotifOpen(true)}
                  onOpenQueue={() => setIsQueueOpen(true)}
                />
          }
          
          {/* Changed 'overflow-hidden' to 'overflow-y-auto' to allow pages to scroll inside viewport safely */}
          <div className="flex-1 px-2 pb-24 overflow-y-auto min-h-0 [&::-webkit-scrollbar]:w-2 [&::-webkit-scrollbar-track]:bg-transparent [&::-webkit-scrollbar-thumb]:bg-stone-500/30 [&::-webkit-scrollbar-thumb]:rounded-full hover:[&::-webkit-scrollbar-thumb]:bg-stone-500/60">
            <Page />
          </div>
        </div>
        
        <Navbar />
        
        <Notification
          isOpen={isNotifOpen}
          onClose={() => setIsNotifOpen(false)}
          notifications={notifications}
          onClearAll={() => setNotifications([])}
        />
        
        <Queue
          isOpen={isQueueOpen}
          onClose={() => setIsQueueOpen(false)}
          jobs={jobs}
          onClearHistory={() =>
            setJobs((prev) => prev.filter((j) => j.status === "running" || j.status === "queued"))
          }
        />
      </Suspense>
    </div>
  );
}