<script setup lang="ts">
import ModalWindow from "@/components/common/ModalWindow.vue";
import { useExportSquad } from "@/composables/api/squads/useExportSquad";

const props = defineProps<{
	modelValue: boolean;
	myOnly?: boolean;
	title?: string;
	squadId?: string | null;
}>();

const emit = defineEmits(["update:modelValue"]);

const { exportAll, exportByID, downloading } = useExportSquad();

const close = () => emit("update:modelValue", false);

const formats = [
	{ key: "pdf", label: "PDF", icon: "ph:file-pdf" },
	{ key: "html", label: "HTML", icon: "ph:file-html" },
	{ key: "md", label: "Markdown", icon: "ph:file-text" },
] as const;

const doExport = async (format: string) => {
	if (props.squadId) {
		await exportByID(props.squadId, format);
	} else {
		await exportAll(format, props.myOnly ?? true);
	}
	close();
};
</script>

<template>
	<ModalWindow
		:model-value="modelValue"
		:title="title"
		center-title
		width="max-w-sm"
		@close="close"
		@update:model-value="(v) => emit('update:modelValue', v)"
	>
		<div class="flex flex-col gap-3 py-2">
			<button
				v-for="fmt in formats"
				:key="fmt.key"
				:disabled="downloading"
				@click="doExport(fmt.key)"
				class="flex items-center gap-3 px-4 py-3 rounded-xl border border-emerald-100 bg-white hover:bg-emerald-50 transition disabled:opacity-50 cursor-pointer"
			>
				<Icon :name="fmt.icon" size="24" class="text-emerald-500" />
				<span class="text-sm font-medium text-gray-700">{{ fmt.label }}</span>
			</button>
		</div>
	</ModalWindow>
</template>
