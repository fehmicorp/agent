import React, { useEffect, useState } from "react";
import { bgMain, cardMain, headingText, mutedText, spanText } from "../../utils/colour";
import App from "../data";
import { scan } from "../data/func";

// This array is fine outside the component since it's just static data
const scanPipeline = [
  { text: "[INIT] Initializing security matrices...", type: "init" as const },
  { text: "[OK] Core architecture loaded securely.", type: "ok" as const },
  { text: "[RUN] Checking system configurations...", type: "run" as const },
  { text: "[SCAN] C:/Windows/System32/drivers/etc/hosts", type: "scan" as const },
  { text: "[RISK] Found unencrypted credential cache!", type: "risk" as const },
  { text: "[SCAN] C:/Users/Admin/AppData/Local/Temp/cache_042.tmp", type: "scan" as const },
  { text: "[SCAN] C:/Program_Files/App/config.json", type: "scan" as const },
  { text: "[RISK] Suspicious outbound outbound port rule detected.", type: "risk" as const },
  { text: "[SCAN] C:/Users/Admin/Documents/invoice_export.csv", type: "scan" as const },
  { text: "[OK] Structural filesystem scan finalized.", type: "ok" as const },
  { text: "[RISK] Found unencrypted credential cache!", type: "risk" as const },
  { text: "[SCAN] C:/Users/Admin/AppData/Local/Temp/cache_042.tmp", type: "scan" as const },
  { text: "[SCAN] C:/Program_Files/App/config.json", type: "scan" as const },
  { text: "[RISK] Suspicious outbound outbound port rule detected.", type: "risk" as const },
  { text: "[SCAN] C:/Users/Admin/Documents/invoice_export.csv", type: "scan" as const },
  { text: "[OK] Structural filesystem scan finalized.", type: "ok" as const },
  { text: "[RISK] Found unencrypted credential cache!", type: "risk" as const },
  { text: "[SCAN] C:/Users/Admin/AppData/Local/Temp/cache_042.tmp", type: "scan" as const },
  { text: "[SCAN] C:/Program_Files/App/config.json", type: "scan" as const },
  { text: "[RISK] Suspicious outbound outbound port rule detected.", type: "risk" as const },
  { text: "[SCAN] C:/Users/Admin/Documents/invoice_export.csv", type: "scan" as const },
  { text: "[OK] Structural filesystem scan finalized.", type: "ok" as const },
];

export default function Scan(): React.JSX.Element {
  // ALL HOOKS MUST LIVE EXACTLY HERE INSIDE THE FUNCTION
  const [scanning, setScanning] = useState(false);
  const [progress, setProgress] = useState(45);
  const [scanTitle, setScanTitle] = useState("Inventory Collection");
  const [logs, setLogs] = useState<{ text: string; type: 'init' | 'run' | 'scan' | 'ok' | 'risk' }[]>([]);

  useEffect(() => {
    let intervalId: ReturnType<typeof setInterval> | undefined;

    if (scanning) {
      setLogs([scanPipeline[0]]); 
      let currentIndex = 1;
      setProgress(5);

      intervalId = setInterval(() => {
        if (currentIndex < scanPipeline.length) {
          const nextLog = scanPipeline[currentIndex];          
          setLogs((prev) => [...prev, nextLog]);
          setProgress(Math.min(Math.floor((currentIndex / (scanPipeline.length - 1)) * 100), 100));          
          currentIndex++;
        }        
        // Clear the interval right away if we've pushed the final item
        if (currentIndex >= scanPipeline.length) {
          clearInterval(intervalId);
        }
      }, 1000);
    } else {
      setLogs([]);
      setProgress(0);
    }

    return () => {
      if (intervalId) clearInterval(intervalId);
    };
  }, [scanning]);

  const handleScan = (id: number) => {
    setScanning(true);
    const clickedScan = scan.find(item => item.id === id);
    if (clickedScan) {
      setScanTitle(clickedScan.title);
    }
  };

  return (
    <div className="pt-4">
      {/* 1. Scan Selection Buttons */}
      {!scanning && (
        <div className="grid grid-cols-2 gap-4">
          {scan.map((item) => (
            <ScanButton
              key={item.id}
              id={item.id}
              title={item.title}
              desc={item.desc}
              onScan={handleScan}
            />
          ))}
        </div>
      )}

      {/* 2. Active Scan Terminal View */}
      {scanning && (
        <div className={`${cardMain} rounded-2xl p-4 mt-4`}>
          <div className="flex justify-between items-center">
            <h2 className={spanText}>Scan Progress</h2>
            <button 
              onClick={() => setScanning(false)}
              className="px-3 py-1 text-xs font-medium text-red-400 border border-red-500/30 bg-red-500/10 rounded-lg transition-all hover:bg-red-500/20 active:scale-95"
            >
              Cancel Scan
            </button>
          </div>

          <div className="mt-4">
            <div className="w-full h-3 rounded-full bg-stone-300/20 overflow-hidden">
              <div 
                className="h-full bg-gradient-to-r from-sky-500 to-cyan-400 rounded-full animate-pulse transition-all duration-300" 
                style={{ width: `${progress}%` }} 
              />
            </div>
            <p className={`${mutedText} text-xs mt-2`}>
              {scanTitle} • {progress}%
            </p>
          </div>

          <hr className="border-stone-700/10 dark:border-stone-300/10 my-4" />

          <div>
            <h3 className={`${spanText} text-xs tracking-wider opacity-60 mb-2`}>
              Scanning Log
            </h3>
            <div className="bg-stone-300 dark:bg-stone-900/40 rounded-xl p-3 font-mono text-[11px] h-52 overflow-y-auto space-y-1.5 flex flex-col">
              <div className="space-y-1.5">
                {logs.map((log, index) => {
                  let textClass = "dark:text-stone-400 text-stone-900 opacity-80";
                  if (log.type === "init" || log.type === "run") textClass = "dark:text-sky-400 text-sky-700";
                  if (log.type === "ok") textClass = "text-emerald-700 dark:text-emerald-400";
                  if (log.type === "risk") textClass = "dark:text-red-400 text-red-600 font-bold animate-pulse";
                  if (log.type === "scan" && index === logs.length - 1) textClass = "dark:text-amber-200 text-amber-800 animate-pulse";

                  return (
                    <div key={index} className={`${textClass} flex justify-between`}>
                      <span>{log.text}</span>
                      {log.type === "ok" && <span className="dark:opacity-60 opacity-80 text-xs dark:text-emerald-300 px-3 text-emerald-700">SUCCESS</span>}
                      {log.type === "risk" && <span className="text-xs dark:bg-red-400/30 bg-red-600/30 px-2 rounded text-red-700 dark:text-red-300">ALERT</span>}
                    </div>
                  );
                })}
              </div>
            </div>
          </div>
        </div>
      )}      

      {/* 3. Global Scan Summary */}
      {!scanning && (
        <div className={`${cardMain} rounded-2xl p-4 mt-4`} >
          <div className="flex justify-between items-end">
            <h2 className={spanText}> Last Scan Summary </h2>
            <span className="text-[12px]">Date: Tuesday June 23, 2026 ( 16:50:52 )</span>
          </div>
          <div className="grid grid-cols-3 gap-4 mt-3">
            <Metric title="Software" value="154" />
            <Metric title="Vulnerabilities" value="3" />
            <Metric title="Missing Patches" value="12" />
          </div>
        </div>
      )}
    </div>
  );
}

function Metric({
  title,
  value,
}: {
  title: string;
  value: string;
}) {
  return (
    <div className="text-center">
      <p className="text-2xl font-bold">{value}</p>
      <p className="text-xs opacity-70">{title}</p>
    </div>
  );
}

function ScanButton({
  id,
  title,
  desc,
  onScan,
}: {
  id: number;
  title: string;
  desc: string;
  onScan: (id: number) => void;
}) {
  return (
    <button onClick={() => onScan(id)} className={`
      ${cardMain} group relative rounded-2xl p-4 text-left overflow-hidden transition-all duration-500 ease-out hover:scale-[1.02]
      hover:-translate-y-1 hover:shadow-[0_12px_40px_rgba(59,130,246,0.15)]`}>
      <div className="absolute inset-0 opacity-0 transition-opacity duration-500 group-hover:opacity-100 bg-gradient-to-br from-slate-100/50 dark:from-slate-900/50 dark:via-blue-400/35 via-blue-400/35 dark:to-violet-500/50 to-violet-500/50"/>
      <div className="relative z-10">
        <h3 className="font-semibold">{title}</h3>
        <p className="text-xs opacity-70 mt-1">{desc}</p>
      </div>
    </button>
  );
}