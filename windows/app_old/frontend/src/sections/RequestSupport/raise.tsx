import React, { useState, useMemo, useEffect, useRef } from "react";
import { cardMain } from "../../utils/colour";
import { support } from "../data/dummy";
import { RaiseProps, SupportTicket } from "../data";

export default function Raise({ onTicketCreated }: RaiseProps): React.JSX.Element {
  const prioritiesList = useMemo(() => {
    return support.priority.length > 0 ? support.priority : ["Low", "Medium", "High", "Critical"];
  }, []);

  // Safely extract limit settings with solid local defaults
  const limits = useMemo(() => {
    return {
      minDesc: support.limit?.minDesc ?? 10,
      maxDesc: support.limit?.maxDesc ?? 200,
      fileConfig: support.limit?.file ?? null
    };
  }, []);

  // Format valid types dynamically into standard mime/extension accept masks: ".png,.jpeg,.jpg"
  const acceptedFormatsString = useMemo(() => {
    if (!limits.fileConfig?.type || limits.fileConfig.type.length === 0) return "image/*";
    return limits.fileConfig.type.map(ext => `.${ext.toLowerCase()}`).join(",");
  }, [limits]);

  // State hooks initialized to placeholder elements
  const [category, setCategory] = useState("null");
  const [subcategory, setSubcategory] = useState("null");
  const [priority, setPriority] = useState("null");
  const [subject, setSubject] = useState("");
  const [description, setDescription] = useState("");
  const [attachedFiles, setAttachedFiles] = useState<File[]>([]);
  const [isSubmitting, setIsSubmitting] = useState(false);
  
  const fileInputRef = useRef<HTMLInputElement>(null);

  // Extract mapped subcategories for whichever root category option is selected
  const availableSubcategories = useMemo(() => {
    if (category === "null") return [];
    const foundCategory = support.categories.find(c => c.title === category);
    return foundCategory ? foundCategory.scategory : [];
  }, [category]);

  // Automatically locate matching subject from recent history logs or mapping schema
  useEffect(() => {
    if (category !== "null" && subcategory !== "null") {
      const matchedTicket = [...support.tickets]
        .reverse()
        .find(t => t.category === category && t.subcategory === subcategory);

      if (matchedTicket) {
        setSubject(matchedTicket.subject);
      } else {
        const currentSubObject = availableSubcategories.find(s => s.title === subcategory);
        if (currentSubObject && currentSubObject.subject && currentSubObject.subject.length > 0) {
          setSubject(currentSubObject.subject[0]);
        } else {
          setSubject(`${category} - ${subcategory} Exception`);
        }
      }
    } else {
      setSubject("");
    }
  }, [category, subcategory, availableSubcategories]);

  // Form Validation Pipeline Matrix Rules
  const isFormValid = useMemo(() => {
    if (category === "null" || subcategory === "null" || priority === "null") return false;
    if (!subject.trim()) return false;
    
    // Validate strict character metric rules
    const descLength = description.trim().length;
    if (descLength < limits.minDesc || descLength > limits.maxDesc) return false;

    // Validate absolute matching file constraints (multi-file arrays)
    if (limits.fileConfig) {
      const isMandatory = (limits.fileConfig as any).manatory ?? limits.fileConfig.mandatory ?? false;
      const fileCount = attachedFiles.length;
      
      if (isMandatory && fileCount === 0) return false;
      
      // Strict upper limit ceiling bounds protection
      const maxLimit = limits.fileConfig.max ?? 10;
      if (fileCount > maxLimit) return false;
    }

    return true;
  }, [category, subcategory, priority, subject, description, attachedFiles, limits]);

  const handleCategoryChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    setCategory(e.target.value);
    setSubcategory("null");
  };

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files && e.target.files.length > 0) {
      const incomingFiles = Array.from(e.target.files);
      const maxAllowed = limits.fileConfig?.max ?? 10;
      
      setAttachedFiles(prev => {
        const updatedList = [...prev, ...incomingFiles];
        // Prevent exceeding explicit max array configuration limits
        if (updatedList.length > maxAllowed) {
          return updatedList.slice(0, maxAllowed);
        }
        return updatedList;
      });
    }
    // Clear field value stream buffer pointer so matching re-uploads trigger cleanly
    if (fileInputRef.current) fileInputRef.current.value = "";
  };

  const removeFile = (indexToRemove: number) => {
    setAttachedFiles(prev => prev.filter((_, idx) => idx !== indexToRemove));
  };

  const handleRaiseTicket = (e: React.FormEvent) => {
    e.preventDefault();
    if (!isFormValid || isSubmitting) return;

    setIsSubmitting(true);

    setTimeout(() => {
      const generatedRandomId = Math.floor(1000 + Math.random() * 9000);
      const formattedTktId = `TKT-${generatedRandomId}`;
      const currentTimestamp = new Date().toISOString().replace('T', ' ').substring(0, 16);

      const generatedTicket: SupportTicket = {
        id: null,
        tktId: formattedTktId,
        subject,
        category,
        subcategory,
        priority,
        status: "Open",
        created_at: currentTimestamp,
        updated_at: null,
        description,
        attachment: undefined
      };

      // Pass actual file payload structure forward inside dispatch context parameters if needed
      onTicketCreated(generatedTicket);
      
      // Flush inputs entirely back to baseline metrics
      setCategory("null");
      setSubcategory("null");
      setPriority("null");
      setSubject("");
      setDescription("");
      setAttachedFiles([]);
      setIsSubmitting(false);
    }, 1500);
  };

  return (
    <div className={`${cardMain} rounded-2xl p-5 border border-stone-200/5 space-y-4`}>
      <h2 className="text-xs font-bold text-stone-900 dark:text-stone-100 font-mono tracking-wider uppercase border-b border-stone-400/10 pb-2">
        Raise New Request
      </h2>

      <form onSubmit={handleRaiseTicket} className="space-y-4 font-mono text-xs">
        <div className="space-y-4">
          
          {/* Main Grid: Category & Priority Selection Fields */}
          <div className="grid grid-cols-2 gap-3">
            <select
              value={category}
              onChange={handleCategoryChange}
              className="w-full bg-stone-100/80 dark:bg-stone-950/90 text-stone-700 dark:text-stone-300 border border-stone-400/30 dark:border-stone-200/10 rounded-md p-2 text-xs font-mono focus:outline-none focus:border-blue-500/50 transition-colors"
            >
              <option value="null">Select Categories</option>
              {support.categories.map((cat) => (
                <option key={cat.id} value={cat.title}>
                  {cat.title}
                </option>
              ))}
            </select>

            <select
              value={priority}
              onChange={(e) => setPriority(e.target.value)}
              className="w-full bg-stone-100/80 dark:bg-stone-950/90 text-stone-700 dark:text-stone-300 border border-stone-400/30 dark:border-stone-200/10 rounded-md p-2 text-xs font-mono focus:outline-none focus:border-blue-500/50 transition-colors"
            >
              <option value="null">Select Priority</option>
              {prioritiesList.map((prio) => (
                <option key={prio} value={prio}>
                  {prio}
                </option>
              ))}
            </select>
          </div>

          {/* Subcategory Select Dropdown Segment Block */}
          <select
            value={subcategory}
            disabled={category === "null"}
            onChange={(e) => setSubcategory(e.target.value)}
            className="w-full bg-stone-100/80 dark:bg-stone-950/90 text-stone-700 dark:text-stone-300 border border-stone-400/30 dark:border-stone-200/10 rounded-md p-2 text-xs font-mono focus:outline-none focus:border-blue-500/50 transition-colors disabled:opacity-40"
          >
            <option value="null">
              {category === "null" ? "Select a category first" : "Select Sub Categories"}
            </option>
            {availableSubcategories.map((subcat) => (
              <option key={subcat.id} value={subcat.title}>
                {subcat.title}
              </option>
            ))}
          </select>

          {/* Prefilled Mapped Subject Display Field Input */}
          <input
            type="text"
            required
            placeholder="Brief summary of system exception..."
            value={subject}
            onChange={(e) => setSubject(e.target.value)}
            className="w-full bg-stone-100/80 dark:bg-stone-950/90 text-stone-700 dark:text-stone-300 border border-stone-400/20 dark:border-stone-200/10 rounded-xl px-3 py-2 focus:outline-none focus:border-blue-500/50 transition-colors"
          />
        </div>

        {/* Textarea Problem Description Form Node Block */}
        <div className="space-y-1">
          <textarea
            required
            rows={4}
            placeholder={`Include full error logs (${limits.minDesc}-${limits.maxDesc} chars)...`}
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            className="w-full bg-stone-100/80 dark:bg-stone-950/90 text-stone-700 dark:text-stone-300 border border-stone-400/20 dark:border-stone-200/10 rounded-xl px-3 py-2.5 focus:outline-none focus:border-blue-500/50 transition-colors font-sans resize-none"
          />
          <div className="flex justify-end text-[10px] text-stone-400/70 font-mono px-1">
            Chars: {description.trim().length} / [{limits.minDesc}-${limits.maxDesc}]
          </div>
        </div>

        {/* Dynamic File Upload Block: Appears only if config rules are present */}
        {limits.fileConfig && (
          <div className="space-y-2">
            <div className="flex items-center justify-between gap-3 bg-stone-100/40 dark:bg-stone-950/40 border border-stone-400/15 rounded-xl px-3 py-1.5">
              <span className="text-[10px] text-stone-500 truncate max-w-[65%]">
                {attachedFiles.length > 0 
                  ? `Selected: ${attachedFiles.length} file(s) [Max: ${limits.fileConfig.max ?? 10}]` 
                  : (
                    ((limits.fileConfig as any).manatory ?? limits.fileConfig.mandatory)
                      ? "⚠️ Schematic diagnostic asset required" 
                      : "No schematic diagnostic asset attached"
                  )
                }
              </span>
              <label className="bg-stone-200 dark:bg-stone-800 hover:opacity-80 text-stone-700 dark:text-stone-300 px-2 py-1 rounded text-[10px] cursor-pointer transition-opacity whitespace-nowrap">
                Choose File
                <input
                  type="file"
                  multiple
                  ref={fileInputRef}
                  accept={acceptedFormatsString}
                  onChange={handleFileChange}
                  className="hidden"
                />
              </label>
            </div>

            {/* Render queue table for multiple attachments preview logs */}
            {attachedFiles.length > 0 && (
              <div className="bg-stone-100/20 dark:bg-stone-950/20 border border-stone-400/10 rounded-xl p-2 space-y-1.5 max-h-28 overflow-y-auto">
                {attachedFiles.map((file, idx) => (
                  <div key={`${file.name}-${idx}`} className="flex items-center justify-between bg-stone-100/60 dark:bg-stone-900/60 rounded px-2 py-1 text-[10px] text-stone-600 dark:text-stone-400">
                    <span className="truncate max-w-[80%]">📎 {file.name}</span>
                    <button
                      type="button"
                      onClick={() => removeFile(idx)}
                      className="text-red-500/80 hover:text-red-500 font-bold px-1 transition-colors"
                      title="Remove asset"
                    >
                      ✕
                    </button>
                  </div>
                ))}
              </div>
            )}
          </div>
        )}

        {/* Dispatch Process Submission Action Button Trigger */}
        <button
          type="submit"
          disabled={!isFormValid || isSubmitting}
          className="w-full py-2.5 bg-stone-800 text-stone-100 dark:bg-stone-100 dark:text-stone-950 font-bold border border-stone-400/20 hover:opacity-90 active:scale-95 rounded-xl transition-all flex items-center justify-center gap-2 disabled:opacity-30 disabled:cursor-not-allowed disabled:scale-100"
        >
          {isSubmitting ? (
            <>
              <span className="h-3 w-3 border-2 border-stone-500 border-t-transparent rounded-full animate-spin" />
              Routing to Queue...
            </>
          ) : (
            "Dispatch Support Ticket"
          )}
        </button>
      </form>
    </div>
  );
}