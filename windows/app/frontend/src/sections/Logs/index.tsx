import React, { useState } from "react";
import { cardMain, spanText, mutedText } from "../../utils/colour";

// Mock historic log entries matching your system structure (with unique IDs fixed)
const mockLogsData = [
  { id: 1, timestamp: "2026-06-23 16:50:52", type: "risk", category: "Vulnerability", title: "Suspicious outbound port rule detected", details: "Port 44321 showing unusual outbound traffic bursts to untrusted IP." },
  { id: 2, timestamp: "2026-06-23 16:48:10", type: "ok", category: "System Scan", title: "Structural filesystem scan finalized", details: "All 14,820 security sectors parsed cleanly without errors." },
  { id: 3, timestamp: "2026-06-23 16:45:22", type: "risk", category: "Credential", title: "Unencrypted credential cache found", details: "Located exposed credentials package in AppData/Local/Temp." },
  { id: 4, timestamp: "2026-06-23 16:41:05", type: "info", category: "Database", title: "Checking database configurations...", details: "Verified automated connection parameters and schema access tokens." },
  { id: 5, timestamp: "2026-06-23 16:39:12", type: "risk", category: "Vulnerability", title: "Suspicious outbound port rule detected", details: "Port 44321 showing unusual outbound traffic bursts to untrusted IP." },
  { id: 6, timestamp: "2026-06-23 16:37:04", type: "ok", category: "System Scan", title: "Structural filesystem scan finalized", details: "All 14,820 security sectors parsed cleanly without errors." },
  { id: 7, timestamp: "2026-06-23 16:35:50", type: "risk", category: "Credential", title: "Unencrypted credential cache found", details: "Located exposed credentials package in AppData/Local/Temp." },
  { id: 8, timestamp: "2026-06-23 16:31:15", type: "info", category: "Database", title: "Checking database configurations...", details: "Verified automated connection parameters and schema access tokens." },
  { id: 9, timestamp: "2026-06-23 16:29:40", type: "risk", category: "Vulnerability", title: "Suspicious outbound port rule detected", details: "Port 44321 showing unusual outbound traffic bursts to untrusted IP." },
  { id: 10, timestamp: "2026-06-23 16:25:33", type: "ok", category: "System Scan", title: "Structural filesystem scan finalized", details: "All 14,820 security sectors parsed cleanly without errors." },
  { id: 11, timestamp: "2026-06-23 16:20:11", type: "risk", category: "Credential", title: "Unencrypted credential cache found", details: "Located exposed credentials package in AppData/Local/Temp." },
  { id: 12, timestamp: "2026-06-23 16:15:02", type: "info", category: "Database", title: "Checking database configurations...", details: "Verified automated connection parameters and schema access tokens." },
  { id: 13, timestamp: "2026-06-23 16:10:00", type: "ok", category: "Initialization", title: "Core architecture loaded securely", details: "Security kernel initialized with system footprint validation signature." },
];

export default function Logs(): React.JSX.Element {
  const [filter, setFilter] = useState<'all' | 'risk' | 'ok' | 'info'>('all');
  const [search, setSearch] = useState("");

  // Filtering and searching logic
  const filteredLogs = mockLogsData.filter(log => {
    const matchesFilter = filter === 'all' || log.type === filter;
    const matchesSearch = log.title.toLowerCase().includes(search.toLowerCase()) || 
                          log.category.toLowerCase().includes(search.toLowerCase());
    return matchesFilter && matchesSearch;
  });

  return (
    <>
      {/* 1. Header Toolbar & Filters */}
      <div className="rounded-2xl px-4 py-2 flex flex-row md:flex-row gap-4 justify-between items-center">
        <div>
          <h2 className={spanText}>System Security Audit Logs</h2>
          <p className={`${mutedText} text-xs mt-0.5`}>Review historic system modifications, threats, and health check actions.</p>
        </div>

        {/* Filter Badges */}
        <div className="flex bg-stone-900/40 border border-stone-200/10 p-1 rounded-xl text-xs font-medium">
          {(['all', 'risk', 'ok', 'info'] as const).map((type) => (
            <button
              key={type}
              onClick={() => setFilter(type)}
              className={`px-3 py-1.5 rounded-lg capitalize transition-all ${
                filter === type 
                  ? "bg-stone-800 text-white shadow-sm" 
                  : "text-stone-400 hover:text-stone-200"
              }`}
            >
              {type === 'ok' ? 'Success' : type === 'risk' ? 'Risks' : type}
            </button>
          ))}
        </div>
      </div>

      {/* 2. Live Search Utility */}
      <div className="relative p-2">
        <input
          type="text"
          placeholder="Search logs by title, category, or description..."
          value={search}
          onChange={(e) => setSearch(e.target.value)}
          className="w-full bg-stone-950/20 text-stone-200 border border-stone-200/10 rounded-xl px-4 py-2.5 text-xs focus:outline-none focus:border-sky-500/50 transition-colors"
        />
      </div>

      {/* 3. The Logs List Wrapper (Set strictly to h-48 with full x and y axis scrolling) */}
      <div className="h-64 overflow-x-auto overflow-y-auto space-y-1 px-2 font-mono text-xs border border-stone-200/5 bg-stone-900/10 rounded-xl scrollbar-thin scrollbar-thumb-stone-800 scrollbar-track-transparent">
        {/* Inner wrapper ensures wide lines expand horizontally instead of breaking lines */}
        <div className="min-w-max p-1"> 
          {filteredLogs.length > 0 ? (
            filteredLogs.map((log) => {
              let typeDot = "bg-stone-500";
              let textClass = "text-stone-300";
              
              if (log.type === "risk") {
                typeDot = "bg-red-500 shadow-[0_0_8px_rgba(239,68,68,0.5)]";
                textClass = "text-red-400 font-medium";
              } else if (log.type === "ok") {
                typeDot = "bg-emerald-500";
                textClass = "text-emerald-400";
              } else if (log.type === "info") {
                typeDot = "bg-sky-500";
                textClass = "text-sky-400";
              }

              return (
                <div 
                  key={log.id} 
                  className="flex items-center justify-between py-1.5 px-3 rounded-lg hover:bg-stone-500/5 transition-colors border-b border-stone-200/5 gap-8"
                >
                  {/* Left section: Status Dot + Time + Category + Title */}
                  <div className="flex items-center gap-3">
                    <span className={`h-2 w-2 rounded-full flex-shrink-0 ${typeDot}`} />
                    
                    <span className="text-stone-500 whitespace-nowrap text-[11px]">
                      [{log.timestamp.split(" ")[1]}]
                    </span>

                    <span className="text-stone-500 uppercase text-[10px] tracking-wider border border-stone-200/10 px-1 rounded bg-stone-900/20 whitespace-nowrap">
                      {log.category}
                    </span>

                    <p className={`${textClass} whitespace-nowrap`}>
                      {log.title} <span className="text-stone-500 opacity-60 font-light text-[11px] ml-1">— {log.details}</span>
                    </p>
                  </div>

                  {/* Right section: Status flags */}
                  <div className="flex-shrink-0 text-[10px] font-bold tracking-wider pl-4">
                    {log.type === "risk" && <span className="text-red-400/90 bg-red-500/10 px-1.5 py-0.5 rounded border border-red-500/20">ALERT</span>}
                    {log.type === "ok" && <span className="text-emerald-400/90 opacity-60">SUCCESS</span>}
                    {log.type === "info" && <span className="text-sky-400/90 opacity-60">INFO</span>}
                  </div>
                </div>
              );
            })
          ) : (
            <div className="py-8 text-center text-stone-500 text-xs">
              No system log records matched your query.
            </div>
          )}
        </div>
      </div>
    </>
  );
}