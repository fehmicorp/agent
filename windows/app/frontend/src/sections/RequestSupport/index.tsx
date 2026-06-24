import React, { useState } from "react";
import { support, ticketTimeline as initialTimeline } from "../data/dummy";
import { SupportTicket } from "../data";
import Raise from "./raise";
import Tickets from "./tickets";

export default function Support(): React.JSX.Element {
  const [tickets, setTickets] = useState<SupportTicket[]>(support.tickets as any);
  const [selectedTicketId, setSelectedTicketId] = useState<string | null>(null);
  const [timelineLogs, setTimelineLogs] = useState(initialTimeline);
  const [isNewRequest, setIsNewRequest] = useState(false);

  const handleAddNewTicket = (newTicket: SupportTicket) => {
    setTickets((prevTickets) => [newTicket, ...prevTickets]);
    setIsNewRequest(false); // Drop back to main workspace list channel view after submission
  };

  const handleAddTimelineComment = (tktId: string, text: string, files: File[]) => {
    const newCommentNode = {
      id: Math.floor(1000 + Math.random() * 9000),
      tktId: tktId,
      author: {
        name: "Current System Operator",
        id: 999
      },
      created_at: new Date().toISOString().replace('T', ' ').substring(0, 16),
      description: text + (files.length > 0 ? ` [📎 Embedded Attachments: ${files.length} asset logs]` : "")
    };

    setTimelineLogs(prev => [newCommentNode, ...prev]);
  };

  return (
    <div className="space-y-4 pt-2 font-mono">
      {/* Top Navigation Control Dashboard Header */}
      <div className="flex items-center justify-between border-b border-stone-400/10 pb-4">
        <div>
          <h1 className="text-sm font-bold text-stone-900 dark:text-stone-100 uppercase tracking-wider">
            Agent Support Desk
          </h1>
          <p className="text-[10px] text-stone-400">Manage infrastructure logs and incident tracking reports.</p>
        </div>
        
        {/* Toggle Actions Panel */}
        {!isNewRequest ? (
          <button
            type="button"
            onClick={() => {
              setSelectedTicketId(null); // Clean out background detail states
              setIsNewRequest(true); // Isolate view layout exclusively for Raise
            }}
            className="px-3 py-1.5 bg-stone-900 text-stone-100 dark:bg-stone-100 dark:text-stone-950 font-bold rounded-xl text-[10px] uppercase tracking-wide hover:opacity-90 transition-all border border-stone-400/20"
          >
            + Create new request
          </button>
        ) : (
          <button
            type="button"
            onClick={() => setIsNewRequest(false)}
            className="px-3 py-1.5 bg-stone-200 dark:bg-stone-800 text-stone-700 dark:text-stone-300 font-bold rounded-xl text-[10px] uppercase tracking-wide hover:opacity-90 transition-all border border-stone-400/10"
          >
            ← Cancel & View Logs
          </button>
        )}
      </div>

      {/* Dynamic View Distribution Switch Router */}
      <div className="w-full">
        {isNewRequest ? (
          /* Isolated full-width container workspace layout for Raise component */
          <div className="w-full max-w-2xl mx-auto animate-fade-in">
            <Raise onTicketCreated={handleAddNewTicket} />
          </div>
        ) : (
          /* Ticket Manifest Tracking Views Row Layer */
          <div className="w-full">
            <Tickets 
              tickets={tickets} 
              selectedTicketId={selectedTicketId}
              onSelectTicket={setSelectedTicketId}
              onAddTimelineComment={handleAddTimelineComment}
              timelineLogs={timelineLogs}
            />
          </div>
        )}
      </div>
    </div>
  );
}