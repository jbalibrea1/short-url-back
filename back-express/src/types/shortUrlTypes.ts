export interface ShortUrlEntry {
  id: string;
  url: string;
  title?: string | null;
  image?: string | null;
  logo?: string | null;
  description?: string | null;
  totalClicks: number;
  shortUrl: string;
  createdAt: Date;
}

export type NewShortUrlEntry = Omit<
  ShortUrlEntry,
  'shortUrl' | 'createdAt' | 'id' | 'totalClicks'
>;
