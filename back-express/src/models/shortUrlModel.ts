/* eslint-disable */
import mongoose from 'mongoose';
const shortUrlSchema = new mongoose.Schema(
  {
    url: String,
    title: String || null,
    logo: String || null,
    description: String || null,
    shortUrl: { type: String, unique: true },
    totalClicks: { type: Number, default: 0 },
  },
  { timestamps: true }
);

// Transform the returned object to a more readable format
shortUrlSchema.set('toJSON', {
  transform: (_document: any, returnedObject: any) => {
    returnedObject.id = returnedObject._id.toString();
    delete returnedObject._id;
    delete returnedObject.__v;
  },
});

const ShortUrlModel = mongoose.model('ShortUrl', shortUrlSchema);

export default ShortUrlModel;
