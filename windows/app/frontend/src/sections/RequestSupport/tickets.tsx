import React, { useState, useMemo, useRef } from "react";
import { cardMain } from "../../utils/colour";
import { support } from "../data/dummy";
import { SupportTicket } from "../data";
import Timeline from "./timeline";

interface TicketsProps {
  tickets: SupportTicket[];
  selectedTicketId: string | null;
  onSelectTicket: (tktId: string | null) => void;
  onAddTimelineComment: (tktId: string, text: string, files: File[]) => void;
  timelineLogs: any[];
}

export default function Tickets({ 
  tickets, 
  selectedTicketId, 
  onSelectTicket,
  onAddTimelineComment,
  timelineLogs
}: TicketsProps): React.JSX.Element {
  const [filterStatus, setFilterStatus] = useState<string>("All");
  const [commentText, setCommentText] = useState("");
  const [commentFiles, setCommentFiles] = useState<File[]>([]);
  const [isCommenting, setIsCommenting] = useState(false);
  
  // --- PAGINATION STATES ---
  const [currentPage, setCurrentPage] = useState<number>(1);
  const [pageSize, setPageSize] = useState<number>(10); // Default items limit tracking array
  
  const fileInputRef = useRef<HTMLInputElement>(null);

  const limits = useMemo(() => {
    return {
      minDesc: support.limit?.minDesc ?? 10,
      maxDesc: support.limit?.maxDesc ?? 200,
      fileConfig: support.limit?.file ?? null
    };
  }, []);

  const acceptedFormatsString = useMemo(() => {
    if (!limits.fileConfig?.type || limits.fileConfig.type.length === 0) return "image/*";
    return limits.fileConfig.type.map(ext => `.${ext.toLowerCase()}`).join(",");
  }, [limits]);

  // Filter Pipeline Array Logic Resetting Page Baseline Counters on Change
  const filteredTickets = useMemo(() => {
    setCurrentPage(1); // Force return to page 1 during filter sweeps
    if (filterStatus === "All") return tickets;
    return tickets.filter((ticket) => ticket.status === filterStatus);
  }, [tickets, filterStatus]);

  // Page Processing Slices Window Pipeline Calculations
  const totalPages = Math.ceil(filteredTickets.length / pageSize) || 1;
  
  const paginatedTickets = useMemo(() => {
    const startIndex = (currentPage - 1) * pageSize;
    return filteredTickets.slice(startIndex, startIndex + pageSize);
  }, [filteredTickets, currentPage, pageSize]);

  const activeTicket = useMemo(() => {
    if (!selectedTicketId) return null;
    return tickets.find(t => t.tktId === selectedTicketId) || null;
  }, [tickets, selectedTicketId]);

  const isCommentValid = useMemo(() => {
    const commentLength = commentText.trim().length;
    if (commentLength < limits.minDesc || commentLength > limits.maxDesc) return false;
    if (limits.fileConfig) {
      const isMandatory = (limits.fileConfig as any).manatory ?? limits.fileConfig.mandatory ?? false;
      if (isMandatory && commentFiles.length === 0) return false;
    }
    return true;
  }, [commentText, commentFiles, limits]);

  const getStatusStyles = (status: string) => {
    switch (status) {
      case "Open": return "bg-blue-500/10 text-blue-600 dark:text-blue-400 border-blue-500/20";
      case "In Progress": return "bg-amber-500/10 text-amber-600 dark:text-amber-400 border-amber-500/20";
      case "Resolved": return "bg-emerald-500/10 text-emerald-600 dark:text-emerald-400 border-emerald-500/20";
      default: return "bg-stone-500/10 text-stone-600 dark:text-stone-400 border-stone-500/20";
    }
  };

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files.length > 0) {
      setCommentFiles(prev => [...prev, ...Array.from(e.target.files!)].slice(0, limits.fileConfig?.max ?? 10));
    }
    if (fileInputRef.current) fileInputRef.current.value = "";
  };

  const handleSubmitComment = (e: React.FormEvent) => {
    e.preventDefault();
    if (!isCommentValid || !selectedTicketId) return;
    onAddTimelineComment(selectedTicketId, commentText, commentFiles);
    setCommentText("");
    setCommentFiles([]);
    setIsCommenting(false);
  };

  // --- SINGLE MODE INNER TICKET TIMELINE VIEW MAPPING ---
  if (activeTicket) {
    return (
      <div className="space-y-4 animate-fade-in">
        <div className="flex items-center gap-4 border-b border-stone-400/10 pb-2">
          <button
            type="button"
            onClick={() => { onSelectTicket(null); setIsCommenting(false); }}
            className="px-3 py-1.5 bg-stone-200 dark:bg-stone-800 text-stone-700 dark:text-stone-300 rounded-lg text-[10px] font-mono border border-stone-400/10 hover:opacity-80 transition-opacity whitespace-nowrap"
          >
            ← Back to Log
          </button>
          <div className="flex items-center gap-2 font-mono text-xs truncate">
            <span className={`text-[10px] font-bold px-1.5 py-0.5 rounded border ${getStatusStyles(activeTicket.status)}`}>
              {activeTicket.tktId}
            </span>
            <span className="font-bold text-stone-900 dark:text-stone-100 truncate">{activeTicket.subject}</span>
          </div>
        </div>

        <Timeline selectedTicket={activeTicket} timelineLogs={timelineLogs} />

        {/* Comment Injection Input Block */}
        <div className={`${cardMain} rounded-2xl p-4 border border-stone-200/5 space-y-3 font-mono text-xs`}>
          <div className="flex items-center justify-between">
            <h4 className="text-[10px] font-bold text-stone-400 uppercase tracking-wider">Post Comment Update</h4>
            {!isCommenting && (
              <button
                type="button"
                onClick={() => setIsCommenting(true)}
                className="px-2.5 py-1 bg-stone-800 text-white dark:bg-stone-200 dark:text-stone-950 font-bold rounded-md text-[10px] hover:opacity-90"
              >
                + Add Comment
              </button>
            )}
          </div>

          {isCommenting && (
            <form onSubmit={handleSubmitComment} className="space-y-3 pt-2 border-t border-stone-400/10">
              <textarea
                required
                rows={3}
                placeholder="Type structural runtime annotations..."
                value={commentText}
                onChange={(e) => setCommentText(e.target.value)}
                className="w-full bg-stone-100/80 dark:bg-stone-950/90 text-stone-700 dark:text-stone-300 border border-stone-400/20 dark:border-stone-200/10 rounded-xl px-3 py-2 text-xs"
              />
              <div className="flex gap-2">
                <button type="button" onClick={() => setIsCommenting(false)} className="px-3 py-1.5 bg-stone-200 dark:bg-stone-800 rounded-lg text-[10px]">Cancel</button>
                <button type="submit" disabled={!isCommentValid} className="px-4 py-1.5 bg-stone-800 text-stone-100 dark:bg-stone-100 dark:text-stone-950 rounded-lg text-[10px] disabled:opacity-40">Commit Update</button>
              </div>
            </form>
          )}
        </div>
      </div>
    );
  }

  // --- RUNTIME SYSTEM LIST OVERVIEW GRID RENDER ---
  return (
    <div className="w-full space-y-4 animate-fade-in">
      {/* Sticky Segment Configuration Navigation Controls Header */}
      <div className="sticky top-0 z-30 py-3 bg-stone-100/90 dark:bg-stone-950/90 backdrop-blur-md border-b border-stone-200/50 dark:border-stone-800/50 flex flex-wrap gap-3 items-center justify-between transition-all">
        <div className="flex gap-1 font-mono text-[10px]">
          {["All", "Open", "In Progress", "Resolved"].map((status) => (
            <button
              key={status}
              type="button"
              onClick={() => setFilterStatus(status)}
              className={`px-2.5 py-1 rounded-md border transition-all ${
                filterStatus === status
                  ? "bg-stone-800 text-white border-transparent dark:bg-stone-200 dark:text-stone-950 font-semibold"
                  : "bg-transparent text-stone-500 border-stone-400/20 hover:bg-stone-400/10"
              }`}
            >
              {status}
            </button>
          ))}
        </div>

        {/* Dynamic Items Limit Dropdown Matrix */}
        <div className="flex items-center gap-1.5 text-[10px] font-mono text-stone-500">
          <span>Display Rows:</span>
          <select
            value={pageSize}
            onChange={(e) => {
              setPageSize(Number(e.target.value));
              setCurrentPage(1); // Jump back to base indexing array matrix rules
            }}
            className="bg-stone-200 dark:bg-stone-800 text-stone-700 dark:text-stone-300 border border-stone-400/20 rounded-md px-1.5 py-0.5 focus:outline-none"
          >
            {[10, 20, 50, 100].map(size => (
              <option key={size} value={size}>{size}</option>
            ))}
          </select>
        </div>
      </div>

      {/* Ticket Rows Output Map */}
      {paginatedTickets.length === 0 ? (
        <div className={`${cardMain} rounded-2xl p-12 text-center border border-dashed dark:border-stone-800 border-stone-300`}>
          <p className="text-xs font-mono text-stone-500">📭 Empty dataset window. No active indices matching parameter masks.</p>
        </div>
      ) : (
        <div className="space-y-3">
          {paginatedTickets.map((ticket) => {
            const catTitle = typeof ticket.category === "object" && ticket.category !== null
              ? (ticket.category as any).title : String(ticket.category);

            return (
              <div 
                key={ticket.id} 
                onClick={() => onSelectTicket(ticket.tktId)}
                className={`${cardMain} rounded-2xl p-4 border cursor-pointer border-stone-200/5 hover:border-stone-400/20 dark:hover:border-stone-200/10 transition-all duration-200 flex items-center justify-between`}
              >
                <div className="space-y-0.5 truncate max-w-[80%]">
                  <div className="flex items-center gap-2 truncate">
                    <span className={`text-[10px] font-bold px-1.5 py-0.2 rounded border ${getStatusStyles(ticket.status)}`}>
                      {ticket.tktId}
                    </span>
                    <h3 className="text-xs font-bold text-stone-900 dark:text-stone-100 truncate">{ticket.subject}</h3>
                  </div>
                  <p className="text-[10px] text-stone-500">Opened: {ticket.created_at} • {catTitle}</p>
                </div>
                <span className={`px-2 py-0.5 rounded text-[9px] font-bold border ${getStatusStyles(ticket.status)}`}>
                  {ticket.status}
                </span>
              </div>
            );
          })}
        </div>
      )}

      {/* --- STANDARDIZED TERMINAL FOOTER PAGINATION PANEL --- */}
      {filteredTickets.length > 0 && (
        <div className="flex items-center justify-between border-t border-stone-400/10 pt-4 text-[10px] font-mono text-stone-500">
          <div>
            Showing {Math.min(filteredTickets.length, (currentPage - 1) * pageSize + 1)}-
            {Math.min(filteredTickets.length, currentPage * pageSize)} of {filteredTickets.length} records
          </div>
          
          <div className="flex items-center gap-1">
            <button
              type="button"
              disabled={currentPage === 1}
              onClick={() => setCurrentPage(prev => Math.max(1, prev - 1))}
              className="px-2 py-1 bg-stone-200 dark:bg-stone-800 rounded disabled:opacity-30 disabled:cursor-not-allowed text-stone-700 dark:text-stone-300 font-bold border border-stone-400/10"
            >
              PREV
            </button>
            
            {/* Render explicit individual incremental page matrix block keys */}
            {Array.from({ length: totalPages }, (_, idx) => idx + 1).map((pageNum) => (
              <button
                key={pageNum}
                type="button"
                onClick={() => setCurrentPage(pageNum)}
                className={`px-2.5 py-1 rounded border font-bold ${
                  currentPage === pageNum
                    ? "bg-stone-800 text-white border-transparent dark:bg-stone-200 dark:text-stone-950"
                    : "bg-transparent border-stone-400/20 hover:bg-stone-400/10 text-stone-500"
                }`}
              >
                {pageNum}
              </button>
            ))}

            <button
              type="button"
              disabled={currentPage === totalPages}
              onClick={() => setCurrentPage(prev => Math.min(totalPages, prev + 1))}
              className="px-2 py-1 bg-stone-200 dark:bg-stone-800 rounded disabled:opacity-30 disabled:cursor-not-allowed text-stone-700 dark:text-stone-300 font-bold border border-stone-400/10"
            >
              NEXT
            </button>
          </div>
        </div>
      )}
    </div>
  );
}