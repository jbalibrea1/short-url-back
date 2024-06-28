const mongoose = require('mongoose');

const connectDB = () => {
  mongoose.set('strictQuery', false);

  const url = process.env.MONGODB_URI;
  console.log('connecting to', url);

  mongoose
    .connect(url)
    .then((_result: any) => {
      console.log('connected to MongoDB');
    })
    .catch((error: Error) => {
      console.log('error connecting to MongoDB:', error.message);
    });
};

export default connectDB;
