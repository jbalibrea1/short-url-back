export interface ShortUrlEntry {
  id: string;
  url: string;
  title?: string | null;
  logo?: string | null;
  description?: string | null;
  totalClicks: number;
  shortUrl: string;
  createdAt: Date;
  updatedAt: Date;
}

export type NewShortUrlEntry = Omit<
  ShortUrlEntry,
  'shortUrl' | 'createdAt' | 'updatedAt' | 'id' | 'totalClicks'
>;
