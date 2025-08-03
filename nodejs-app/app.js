// Import the express library, a minimal and flexible Node.js web application framework.
const express = require('express');

// --- Constants ---
// The port the server will listen on. 8080 is a common port for web development.
const PORT = 8080;
// The host the server will bind to. '0.0.0.0' means it will accept connections on all available network interfaces.
const HOST = '0.0.0.0';

// --- App ---
// Create an instance of an Express application.
const app = express();

// Define a route handler for GET requests to the root URL ('/').
app.get('/', (req, res) => {
  // Send a string response to the client.
  res.send('Hello from Node.js!');
});

// Start the server and have it listen for connections on the specified host and port.
app.listen(PORT, HOST, () => {
  // Log a message to the console once the server is running.
  console.log(`Node.js server running on http://${HOST}:${PORT}`);
});
