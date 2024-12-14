// Retrieve username from localStorage or prompt for it
let username = localStorage.getItem("username");
if (!username) {
  username = prompt("Enter your name") || "Anonymous";
  localStorage.setItem("username", username);
}

// Display the username on the page
const welcomeMessage = document.createElement("h2");
welcomeMessage.textContent = `Welcome, ${username}`;
document.body.appendChild(welcomeMessage);

// WebSocket setup
const wsUrl = `ws://${window.location.hostname}:8080/ws`;
const socket = new WebSocket(wsUrl);

// Create UI elements
const messagesDiv = document.createElement("div");
messagesDiv.id = "messages";
messagesDiv.style.border = "1px solid #ccc";
messagesDiv.style.padding = "10px";
messagesDiv.style.height = "300px";
messagesDiv.style.overflowY = "auto";
document.body.appendChild(messagesDiv);

const input = document.createElement("input");
input.type = "text";
input.id = "messageInput";
input.placeholder = "Type a message...";
document.body.appendChild(input);

const sendBtn = document.createElement("button");
sendBtn.textContent = "Send";
document.body.appendChild(sendBtn);

// Fetch previous messages for the user
fetch(`/messages?username=${encodeURIComponent(username)}`)
  .then((res) => res.json())
  .then((messages) => {
    // Display the user's message history
    messages.forEach((msg: { username: string; message: string; timestamp: string }) => {
      displayMessage(msg.username, msg.message, msg.timestamp);
    });
  })
  .catch((err) => console.error("Failed to fetch messages:", err));

// Listen for incoming WebSocket messages
socket.onmessage = (event) => {
  const msg = JSON.parse(event.data);
  displayMessage(msg.username, msg.message, msg.timestamp);
};

// Handle sending messages
sendBtn.onclick = () => sendMessage();
input.addEventListener("keyup", (e) => {
  if (e.key === "Enter") sendMessage();
});

function sendMessage() {
  const message = input.value.trim();
  if (!message) return;

  const msg = { username, message };
  socket.send(JSON.stringify(msg));
  input.value = "";
}

function displayMessage(username: string, message: string, timestamp: string) {
  const p = document.createElement("p");
  const time = new Date(timestamp).toLocaleTimeString();
  p.textContent = `${username}: ${message} (${time})`;
  messagesDiv.appendChild(p);
  messagesDiv.scrollTop = messagesDiv.scrollHeight;
}
