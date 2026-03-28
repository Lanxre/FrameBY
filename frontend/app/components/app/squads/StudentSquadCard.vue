<script setup lang="ts">
import { SQUAD_STATUS_LABELS, SQUAD_STATUS_COLORS } from "@/const/squad";
import type { StudentSquad } from "@/types/frontend/student-squad";
import type { FramebyAppRole } from "~/types/frontend/enums/role";
import ToolTip from "@/components/ui/ToolTip.vue";

const props = withDefaults(
	defineProps<{
		squad: StudentSquad;
		isJoining?: boolean;
		isLeaving?: boolean;
		isLoading?: boolean;
		isUniversityMode?: boolean;
		onJoin?: (id: string) => void;
		onLeave?: (id: string) => void;
		onEdit?: (squad: StudentSquad) => void;
		onClose?: (squadId: string) => void;
		onOpen?: (squadId: string) => void;
		onApprove?: (squad: StudentSquad) => void;
		onReject?: (squad: StudentSquad) => void;
	}>(),
	{
		isLoading: false,
		isUniversityMode: false,
	},
);

const canJoin = computed(
	() =>
		props.squad.status === "recruitment_open" &&
		props.squad.current_count < props.squad.max_participants &&
		props.onJoin,
);

const canLeave = computed(
	() => props.squad.status === "recruitment_open" && props.onLeave,
);

const canEdit = computed(() => props.onEdit !== undefined);
const canClose = computed(
	() =>
		!props.isUniversityMode && props.squad.status !== "closed" && props.onClose,
);
const canOpen = computed(() => props.squad.status === "closed" && props.onOpen);
const canApprove = computed(
	() =>
		props.isUniversityMode &&
		props.squad.status === "pending" &&
		props.onApprove,
);
const canReject = computed(
	() =>
		props.isUniversityMode &&
		props.squad.status === "pending" &&
		props.onReject,
);

const isPending = computed(() => props.squad.status === "pending");
const isRejected = computed(() => props.squad.status === "rejected");
</script>

<template>
	<div class="p-5 rounded-2xl border border-emerald-100 bg-white shadow-md hover:shadow-lg transition-shadow duration-300">
		<div class="flex items-center justify-between gap-2 mb-4">
			<div class="flex items-center gap-2 min-w-0">
				<div class="w-1 h-8 rounded-full bg-gradient-to-b from-emerald-400 to-green-600 shrink-0"></div>
				<h3 class="font-bold text-lg text-gray-800 truncate">
					{{ squad.title }}
				</h3>
				<span
					class="px-2.5 py-1 text-xs font-medium rounded-full shrink-0"
					:class="[
						SQUAD_STATUS_COLORS[squad.status]?.bg,
						SQUAD_STATUS_COLORS[squad.status]?.text,
					]"
				>
					{{ SQUAD_STATUS_LABELS[squad.status] }}
				</span>
			</div>
			<div class="flex items-center justify-center gap-1 shrink-0">
				<template v-if="isUniversityMode">
					<ToolTip text="Одобрить">
						<button
							v-if="canApprove"
							@click="onApprove?.(squad)"
							:disabled="isLoading"
							class="p-2 rounded-xl text-gray-400 hover:text-emerald-600 hover:bg-emerald-50 transition-all disabled:opacity-50"
						>
							<Icon name="ph:check-circle" size="20" />
						</button>
					</ToolTip>

					<ToolTip text="Отклонить">
						<button
							v-if="canReject"
							@click="onReject?.(squad)"
							:disabled="isLoading"
							class="p-2 rounded-xl text-gray-400 hover:text-red-600 hover:bg-red-50 transition-all disabled:opacity-50"
						>
							<Icon name="ph:x-circle" size="20" />
						</button>
					</ToolTip>

					<ToolTip text="Закрыть набор">
						<button
							v-if="squad.status === 'recruitment_open'"
							@click="onClose?.(squad.id)"
							:disabled="isLoading"
							class="p-2 rounded-xl text-gray-400 hover:text-red-600 hover:bg-red-50 transition-all disabled:opacity-50"
						>
							<Icon name="ph:x" size="20" />
						</button>
					</ToolTip>
				</template>

				<template v-else>
					<ToolTip text="Редактировать">
						<button
							v-if="canEdit"
							@click="onEdit?.(squad)"
							class="p-2 rounded-xl text-gray-400 hover:text-emerald-600 hover:bg-emerald-50 transition-all"
						>
							<Icon name="ph:pencil" size="20" />
						</button>
					</ToolTip>

					<ToolTip text="Закрыть">
						<button
							v-if="canClose"
							@click="onClose?.(squad.id)"
							class="p-2 rounded-xl text-gray-400 hover:text-red-600 hover:bg-red-50 transition-all"
						>
							<Icon name="ph:x" size="20" />
						</button>
					</ToolTip>

					<ToolTip text="Открыть">
						<button
							v-if="canOpen"
							@click="onOpen?.(squad.id)"
							class="p-2 rounded-xl text-gray-400 hover:text-green-600 hover:bg-green-50 transition-all"
						>
							<Icon name="ph:plus" size="20" />
						</button>
					</ToolTip>
				</template>
			</div>
		</div>

		<div class="bg-gray-50 rounded-xl p-3 mb-4">
			<div class="flex items-center gap-3">
				<div class="w-10 h-10 rounded-xl bg-gradient-to-br from-emerald-400 to-green-600 flex items-center justify-center shrink-0">
					<Icon name="ph:user" class="text-white" size="20" />
				</div>
				<div class="min-w-0 flex-1">
					<p class="font-semibold text-gray-800 truncate">
						{{ squad.organizer.name || "Неизвестно" }}
					</p>
					<p v-if="squad.organizer.role" class="text-xs text-gray-500">
						{{ formatRole(squad.organizer.role as FramebyAppRole) }}
					</p>
				</div>
			</div>
			<div v-if="squad.organizer.position || squad.organizer.phone" class="mt-2 pt-2 border-t border-gray-200 grid grid-cols-2 gap-2 text-xs text-gray-500">
				<div v-if="squad.organizer.position" class="flex items-center gap-1 truncate">
					<Icon name="ph:briefcase" size="12" class="shrink-0" />
					<span class="truncate">{{ squad.organizer.position }}</span>
				</div>
				<div v-if="squad.organizer.phone" class="flex items-center gap-1 truncate">
					<Icon name="ph:phone" size="12" class="shrink-0" />
					<span class="truncate">{{ squad.organizer.phone }}</span>
				</div>
			</div>
		</div>

		<div v-if="squad.profile" class="flex items-center gap-2 mb-3 text-sm">
			<div class="w-6 h-6 rounded-lg bg-teal-50 flex items-center justify-center shrink-0">
				<Icon name="ph:briefcase" size="14" class="text-teal-600" />
			</div>
			<span class="text-emerald-600 font-medium truncate">{{ squad.profile }}</span>
		</div>

		<p v-if="squad.description" class="text-sm text-gray-500 line-clamp-2 mb-4">
			{{ squad.description }}
		</p>

		<div class="flex items-center justify-between text-sm mb-2">
			<div class="flex items-center gap-2">
				<div class="w-7 h-7 rounded-lg bg-emerald-50 flex items-center justify-center shrink-0">
					<Icon name="ph:users" size="14" class="text-emerald-600" />
				</div>
				<span class="text-gray-600">
					<span class="font-semibold text-emerald-600">{{ squad.current_count }}</span>
					<span class="text-gray-400"> / {{ squad.max_participants }}</span>
				</span>
			</div>
			<span class="text-xs font-medium" :class="squad.current_count >= squad.max_participants ? 'text-green-600' : 'text-gray-400'">
				{{ Math.round((squad.current_count / squad.max_participants) * 100) }}%
			</span>
		</div>

		<div v-if="squad.status === 'recruitment_open'" class="mb-5">
			<div class="h-2 bg-gray-100 rounded-full overflow-hidden">
				<div
					class="h-full bg-gradient-to-r from-emerald-400 via-green-500 to-teal-400 transition-all duration-500 rounded-full"
					:style="{
						width: `${Math.min((squad.current_count / squad.max_participants) * 100, 100)}%`,
					}"
				/>
			</div>
		</div>

		<div v-if="isPending" class="mb-4 p-3 bg-amber-50 rounded-xl border border-amber-100">
			<div class="flex items-center gap-2 text-amber-700 text-sm">
				<Icon name="ph:hourglass-medium" size="18" />
				<span class="font-medium">Ожидает подтверждения</span>
			</div>
		</div>

		<div class="flex gap-2">
			<button
				v-if="canJoin"
				@click="onJoin?.(squad.id)"
				:disabled="isJoining"
				class="flex-1 px-4 py-2.5 rounded-xl text-sm font-semibold bg-gradient-to-r from-emerald-500 to-green-600 text-white flex items-center justify-center gap-2 hover:from-emerald-600 hover:to-green-700 active:scale-[0.98] transition-all disabled:opacity-50 shadow-md shadow-emerald-100"
			>
				<Icon v-if="!isJoining" name="ph:user-plus" size="18" />
				<div v-else class="w-4 h-4 border-2 border-white border-t-transparent rounded-full animate-spin"></div>
				Вступить
			</button>

			<button
				v-if="canLeave"
				@click="onLeave?.(squad.id)"
				:disabled="isLeaving"
				class="flex-1 px-4 py-2.5 rounded-xl text-sm font-semibold bg-white border border-red-200 text-red-600 flex items-center justify-center gap-2 hover:bg-red-50 active:scale-[0.98] transition-all disabled:opacity-50"
			>
				<Icon v-if="!isLeaving" name="ph:user-minus" size="18" />
				<div v-else class="w-4 h-4 border-2 border-red-600 border-t-transparent rounded-full animate-spin"></div>
				Покинуть
			</button>
		</div>
	</div>
</template>
