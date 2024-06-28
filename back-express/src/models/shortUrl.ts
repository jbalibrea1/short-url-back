const mongoose = require('mongoose');

const shortUrlSchema = new mongoose.Schema({
  url: String,
  image: String || null,
  title: String || null,
  logo: String || null,
  description: String || null,
  shortUrl: { type: String, unique: true },
  totalClicks: { type: Number, default: 0 },
  createdAt: { type: Date, default: Date.now },
});

/**
 * Transform the returned object to a more readable format
 *  {
 *  id: '60f1b26b4f3b4b001f3f3b4b',
 *  url: 'https://www.google.com',
 *  shortUrl: 'http://localhost:3000/abc123',
 *  createdAt: '2021-07-17T14:00:59.000Z'
 * }
 *
 * @param {Object} document - The document object
 * @param {Object} returnedObject - The returned object
 * @returns {Object} - The transformed object
 * trsanform like this:
 */
shortUrlSchema.set('toJSON', {
  transform: (_document: any, returnedObject: any) => {
    returnedObject.id = returnedObject._id.toString();
    delete returnedObject._id;
    delete returnedObject.__v;
  },
});

const ShortUrl = mongoose.model('ShortUrl', shortUrlSchema);

export default ShortUrl;
