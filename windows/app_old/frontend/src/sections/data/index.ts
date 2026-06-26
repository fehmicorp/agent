
const App = {
  name: "Fehmi Endpoint Agent",
  version: "v1.0.1",
  tag: "stable"
}

export default App;

export interface SoftwareItem {
  id: number;
  tag: string;
  name: string;
  developer: string;
  category: string;
  version: string;
  size: string;
  status: string;
  architecture: string;
  description: string;
}

export interface SupportTicket {
  id: number | null;
  tktId: string;
  subject: string;
  category: string;
  subcategory: string;
  priority: string;
  status: string;
  created_at: string;
  updated_at: string | null;
  description: string;
  attachment: undefined | null;
}


export interface RaiseProps {
  onTicketCreated: (newTicket: SupportTicket) => void;
}

export interface SubCategoryItem {
  id: number;
  tag: string;
  title: string;
}

export interface CategoryItem {
  id: number;
  tag: string;
  title: string;
  scategory: SubCategoryItem[];
}

export interface TicketsProps {
  tickets: SupportTicket[];
  selectedTicketId: string | null;
  onSelectTicket: (tktId: string | null) => void;
  onAddTimelineComment: (tktId: string, text: string, files: File[]) => void;
  timelineLogs: any[];
}