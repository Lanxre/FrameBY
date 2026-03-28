import type { ChatMessage } from "@/types/frontend/chat";

interface WSMessage {
	type: string;
	payload: any;
}

interface NewMessagePayload {
	chat_id: string;
	message: ChatMessage;
}

interface ChatWebSocket {
	isConnected: Ref<boolean>;
	onNewMessage: Ref<((chatId: string, message: ChatMessage) => void) | null>;
	connect: () => void;
	disconnect: () => void;
	joinChat: (chatId: string) => void;
	leaveChat: (chatId: string) => void;
	sendTyping: (chatId: string) => void;
}

let wsInstance: ChatWebSocket | null = null;

function createWebSocket(): ChatWebSocket {
	const ws = ref<WebSocket | null>(null);
	const isConnected = ref(false);
	const onNewMessage = ref<((chatId: string, message: ChatMessage) => void) | null>(null);

	const send = (type: string, payload: any) => {
		if (ws.value && ws.value.readyState === WebSocket.OPEN) {
			ws.value.send(JSON.stringify({ type, payload }));
		} else {
			const checkAndSend = setInterval(() => {
				if (ws.value && ws.value.readyState === WebSocket.OPEN) {
					clearInterval(checkAndSend);
					ws.value.send(JSON.stringify({ type, payload }));
				}
			}, 50);
		}
	};

	const connect = () => {
		if (ws.value && isConnected.value) return;

		const tokenCookie = useCookie("FRAMEBY_ACCESS_TOKEN");
		const token = tokenCookie.value;
		if (!token) {
			console.log("WS: No token");
			return;
		}

		const wsUrl = `${window.location.protocol === "https:" ? "wss:" : "ws:"}//localhost:8080/api/ws/chat?token=${encodeURIComponent(token)}`;
		console.log("WS connecting to:", wsUrl);
		ws.value = new WebSocket(wsUrl);

		ws.value.onopen = () => {
			console.log("WS connected");
			isConnected.value = true;
		};

		ws.value.onclose = () => {
			isConnected.value = false;
			setTimeout(connect, 3000);
		};

		ws.value.onerror = () => {
			isConnected.value = false;
		};

		ws.value.onmessage = (event) => {
			try {
				const data = JSON.parse(event.data);
				console.log("WS received:", data);
				if (data.type === "new_message" && data.chat_id && data.message) {
					if (onNewMessage.value) {
						onNewMessage.value(data.chat_id, data.message);
					}
				}
			} catch (e) {
				console.error("WebSocket parse error:", e);
			}
		};
	};

	const disconnect = () => {
		if (ws.value) {
			ws.value.close();
			ws.value = null;
			isConnected.value = false;
		}
	};

	const joinChat = (chatId: string) => {
		send("join_chat", chatId);
	};

	const leaveChat = (chatId: string) => {
		send("leave_chat", chatId);
	};

	const sendTyping = (chatId: string) => {
		send("typing", { chat_id: chatId });
	};

	return {
		isConnected,
		onNewMessage,
		connect,
		disconnect,
		joinChat,
		leaveChat,
		sendTyping,
	};
}

export function useChatWebSocket(): ChatWebSocket {
	if (!wsInstance) {
		wsInstance = createWebSocket();
	}

	return wsInstance;
}
