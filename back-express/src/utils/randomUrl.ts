import { nanoid } from 'nanoid';

const generateUniqueShortUrl = (): string => {
  return nanoid(5);
};

export default generateUniqueShortUrl;
