import React, { useMemo } from "react";
import { SupportTicket } from "../data";

interface TimelineProps {
  selectedTicket: SupportTicket | null;
  timelineLogs: any[];
}

export default function Timeline({ selectedTicket, timelineLogs }: TimelineProps): React.JSX.Element {
  const filteredTimeline = useMemo(() => {
    if (!selectedTicket) return [];
    return timelineLogs
      .filter((item) => item.tktId === selectedTicket.tktId)
      .sort((a, b) => new Date(b.created_at).getTime() - new Date(a.created_at).getTime());
  }, [selectedTicket, timelineLogs]);

  return (
    <div className="space-y-4 font-mono text-xs">
      <div className="bg-stone-100/50 dark:bg-stone-900/40 border border-stone-400/10 rounded-xl p-3 space-y-1">
        <div className="text-[10px] text-stone-400 font-bold uppercase tracking-wide">Root Ticket Description</div>
        <p className="text-[11px] font-sans leading-relaxed text-stone-700 dark:text-stone-300">
          {selectedTicket?.description}
        </p>
      </div>

      <div className="space-y-4 relative before:absolute before:top-2 before:bottom-2 before:left-2.5 before:w-px before:bg-stone-300 dark:before:bg-stone-800">
        {filteredTimeline.length === 0 ? (
          <div className="pl-7 text-[10px] text-stone-500 italic py-1">
            No append timeline logs posted against this container manifest node yet.
          </div>
        ) : (
          filteredTimeline.map((update) => (
            <div key={update.id} className="relative pl-7 space-y-1 group">
              <div className="absolute left-1.5 top-1.5 h-2 w-2 rounded-full bg-stone-400 dark:bg-stone-600 group-hover:bg-blue-500 transition-colors" />
              <div className="flex items-center justify-between text-[10px] text-stone-500">
                <span className="font-bold text-stone-700 dark:text-stone-300">{update.author.name}</span>
                <span className="text-[9px]">{update.created_at}</span>
              </div>
              <p className="text-[11px] text-stone-600 dark:text-stone-400 font-sans leading-relaxed bg-stone-100/30 dark:bg-stone-950/20 px-2.5 py-1.5 rounded-lg border border-stone-400/5">
                {update.description}
              </p>
            </div>
          ))
        )}
      </div>
    </div>
  );
}