import { NewShortUrlEntry, ShortUrlEntry } from '../interfaces/shortUrlTypes';
import ShortUrlModel from '../models/shortUrlModel';
import addMetadata from '../utils/metadata';
import generateUniqueShortUrl from '../utils/randomUrl';
import toNewShortUrlEntry from '../utils/shortEntry';

const getAll = async () => {
  const allUrls: ShortUrlEntry[] = await ShortUrlModel.find({});
  return allUrls;
};

//TODO add type for and destructuring for url
const create = async (url: string) => {
  let newShortUrlEntry: NewShortUrlEntry = toNewShortUrlEntry({ url });

  // add metadata to the entry
  newShortUrlEntry = await addMetadata(newShortUrlEntry);

  // generate a unique short url while not clashing with existing ones
  let shortUrl = generateUniqueShortUrl();
  while (await ShortUrlModel.findOne({ shortUrl })) {
    shortUrl = generateUniqueShortUrl();
  }

  // create a new entry
  const newEntry = new ShortUrlModel({
    ...newShortUrlEntry,
    shortUrl,
  });

  const savedEntry = await newEntry.save();

  return savedEntry;
};

const get = async (shortUrl: string) => {
  const entry = await ShortUrlModel.findOne({ shortUrl });

  if (!entry) {
    throw new Error(`Short URL not found for ${shortUrl}`);
  }

  entry.totalClicks += 1;
  await entry.save();
  return entry;
};

export default { getAll, create, get };
