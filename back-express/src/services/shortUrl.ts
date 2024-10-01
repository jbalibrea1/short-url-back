import ShortUrlModel from '../models/shortUrlModel';
import getMetadata from '../utils/getMetadata';
import generateUniqueShortUrl from '../utils/randomUrl';
import toNewShortUrlEntry from '../utils/shortEntry';
import truncateString from '../utils/truncateString';

const getAll = async () => {
  const allUrls = await ShortUrlModel.find({});
  return allUrls;
};

const create = async (url: String) => {
  const newShortUrlEntry = toNewShortUrlEntry(url);

  //get metadata from  url
  const metadata = await getMetadata(newShortUrlEntry.url);

  // generate a unique short url while not clashing with existing ones
  let shortUrl = generateUniqueShortUrl();
  while (await ShortUrlModel.findOne({ shortUrl })) {
    shortUrl = generateUniqueShortUrl();
  }

  if (metadata.title) newShortUrlEntry.title = metadata.title;
  if (metadata.logo) newShortUrlEntry.logo = metadata.logo;
  if (metadata.description)
    newShortUrlEntry.description = truncateString(metadata.description, 50);

  // create a new entry
  const newEntry = new ShortUrlModel({
    ...newShortUrlEntry,
    shortUrl,
  });
  const savedEntry = await newEntry.save();

  return savedEntry;
};

const get = async (shortUrl: String) => {
  const entry = await ShortUrlModel.findOne({ shortUrl });

  if (!entry) {
    throw new Error(`Short URL not found for ${shortUrl}`);
  }

  entry.totalClicks += 1;
  await entry.save();
  return entry;
};

export default { getAll, create, get };
