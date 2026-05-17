<script setup lang="ts">
import type { Chat, ChatMessage, SearchUser } from "@/types/frontend/chat";
import { useChats, useCreateChat, useChatWebSocket, useMessages, useSearchUsers } from "~/composables/api/chat";
import { $api } from "~/composables/api/useApi";
import { $fetch } from "ofetch";
const { notify } = useNotificationStore();

const { chats, isLoading: isLoadingChats, fetchChats } = useChats();
const { createChat, isLoading: isCreatingChat } = useCreateChat();
const { messages, fetchMessages, sendMessage } = useMessages();
const { results: searchResults, searchUsers, clearResults, isLoading: isSearching } = useSearchUsers();
const ws = useChatWebSocket();

const activeChatId = ref<string | null>(null);
const activeMessages = ref<ChatMessage[]>([]);
const message = ref("");
const isSending = ref(false);
const fileInputRef = ref<HTMLInputElement | null>(null);
const messagesContainerRef = ref<HTMLElement | null>(null);
const searchQuery = ref("");
const showSearchResults = ref(false);
const typingTimeout = ref<ReturnType<typeof setTimeout> | null>(null);

const activeChat = computed(() => chats.value.find((c) => c.id === activeChatId.value));

const getOtherUserName = (chat: Chat): string => {
	if (chat.name) return chat.name;
	return "Администратор";
};

const scrollToBottom = () => {
	nextTick(() => {
		if (messagesContainerRef.value) {
			messagesContainerRef.value.scrollTo({
				top: messagesContainerRef.value.scrollHeight,
				behavior: "smooth",
			});
		}
	});
};

watch(activeMessages, () => scrollToBottom(), { deep: true });
watch(activeChatId, () => scrollToBottom());

onMounted(async () => {
	await fetchChats();
	ws.connect();

	ws.onNewMessage.value = (chatId: string, msg: ChatMessage) => {
		if (chatId === activeChatId.value) {
			const exists = activeMessages.value.some((m) => m.id === msg.id);
			if (!exists) {
				activeMessages.value.push(msg);
			}
			$api(`/chats/${chatId}/read`, { method: "POST" }).catch(() => {});
		}

		const chat = chats.value.find((c) => c.id === chatId);
		if (chat) {
			if (chatId !== activeChatId.value) {
				chat.unread_count = (chat.unread_count || 0) + 1;
			} else {
				chat.unread_count = 0;
			}
			chat.last_message = msg.content;
			chat.last_message_at = msg.created_at;
		}
	};
});

onUnmounted(() => {
	ws.disconnect();
});

watch(searchQuery, (query) => {
	if (query.length >= 2) {
		showSearchResults.value = true;
		searchUsers(query);
	} else {
		showSearchResults.value = false;
		clearResults();
	}
});

const selectChat = async (chat: Chat) => {
	if (activeChatId.value) {
		ws.leaveChat(activeChatId.value);
	}

	activeChatId.value = chat.id;
	chat.unread_count = 0;

	await fetchMessages(chat.id);
	activeMessages.value = messages.value;

	await $api(`/chats/${chat.id}/read`, { method: "POST" }).catch(() => {});

	if (ws.isConnected.value) {
		ws.joinChat(chat.id);
	} else {
		const checkConnection = setInterval(() => {
			if (ws.isConnected.value) {
				clearInterval(checkConnection);
				ws.joinChat(chat.id);
			}
		}, 100);
	}
};

const openChatWithUser = async (user: SearchUser) => {
	showSearchResults.value = false;
	searchQuery.value = "";
	clearResults();

	try {
		const chatData = await $fetch<{ id: string; type: string; name?: string }>("/chats/find", {
			baseURL: "/api",
			method: "GET",
			query: { user_id: user.user_id },
		});
		console.log("Found chat:", chatData);
		
		const chatFromList = chats.value.find((c) => c.id === chatData.id);
		if (chatFromList) {
			selectChat(chatFromList);
			return;
		}
		
		await fetchChats();
		const chat = chats.value.find((c) => c.id === chatData.id);
		if (chat) {
			selectChat(chat);
			return;
		}
	} catch (e: any) {
		if (e.response?.status === 404) {
			console.log("No existing chat, creating new");
		} else {
			console.error("Error finding chat:", e);
		}
	}

	const newChat = await createChat(user.user_id);
	if (newChat && newChat.id) {
		await fetchChats();
		const createdChat = chats.value.find((c) => c.id === newChat.id);
		if (createdChat) {
			selectChat(createdChat);
		} else {
			selectChat(newChat);
		}
		notify({
			title: "Успех",
			content: "Чат создан",
			type: "success",
		});
	}
};

const sendChatMessage = async () => {
	if (!message.value.trim() || !activeChatId.value || isSending.value) return;

	const text = message.value;
	message.value = "";
	isSending.value = true;

	await sendMessage(activeChatId.value, "text", text);

	const chat = chats.value.find((c) => c.id === activeChatId.value);
	if (chat) {
		chat.last_message = text;
		chat.last_message_at = new Date().toISOString();
	}

	isSending.value = false;
};

const handleTyping = () => {
	if (activeChatId.value) {
		ws.sendTyping(activeChatId.value);
	}

	if (typingTimeout.value) {
		clearTimeout(typingTimeout.value);
	}
};

const openFileDialog = () => {
	fileInputRef.value?.click();
};

const handleFileChange = async (e: Event) => {
	const files = (e.target as HTMLInputElement).files;
	if (!files || !activeChatId.value) return;

	for (const file of Array.from(files)) {
		const reader = new FileReader();
		reader.onload = async (event) => {
			const content = event.target?.result as string;
			const type = file.type.startsWith("image") ? "image" : "file";

			await sendMessage(activeChatId.value!, type, content);

			const chat = chats.value.find((c) => c.id === activeChatId.value);
			if (chat) {
				chat.last_message = file.name;
				chat.last_message_at = new Date().toISOString();
			}
		};
		reader.readAsDataURL(file);
	}

	if (fileInputRef.value) {
		fileInputRef.value.value = "";
	}
};

const formatTime = (dateStr: string | undefined): string => {
	if (!dateStr) return "";
	const date = new Date(dateStr);
	return date.toLocaleTimeString([], { hour: "2-digit", minute: "2-digit" });
};
</script>

<template>
	<div class="grid grid-cols-[320px_1fr] h-[600px] border border-emerald-200 rounded-2xl overflow-hidden bg-white shadow-sm">
		<div class="flex flex-col border-r border-emerald-100 bg-white h-full overflow-hidden">
			<div class="p-4 border-b border-emerald-100 shrink-0">
				<input
					v-model="searchQuery"
					type="text"
					placeholder="Поиск пользователя..."
					class="w-full px-3 py-2 rounded-xl text-sm bg-gray-50 border border-emerald-100 focus:outline-none focus:ring-2 focus:ring-emerald-400/40"
					@focus="showSearchResults = searchQuery.length >= 2"
				/>
			</div>

			<div v-if="showSearchResults && searchResults.length > 0" class="border-b border-emerald-100 max-h-64 overflow-y-auto custom-scrollbar">
				<div
					v-for="user in searchResults"
					:key="user.user_id"
					@click="openChatWithUser(user)"
					class="px-4 py-3 cursor-pointer hover:bg-emerald-50 border-b border-emerald-50 transition"
				>
					<div class="flex items-center gap-3">
						<div class="w-8 h-8 rounded-full bg-emerald-100 flex items-center justify-center">
    						<img v-if="user!.avatar" :src="formatAvatar(user!.login, user!.avatar)" alt="Avatar" class="w-full h-full object-cover rounded-full" />
                            <span v-else>{{ user!.login.charAt(0).toUpperCase() }}</span>
						</div>
						<div class="min-w-0 flex-1">
							<p class="text-sm font-medium text-gray-800 truncate">{{ user.profile!.full_name }}</p>
							<p class="text-xs text-gray-500 truncate">@{{ user.login }}</p>
						</div>
					</div>
				</div>
			</div>

			<div class="flex-1 overflow-y-auto custom-scrollbar">
				<div v-if="isLoadingChats && chats.length === 0" class="flex justify-center py-8">
					<div class="animate-spin rounded-full h-6 w-6 border-b-2 border-emerald-600"></div>
				</div>

				<div v-else-if="chats.length === 0" class="p-4 text-center text-gray-500 text-sm">
					Нет чатов. Начните поиск пользователя слева.
				</div>

				<div v-else>
					<div
						v-for="chat in chats"
						:key="chat.id"
						@click="selectChat(chat)"
						:class="[
							'px-4 py-3 cursor-pointer border-b border-emerald-50 relative',
							activeChatId === chat.id ? 'bg-emerald-50' : 'hover:bg-emerald-50/50',
						]"
					>
						<div class="flex justify-between items-center">
							<div class="text-sm font-medium text-gray-800 truncate flex items-center gap-2">
    				            <div class="w-6 h-6 rounded-full bg-emerald-100 flex items-center justify-center">
                                    <img v-if="chat.avatar" :src="formatAvatar(chat.name!, chat!.avatar)" alt="Avatar" class="object-cover rounded-full" />
                                    <span v-else class="text-[14px]">{{ chat.name!.charAt(0).toUpperCase() }}</span>     
                                </div>
							    <span class="truncate w-[125px]">
									{{ getOtherUserName(chat) }}
								</span>
								<span
									v-if="(chat.unread_count || 0) > 0"
									class="px-1.5 py-0.5 text-xs rounded-full bg-emerald-500 text-white"
								>
									{{ chat.unread_count }}
								</span>
							</div>
							<span class="text-xs text-gray-400">{{ formatTime(chat.last_message_at) }}</span>
						</div>
						<p class="text-xs text-gray-500 truncate mt-0.5 ml-8">{{ chat.last_message || "Начните общение" }}</p>
					</div>
				</div>
			</div>
		</div>

		<div class="grid grid-rows-[auto_1fr_auto] h-full max-h-full overflow-hidden bg-emerald-50/20">
			<div class="px-4 py-3 border-b border-emerald-100 bg-white/80 backdrop-blur z-10">
                <div class="flex flex-row items-center gap-4">
                    <div v-if="activeChat" class="w-8 h-8 rounded-full bg-emerald-100 flex items-center justify-center">
        				<img v-if="activeChat.avatar" :src="formatAvatar(activeChat.name!, activeChat.avatar!)" alt="Avatar" class="w-full h-full object-cover rounded-full" />
                        <span v-else class="text-[14px]">{{ activeChat.name!.charAt(0).toUpperCase() }}</span>    
         			</div>
    			    <p class="text-sm font-semibold text-gray-800">
    					{{ activeChat ? getOtherUserName(activeChat) : "Чат не выбран" }}
    				</p>
                </div>
			</div>

			<div ref="messagesContainerRef" class="min-h-0 overflow-y-auto px-4 py-4 space-y-4 custom-scrollbar">
				<template v-if="activeChatId">
					<div v-if="activeMessages.length === 0" class="flex flex-col items-center justify-center h-full text-gray-400">
						<Icon name="ph:chat-circle-dots" size="48" class="mb-2" />
						<p class="text-sm">Начните общение</p>
					</div>

					<div
						v-for="msg in activeMessages"
						:key="msg.id"
						class="flex"
						:class="msg.user_id === useAuthStore().user?.id ? 'justify-end' : 'justify-start'"
					>
						<div
							class="max-w-[85%] px-3 py-2 rounded-2xl text-sm break-words shadow-sm"
							:class="
								msg.user_id === useAuthStore().user?.id
									? 'bg-emerald-500 text-white rounded-tr-none'
									: 'bg-white border border-emerald-100 text-gray-700 rounded-tl-none'
							"
						>
							<template v-if="msg.type === 'image'">
								<img :src="msg.content" class="rounded-lg max-w-full max-h-64" />
							</template>
							<template v-else-if="msg.type === 'file'">
								<a :href="msg.content" download class="flex items-center gap-2">
									<Icon name="ph:file" size="18" />
									<span class="truncate">Файл</span>
								</a>
							</template>
							<template v-else>{{ msg.content }}</template>
							<p
								class="text-xs mt-1 opacity-60"
								:class="msg.user_id === useAuthStore().user?.id ? 'text-right' : ''"
							>
								{{ formatTime(msg.created_at) }}
							</p>
						</div>
					</div>
				</template>

				<div v-else class="flex flex-col items-center justify-center h-full text-gray-400">
					<Icon name="ph:hand-pointing" size="48" class="mb-2" />
					<p class="text-sm">Выберите чат или найдите пользователя</p>
				</div>
			</div>

			<div v-if="activeChatId" class="p-4 border-t border-emerald-100 bg-white z-10">
				<div class="flex items-center gap-2 h-10">
					<button
						@click="openFileDialog"
						class="p-2 rounded-xl bg-gray-50 text-gray-500 hover:bg-emerald-50 transition shrink-0 flex items-center justify-center h-full w-10 cursor-pointer"
					>
						<Icon name="ph:paperclip" size="20" />
					</button>

					<input ref="fileInputRef" type="file" multiple class="hidden" @change="handleFileChange" />

					<input
						v-model="message"
						@input="handleTyping"
						@keydown.enter.prevent="sendChatMessage"
						type="text"
						placeholder="Введите сообщение..."
						class="flex-1 min-w-0 px-4 h-full rounded-xl text-sm bg-gray-50 border border-emerald-400 focus:outline-none focus:ring-2 focus:ring-emerald-400/40 focus:bg-white transition-all"
					/>

					<button
						@click="sendChatMessage"
						:disabled="isSending || !message.trim()"
						class="rounded-xl text-white bg-emerald-500 hover:bg-emerald-600 transition shrink-0 flex items-center justify-center h-full w-10 cursor-pointer disabled:opacity-50"
					>
						<Icon v-if="!isSending" name="ph:paper-plane-right-fill" size="20" />
						<div v-else class="w-5 h-5 border-2 border-white border-t-transparent rounded-full animate-spin"></div>
					</button>
				</div>
			</div>
		</div>
	</div>
</template>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
	width: 6px;
}

.custom-scrollbar::-webkit-scrollbar-track {
	background: transparent;
}

.custom-scrollbar::-webkit-scrollbar-thumb {
	background: #d1d5db;
	border-radius: 10px;
}

.custom-scrollbar::-webkit-scrollbar-thumb:hover {
	background: #10b981;
}

.custom-scrollbar {
	scrollbar-width: thin;
	scrollbar-color: #d1d5db transparent;
}
</style>
