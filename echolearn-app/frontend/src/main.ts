// Function to fetch user location based on IP
async function fetchUserLocation(): Promise<string> {
    try {
      const response = await fetch("https://ipwhois.app/json/");
      const data = await response.json();
      return data.country || "Unknown";
    } catch (error) {
      console.error("Failed to fetch location:", error);
      return "Unknown";
    }
  }
  
  // Retrieve or prompt for username
  let username = localStorage.getItem("username");
  if (!username) {
    username = prompt("Enter your name") || "Anonymous";
    localStorage.setItem("username", username);
  }
  
  // Fetch user location
  let userLocation = localStorage.getItem("userLocation");
  if (!userLocation) {
    userLocation = await fetchUserLocation();
    localStorage.setItem("userLocation", userLocation);
  }
  
  // Update the welcome message with username and location
  const welcomeMsgEl = document.querySelector("#welcome-message");
  if (welcomeMsgEl) {
    welcomeMsgEl.innerHTML = `Welcome, <span>${username} (${userLocation})</span>`;
  }
  
  // Set up WebSocket connection
  const wsUrl = `ws://${window.location.hostname}:8080/ws`;
  const socket = new WebSocket(wsUrl);
  
  // UI elements
  const messagesDiv = document.querySelector("#messages") as HTMLDivElement;
  const input = document.querySelector("#messageInput") as HTMLInputElement;
  const sendBtn = document.querySelector("#sendBtn") as HTMLButtonElement;
  
  // Fetch user's historical messages
  fetch(`/messages?username=${encodeURIComponent(username!)}`)
    .then((res) => res.json())
    .then((messages) => {
      messages.forEach((msg: { username: string; location: string; message: string; timestamp: string }) => {
        displayMessage(msg.username, msg.location, msg.message, msg.timestamp);
      });
    })
    .catch((err) => console.error("Failed to fetch messages:", err));
  
  // Listen for incoming WebSocket messages
  socket.onmessage = (event) => {
    const msg = JSON.parse(event.data);
    displayMessage(msg.username, msg.location, msg.message, msg.timestamp);
  };
  
  // Handle sending messages
  sendBtn.onclick = () => sendMessage();
  input.addEventListener("keyup", (e) => {
    if (e.key === "Enter") sendMessage();
  });
  
  function sendMessage() {
    const message = input.value.trim();
    if (!message) return;
  
    // Include username and userLocation in the message payload
    const msg = { username, location: userLocation, message };
    socket.send(JSON.stringify(msg));
    input.value = "";
  }
  
  function displayMessage(username: string, location: string, message: string, timestamp: string) {
    const messageBubble = document.createElement("p");
    const formattedTime = new Date(timestamp).toLocaleTimeString();
    messageBubble.textContent = `${username} (${location}): ${message} (${formattedTime})`;
  
    // Highlight the user's own messages differently
    if (username === localStorage.getItem("username")) {
      messageBubble.classList.add("self");
    }
  
    messagesDiv.appendChild(messageBubble);
    messagesDiv.scrollTop = messagesDiv.scrollHeight;
  }

export { };
  