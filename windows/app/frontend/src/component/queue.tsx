import React, { useMemo, useState } from "react";
import { X, CheckCircle2, AlertCircle, Loader2, Clock, Terminal } from "lucide-react";

export interface QueueJob {
  id: string;
  name: string;
  type: "installation" | "update" | "scan" | "backup";
  status: "running" | "queued" | "success" | "failed";
  progress: number; // Percent 0-100
  timestamp: string;
  details?: string;
}

interface QueueProps {
  isOpen: boolean;
  onClose: () => void;
  jobs: QueueJob[];
  onClearHistory: () => void;
}

export default function Queue({
  isOpen,
  onClose,
  jobs,
  onClearHistory,
}: QueueProps): React.JSX.Element {
  const [activeTab, setActiveTab] = useState<"all" | "active" | "history">("all");

  // Pipeline Filter logic for Jobs Manifest
  const filteredJobs = useMemo(() => {
    if (activeTab === "active") {
      return jobs.filter((j) => j.status === "running" || j.status === "queued");
    }
    if (activeTab === "history") {
      return jobs.filter((j) => j.status === "success" || j.status === "failed");
    }
    return jobs;
  }, [jobs, activeTab]);

  const getStatusIcon = (status: QueueJob["status"]) => {
    switch (status) {
      case "running":
        return <Loader2 size={13} className="animate-spin text-amber-500" />;
      case "queued":
        return <Clock size={13} className="text-stone-400" />;
      case "success":
        return <CheckCircle2 size={13} className="text-emerald-500" />;
      case "failed":
        return <AlertCircle size={13} className="text-red-500" />;
    }
  };

  const getStatusBadgeStyle = (status: QueueJob["status"]) => {
    switch (status) {
      case "running":
        return "bg-amber-500/10 text-amber-600 dark:text-amber-400 border-amber-500/20";
      case "queued":
        return "bg-stone-500/10 text-stone-500 dark:text-stone-400 border-stone-500/20";
      case "success":
        return "bg-emerald-500/10 text-emerald-600 dark:text-emerald-400 border-emerald-500/20";
      case "failed":
        return "bg-red-500/10 text-red-600 dark:text-red-400 border-red-500/20";
    }
  };

  return (
    <>
      {/* Backdrop overlay blur sheet element */}
      <div
        onClick={onClose}
        className={`fixed inset-0 bg-stone-950/40 backdrop-blur-sm z-50 transition-opacity duration-300 ${
          isOpen ? "opacity-100 pointer-events-auto" : "opacity-0 pointer-events-none"
        }`}
      />

      {/* Right Drawer Layout matching 75% display standard bounds */}
      <div
        className={`fixed top-0 right-0 h-full w-3/4 max-w-lg bg-stone-50 dark:bg-stone-900 border-l border-stone-200 dark:border-stone-800 shadow-2xl z-50 flex flex-col justify-between transform transition-transform duration-300 ease-out font-mono ${
          isOpen ? "translate-x-0" : "translate-x-full"
        }`}
      >
        {/* Header Block Section */}
        <div className="p-4 border-b border-stone-200 dark:border-stone-800">
          <div className="flex items-center justify-between mb-3">
            <div>
              <h2 className="text-xs font-bold uppercase tracking-wider text-stone-900 dark:text-stone-100 flex items-center gap-1.5">
                <Terminal size={14} className="text-stone-500" /> Task Engine Workspace
              </h2>
              <p className="text-[10px] text-stone-500">Infrastructure operational runtime stacks</p>
            </div>
            <button
              type="button"
              onClick={onClose}
              className="p-1 rounded-md text-stone-500 hover:bg-stone-200 dark:hover:bg-stone-800 transition-colors"
            >
              <X size={16} />
            </button>
          </div>

          {/* Segmented Controller Navigation Menu Filters Matrix */}
          <div className="flex gap-1 bg-stone-200/60 dark:bg-stone-950/40 p-0.5 rounded-lg text-[10px]">
            {(["all", "active", "history"] as const).map((tab) => (
              <button
                key={tab}
                type="button"
                onClick={() => setActiveTab(tab)}
                className={`flex-1 py-1 rounded-md text-center uppercase tracking-wide transition-all ${
                  activeTab === tab
                    ? "bg-white dark:bg-stone-800 text-stone-950 dark:text-stone-50 shadow-sm font-bold"
                    : "text-stone-500 hover:text-stone-800 dark:hover:text-stone-300"
                }`}
              >
                {tab}
              </button>
            ))}
          </div>
        </div>

        {/* Dynamic List Content Engine Render Matrix */}
        <div className="flex-1 p-4 overflow-y-auto space-y-3">
          {filteredJobs.length === 0 ? (
            <div className="h-full flex flex-col items-center justify-center text-center opacity-60 py-12">
              <span className="text-sm">🪐</span>
              <p className="text-[10px] text-stone-500 mt-1">No execution parameters matching view indices.</p>
            </div>
          ) : (
            filteredJobs.map((job) => (
              <div
                key={job.id}
                className="p-3 rounded-xl bg-white dark:bg-stone-950 border border-stone-200/60 dark:border-stone-800/60 shadow-sm space-y-2.5"
              >
                {/* Meta details node */}
                <div className="flex items-start justify-between gap-2">
                  <div className="space-y-0.5 max-w-[70%]">
                    <h4 className="text-[11px] font-bold text-stone-900 dark:text-stone-100 truncate">
                      {job.name}
                    </h4>
                    <p className="text-[9px] text-stone-400 uppercase tracking-tight">
                      ID: {job.id} • {job.timestamp}
                    </p>
                  </div>
                  <span
                    className={`px-1.5 py-0.5 rounded text-[8px] font-bold border uppercase tracking-wider flex items-center gap-1 ${getStatusBadgeStyle(
                      job.status
                    )}`}
                  >
                    {getStatusIcon(job.status)}
                    {job.status}
                  </span>
                </div>

                {/* Conditional progress monitor line output parameters */}
                {job.status === "running" && (
                  <div className="space-y-1">
                    <div className="flex justify-between items-center text-[9px] text-stone-400">
                      <span>Processing cluster dependencies...</span>
                      <span className="font-bold text-stone-600 dark:text-stone-300">{job.progress}%</span>
                    </div>
                    <div className="w-full h-1 bg-stone-100 dark:bg-stone-800 rounded-full overflow-hidden">
                      <div
                        className="h-full bg-amber-500 transition-all duration-300 ease-out"
                        style={{ width: `${job.progress}%` }}
                      />
                    </div>
                  </div>
                )}

                {/* Subtext execution errors string logging */}
                {job.details && (
                  <p className="text-[9px] bg-stone-100/60 dark:bg-stone-900/60 p-1.5 rounded border border-stone-200/40 dark:border-stone-800/40 text-stone-500 max-h-16 overflow-y-auto leading-relaxed">
                    $ {job.details}
                  </p>
                )}
              </div>
            ))
          )}
        </div>

        {/* Clear Execution Logs Base Controls Deck */}
        <div className="p-4 bg-stone-100/80 dark:bg-stone-900/80 backdrop-blur border-t border-stone-200 dark:border-stone-800">
          <button
            type="button"
            disabled={!jobs.some((j) => j.status === "success" || j.status === "failed")}
            onClick={onClearHistory}
            className="w-full flex items-center justify-center gap-1.5 py-2 border border-stone-300 dark:border-stone-700 hover:bg-stone-200/50 dark:hover:bg-stone-800/50 text-stone-700 dark:text-stone-300 font-bold rounded-xl text-[10px] uppercase tracking-wider transition-all disabled:opacity-40 disabled:cursor-not-allowed"
          >
            Clear Historical Logs
          </button>
        </div>
      </div>
    </>
  );
}