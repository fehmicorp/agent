import React, { useEffect, useState } from "react";
import { cardMain, anim_1, spanText } from "../../utils/colour";
import { Card, Info, Status } from "../../utils/func";

// Import Wails runtime binds generated during standard go compilation routines
import { DashboardUpdate } from "../../../wailsjs/go/main/App";
import { EventsOn, EventsOff } from "../../../wailsjs/runtime/runtime";

interface MetricCard {
  title: string;
  value: string | number; 
}

interface InfoItem {
  title: string;
  value: string;
}

interface StatusItem {
  title: string;
  value: boolean;
}

// Moving getStatTheme and Card here to keep everything clean and compiled
export function getStatTheme(value: number, isStringFallback = false) {
  if (isStringFallback) {
    return { text: "text-blue-500", shadow: "hover:shadow-blue-500/30" };
  }
  if (value <= 50) return { text: "text-blue-500", shadow: "hover:shadow-blue-500/30" };
  if (value <= 60) return { text: "text-green-500", shadow: "hover:shadow-green-500/30" };
  if (value <= 70) return { text: "text-lime-500", shadow: "hover:shadow-lime-500/30" };
  if (value <= 80) return { text: "text-orange-500", shadow: "hover:shadow-orange-500/30" };
  if (value <= 90) return { text: "text-red-500", shadow: "hover:shadow-red-500/30" };
  
  return { text: "text-red-900 dark:text-red-700", shadow: "hover:shadow-red-900/40" };
}

export function CustomCard({ title, value }: { title: string; value: string | number }) {
  const isNumeric = typeof value === "number";
  const numericValue = isNumeric ? value : 0;
  const colour = getStatTheme(numericValue, !isNumeric);

  return (
    <div className={`${cardMain} rounded-xl items-center px-2 py-3 ${anim_1} hover:shadow-lg ${colour.shadow}`}>
      <p className={spanText}>{title}</p>
      <p className={`text-xl font-semibold mt-2 whitespace-nowrap truncate ${colour.text}`}>
        {isNumeric ? `${value}%` : value}
      </p>
    </div>
  );
}

function MetricCardSkeleton() {
    return (
        <div className={`${cardMain} rounded-xl p-3 animate-pulse`}>
            <div className="h-3 w-24 rounded bg-stone-300 dark:bg-stone-700"/>
            <div className="h-8 w-16 mt-4 rounded bg-stone-300 dark:bg-stone-700"/>
        </div>
    );
}

function InfoSkeleton() {

    return (
        <>
            {Array.from({ length: 5 }).map((_, i) => (
                <div
                    key={i}
                    className="flex justify-between items-center py-2 border-b border-stone-500/10"
                >
                    <div className="h-4 w-28 rounded bg-stone-300 dark:bg-stone-700 animate-pulse"/>
                    <div className="h-4 w-24 rounded bg-stone-300 dark:bg-stone-700 animate-pulse"/>
                </div>
            ))}
        </>
    );
}

function SecuritySkeleton() {

    return (
        <>
            {Array.from({ length: 4 }).map((_, i) => (
                <div
                    key={i}
                    className="flex justify-between items-center py-2 border-b border-stone-500/10"
                >
                    <div className="h-4 w-36 rounded bg-stone-300 dark:bg-stone-700 animate-pulse"/>
                    <div className="h-5 w-14 rounded-full bg-stone-300 dark:bg-stone-700 animate-pulse"/>
                </div>
            ))}
        </>
    );
}

export default function Dashboard(): React.JSX.Element {
  const [stats, setStats] = useState<MetricCard[]>([
    { title: "CPU Usage", value: 0 },
    { title: "RAM Usage", value: 0 },
    { title: "Disk Usage", value: 0 },
    { title: "Network", value: 0 },
  ]);

  const [deviceInfo, setDeviceInfo] = useState<InfoItem[]>([]);
  const [securityStatus, setSecurityStatus] = useState<StatusItem[]>([]);
  const [statsLoading, setStatsLoading] = useState(true);
  const [deviceLoading, setDeviceLoading] = useState(true);
  const [securityLoading, setSecurityLoading] = useState(true);

  useEffect(() => {
    DashboardUpdate()
      .then((info) => {

        setDeviceInfo([
          { title: "Hostname", value: info.hostname },
          { title: "Domain", value: info.domain },
          { title: "User", value: info.user },
          { title: "OS", value: info.os },
          { title: "Agent Version", value: info.agentVersion },
        ]);

        setDeviceLoading(false);

        setSecurityStatus([
          { title: "Windows Defender", value: info.windowsDefender },
          { title: "Firewall", value: info.firewall },
          { title: "TPM", value: info.tpm },
          { title: "BitLocker", value: info.bitLocker },
        ]);

        setSecurityLoading(false);

      })
      .catch(console.error);

    const handleMetricsUpdate = (data: any) => {

        const cleanInt = (val: string) =>
            parseInt(val.replace("%",""),10) || 0;

        setStats([
            { title:"CPU Usage", value:cleanInt(data.cpu)},
            { title:"RAM Usage", value:cleanInt(data.ram)},
            { title:"Disk Usage", value:cleanInt(data.disk)},
            { title:"Network", value:data.network},
        ]);

        setStatsLoading(false);
    }

    EventsOn("metrics_update", handleMetricsUpdate);
    return () => EventsOff("metrics_update");
  }, []);

  return (
    <>
      <div className="grid grid-cols-4 pt-4 gap-4 mb-6">
        {statsLoading
            ? Array.from({ length: 4 }).map((_, i) => (
                <MetricCardSkeleton key={i} />
              ))
            : stats.map((item) => (
                <CustomCard
                    key={item.title}
                    title={item.title}
                    value={item.value}
                />
              ))
        }
      </div>

      <div className="grid grid-cols-2 gap-6">
        <div className={`${cardMain} rounded-xl p-5 border dark:border-stone-500/20 border-stone-500/20`}>
          <h2 className={spanText}>Device Information</h2>
          {deviceLoading
              ? <InfoSkeleton/>
              : deviceInfo.map((item)=>(
                  <Info
                      key={item.title}
                      title={item.title}
                      value={item.value}
                  />
              ))
          }
        </div>
        
        <div className={`${cardMain} rounded-xl p-5 border dark:border-stone-500/20 border-stone-500/20`}>
          <h2 className={spanText}>Security Status</h2>
          {securityLoading
              ? <SecuritySkeleton/>
              : securityStatus.map((item)=>(
                  <Status
                      key={item.title}
                      title={item.title}
                      value={item.value}
                  />
              ))
          }
        </div>
      </div>
    </>
  );
}