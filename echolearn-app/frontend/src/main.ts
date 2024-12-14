const messagesDiv = document.getElementById("messages") as HTMLDivElement;
const messageInput = document.getElementById("messageInput") as HTMLInputElement;
const sendBtn = document.getElementById("sendBtn") as HTMLButtonElement;

// Connect to WebSocket
const protocol = window.location.protocol === "https:" ? "wss:" : "ws:";
const socket = new WebSocket(`${protocol}//localhost:8080/ws`);

socket.onmessage = (event) => {
    const msg = document.createElement("p");
    msg.textContent = event.data;
    messagesDiv.appendChild(msg);
    messagesDiv.scrollTop = messagesDiv.scrollHeight;
};

sendBtn.addEventListener("click", () => {
    sendMessage();
});

messageInput.addEventListener("keyup", (e) => {
    if (e.key === "Enter") {
        sendMessage();
    }
});

function sendMessage() {
    const message = messageInput.value.trim();
    if (message) {
        socket.send(message);
        messageInput.value = "";
    }
}
